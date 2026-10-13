// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import (
	"fmt"
	"strconv"
	"strings"
)

// Up_precedent transforme les N mots précédents en majuscules avec le motif (up, N).
func Up_precedent(phrase []string) []string {
	var res []string
	for i := 0; i < len(phrase); i++ {
		if phrase[i] == "(up," && i+1 < len(phrase) {
			nbr_string := Split(phrase[i+1], ")")
			nombre_int, err := strconv.Atoi(nbr_string[0])
			if err != nil{
				fmt.Println("Il semble y avoir un problème avec la conversion en chiffre")
			}else if nombre_int >len(res){
				fmt.Println("la grandeur est trop élevé, le système ne peut la traiter")
			}else{
				for word:= len(res)- nombre_int; word<len(res) ;word++{
					res[word] = strings.ToUpper(res[word])
					
				}
				i++
			}
		} else{
			res = append(res, phrase[i])
		}
	}
	return res
}

// Split sépare la valeur numérique du motif de déclenchement et renvoie les éléments extraits.
func Split(sentence string, sep string) []string {
	var result []string
	var tempo string
	var result1 string
	for i := 0; i < len(sentence); i++ {
		tempo = string(sentence[i])

		if sentence[i] == 41 {
			result = append(result, result1)
			result1 = ""
		} else {
			tempo = string(sentence[i])
			result1 += tempo
		}
	}
	return result

}
