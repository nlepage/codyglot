package filestore

import (
	"io"
	"os"
	"path/filepath"
)

type Writer interface {
	Open(path string, chmod os.FileMode) (io.WriteCloser, error)
}

func FsWriter(dir string) Writer {
	return fsWriter(dir)
}

type fsWriter string

var _ Writer = fsWriter("")

func (rootDir fsWriter) Open(relPath string, chmod os.FileMode) (io.WriteCloser, error) {
	path := filepath.Join(string(rootDir), relPath)
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, dirMode); err != nil {
		return nil, err
	}

	return os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, chmod)
}

type Reader interface {
	Copy(Writer) error
}

func FsReader(path string, config Config, includeDirName bool) Reader {
	return fsReader{path, config, includeDirName}
}

type fsReader struct {
	path           string
	config         Config
	includeDirName bool
}

var _ Reader = fsReader{}

func (fr fsReader) Copy(w Writer) error {
	// FIXME wrap errors
	info, err := os.Stat(fr.path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fr.copyDir(w)
	}

	return fr.copy(w, fr.path, info.Name(), info.Mode())
}

func (fr fsReader) copyDir(w Writer) error {
	dir, err := filepath.Abs(fr.path)
	if err != nil {
		return err
	}

	dirName := filepath.Base(dir)

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if fr.includeDirName {
			relPath = filepath.Join(dirName, relPath)
		}

		if !info.IsDir() {
			if err := fr.copy(w, path, relPath, info.Mode()); err != nil {
				return err
			}
		}

		return nil
	})
}

func (fr fsReader) copy(w Writer, path string, relPath string, chmod os.FileMode) error {
	// FIXME wrap errors

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	iow, err := w.Open(relPath, chmod)
	if err != nil {
		return err
	}
	defer iow.Close()

	// FIXME reuse buffer
	if _, err := io.CopyBuffer(iow, f, make([]byte, fr.config.ChunkSize)); err != nil {
		return err
	}

	return nil
}
