package image

type Image interface {
	Name() string
	Tags() []string
	ContainerfileName() string
	Context() string
	BuildArgs() map[string]string
	Parents() []string
}

type image struct {
	name          string
	tags          []string
	containerFile string
	context       string
	buildArgs     map[string]string
	parents       []string
}

func NewImage(n string, ts []string, cf, c string, ba map[string]string, ps []string) Image {
	return &image{
		name:          n,
		tags:          ts,
		containerFile: cf,
		context:       c,
		buildArgs:     ba,
		parents:       ps,
	}
}

func (i *image) Name() string {
	return i.name
}

func (i *image) Tags() []string {
	return i.tags
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
