package shanghai

import "io"

type LogWriters struct {
	Err io.Writer
	Out io.Writer
}
