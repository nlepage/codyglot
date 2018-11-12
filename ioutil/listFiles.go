package ioutil

import (
	"os"
	"path/filepath"
)

func ListFiles(dir string) ([]string, error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	for i, file := range files {
		files[i] = filepath.Join(dir, file)
	}

	return files, nil
}
