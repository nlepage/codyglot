package routerserver

import (
	"fmt"
	"strings"

	"github.com/Zenika/codyglot/cmd/router"
	"github.com/Zenika/codyglot/router/server"
	"github.com/spf13/cobra"
)

const (
	defaultPort = 9090
)

var (
	port      int
	languages []string
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "Start Codyglot router gRPC server",
	RunE: func(cmd *cobra.Command, args []string) error {
		languages, err := getLanguagesMap()
		if err != nil {
			return err
		}

		s := &server.Server{
			Port:         port,
			LanguagesMap: languages,
		}

		return s.Serve()
	},
}

func init() {
	cmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Listening port")
	cmd.Flags().StringSliceVarP(&languages, "language", "l", nil, "Language")
	router.Cmd.AddCommand(cmd)
}

func getLanguagesMap() (map[string]string, error) {
	languagesMap := make(map[string]string, len(languages))

	for _, language := range languages {
		languageSplit := strings.SplitN(language, ":", 2)

		if len(languageSplit) != 2 {
			return nil, fmt.Errorf("Invalid language: %s", language)
		}

		languagesMap[languageSplit[0]] = languageSplit[1]
	}

	return languagesMap, nil
}
