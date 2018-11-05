package server

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/server/config"
	service "github.com/nlepage/codyglot/service/filestore"
	"google.golang.org/grpc"
)

// Server is the file store server
type Server struct{}

// Init initializes file store server
func Init() error {
	return os.MkdirAll(config.Config.Root, filestore.DirMode)
}

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Config.Port))
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
func (*Server) Put(srv service.FileStore_PutServer) error {
	// FIXME wrap errors

	id := uuid.New().String()

	dir := filepath.Join(config.Config.Root, id)

	if err := os.Mkdir(dir, filestore.DirMode); err != nil {
		return err
	}

	if err := filestore.Write(srv, dir); err != nil {
		return err
	}

	return srv.SendAndClose(&service.Id{Id: id})
}

// Get handles a get request
func (*Server) Get(id *service.Id, srv service.FileStore_GetServer) error {
	// FIXME wrap errors

	return filestore.SendDir(srv, filepath.Join(config.Config.Root, id.Id), config.Config.FileStoreConfig, false)
}

var _ service.FileStoreServer = &Server{}
