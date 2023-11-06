package shanghai

import (
	"fmt"
	"os/exec"
)

func PushImages(c *Config, f *Shanghaifile, lw LogWriters, i string) error {
	st, stExists := findSubtree(f, i)

	if err := pushImage(lw, f, i, c.Engine); err != nil {
		return fmt.Errorf("failed to push image '%s': %w", i, err)
	}

	if stExists {
		if err := walkTreeAction(lw, f, st, f.Images, c.Engine, pushImage); err != nil {
			return fmt.Errorf("failed to push images: %w", err)
		}
	}

	return nil
}

func pushImage(lw LogWriters, f *Shanghaifile, i string, e string) error {
	im := f.Images[i]

	cmd := exec.Command(e, "push", im.Tag)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Pushing %s\n", im.Tag)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run push command for '%s': %s\n", im.Tag, err.Error())))
		return fmt.Errorf("failed to run push command for '%s': %w", im.Tag, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Push done for %s\n", im.Tag)))

	return nil
}
