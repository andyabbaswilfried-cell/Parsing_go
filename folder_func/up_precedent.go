package folderfunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Up_p(phrase []string) []string {
	var res []string
	for i := 0; i < len(phrase); i++ {
		if phrase[i] == "(up," &&  i+1<len(phrase) {
			nbr := Split(phrase[i+1], ")")
			NBR, err := strconv.Atoi(nbr[0])
			if err != nil {
				fmt.Println("il y a un problème avec up")
			} else {
				for word := len(res) - NBR; word < len(res); word++ {
					res[word] = strings.ToUpper(res[word])				
				}
				i++
			}
		} else {
			res = append(res, phrase[i])
		}
	}
	return res
}

func Split(sentence string, sep string) []string {
	var result []string
	var tempo string
	var result1 string
	for i := 0; i < len(sentence); i++ {
		tempo = string(sentence[i])
		
		if sentence[i] == 41 {
			result = append(result, result1)
			result1=""
		} else {
			tempo = string(sentence[i])
			result1 += tempo
		}
	}
	return result

}


