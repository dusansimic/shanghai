package shanghai

type Image interface {
	Name() string
	Tag() string
	ContainerfileName() string
	Context() string
	BuildArgs() map[string]string
	Parents() []string
}

type image struct {
	name          string
	tag           string
	containerFile string
	context       string
	buildArgs     map[string]string
	parents       []string
}

func NewImage(n, t, cf, c string, ba map[string]string, ps []string) Image {
	return &image{
		name:          n,
		tag:           t,
		containerFile: cf,
		context:       c,
		buildArgs:     ba,
		parents:       ps,
	}
}

func (i *image) Name() string {
	return i.name
}

func (i *image) Tag() string {
	return i.tag
}

func (i *image) ContainerfileName() string {
	return i.containerFile
}

func (i *image) Context() string {
	return i.context
}

func (i *image) BuildArgs() map[string]string {
	return i.buildArgs
}

func (i *image) Parents() []string {
	return i.parents
}

func (i *image) String() string {
	return i.name
}
