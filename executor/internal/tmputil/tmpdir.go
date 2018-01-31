package tmputil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	tempDirPrefix         = "codyglot"
	defaultClosedBuffer   = 10
	defaultCloserRoutines = 2
)

var (
	closed = make(chan string, defaultClosedBuffer)
)

type TmpDir struct {
	path   string
	closed bool
}

func init() {
	for i := 0; i < defaultCloserRoutines; i++ {
		go func() {
			for path := range closed {
				if err := os.RemoveAll(path); err != nil {
					log.Errorln("Failed to remove temp dir", path)
				}
			}
		}()
	}
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
	td.checkClosed("Join")

	fullPaths := make([]string, len(paths)+1)
	fullPaths[0] = td.path
	copy(fullPaths[1:], paths)
	return path.Join(fullPaths...)
}

func (td *TmpDir) WriteFile(name string, content string) (string, error) {
	td.checkClosed("WriteFile")

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

func (td *TmpDir) checkClosed(name string) {
	if td.closed {
		panic(fmt.Sprintf("TmpDir: Invalid state, %s should not be called after closing temp dir"))
	}
}

func (td *TmpDir) Close() {
	td.checkClosed("Close")

	td.closed = true
	closed <- td.path
}
