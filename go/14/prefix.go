package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	lenCMP := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(strs, lenCMP)
	var resultPrefix []string
	var count int
	comparative_term := strs[0]
	var new_prefix string
	for _, r := range comparative_term {
		count = 0
		new_prefix = new_prefix + string(r)
		for _, val := range strs {
			if strings.HasPrefix(val, new_prefix) {
				count++
			}
		}
		if count == len(strs) {
			resultPrefix = append(resultPrefix, new_prefix)
		}
	}
	if len(resultPrefix) == 0 {
		return ""
	}
	return resultPrefix[len(resultPrefix)-1]
}

func main() {
	test := []string{"flower", "flow", "flight"}
	oi := longestCommonPrefix(test)
	fmt.Println(oi)
}
