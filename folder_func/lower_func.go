package folderfunc

import (
	
	"strings"
)

func Lower(res []string) []string {
	var resultat []string

	for i := 0; i < len(res); i++ {
		if res[i] == "(low)" && i > 0 {			
				resultat[len(resultat)-1] = strings.ToLower(res[i-1])
		} else {
			resultat = append(resultat, res[i])
		}

		

	}
	return resultat
}
