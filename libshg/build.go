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
		if err := walkTreeToBuildNodes(st, f.Images, c.Engine); err != nil {
			return fmt.Errorf("failed to build images: %w", err)
		}
	}

	return nil
}

func findSubtree(f *Shanghaifile, i string) (Node, bool) {
	return walkTreeToFindSubtree(f.Tree, i)
}

// walkTreeToFindSubtree walks the tree to find a node
func walkTreeToFindSubtree(t Node, i string) (Node, bool) {
	if _, ok := t[i]; ok {
		if t[i] == nil {
			return nil, false
		}
		return t[i].(Node), true
	}

	for st := range t {
		if _, ok := t[st].(Node); ok {
			return walkTreeToFindSubtree(t[st].(Node), i)
		}
	}

	return nil, false
}

func walkTreeToBuildNodes(t Node, is MapOfImages, e string) error {
	for k := range t {
		if err := buildImage(is[k], e); err != nil {
			return fmt.Errorf("failed to build image '%s': %w", k, err)
		}

		// Check if this node is also a subtree
		if _, ok := t[k].(Node); ok {
			walkTreeToBuildNodes(t[k].(Node), is, e)
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
