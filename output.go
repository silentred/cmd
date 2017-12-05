package cmd

import "io"

type LinerWriter interface {
	io.Writer
	io.Closer
	Lines() []string
}
