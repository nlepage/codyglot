package filestore

import (
	"io"
	"os"
	"path/filepath"

	service "github.com/nlepage/codyglot/service/filestore"
)

type fileMessageSender interface {
	Send(*service.FileMessage) error
}

func sendFile(sender fileMessageSender, path string, relPath string, info os.FileInfo, config Config) error {
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
	b := make([]byte, config.ChunkSize)

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

func sendDir(sender fileMessageSender, dir string, config Config, includeDirName bool) error {
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
			if err := sendFile(sender, path, relPath, info, config); err != nil {
				return err
			}
		}

		return nil
	})
}
