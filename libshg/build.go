package libshg

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/juxR/usg"
	"github.com/wzshiming/ctc"
)

// BuildImages builds image subtree
func BuildImages(c *Config, f *Shanghaifile, i string) error {
	st, stExists := findSubtree(f, i)

	if err := buildImage(f.Images[i], c.Engine); err != nil {
		return fmt.Errorf("failed to build image '%s': %w", i, err)
	}

	if stExists {
		if err := walkTreeAction(st, f.Images, c.Engine, buildImage); err != nil {
			return fmt.Errorf("failed to build images: %w", err)
		}
	}

	return nil
}

func buildImage(i Image, e string) error {
	cmd := exec.Command(e, "build", "-t", i.Tag, "-f", i.ContainerFile, i.Context)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Printf("Building '%s' ", i.Tag)
	if err := cmd.Run(); err != nil {
		fmt.Println(ctc.ForegroundRed, usg.Get.Cross, ctc.Reset)
		return fmt.Errorf("failed to build image '%s': %w", i.Tag, err)
	}
	fmt.Println(ctc.ForegroundGreen, usg.Get.Tick, ctc.Reset)

	return nil
}
