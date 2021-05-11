package main

import (
	"fmt"
	"urltree"
)

func main() {
	utree := urltree.NewTree("auth")
	utree.Make("GET", "/api/v1/user/:username")
	utree.Make("POST", "/api/v1/user/:username")
	utree.Make("DELETE", "/api/v1/user/:username")
	utree.Make("PATCH", "/api/v1/user/:username")
	utree.Make("GET", "/api/v1/user")
	utree.Make("DELETE", "/api/v1/user")
	utree.Make("GET", "/api/v1/config")
	utree.Make("GET", "/api/v1/config/*")
	utree.Make("*", "/api/v1/alert/:alertname")
	utree.Make("POST", "/api/v1/alert/:alertname/*")
	utree.Printing()

	var Method = []string{"POST", "GET", "DELETE", "PATCH"}
	var InputURL = []string{
		"/api/v1/user",
		"/api/v1/user/yiaw",
		"/api/v1/alert/status",
		"/api/v1/config",
		"/api/v1/config/setting"}

	for i := 0; i < len(InputURL); i++ {
		for j := 0; j < len(Method); j++ {
			ok, method, url := utree.MatchURL(Method[j], InputURL[i])
			fmt.Printf("Bool = %-15t Method = %-7s URL = %s\n", ok, method, url)
		}
	}
}
