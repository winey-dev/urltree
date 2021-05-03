package main

import (
	"fmt"
	"urltree"
)

func main() {
	//var strlist = []string{"DELETE", "PATCH"}
	//p := strlist[len(strlist)-2:]
	//fmt.Println(p)

	/*
		// TEST Make Tree And Get Tree Start ..
		tree := urltree.NewTree("Default")
		tree.Make("GET", "/em-ui/v1/user/:username")
		tree.Make("PATCH", "/em-ui/v1/user/:username")
		tree.Make("PUT", "/em-ui/v1/user/:username")
		tree.Make("PUT", "/em-ui/v1/user/:username/:mapping")
		tree.Make("PUT", "/em-ui/v1/user")
		tree.Make("PUT", "/em-ui/v1/user/*")
		tree.Make("*", "/em-ui/v1/*")
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
		// TEST Make Tree And Get Tree END ..
	*/
	// TEST Matcing Tree

	admintree := urltree.NewTree("Admin")
	admintree.Make("*", "/*")
	//admintree.Printing()
	//fmt.Println()

	utree := urltree.NewTree("User")
	utree.Make("*", "/em-mmc/v1/*")
	utree.Make("GET", "/em-config/v1/nf")
	utree.Make("POST", "/em-alert/v1/alert/*")
	utree.Make("GET", "/em-alert/v1/history/alarm")
	utree.Printing()

	var inputMethod = []string{"GET", "PUT", "POST", "DELETE", "PATCH"}

	var inputURL = []string{
		"/em-mmc/v1/command",
		"/em-mmc/v1/help",
		"/em-config/v1/pod",
		"/em-config/v1/nf",
		"/em-alert/v1/alert/alartname",
		"/em-alert/v1/alert/history",
		"/em-alert/v1/history/alert",
		"/em-mmc/v1/config"}

	for i := 0; i < len(inputMethod); i++ {
		for j := 0; j < len(inputURL); j++ {
			ok, method, url := utree.MatchURL(inputMethod[i], inputURL[j])
			fmt.Printf("Bool = %-15t Method = %-7s URL = %s\n", ok, method, url)
		}
	}
}
