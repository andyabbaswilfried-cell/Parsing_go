package folderfunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Up_p(phrase []string) []string {

	var res []string
	
	for i, r := range phrase {
		if res[i] == "(cap)" && i > 0 {
		if r == "(up," {
			if len(phrase) > 1 {
				data, err := strconv.Atoi(strings.TrimSuffix(res[i+1], ")\u200b"))
				if err != nil {
					fmt.Println("il y a une erreur")
				}
				i++
				res = up_precedent(phrase, data)
			}

		} else{
			res = append(res, r)
		}
		
	} //
	return res
	}
}


func Bin(res []string) []string {
	var resultat []string
	for i := 0; i < len(res); i++ {
		if res[i] == "(bin)" && i > 0 {
			decimal, err := strconv.ParseInt(res[i-1], 2, 64)
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