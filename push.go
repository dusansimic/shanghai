package shanghai

import (
	"fmt"
	"os/exec"
	"strings"
)

func PushImages(c *Config, f *Shanghaifile, this bool, lw LogWriters, i string) error {
	if this {
		im := f.Tree.Get(i)

		if !strings.HasPrefix(im.Tag(), "localhost/") {
			return nil
		}

		if err := pushImage(lw, f, im, c.Engine); err != nil {
			return fmt.Errorf("failed to push image '%s': %w", i, err)
		}
	} else {
		is := f.Tree.Topological(i)

		for _, im := range is {
			if strings.HasPrefix(im.Tag(), "localhost/") {
				continue
			}

			if err := pushImage(lw, f, im, c.Engine); err != nil {
				return fmt.Errorf("failed to push image '%s': %w", i, err)
			}
		}
	}

	return nil
}

func pushImage(lw LogWriters, f *Shanghaifile, im Image, e string) error {
	cmd := exec.Command(e, "push", im.Tag())

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Pushing %s\n", im.Name())))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run push command for '%s': %s\n", im.Name(), err.Error())))
		return fmt.Errorf("failed to run push command for '%s': %w", im.Name(), err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Push done for %s\n", im.Name())))

	return nil
}
