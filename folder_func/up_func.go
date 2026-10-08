package folderfunc

import "strings"



func Up(res []string) []string {
	var resultat []string

	for i := 0; i < len(res); i++ {
		if res[i] == "(up)" && i > 0 {			
				resultat[len(resultat)-1] = strings.ToUpper(res[i-1])
		} else {
			resultat = append(resultat, res[i])
		}

		

	}
	return resultat
}