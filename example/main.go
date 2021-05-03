package main

import (
	"fmt"
	"urltree"
)

func main() {
	tree := urltree.NewTree("Default")
	tree.Make("GET", "/em-ui/v1/user/:username")
	tree.Make("PATCH", "/em-ui/v1/user/:username")
	tree.Make("PUT", "/em-ui/v1/user/:username")
	tree.Make("PUT", "/em-ui/v1/user/:username/:mapping")
	tree.Make("PUT", "/em-ui/v1/user")
	tree.Make("PUT", "/em-ui/v1/user/*")
	tree.Printing()

	fmt.Println()
	tree2 := urltree.NewTree("User")
	tree2.Make("GET", "/em-ui/v1/user/:username")
	tree2.Make("PATCH", "/em-ui/v1/user/:username")
	tree2.Make("PUT", "/em-ui/v1/user/:username")
	tree2.Make("PUT", "/em-ui/v1/user/:username/:mapping")
	tree2.Make("PUT", "/em-ui/v1/user")
	tree2.Make("PUT", "/em-ui/v1/user/*")
	tree2.Printing()

	t1 := urltree.GetTree("Default")
	t2 := urltree.GetTree("User")

	t1.Printing()
	t2.Printing()

}
