package shanghai

import (
	"fmt"
	"os/exec"
)

// BuildImages builds image subtree
func BuildImages(c *Config, f *Shanghaifile, lw LogWriters, i string) error {
	is := f.Tree.PreorderFrom(i)

	for _, im := range is {
		if err := buildImage(lw, f, im, c.Engine); err != nil {
			return fmt.Errorf("failed to build image '%s': %w", i, err)
		}
	}

	return nil
}

func buildImage(lw LogWriters, f *Shanghaifile, im Image, e string) error {
	buildArgs := []string{}
	for k, v := range f.BuildArguments {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	for k, v := range im.BuildArgs() {
		buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	envVars := []string{}
	for k, v := range f.EnvironmentVariables {
		envVars = append(envVars, fmt.Sprintf("%s=%s", k, v))
	}

	cmdArgs := []string{"build"}
	cmdArgs = append(cmdArgs, buildArgs...)
	cmdArgs = append(cmdArgs, envVars...)
	cmdArgs = append(cmdArgs, "-t", im.Tag(), "-f", im.ContainerfileName(), im.Context())

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
