package exec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

// Cmd is a wrapper for exec.Cmd
type Cmd struct {
	cmd       *exec.Cmd
	wg        sync.WaitGroup
	stdin     string
	stdout    *bytes.Buffer
	stderr    *bytes.Buffer
	errs      []error
	exitErr   *exec.ExitError
	startTime time.Time
	duration  time.Duration
}

// Command creates a new Cmd
func Command(ctx context.Context, name string, args ...string) *Cmd {
	cmd := exec.CommandContext(ctx, name, args...)

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout, cmd.Stderr = stdout, stderr

	return &Cmd{
		cmd:    cmd,
		stdout: stdout,
		stderr: stderr,
	}
}

// WithStdin sets the text to pipe on command stdin
func (cmd *Cmd) WithStdin(stdin string) *Cmd {
	cmd.stdin = stdin
	return cmd
}

// WithDir sets the working directory of the command
func (cmd *Cmd) WithDir(dir string) *Cmd {
	cmd.cmd.Dir = dir
	return cmd
}

// Run runs the command
func (cmd *Cmd) Run() error {
	cmd.writeStdin()

	cmd.start()
	cmd.wait()

	return cmd.error()
}

// Status returns exit status
func (cmd *Cmd) Status() uint32 {
	cmd.checkErrors("Status")

	if cmd.exitErr == nil {
		return 0
	}

	var exitStatus uint32 = 255
	if waitStatus, ok := cmd.exitErr.Sys().(syscall.WaitStatus); ok {
		exitStatus = uint32(waitStatus)
	}

	return exitStatus
}

// Stdout returns standard output
func (cmd *Cmd) Stdout() string {
	cmd.checkErrors("Stdout")

	return string(cmd.stdout.Bytes())
}

// Stderr returns standard error output
func (cmd *Cmd) Stderr() string {
	cmd.checkErrors("Stderr")

	return string(cmd.stderr.Bytes())
}

// Duration returns duration of execution
func (cmd *Cmd) Duration() int64 {
	cmd.checkErrors("Duration")

	return int64(cmd.duration)
}

// CommandResult returns result of execution
func (cmd *Cmd) CommandResult() *service.CommandResult {
	return &service.CommandResult{
		Status:   cmd.Status(),
		Stdout:   cmd.Stdout(),
		Stderr:   cmd.Stderr(),
		Duration: cmd.Duration(),
	}
}

func (cmd *Cmd) writeStdin() {
	if len(cmd.errs) != 0 || len(cmd.stdin) == 0 {
		return
	}

	stdinPipe, err := cmd.cmd.StdinPipe()
	if err != nil {
		cmd.addError(err, "Failed to retrieve stdin pipe")
		return
	}

	cmd.wg.Add(1)
	go func() {
		defer cmd.wg.Done()
		defer func() {
			if err := stdinPipe.Close(); err != nil {
				cmd.addError(err, "Failed to close stdin pipe")
			}
		}()

		if _, err := io.WriteString(stdinPipe, cmd.stdin); err != nil {
			cmd.addError(err, "Failed to write stdin")
		}
	}()
}

func (cmd *Cmd) captureStdxxx(stdxxxPipeProvider func() (io.ReadCloser, error)) *bytes.Buffer {
	if cmd.hasErrors() {
		return nil
	}

	stdxxxPipe, err := stdxxxPipeProvider()
	if err != nil {
		cmd.addError(err, "Failed to retrieve std(out/err) pipe")
		return nil
	}

	var buf bytes.Buffer

	cmd.wg.Add(1)
	go func() {
		defer cmd.wg.Done()

		if _, err := io.Copy(&buf, stdxxxPipe); err != nil {
			cmd.addError(err, "Failed to read std(out/err)")
		}
	}()

	return &buf
}

func (cmd *Cmd) start() {
	if cmd.hasErrors() {
		return
	}

	cmd.startTime = time.Now()

	if err := cmd.cmd.Start(); err != nil {
		cmd.addError(err, "Failed to start command")
	}
}

func (cmd *Cmd) wait() {
	if cmd.hasErrors() {
		return
	}

	defer func() {
		cmd.duration = time.Since(cmd.startTime)
	}()

	if err := cmd.cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			cmd.exitErr = exitErr
		} else {
			cmd.addError(err, "Failed to run command")
		}
	}

	cmd.wg.Wait()
}

func (cmd *Cmd) addError(err error, message string) {
	cmd.errs = append(cmd.errs, errors.Wrap(err, "Command: "+message))
}

func (cmd *Cmd) hasErrors() bool {
	return len(cmd.errs) != 0
}

func (cmd *Cmd) error() error {
	if !cmd.hasErrors() {
		return nil
	}

	if len(cmd.errs) == 1 {
		return cmd.errs[0]
	}

	message := "Command: Multiple errors:"
	for _, err := range cmd.errs {
		message += "\n" + err.Error()
	}
	return errors.New(message)
}

func (cmd *Cmd) checkErrors(name string) {
	if cmd.hasErrors() {
		panic(fmt.Sprintf("Command: Invalid state, %s should not be called when there are errors", name))
	}
}
