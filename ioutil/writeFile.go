package ioutil

import (
	"os"
	"path/filepath"

	svc "github.com/nlepage/codyglot/service"
)

// WriteFile writes a file in dir
func WriteFile(dir string, name string, content string) (string, error) {
	p := filepath.Join(dir, name)

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return "", err
	}

	f, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}

	return p, nil
}

// WriteSources writes source files in dir
func WriteSources(dir string, sources []*svc.SourceFile) error {
	for _, srcFile := range sources {
		if _, err := WriteFile(dir, srcFile.Path, srcFile.Content); err != nil {
			return err
		}
	}

	return nil
}
