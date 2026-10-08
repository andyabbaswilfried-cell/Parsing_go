package folderfunc

import "strings"


func Cap(res []string) []string {
	var resultat []string
	for i := 0; i < len(res); i++ {
		var cap string
		if res[i] == "(cap)" && i > 0 {
			word := []rune(res[i-1])
			if len(word) > 0 {
				
				for i,r := range word{
					if i==0{
						cap += strings.ToUpper(string(r))
					} else{
						cap += string(r)
					}
				}
				
			}
			val_cap := string(cap)
			 resultat[len(resultat)-1]= val_cap
		} else{
			resultat= append(resultat, res[i])
		}
	}
	
	return resultat
}
