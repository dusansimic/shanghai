package shanghai

type Session interface {
	Build(i string) error
	Push(i string) error
}

type session struct {
	c    *Config
	f    *Shanghaifile
	l    LogWriters
	this bool
}

func NewSession(c *Config, f *Shanghaifile, this bool, l LogWriters) Session {
	return &session{
		c:    c,
		f:    f,
		l:    l,
		this: this,
	}
}

func (s *session) Build(i string) error {
	return BuildImages(s.c, s.f, s.this, s.l, i)
}

func (s *session) Push(i string) error {
	return PushImages(s.c, s.f, s.l, i)
}
