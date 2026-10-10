package folderfunc

import "strings"
func Ponctuation_simple(phrase []string) []string {
	var res []string

	for i := 0; i < len(phrase); i++ {
		if Is_ponctuation(phrase[i]) {
			if len(res) >= 1 {
				res[len(res)-1] = res[len(res)-1] + phrase[i]

			}

		} else {
			switch rune(phrase[i][0]) {
			case '!', '?', ':', ',', '.',';' :
				after_word:= strings.TrimLeft(phrase[i], "!?,;:.")
				
				length_ponctuation:= len(phrase[i]) - len(after_word)
				if len(res) > 0{
				 res[len(res)-1] = res[len(res)-1]+ phrase[i][:length_ponctuation]
				
				}else if len(res)==0{
				res =append(res, phrase[i])
				}else{
					res =append(res, after_word)
				
				}
				 
				
				
			default:
				res = append(res, phrase[i])
			}
		}

	}
	return res
}
func Is_ponctuation(mot string) bool {
	for _, r := range mot {
		if r != '.' && r != '!' && r != ':' && r != ';' && r != '?' && r != ',' {
			return false

		}

	}
	return true

}
