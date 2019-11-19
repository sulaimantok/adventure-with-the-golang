package main

import (
	"fmt"
	"strings"
  "encoding/json"
   "strconv"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}