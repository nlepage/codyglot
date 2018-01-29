package tmputil

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

const (
	tempDirPrefix = "codyglot"
)

type TmpDir struct {
	path   string
	closed bool
}

func NewTmpDir() (*TmpDir, error) {
	path, err := ioutil.TempDir("", tempDirPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "TmpDir: Failed to create temp dir")
	}

	return &TmpDir{
		path: path,
	}, nil
}

func (td *TmpDir) Join(paths ...string) string {
	fullPaths := make([]string, len(paths)+1)
	fullPaths[0] = td.path
	copy(fullPaths[1:], paths)
	return path.Join(fullPaths...)
}

func (td *TmpDir) WriteFile(name string, content string) (string, error) {
	p := td.Join(name)

	f, err := os.Create(p)
	if err != nil {
		return "", errors.Wrap(err, "TmpDir: Failed to create file")
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return "", errors.Wrap(err, "TmpDir: Failed to write file")
	}

	return p, nil
}

func (td *TmpDir) Close() {
	td.closed = true
}
