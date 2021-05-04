package main

import (
	"fmt"
	"urltree"
)

func main() {
	utree := urltree.NewTree("User")
	utree.Make("PUT", "/em-mmc/v1/command/dis-*")
	utree.Make("PUT", "/em-mmc/v1/command/chg-sys-*")
	utree.Printing()

	var Method string = "PUT"
	var inputURL = []string{
		"/em-mmc/v1/command/dis-sys-rsc",
		"/em-mmc/v1/command/dis-alm-info",
		"/em-mmc/v1/command/chg-sys-info",
		"/em-mmc/v1/command/chg-info-rsc",
		"/em-mmc/v1/command/del-sys-rsc",
		"/em-mmc/v1/command/chg-proc-rsc",
		"/em-mmc/v1/command/del-sys-rsc",
		"/em-mmc/v1/command/chg-proc-rsc"}

	for j := 0; j < len(inputURL); j++ {
		ok, method, url := utree.MatchURL(Method, inputURL[j])
		fmt.Printf("Bool = %-15t Method = %-7s URL = %s\n", ok, method, url)
	}

}
