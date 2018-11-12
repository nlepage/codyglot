package filestore

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	service "github.com/nlepage/codyglot/service/filestore"
	"google.golang.org/grpc"
)

// Server is the file store server
type Server struct {
	Config ServerConfig
}

// Init initializes file store server
func (s *Server) Init() error {
	return os.MkdirAll(s.Config.Root, dirMode)
}

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	service.RegisterFileStoreServer(grpcSrv, s)
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

// TODO reverse channel for errors ?

// Put handles a put request
func (s *Server) Put(srv service.FileStore_PutServer) error {
	// FIXME wrap errors

	id := uuid.New().String()

	dir := filepath.Join(s.Config.Root, id)

	if err := os.Mkdir(dir, dirMode); err != nil {
		return err
	}

	if err := recv(srv, dir); err != nil {
		return err
	}

	return srv.SendAndClose(&service.Id{Id: id})
}

// Get handles a get request
func (s *Server) Get(id *service.Id, srv service.FileStore_GetServer) error {
	// FIXME wrap errors

	return sendDir(srv, filepath.Join(s.Config.Root, id.Id), s.Config.Config, false)
}

var _ service.FileStoreServer = &Server{}
