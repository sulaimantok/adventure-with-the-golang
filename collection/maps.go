package main

import "fmt"

func main() {
	friends := map[string]int{ //mapping strings to integers
		"Abdullah":       50,
		"Ahmad":     	  21,
		"Hasan":          41,
		"husein": 		  29,
	}

	fmt.Printf("%#v", friends)
}