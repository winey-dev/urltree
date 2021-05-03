package main

import (
	"urltree"
)

func main() {
	tree := urltree.NewTree("Sample")
	tree.Make("GET", "/api/v1/user")
	tree.Make("DELETE", "/api/v1/user")
	tree.Make("GET", "/api/v1/user/:name")
	tree.Make("POST", "/api/v1/user/:name")
	tree.Make("DELETE", "/api/v1/user/:name")
	tree.Make("GET", "/api/v1/rules")
	tree.Make("DELETE", "/api/v1/rules")
	tree.Make("GET", "/api/v1/rules/:name")
	tree.Make("POST", "/api/v1/rules/:name")
	tree.Make("DELETE", "/api/v1/rules/:name")

	tree.Printing()

	/*
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
	*/
}
