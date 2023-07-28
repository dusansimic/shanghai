package shanghai

import "fmt"

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

func walkTreeAction(lw LogWriters, t Node, is MapOfImages, e string, a func(LogWriters, Image, string) error) error {
	for k := range t {
		if err := a(lw, is[k], e); err != nil {
			return fmt.Errorf("failed to complete action on image '%s': %w", k, err)
		}

		// Check if this node is also a subtree
		if _, ok := t[k].(Node); ok {
			walkTreeAction(lw, t[k].(Node), is, e, a)
		}
	}

	return nil
}
