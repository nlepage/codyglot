package srcutil

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

// WriteSourceFile writes a source file to disk and returns its absolute path
func WriteSourceFile(name string, source string) (string, error) {
	dir, err := ioutil.TempDir("", "codyglot")
	if err != nil {
		return "", errors.Wrap(err, "WriteSourceFile: Failed to create temp dir")
	}

	p := path.Join(dir, name)

	f, err := os.Create(p)
	if err != nil {
		return "", errors.Wrap(err, "WriteSourceFile: Failed to create file")
	}
	defer f.Close()

	_, err = f.WriteString(source)
	if err != nil {
		return "", errors.Wrap(err, "WriteSourceFile: Failed to write file")
	}

	return p, nil
}
