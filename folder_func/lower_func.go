// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import "strings"

// Lower transforme le mot précédent en minuscules lorsque le token (low) est trouvé.
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
