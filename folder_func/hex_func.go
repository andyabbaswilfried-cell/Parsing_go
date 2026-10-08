package folderfunc

import (
	"fmt"
	"strconv"
)

func Hex(res []string) []string {
	var resultat []string
	for i := 0; i < len(res); i++ {
		if res[i] == "(hex)" && i > 0 {
			decimal, err := strconv.ParseInt(res[i-1], 16, 32)
			if err != nil {
				fmt.Println("la fonction binaire")
			} else {

				resultat[len(resultat)-1] = strconv.Itoa(int(decimal))
			}
		} else {
			resultat = append(resultat, res[i])
		}

	}
	return resultat
}
