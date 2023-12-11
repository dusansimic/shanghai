package shanghai

import (
	"fmt"
	"shanghai/stack"
	"slices"
	"strings"
)

type PolyTree interface {
	Add(Image) error
	Get(string) Image
	Topological(string) []Image
	Nodes() []Image
}

type ptNode struct {
	image    Image
	children []*ptNode
	parents  []*ptNode
}

type polytree struct {
	m map[string]*ptNode
}

func NewPolyTree() PolyTree {
	return &polytree{
		m: make(map[string]*ptNode),
	}
}

func (pt *polytree) Add(i Image) error {
	// If image already exists in a tree, return an error
	if inTree(pt, i.Name()) {
		return fmt.Errorf("image already exists in tree")
	}

	// Create image node in tree and add it to a map
	pt.m[i.Name()] = &ptNode{
		image:    i,
		children: []*ptNode{},
		parents:  []*ptNode{},
	}

	// For each parent of the image, if it's in the tree, add the node to the
	// slice of children and add the parent to the slice of parents.
	for _, p := range i.Parents() {
		if inTree(pt, p) {
			pt.m[p].children = append(pt.m[p].children, pt.m[i.Name()])
			pt.m[i.Name()].parents = append(pt.m[i.Name()].parents, pt.m[p])
		}
	}

	// Iterate over all available nodes in the tree (map) in order to find
	// children of the new image. If there is a child (it contains the image name)
	// in the parents slice, add it to the children slice.
	for _, n := range pt.m {
		if slices.Contains(n.image.Parents(), i.Name()) {
			pt.m[i.Name()].children = append(pt.m[i.Name()].children, n)
		}

	}

	return nil
}

func inTree(pt *polytree, i string) bool {
	_, ok := pt.m[i]
	return ok
}

func (pt *polytree) Get(i string) Image {
	if inTree(pt, i) {
		n := pt.m[i]
		return n.image
	}
	return nil
}

func (pt *polytree) Topological(i string) []Image {
	if !inTree(pt, i) {
		return nil
	}

	is := []Image{}
	s := stack.NewStack[Image]()
	topological(pt.m[i], s)

	for !s.Empty() {
		n, _ := s.Pop()
		is = append(is, n)
	}

	return is
}

func topological(n *ptNode, s stack.Stack[Image]) {
	for _, c := range n.children {
		topological(c, s)
	}
	s.Push(n.image)
}

func (pt *polytree) Nodes() []Image {
	is := []Image{}

	for _, n := range pt.m {
		is = append(is, n.image)
	}

	return is
}

func (pt *polytree) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for k, v := range pt.m {
		sb.WriteString(" ")
		sb.WriteString(k)
		sb.WriteString(" c{")
		for _, c := range v.children {
			sb.WriteString(" ")
			sb.WriteString(c.image.Name())
		}
		sb.WriteString(" }")
		sb.WriteString(" p{")
		for _, p := range v.parents {
			sb.WriteString(" ")
			sb.WriteString(p.image.Name())
		}
		sb.WriteString(" }")
	}

	sb.WriteString(" ]")

	return sb.String()
}
