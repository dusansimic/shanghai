package shanghai

import "github.com/dusansimic/shanghai/file"

type Session interface {
	Build(i string) error
	Push(i string) error
}

type session struct {
	c     *Config
	f     *file.File
	l     LogWriters
	this  bool
	group bool
}

func NewSession(c *Config, f *file.File, this, group bool, l LogWriters) Session {
	return &session{
		c:     c,
		f:     f,
		l:     l,
		this:  this,
		group: group,
	}
}

func (s *session) Build(n string) error {
	if s.group {
		return BuildGroup(s, n)
	} else {
		return BuildImages(s, n)
	}
}

func (s *session) Push(n string) error {
	if s.group {
		return PushGroup(s, n)
	} else {
		return PushImages(s, n)
	}
}
