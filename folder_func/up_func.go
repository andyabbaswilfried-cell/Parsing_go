package folderfunc

import "strings"



func Up(phrase []string) []string {
	var resultat []string

	for i := 0; i < len(phrase); i++ {
		if phrase[i] == "(up)" && i > 0 {			
				resultat[len(resultat)-1] = strings.ToUpper(phrase[i-1])
		} else {
			resultat = append(resultat, phrase[i])
		}
	}
	return resultat
}