// Package stdcapture used in tests to capture the result of the script outputs to std console
package stdcapture

import (
	"bytes"
	"os"
)

// New Std capturer
func New() *StdoutCapture {
	return &StdoutCapture{}
}

// This captures what the code prints to the standard output
type StdoutCapture struct {
	originalStdout *os.File
	r              *os.File
	w              *os.File
	buf            bytes.Buffer
}

// StartCapture redirects standard output to a stream
func (c *StdoutCapture) StartCapture() {
	c.originalStdout = os.Stdout
	c.r, c.w, _ = os.Pipe()
	os.Stdout = c.w
}

// StopCapture stops capturing and closes the stream, re set original standard output
func (c *StdoutCapture) StopCapture() string {
	c.w.Close()
	os.Stdout = c.originalStdout
	c.buf.ReadFrom(c.r)
	c.r.Close()

	return c.buf.String()
}
