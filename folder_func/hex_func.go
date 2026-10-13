// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import (
	"fmt"
	"strconv"
)

// Hex convertit le mot précédent de l'hexadécimal vers le décimal lorsque le token (hex) est trouvé.
func Hex(phrase []string) []string {
	var res []string
	for i := 0; i < len(phrase); i++ {
		if phrase[i] == "(hex)" && i > 0 {
			data, err := strconv.ParseInt(res[i-1], 16, 64)
			if err != nil {
				fmt.Println("Il y a un problème avec hex")
			} else {
				res[len(res)-1] = strconv.Itoa(int(data))
				
			}
		} else {
			res = append(res, phrase[i])
		}
	}
	return res
}
