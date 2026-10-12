package folderfunc

import "strings"

func Lower(phrase []string) []string {
	var res []string
	for i := 0; i < len(phrase); i++ {
		if phrase[i] == "(low)" && i > 0 {
			res[len(res)-1] = strings.ToLower(res[len(res)-1])
			
		} else {
			res = append(res, phrase[i])
		}
	}
	return res
}
