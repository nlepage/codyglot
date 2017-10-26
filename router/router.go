package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	router "github.com/Zenika/codyglot/router/service"
	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

const (
	defaultPort = 8080
)

var (
	port       int
	_languages []string
	languages  map[string]string
)

func init() {
	cmd.PersistentFlags().IntVarP(&port, "port", "p", defaultPort, "Listening port, default 8080")
	cmd.PersistentFlags().StringSliceVarP(&_languages, "language", "l", nil, "Language")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var cmd = &cobra.Command{
	Use:  "router",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	languages = make(map[string]string, len(_languages))
	for _, language := range _languages {
		languageSplit := strings.SplitN(language, ":", 2)

		if len(languageSplit) != 2 {
			return fmt.Errorf("Invalid language: %s", language)
		}

		languages[languageSplit[0]] = languageSplit[1]
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	router.RegisterRouterServer(grpcSrv, &server{})
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

type server struct{}

func (*server) Execute(ctx context.Context, req *router.ExecuteRequest) (*router.ExecuteResponse, error) {
	return nil, nil
}

func (*server) Languages(ctx context.Context, req *router.LanguagesRequest) (*router.LanguagesResponse, error) {
	res := router.LanguagesResponse{
		Languages: make([]string, 0, len(languages)),
	}

	for language := range languages {
		res.Languages = append(res.Languages, language)
	}

	return &res, nil
}

var _ router.RouterServer = (*server)(nil)
