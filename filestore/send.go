package filestore

import (
	"io"
	"os"

	service "github.com/nlepage/codyglot/service/filestore"
)

type fileMessageSender interface {
	Send(*service.FileMessage) error
}

func send(sender fileMessageSender, fr Reader) error {
	return fr.Copy(filesSender{sender})
}

type filesSender struct {
	ms fileMessageSender
}

var _ Writer = filesSender{}

func (fs filesSender) Open(path string, chmod os.FileMode) (io.WriteCloser, error) {
	if err := fs.ms.Send(&service.FileMessage{
		FileMessage: &service.FileMessage_FileInfo{
			FileInfo: &service.FileInfo{
				Path:  path,
				Chmod: int32(chmod),
			},
		},
	}); err != nil {
		return nil, err
	}

	return fileSender{fs.ms}, nil
}

type fileSender struct {
	ms fileMessageSender
}

var _ io.WriteCloser = fileSender{}

func (fs fileSender) Write(p []byte) (n int, err error) {
	if err := fs.ms.Send(&service.FileMessage{
		FileMessage: &service.FileMessage_FileContent{
			FileContent: &service.FileContent{
				Content: p,
			},
		},
	}); err != nil {
		return 0, err
	}

	return len(p), nil
}

func (fs fileSender) Close() error {
	return nil
}
