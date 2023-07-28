package shanghai

import (
	"fmt"
	"os/exec"
)

// BuildImages builds image subtree
func BuildImages(c *Config, f *Shanghaifile, lw LogWriters, i string) error {
	st, stExists := findSubtree(f, i)

	if err := buildImage(lw, f.Images[i], c.Engine); err != nil {
		return fmt.Errorf("failed to build image '%s': %w", i, err)
	}

	if stExists {
		if err := walkTreeAction(lw, st, f.Images, c.Engine, buildImage); err != nil {
			return fmt.Errorf("failed to build images: %w", err)
		}
	}

	return nil
}

func buildImage(lw LogWriters, i Image, e string) error {
	buildArgs := []string{}
	for k, v := range i.BuildArgs {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}
	cmdArgs := []string{"build"}
	cmdArgs = append(cmdArgs, buildArgs...)
	cmdArgs = append(cmdArgs, "-t", i.Tag, "-f", i.ContainerFile, i.Context)

	cmd := exec.Command(e, cmdArgs...)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Building %s\n", i.Tag)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run build command for '%s': %s\n", i.Tag, err.Error())))
		return fmt.Errorf("failed to run command build command for '%s': %w", i.Tag, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Build done for %s\n", i.Tag)))

	return nil
}
