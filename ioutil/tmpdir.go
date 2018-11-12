package ioutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/nlepage/codyglot/config"
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

// Path returns the path of the TmpDir
func (td *TmpDir) Path() string {
	return td.path
}

// Join joins sub paths to the path of TmpDir
func (td *TmpDir) Join(paths ...string) string {
	td.checkClosed("Join")

	fullPaths := make([]string, len(paths)+1)
	fullPaths[0] = td.path
	copy(fullPaths[1:], paths)
	return path.Join(fullPaths...)
}

// Close marks the TmpDir closed and sends it to cleanup routines
func (td *TmpDir) Close() {
	td.checkClosed("Close")

	td.closed = true
	closed <- td.path
}

func (td *TmpDir) checkClosed(name string) {
	if td.closed {
		panic(fmt.Sprintf("TmpDir: Invalid state, %s should not be called after closing temp dir", name))
	}
}
