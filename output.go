package cmd

import "io"

type LinerWriter interface {
	io.Writer
	Lines() []string
}
