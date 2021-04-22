package main

import (
	"fmt"
	"pathtree/pathtree"
)

/*
var m = mark{
	m1: []byte{226, 148, 156},                //├
	m2: []byte{226, 148, 128, 226, 148, 128}, //──
	m3: []byte{226, 148, 148},                //└
	m4: []byte{226, 148, 130},                //│
}
*/
func main() {

	tree := pathtree.NewTree("RouterAPI")
	var str = []string{
		"/ui/v1/user/name/DELETE",
		"/ui/v1/user/name/PUT",
		"/ui/v1/user/path/DELETE",
		"/ui/v1/message/GET",
		"/ui/v1/message/POST",
		"/ui/v1/message/PUT",
	}

	for i := 0; i < len(str); i++ {
		tree.Add(str[i])
	}

	tree.Tprinting()
	fmt.Println()
}
