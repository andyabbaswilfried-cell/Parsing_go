// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import "strings"

// Ponctuation_simple normalise la ponctuation et les apostrophes autour des mots.
func Ponctuation_simple(phrase []string) []string {
	var res []string
	comp_apostrofe:=0
	var apostre_start int

	for i := 0; i < len(phrase); i++ { 
		if Is_ponctuation(phrase[i]) {// On s'occupe du cas où la ponctuation est isolé par un espace.
			if len(res) >= 1 {
				res[len(res)-1] = res[len(res)-1] + phrase[i]

			}

		} else {
			switch rune(phrase[i][0]) { // On s'occupe du cas où la ponctuation est collée au début d'un mot.
			case '!', '?', ':', ',', '.',';' :
				after_word:= strings.TrimLeft(phrase[i], "!?,;:.")
				
				length_ponctuation:= len(phrase[i]) - len(after_word)
				if len(res) > 0{
					 res[len(res)-1] = res[len(res)-1]+ phrase[i][:length_ponctuation]
				}
					res =append(res, after_word)
				
				
				 
			case '\'':// On s'occupe du cas des guillemets.
				comp_apostrofe++
				if comp_apostrofe ==1{
					
					res = append(res, phrase[i])
					apostre_start = len(res)-1
					
				}else	if comp_apostrofe ==2{	
						if  apostre_start+1<len(res) {
						res[len(res)-1] = res[len(res)-1]+phrase[i] 
						res[apostre_start] = res[apostre_start]+  res[apostre_start+1]
						  res = append(res[:apostre_start+1],
                   res[apostre_start+2:]...)
				    comp_apostrofe=0
						  } else{
							res = append(res, phrase[i])
							 comp_apostrofe=0
						  }
						 
					} else{
						res = append(res, phrase[i])
					}					
			default:
				res = append(res, phrase[i])
			}
		}

	}
	return res
}
// Is_ponctuation vérifie si un token ne contient que des signes de ponctuation.
func Is_ponctuation(mot string) bool {
	for _, r := range mot {
		if r != '.' && r != '!' && r != ':' && r != ';' && r != '?' && r != ',' {
			return false
		}
	}
	return true

}
