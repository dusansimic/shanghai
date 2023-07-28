package shanghai

import (
	"fmt"
	"reflect"
)

func walkTreeToFindDuplicates(ns Node, is *map[string]interface{}) error {
	for k, v := range ns {
		if _, ok := (*is)[k]; ok {
			return fmt.Errorf("duplicate image '%s' in tree", k)
		}
		(*is)[k] = nil

		if v == nil {
			continue
		}

		vt := reflect.TypeOf(v)
		if vt.Kind() == reflect.Map {
			if err := walkTreeToFindDuplicates(v.(Node), is); err != nil {
				return err
			}
		}
	}

	return nil
}

// ValidateShanghaifile validates a Shanghaifile
func ValidateShanghaifile(f *Shanghaifile, i string) error {
	if f.Tree == nil {
		return fmt.Errorf("tree is empty")
	}

	if f.Images == nil {
		return fmt.Errorf("images is empty")
	}

	// Check for duplicate is in tree
	tis := map[string]interface{}{}
	if err := walkTreeToFindDuplicates(f.Tree, &tis); err != nil {
		return fmt.Errorf("failed to walk tree: %w", err)
	}

	// Check for duplicate images in images
	iis := map[string]interface{}{}
	for i := range f.Images {
		if _, ok := iis[i]; ok {
			return fmt.Errorf("duplicate image '%s' in images list", i)
		}
		iis[i] = nil
	}

	// Check for existance of all tree images in images list
	for i := range tis {
		if _, ok := iis[i]; !ok {
			return fmt.Errorf("image '%s' in tree is not in images list", i)
		}
	}

	// Check for existance of image to build in images list
	if _, ok := iis[i]; !ok {
		return fmt.Errorf("image '%s' to build is not in images list", i)
	}

	return nil
}
