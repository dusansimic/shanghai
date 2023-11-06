package shanghai

type Session interface {
	Build(i string) error
	Push(i string) error
	ValidateShanghaifile(i string) error
}

type session struct {
	c *Config
	f *Shanghaifile
	l LogWriters
}

func NewSession(c *Config, f *Shanghaifile, l LogWriters) Session {
	return &session{
		c: c,
		f: f,
		l: l,
	}
}

func (s *session) Build(i string) error {
	return BuildImages(s.c, s.f, s.l, i)
}

func (s *session) Push(i string) error {
	return PushImages(s.c, s.f, s.l, i)
}

func (s *session) ValidateShanghaifile(i string) error {
	return ValidateShanghaifile(s.f, i)
}
