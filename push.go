package shanghai

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/dusansimic/shanghai/image"
)

func PushImages(s *session, i string) error {
	var ims []image.Image
	if s.this {
		ims = []image.Image{s.f.Tree.Get(i)}
	} else {
		ims = s.f.Tree.Topological(i)
	}

	for _, im := range ims {
		for _, tag := range im.Tags() {
			if strings.HasPrefix(tag, "localhost/") {
				s.l.Out.Write([]byte(fmt.Sprintf("Skipping tag '%s'\n", tag)))
				continue
			}

			if err := pushImage(s.l, tag, s.c.Engine); err != nil {
				return fmt.Errorf("failed to push tag '%s': %w", tag, err)
			}
		}
	}

	return nil
}

func PushGroup(s *session, g string) error {
	names := s.f.Groups[g]
	for _, name := range names {
		if err := PushImages(s, name); err != nil {
			return fmt.Errorf("failed to push image from group '%s': %w", g, err)
		}
	}

	return nil
}

func pushImage(lw LogWriters, t string, e string) error {
	cmd := exec.Command(e, "push", t)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Pushing %s\n", t)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run push command for '%s': %s\n", t, err.Error())))
		return fmt.Errorf("failed to run push command for '%s': %w", t, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Push done for %s\n", t)))

	return nil
}
