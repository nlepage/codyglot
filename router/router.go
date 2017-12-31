package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	router "github.com/Zenika/codyglot/router/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

const (
	defaultPort     = 9090
	defaultHTTPPort = 8080
	defaultEndpoint = "localhost:9090"
)

var (
	port       int
	httpPort   int
	_languages []string
	languages  map[string]string
	endpoint   string
)

var routerCmd = &cobra.Command{
	Use: "router",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start gRPC server",
	RunE:  serve,
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start REST gateway",
	RunE:  serveHTTP,
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Listening port")
	serverCmd.Flags().StringSliceVarP(&_languages, "language", "l", nil, "Language")
	routerCmd.AddCommand(serverCmd)

	gatewayCmd.Flags().IntVarP(&httpPort, "port", "p", defaultHTTPPort, "Listening port")
	gatewayCmd.Flags().StringVarP(&endpoint, "endpoint", "e", defaultEndpoint, "gRPC endpoint")
	routerCmd.AddCommand(gatewayCmd)
}

func main() {
	if err := routerCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func serve(cmd *cobra.Command, args []string) error {
	if err := initLanguages(); err != nil {
		return err
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

func serveHTTP(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := router.RegisterRouterHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

func initLanguages() error {
	languages = make(map[string]string, len(_languages))
	for _, language := range _languages {
		languageSplit := strings.SplitN(language, ":", 2)

		if len(languageSplit) != 2 {
			return fmt.Errorf("Invalid language: %s", language)
		}

		languages[languageSplit[0]] = languageSplit[1]
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
