package golang

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path"
	"syscall"

	executor "github.com/nlepage/codyglot/executor/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Executor is Codyglot Go(lang) executor
type Executor struct {
	Port int
}

// Serve starts listening for gRPC requests
func (e *Executor) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.Port))
	if err != nil {
		return errors.Wrap(err, "Executor: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	executor.RegisterExecutorServer(grpcSrv, e)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Executor: Failed to serve")
	}

	return nil
}

// Execute executes Go(lang) code
func (e *Executor) Execute(ctx context.Context, req *executor.ExecuteRequest) (*executor.ExecuteResponse, error) {
	p, err := writeSourceFile(req.Code)
	if err != nil {
		return nil, err
	}

	cmd, stdin, stdout, stderr, err := command(ctx, "go", "run", p)
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()              // TODO Manage error ?
		io.WriteString(stdin, req.StdIn) // TODO Manage error ?
	}()

	var stdoutBuf, stderrBuf bytes.Buffer

	go func() {
		io.Copy(&stdoutBuf, stdout) // TODO Manage error ?
	}()

	go func() {
		io.Copy(&stderrBuf, stderr) // TODO Manage error ?
	}()

	err = cmd.Start()
	if err != nil {
		return nil, errors.Wrap(err, "Executor: Could not start command")
	}

	var exitCode syscall.WaitStatus
	err = cmd.Wait()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return nil, errors.Wrap(err, "Executor: Something went wrong during command execution")
		}
		exitCode = 255
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			exitCode = status
		}
	}

	return &executor.ExecuteResponse{
		ExitCode: uint32(exitCode),
		StdOut:   string(stdoutBuf.Bytes()),
		StdErr:   string(stderrBuf.Bytes()),
	}, nil
}

var _ executor.ExecutorServer = (*Executor)(nil)

func writeSourceFile(code string) (string, error) {
	dir, err := ioutil.TempDir("", "codyglot")
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to create temp dir")
	}

	p := path.Join(dir, "main.go")

	f, err := os.Create(p)
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to create source file")
	}
	defer f.Close()

	_, err = f.WriteString(code)
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to write source file")
	}

	return p, nil
}

func command(ctx context.Context, name string, args ...string) (*exec.Cmd, io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	cmd := exec.Command(name, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Executor: Failed to retrieve command stdin")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Executor: Failed to retrieve command stdout")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Executor: Failed to retrieve command stderr")
	}

	return cmd, stdin, stdout, stderr, nil
}
