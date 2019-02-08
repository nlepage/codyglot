package filestore

import (
	"io"
	"os"
	"sync"

	service "github.com/nlepage/codyglot/service/filestore"
)

type fileMessageRecver interface {
	Recv() (*service.FileMessage, error)
}

func recv(rcver fileMessageRecver, fw Writer) error {
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

			// TODO buffered channel
			ch = make(chan []byte)
			wg.Add(1)
			go recvFile(fw, fx.FileInfo.Path, os.FileMode(fx.FileInfo.Chmod), ch, &wg)
		case *service.FileMessage_FileContent:
			ch <- fx.FileContent.Content
		}
	}

	close(ch)

	wg.Wait()

	return nil
}

func recvFile(fw Writer, path string, chmod os.FileMode, ch <-chan []byte, wg *sync.WaitGroup) {
	// FIXME wrap errors
	defer wg.Done()

	defer func() {
		// Drain channel to avoid blocking sender goroutine
		for range ch {
		}
	}()

	w, err := fw.Open(path, chmod)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	for p := range ch {
		if _, err := w.Write(p); err != nil {
			panic(err)
		}
	}
}
