package shanghai

type Tree interface {
	Add(i Image, p string) bool
	Find(i string) Image
	FindChildren(i string) []Image
	Preorder() []Image
	PreorderFrom(i string) []Image
}

type node struct {
	image    Image
	children []*node
	parent   *node
}

type tree struct {
	root *node
	// implement using an index instead of traversing a tree
	index map[string]*node
}

func NewTree() Tree {
	return &tree{
		root:  nil,
		index: make(map[string]*node),
	}
}

func (t *tree) Add(i Image, p string) bool {
	if t.root == nil {
		t.root = &node{
			image:    i,
			children: []*node{},
			parent:   nil,
		}
		return true
	}

	if p == t.root.image.Name() {
		n := &node{
			image:    i,
			children: []*node{t.root},
			parent:   nil,
		}

		t.root = n
		return true
	}

	return bfsAndAdd(t, i, p)
}

func bfsAndAdd(t *tree, i Image, p string) bool {
	q := NewQueue()
	q.Enqueue(t.root)

	for !q.Empty() {
		n, err := q.Dequeue()
		if err != nil {
			return false
		}

		if n.image.Name() == p {
			n.children = append(n.children, &node{
				image:    i,
				children: []*node{},
				parent:   n,
			})

			return true
		}

		for _, n := range n.children {
			q.Enqueue(n)
		}
	}

	return false
}

func (t *tree) Find(i string) Image {
	return bfsAndGet(t, i).image
}

func (t *tree) FindChildren(i string) []Image {
	is := make([]Image, 0)
	for _, n := range bfsAndGet(t, i).children {
		if n != nil {
			is = append(is, n.image)
		}
	}
	return is
}

func bfsAndGet(t *tree, i string) *node {
	q := NewQueue()
	q.Enqueue(t.root)

	for !q.Empty() {
		n, err := q.Dequeue()
		if err != nil {
			return nil
		}

		if n.image.Name() == i {
			return n
		}

		for _, n := range n.children {
			q.Enqueue(n)
		}
	}

	return nil
}

func (t *tree) Preorder() []Image {
	is := make([]Image, 0)
	preorder(t.root, is)
	return is
}

func (t *tree) PreorderFrom(i string) []Image {
	is := make([]Image, 0)
	preorder(bfsAndGet(t, i), is)
	return is
}

func preorder(n *node, is []Image) {
	is = append(is, n.image)

	for _, n := range n.children {
		preorder(n, is)
	}
}
