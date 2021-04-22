package pathtree

import (
	"fmt"
	"strings"
)

// Special string
type mark struct {
	m1 []byte
	m2 []byte
	m3 []byte
	m4 []byte
}

var m = mark{
	m1: []byte{226, 148, 156},                //├
	m2: []byte{226, 148, 128, 226, 148, 128}, //──
	m3: []byte{226, 148, 148},                //└
	m4: []byte{226, 148, 130},                //│
}

func (m mark) String(i int) string {
	switch i {
	case 1:
		return string(m.m1)
	case 2:
		return string(m.m2)
	case 3:
		return string(m.m3)
	case 4:
		return string(m.m4)
	}

	return ""
}

type Node struct {
	Path     string
	Printing string
	Child    []Node
}

type Tree struct {
	TreeName string
	root     []Node
}

func NewTree(name string) *Tree {
	tree := &Tree{
		TreeName: name,
	}
	return tree
}

func (t *Tree) Add(fullpath string) {
	var i int = 0
	if fullpath == "" {
		return
	}

	b := []byte(fullpath)
	if b[0] == '/' {
		b = b[1:]
		fullpath = string(b)
	}

	path := strings.Split(fullpath, "/")
	for i = 0; i < len(t.root); i++ {
		if t.root[i].Path == path[0] {
			break
		}
	}

	if i == len(t.root) {
		t.root = append(t.root, Node{Path: path[0]})
	}

	t.root[i].Child = add(t.root[i].Child, path[1:])
}

func add(root []Node, p []string) []Node {
	var i int
	for i = 0; i < len(root); i++ {
		if root[i].Path == p[0] {
			break
		}
	}
	if i == len(root) {
		root = append(root, Node{Path: p[0]})
	}

	if len(p) != 1 {
		root[i].Child = add(root[i].Child, p[1:])
	}

	return root
}

func (t *Tree) Tprinting() {
	if len(t.root) == 0 {
		return
	}

	treeprinting(t.root, -1)
}

func treeprinting(root []Node, d int) {
	var i int
	var cur int /* Current Node Len */
	var chl int /* Child Node Len */

	d++
	cur = len(root)

	for i = 0; i < cur; i++ {
		chl = len(root[i].Child)

		if d == 0 {
			fmt.Println(root[i].Path)
		} else {
			if i == cur-1 {
				fmt.Printf("%s%s%s %s\n", root[i].Printing, m.String(3), m.String(2), root[i].Path)
			} else {
				fmt.Printf("%s%s%s %s\n", root[i].Printing, m.String(1), m.String(2), root[i].Path)
			}
		}

		if chl != 0 {
			treeprinting(root[i].Child, d)
		}
	}
}

func (t *Tree) Printing() {
	if len(t.root) == 0 {
		return
	}
	printing(t.root, 0)
}

func printing(root []Node, d int) {
	var i int
	for i = 0; i < len(root); i++ {
		fmt.Printf("[%d]%s ", d, root[i].Path)
		fmt.Println()
	}

	d++
	for i = 0; i < len(root); i++ {
		printing(root[i].Child, d)
	}
}
