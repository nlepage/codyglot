package ioutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/nlepage/codyglot/config"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	tempDirPrefix = "codyglot"
)

var (
	closed           chan string
	startCleanupOnce sync.Once
)

// TmpDir is a temporary directory
type TmpDir struct {
	path   string
	closed bool
}

func startCleanup() {
	closed = make(chan string, config.CleanupBuffer)

	for i := 0; i < config.CleanupRoutines; i++ {
		go func() {
			for path := range closed {
				if err := os.RemoveAll(path); err != nil {
					log.Errorln("Failed to remove temp dir", path)
				}
			}
		}()
	}
}

// NewTmpDir creates a new TmpDir
func NewTmpDir() (*TmpDir, error) {
	startCleanupOnce.Do(startCleanup)

	path, err := ioutil.TempDir("", tempDirPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "TmpDir: Failed to create temp dir")
	}

	return &TmpDir{
		path: path,
	}, nil
}

// Join joins sub paths to the path of TmpDir
func (td *TmpDir) Join(paths ...string) string {
	td.checkClosed("Join")

	fullPaths := make([]string, len(paths)+1)
	fullPaths[0] = td.path
	copy(fullPaths[1:], paths)
	return path.Join(fullPaths...)
}

// WriteFile writes a file in TmpDir
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
		panic(fmt.Sprintf("TmpDir: Invalid state, %s should not be called after closing temp dir", name))
	}
}

// Close marks the TmpDir closed and sends it to cleanup routines
func (td *TmpDir) Close() {
	td.checkClosed("Close")

	td.closed = true
	closed <- td.path
}

// Path returns the path of the TmpDir
func (td *TmpDir) Path() string {
	return td.path
}

// WriteSources writes source files to the TmpDir
func (td *TmpDir) WriteSources(sources []*service.SourceFile) error {
	for _, srcFile := range sources {
		if err := os.MkdirAll(filepath.Dir(td.Join(srcFile.Path)), 0755); err != nil {
			return errors.Wrap(err, "execute: Failed to create directory")
		}

		if _, err := td.WriteFile(srcFile.Path, srcFile.Content); err != nil {
			// FIXME format error with file information
			return errors.Wrap(err, "execute: Failed to write source file")
		}
	}

	return nil
}