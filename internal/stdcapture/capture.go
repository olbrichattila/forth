// Package stdcapture is used in tests to capture the result of the script outputs to the standard console.
package stdcapture

import (
	"bytes"
	"os"
)

// New creates a new StdoutCapture instance.
func New() *StdoutCapture {
	return &StdoutCapture{}
}

// StdoutCapture captures output written to the standard output.
type StdoutCapture struct {
	originalStdout *os.File
	readPipe       *os.File
	writePipe      *os.File
	buf            bytes.Buffer
}

// StartCapture redirects the standard output to a pipe for capturing.
func (c *StdoutCapture) StartCapture() error {
	var err error
	c.originalStdout = os.Stdout
	c.readPipe, c.writePipe, err = os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = c.writePipe
	return nil
}

// StopCapture stops capturing and restores the original standard output.
func (c *StdoutCapture) StopCapture() (string, error) {
	c.writePipe.Close()
	os.Stdout = c.originalStdout

	_, err := c.buf.ReadFrom(c.readPipe)
	c.readPipe.Close()
	if err != nil {
		return "", err
	}

	return c.buf.String(), nil
}