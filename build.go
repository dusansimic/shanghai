package shanghai

import (
	"fmt"
	"os/exec"
)

// BuildImages builds image subtree
func BuildImages(c *Config, f *Shanghaifile, lw LogWriters, i string) error {
	st, stExists := findSubtree(f, i)

	if err := buildImage(lw, f, i, c.Engine); err != nil {
		return fmt.Errorf("failed to build image '%s': %w", i, err)
	}

	if stExists {
		if err := walkTreeAction(lw, f, st, f.Images, c.Engine, buildImage); err != nil {
			return fmt.Errorf("failed to build images: %w", err)
		}
	}

	return nil
}

func buildImage(lw LogWriters, f *Shanghaifile, i string, e string) error {
	im := f.Images[i]

	buildArgs := []string{}
	for k, v := range f.BuildArguments {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	for k, v := range im.BuildArgs {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	envVars := []string{}
	for k, v := range f.EnvironmentVariables {
		envVars = append(envVars, fmt.Sprintf("%s=%s", k, v))
	}

	cmdArgs := []string{"build"}
	cmdArgs = append(cmdArgs, buildArgs...)
	cmdArgs = append(cmdArgs, envVars...)
	cmdArgs = append(cmdArgs, "-t", im.Tag, "-f", im.ContainerFile, im.Context)

	cmd := exec.Command(e, cmdArgs...)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Building %s\n", im.Tag)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run build command for '%s': %s\n", im.Tag, err.Error())))
		return fmt.Errorf("failed to run command build command for '%s': %w", im.Tag, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Build done for %s\n", im.Tag)))

	return nil
}
