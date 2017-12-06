package cmd

import "io"

type Outputer interface {
	io.Writer
	io.Closer
	Lines() []string
}
