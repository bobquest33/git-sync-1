package syncer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Giter is strcut
type Giter struct {
	writer    io.Writer
	errWriter io.Writer
}

// NewGiter is constructor
func NewGiter(writer, errWriter io.Writer) *Giter {
	return &Giter{
		writer:    writer,
		errWriter: errWriter,
	}
}

// IsExist is check git command
func (*Giter) IsExist() error {
	s, err := exec.LookPath("git")
	fmt.Fprintf(os.Stdout, "%s", s)
	return err
}

// Git is execute git command
func (g *Giter) Git(command string, args ...string) (msg, errMsg string, err error) {
	wr := new(bytes.Buffer)
	errWr := new(bytes.Buffer)

	cmdArgs := append([]string{command}, args...)
	c := exec.Command("git", cmdArgs...)
	c.Stdin = os.Stdin
	c.Stdout = wr
	c.Stderr = errWr

	err = c.Run()
	msg = wr.String()
	errMsg = errWr.String()
	return
}
