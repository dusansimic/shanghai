package shanghai

import (
	"fmt"
	"os/exec"

	"github.com/dusansimic/shanghai/file"
	"github.com/dusansimic/shanghai/image"
)

// BuildImages builds image subtree
func BuildImages(c *Config, f *file.File, this bool, lw LogWriters, i string) error {
	var ims []image.Image
	if this {
		ims = []image.Image{f.Tree.Get(i)}
	} else {
		ims = f.Tree.Topological(i)
	}

	for _, im := range ims {
		tags := im.Tags()
		if err := buildImage(lw, f, im, tags[0], c.Engine); err != nil {
			return fmt.Errorf("failed to build image '%s': %w", i, err)
		}

		for _, tag := range tags[1:] {
			if err := tagImage(lw, tags[0], tag, c.Engine); err != nil {
				return fmt.Errorf("failed to tag image: '%s': %w", tag, err)
			}
		}
	}

	return nil
}

func buildImage(lw LogWriters, f *file.File, im image.Image, t string, e string) error {
	buildArgs := []string{}
	for k, v := range f.BuildArgs {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	for k, v := range im.BuildArgs() {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	envVars := []string{}
	for k, v := range f.EnvVars {
		envVars = append(envVars, fmt.Sprintf("%s=%s", k, v))
	}

	cmdArgs := []string{"build"}
	cmdArgs = append(cmdArgs, buildArgs...)
	cmdArgs = append(cmdArgs, envVars...)
	cmdArgs = append(cmdArgs, "-t", t, "-f", im.ContainerfileName(), im.Context())

	cmd := exec.Command(e, cmdArgs...)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Building %s\n", im.Name())))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run build command for '%s': %s\n", im.Name(), err.Error())))
		return fmt.Errorf("failed to run command build command for '%s': %w", im.Name(), err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Build done for %s\n", im.Name())))

	return nil
}

func tagImage(lw LogWriters, t string, nt string, e string) error {
	cmd := exec.Command(e, "tag", t, nt)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Tagging %s\n", nt)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to add tag '%s' for image '%s': %s\n", nt, t, err.Error())))
		return fmt.Errorf("failed to add tag '%s' for image '%s': %w", nt, t, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Tagging done for %s\n", nt)))

	return nil
}
