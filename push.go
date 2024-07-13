package shanghai

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/dusansimic/shanghai/file"
	"github.com/dusansimic/shanghai/image"
)

func PushImages(c *Config, f *file.File, this bool, lw LogWriters, i string) error {
	var ims []image.Image
	if this {
		ims = []image.Image{f.Tree.Get(i)}
	} else {
		ims = f.Tree.Topological(i)
	}

	for _, im := range ims {
		for _, tag := range im.Tags() {
			if strings.HasPrefix(tag, "localhost/") {
				lw.Out.Write([]byte(fmt.Sprintf("Skipping tag '%s'\n", tag)))
				continue
			}

			if err := pushImage(lw, tag, c.Engine); err != nil {
				return fmt.Errorf("failed to push tag '%s': %w", tag, err)
			}
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
