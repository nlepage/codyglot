package filestore

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	service "github.com/nlepage/codyglot/service/filestore"
)

type fileMessageRecver interface {
	Recv() (*service.FileMessage, error)
}

func recv(rcver fileMessageRecver, dir string) error {
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

	if err := os.MkdirAll(dir, dirMode); err != nil {
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
