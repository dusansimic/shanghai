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
				continue
			}

			if err := pushImage(lw, im, c.Engine); err != nil {
				return fmt.Errorf("failed to push image '%s': %w", i, err)
			}
		}
	}

	return nil
}

func pushImage(lw LogWriters, im image.Image, e string) error {
	for _, tag := range im.Tags() {
		cmd := exec.Command(e, "push", tag)

		cmd.Stderr = lw.Err
		cmd.Stdout = lw.Out

		lw.Out.Write([]byte(fmt.Sprintf("Pushing %s\n", im.Name())))
		if err := cmd.Run(); err != nil {
			lw.Err.Write([]byte(fmt.Sprintf("failed to run push command for '%s': %s\n", im.Name(), err.Error())))
			return fmt.Errorf("failed to run push command for '%s': %w", im.Name(), err)
		}
		lw.Out.Write([]byte(fmt.Sprintf("Push done for %s\n", im.Name())))
	}

	return nil
}
