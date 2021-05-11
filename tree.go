package urltree

import (
	"regexp"
	"strings"
)

var treeArray map[string]*Tree

type Node struct {
	path  string
	child []Node
}

type Tree struct {
	name string
	root []Node
}

func GetTree(name string) *Tree {
	if name == "" {
		return nil
	}
	return treeArray[name]
}

func NewTree(name string) *Tree {
	tree := &Tree{
		name: name,
	}

	if len(treeArray) == 0 {
		treeArray = make(map[string]*Tree)
	}
	treeArray[name] = tree
	return tree
}

func (t *Tree) Make(method, path string) {

	if path == "" || method == "" {
		return
	}

	// path : /path/path => path/path
	b := []byte(path)
	if b[0] == '/' {
		b = b[1:]
		path = string(b)
	}

	// POST path/path => path/path/$POST
	newPath := path + "/" + "$" + strings.ToUpper(method)

	p := strings.Split(newPath, "/")

	t.root = tmake(t.root, p)

}

func tmake(node []Node, p []string) []Node {
	var i int
	for i = 0; i < len(node); i++ {
		if node[i].path == p[0] {
			break
		}
	}
	if i == len(node) {
		node = append(node, Node{path: p[0]})
	}

	if len(p) != 1 {
		node[i].child = tmake(node[i].child, p[1:])
	}
	return node
}

func (t *Tree) MatchURL(method, path string) (bool, string, string) {
	var ok bool
	var npath string

	b := []byte(path)
	if b[0] == '/' {
		b = b[1:]
		npath = string(b)
	}

	newPath := npath + "/" + "$" + strings.ToUpper(method)

	p := strings.Split(newPath, "/")
	if len(p) > 1 {
		ok = matchurl(t.root, p)
	} else {
		ok = false
	}

	return ok, method, path
}

func matchurl(node []Node, p []string) bool {
	var ok bool
	var i int

	for i = 0; i < len(node); i++ {
		if node[i].path[0] == '$' { // $ Method Field
			if node[i].path[1] == '*' || // * All Allowed Method
				node[i].path == p[0] { // Perfect Method Matching
				if len(p) > 1 {
					continue
				} else {
					ok = true

				}
				break
			}
		} else if node[i].path[0] == ':' || // : Dynamic Path
			node[i].path[0] == '*' {
			if len(p) > 1 {
				ok = true
			} else {
				continue
			}
			break
		} else if node[i].path == p[0] { // Perfect Path Matching
			ok = true
			break
		} else if node[i].path[0] != '*' && strings.Contains(node[i].path, "*") {
			// node path set '/user*'  match /username
			if ok, _ = regexp.MatchString(node[i].path, p[0]); !ok {
				continue
			} else {
				break
			}
		}
	}

	if i == len(node) {
		return false
	}

	if len(p) > 1 {
		if node[i].path == "*" {
			p = p[len(p)-2:]
		}
		ok = matchurl(node[i].child, p[1:])
	}

	return ok
}

func (t *Tree) Destroy() {
	if treeArray == nil {
		return
	}
	delete(treeArray, t.name)
}
