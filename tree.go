package tree

import (
	"strings"
)

type Node struct {
	Path  string
	Data  interface{}
	Child []Node
}

type Tree struct {
	Name string
	root []Node
}

func NewTree(name string) *Tree {
	tree := &Tree{
		Name: name,
	}
	return tree
}

// func Make Not Data Set
func (t *Tree) Make(method, path string) {
	var i int = 0
	if path == "" {
		return
	}

	// path : path/path => /path/path
	b := []byte(path)
	if b[0] == '/' {
		b = b[1:]
		path = string(b)
	}

	// POST /path/path => /path/path/POST
	newPath := path + "/" + method

	p := strings.Split(newPath, "/")
	for i = 0; i < len(t.root); i++ {
		if t.root[i].Path == p[0] {
			break
		}
	}

	if i == len(t.root) {
		t.root = append(t.root, Node{Path: p[0]})
	}

	if len(p) != 1 {
		t.root[i].Child = tmake(t.root[i].Child, p[1:])
	}
}

func tmake(node []Node, p []string) []Node {
	var i int
	for i = 0; i < len(node); i++ {
		if node[i].Path == p[0] {
			break
		}
	}
	if i == len(node) {
		node = append(node, Node{Path: p[0]})
	}

	if len(p) != 1 {
		node[i].Child = tmake(node[i].Child, p[1:])
	}
	return node
}

/*
func (t *Tree) Del(path string) error   {}
func del(node []Node, p []string) error {}
func (t *Tree) Draw(path string) Node {}
func draw(node []Node, p []string) Node {}
*/
