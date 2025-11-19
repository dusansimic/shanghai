package shanghai

import "github.com/dusansimic/shanghai/image"

func getImages(s *session, i string) []image.Image {
	var ims []image.Image
	if s.this {
		ims = []image.Image{s.f.Tree.Get(i)}
	} else {
		ims = s.f.Tree.Topological(i)
	}
	return ims
}
