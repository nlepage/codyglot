package config

import (
	"fmt"
	"strings"
)

const (
	// DefaultPort is the default router listening port
	DefaultPort = 9090
)

var (
	// Port is the router listening port
	Port int
	// Languages is the raw list of executor endpoints
	Languages []string
	// LanguagesMap is the map of executor endpoints by language
	LanguagesMap map[string]string
)

// InitLanguagesMap initializes LanguagesMap
func InitLanguagesMap() error {
	LanguagesMap = make(map[string]string, len(Languages))

	for _, language := range Languages {
		languageSplit := strings.SplitN(language, ":", 2)

		if len(languageSplit) != 2 {
			return fmt.Errorf("Invalid language: %s", language)
		}

		LanguagesMap[languageSplit[0]] = languageSplit[1]
	}

	return nil
}
