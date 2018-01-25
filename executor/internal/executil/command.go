package executil

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

// Cmd is a wrapper for exec.Cmd
type Cmd struct {
	cmd       *exec.Cmd
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
	return &Cmd{
		cmd: exec.CommandContext(ctx, name, args...),
	}
}

// WithStdin sets the text to pipe on command stdin
func (cmd *Cmd) WithStdin(stdin string) *Cmd {
	cmd.stdin = stdin
	return cmd
}

// Run runs the command
func (cmd *Cmd) Run() error {
	cmd.writeStdin()
	cmd.stdout = cmd.captureStdxxx(cmd.cmd.StdoutPipe)
	cmd.stderr = cmd.captureStdxxx(cmd.cmd.StderrPipe)

	cmd.start()
	cmd.wait()

	return cmd.error()
}

// ExitStatus returns exit status
func (cmd *Cmd) ExitStatus() uint32 {
	cmd.checkErrors("ExitStatus")

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

func (cmd *Cmd) Duration() time.Duration {
	cmd.checkErrors("Duration")

	return cmd.duration
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

	go func() {
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
	var buf bytes.Buffer

	if cmd.hasErrors() {
		return &buf
	}

	stdxxxPipe, err := stdxxxPipeProvider()
	if err != nil {
		cmd.addError(err, "Failed to retrieve std(out/err) pipe")
		return &buf
	}

	go func() {
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

	if err := cmd.cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			cmd.exitErr = exitErr
		} else {
			cmd.addError(err, "Failed to run command")
		}
	}

	defer func() {
		cmd.duration = time.Since(cmd.startTime)
	}()
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
