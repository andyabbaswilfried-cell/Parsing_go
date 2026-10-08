package folderfunc

import (
	"fmt"
	"strconv"
	"strings"
)

func Up_p(phrase []string) []string {

	var res []string

	for i:=0; i<len(phrase);i++ {
		if phrase[i] == "(up," && i > 0 {
			if len(phrase) > 1 {
				data, err := strconv.Atoi(strings.TrimSuffix(string(phrase[i+1]), ")\u200b"))
				if err != nil {
					fmt.Println("il y a une erreur")
				} else {


					for i:= len(res)-data; i<len(res);i++{
						res[i] = strings.ToUpper(res[i])
					
					}
					i++

				}

			}

		} else {
			res = append(res, phrase[i])

		}
	}
	return res
}
