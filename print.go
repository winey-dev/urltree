package urltree

import (
	"fmt"
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
	case 5:
		return string("   ")
	}

	return ""
}
func (t *Tree) Printing() {
	if len(t.root) == 0 {
		return
	}

	treeprinting(t.root, -1)
}

func de(d, idx, cur, chl int) {

	if d < 1 {
		return
	}
	for i := 0; i < d-1; i++ {

		fmt.Printf("%s", m.String(5))
	}
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
			de(d, i, cur, chl)
			if i == cur-1 {
				fmt.Printf("%s%s %s\n", m.String(3), m.String(2), root[i].Path)
			} else {
				fmt.Printf("%s%s %s\n", m.String(1), m.String(2), root[i].Path)
			}
		}

		if chl != 0 {
			treeprinting(root[i].Child, d)
		}
	}
}
