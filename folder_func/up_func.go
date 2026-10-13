// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import "strings"

// Up transforme le mot précédent en majuscules lorsque le token (up) est trouvé.
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