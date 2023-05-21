package libshg

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/juxR/usg"
	"github.com/wzshiming/ctc"
)

func PushImages(c *Config, f *Shanghaifile, i string) error {
	st, stExists := findSubtree(f, i)

	if err := pushImage(f.Images[i], c.Engine); err != nil {
		return fmt.Errorf("failed to push image '%s': %w", i, err)
	}

	if stExists {
		if err := walkTreeAction(st, f.Images, c.Engine, pushImage); err != nil {
			return fmt.Errorf("failed to push images: %w", err)
		}
	}

	return nil
}

func pushImage(i Image, e string) error {
	cmd := exec.Command(e, "push", i.Tag)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Printf("Pushing '%s' ", i.Tag)
	if err := cmd.Run(); err != nil {
		fmt.Println(ctc.ForegroundRed, usg.Get.Cross, ctc.Reset)
		return fmt.Errorf("failed to push image '%s': %w", i.Tag, err)
	}
	fmt.Println(ctc.ForegroundGreen, usg.Get.Tick, ctc.Reset)

	return nil
}
