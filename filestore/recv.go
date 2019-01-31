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

type FilesWriter interface {
	Open(path string, chmod os.FileMode) io.WriteCloser
}

func recv(rcver fileMessageRecver, fw FilesWriter) error {
	// FIXME wrap errors

	var (
		ch = make(chan []byte)
		wg sync.WaitGroup
	)

	for message, err := rcver.Recv(); err != io.EOF; message, err = rcver.Recv() {
		if err != nil {
			return err
		}

		switch fx := message.GetFileMessage().(type) {
		case *service.FileMessage_FileInfo:
			close(ch)
			ch = make(chan []byte)
			wg.Add(1)
			go writeFile(fw.Open(fx.FileInfo.Path, os.FileMode(fx.FileInfo.Chmod)), ch, &wg)
		case *service.FileMessage_FileContent:
			ch <- fx.FileContent.Content
		}
	}

	close(ch)

	wg.Wait()

	return nil
}

func writeFile(w io.WriteCloser, ch <-chan []byte, wg *sync.WaitGroup) {
	// FIXME wrap errors

	defer w.Close()
	defer wg.Done()

	for p := range ch {
		if _, err := w.Write(p); err != nil {
			panic(err)
		}
	}
}

func DiskFilesWriter(dir string) FilesWriter {
	return diskFilesWriter(dir)
}

type diskFilesWriter string

var _ FilesWriter = diskFilesWriter("")

func (dir diskFilesWriter) Open(path string, chmod os.FileMode) io.WriteCloser {
	return &diskWriter{
		dir:   string(dir),
		path:  path,
		chmod: chmod,
	}
}

type diskWriter struct {
	dir   string
	path  string
	chmod os.FileMode
	f     *os.File
}

var _ io.WriteCloser = (*diskWriter)(nil)

func (dw *diskWriter) Write(p []byte) (n int, err error) {
	if dw.f == nil {
		path := filepath.Join(dw.dir, dw.path)

		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, dirMode); err != nil {
			return 0, err
		}

		dw.f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, dw.chmod)
		if err != nil {
			return 0, err
		}
	}

	return dw.f.Write(p)
}

func (dw *diskWriter) Close() error {
	return dw.f.Close()
}
