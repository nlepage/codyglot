package filestore

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/nlepage/codyglot/filestore/config"
	service "github.com/nlepage/codyglot/service/filestore"
)

const (
	DirMode os.FileMode = 0755
)

type FileMessageRecver interface {
	Recv() (*service.FileMessage, error)
}

func Write(rcver FileMessageRecver, dir string) error {
	// FIXME wrap errors

	var (
		path string
		ch   = make(chan []byte)
		wg   sync.WaitGroup
	)

	for message, err := rcver.Recv(); err != io.EOF; message, err = rcver.Recv() {
		if err != nil {
			return err
		}

		switch fx := message.GetFileMessage().(type) {
		case *service.FileMessage_FileInfo:
			close(ch)
			path = fx.FileInfo.Path
			ch = make(chan []byte)
			wg.Add(1)
			go writeFile(filepath.Join(dir, path), os.FileMode(fx.FileInfo.Chmod), ch, &wg)
		case *service.FileMessage_FileContent:
			ch <- fx.FileContent.Content
		}
	}

	close(ch)

	wg.Wait()

	return nil
}

func writeFile(path string, chmod os.FileMode, ch <-chan []byte, wg *sync.WaitGroup) {
	// FIXME wrap errors

	defer wg.Done()

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, DirMode); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, chmod)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b := range ch {
		if _, err := f.Write(b); err != nil {
			panic(err)
		}
	}
}

type FileMessageSender interface {
	Send(*service.FileMessage) error
}

func SendFile(sender FileMessageSender, path string, relPath string, info os.FileInfo, cfg config.FileStoreConfig) error {
	// FIXME wrap errors

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	sender.Send(&service.FileMessage{
		FileMessage: &service.FileMessage_FileInfo{
			FileInfo: &service.FileInfo{
				Path:  relPath,
				Chmod: int32(info.Mode()),
			},
		},
	})

	// FIXME allocate smaller cap if file is small
	b := make([]byte, cfg.ChunkSize)

	for {
		i, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		sender.Send(&service.FileMessage{
			FileMessage: &service.FileMessage_FileContent{
				FileContent: &service.FileContent{
					Content: b[:i],
				},
			},
		})
	}

	return nil
}

func SendDir(sender FileMessageSender, dir string, cfg config.FileStoreConfig, includeDirName bool) error {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	dirName := filepath.Base(dir)

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// FIXME wrap errors

		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if includeDirName {
			relPath = filepath.Join(dirName, relPath)
		}

		if !info.IsDir() {
			if err := SendFile(sender, path, relPath, info, cfg); err != nil {
				return err
			}
		}

		return nil
	})
}
