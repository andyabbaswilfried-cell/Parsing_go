// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

// Cap applique la transformation (cap) au mot précédent.
func Cap(phrase []string) []string {
	var res []string
	for i:=0; i<len(phrase);i++{
		if phrase[i]=="(cap)" && i>0{
			res[len(res)-1]= Capitalize(res[len(res)-1])
		
		}else{
			res = append(res, phrase[i])
		}
	}
	return res
}


// Capitalize met la première lettre du mot en majuscule en conservant le reste.
func Capitalize(sentence string) string {
	var result string
	var tempo rune
	if Islower(string(sentence[0])) == true {
		tempo = rune(sentence[0]) - 32
		result += string(tempo)
	} else {
		result += string(rune(sentence[0]))
	}
	for i := 1; i < len(sentence); i++ {
		if isprintable(string(sentence[i-1])) == false && isprintable(string(sentence[i])) == true && sentence[i] >= 97 && sentence[i] <= 122 {
			tempo = rune(sentence[i])
			tempo -= 32
			result += string(tempo)
		} else {
			tempo = rune(sentence[i])
			result += string(tempo)
		}
	}
	return result
}



// isprintable vérifie si tous les caractères d'une chaîne sont alphanumériques imprimables.
func isprintable(sentence string) bool {
	result := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] > 47 && sentence[i] < 58 {
			result = true
			continue
		}
		if sentence[i] > 64 && sentence[i] < 91 {
			result = true
			continue
		}
		if sentence[i] > 96 && sentence[i] < 123 {
			result = true
			continue
		}
		return false
	}
	return result
}

// Islower renvoie true si toute la chaîne est en minuscules.
func Islower(sentence string) bool {
	result := false
	var sentence2 = []byte(sentence)
	for i := 0; i < len(sentence); i++ {
		result = false
		if sentence2[i] >= 97 && sentence2[i] <= 122 {
			result = true
		}
		if result == false {
			return false
		}
	}
	return result
}




