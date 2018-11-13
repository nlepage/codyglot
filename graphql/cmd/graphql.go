package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/handler"
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const defaultPort = 8080
const defaultPlayground = false
const defaultEndpoint = "/query"

var (
	endpoint   string
	playground bool
	port       int
	graphqlCmd = &cobra.Command{
		Short: "Starts a graphql server",
		Use:   "graphql",
		RunE: func(_ *cobra.Command, _ []string) error {
			if !strings.HasPrefix(endpoint, "/") {
				endpoint = "/" + endpoint
			}

			http.Handle(endpoint, handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}})))
			log.Infof("Listening on http://localhost:%d%s", port, endpoint)

			if playground {
				http.Handle("/", handler.Playground("GraphQL playground", endpoint))
				log.Infof("Connect to http://localhost:%d/ for GraphQL playground", port)
			}

			return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		},
	}
)

func init() {
	graphqlCmd.Flags().StringVarP(&endpoint, "endpoint", "e", defaultEndpoint, "Endpoint")
	graphqlCmd.Flags().BoolVar(&playground, "playground", defaultPlayground, "Enable playground")
	graphqlCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Listening port")
	cmd.Cmd.AddCommand(graphqlCmd)
}
