package main

import (
	"os"
	"strconv"
	"strings"

	"fmt"
)

func main() {
	entre := os.Args[1]

	txt, err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	} else {

		res := strings.Fields(string(txt))
		var resultat []string

		for i := 0; i < len(res); i++ {

			switch res[i] {
			case "(hex)":
				if len(resultat) > 1 {

					resultat[len(resultat)-1] = hex(resultat[len(resultat)-1])
				}
			case "(bin)":
				if len(resultat) > 1 {

					resultat[len(resultat)-1] = bin(resultat[len(resultat)-1])
				}
			case "(up)":
				if len(resultat) > 1 {

					resultat[len(resultat)-1] = up(resultat[len(resultat)-1])
				}
			case "(low)":
				if len(resultat) > 1 {

					resultat[len(resultat)-1] = lower(resultat[len(resultat)-1])
				}
			case "(cap)":
				if len(resultat) > 1 {

					resultat[len(resultat)-1] = cap(resultat[len(resultat)-1])
				}
			case "(up,":
				if len(resultat) > 1{
					data,err:= strconv.Atoi(strings.TrimSuffix(res[i+1], ")\u200b"))
					if err!= nil {
						fmt.Println("il y a une erreur")
					}
					i++
					resultat = up_p(resultat, data)
				}
			default:
				// If it's not a special tag, just add the word to our result
				resultat = append(resultat, res[i])

			}

		}
		fmt.Println(resultat)
		///var phrase string
		//for _,r := range resultat{
		///	phrase += r + " "
		///}
		///resultat_final:= []byte(phrase)
		///os.WriteFile("test.txt",resultat_final,0644)
		//fmt.Print(resultat)

	}
}

//os.WriteFile("test.txt",[]byte(txt),0644)//

func hex(nbr string) string {
	decimal, err := strconv.ParseInt(nbr, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(decimal))
}

func bin(nbr string) string {
	decimal, err := strconv.ParseInt(nbr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(decimal))
}

func up(mot string) string {
	return strings.ToUpper(mot)
}

func lower(mot string) string {
	return strings.ToLower(mot)
}

func cap(mot string) string {
	var res string
	for i, r := range mot {
		if i == 0 {
			res += string(r-32) 
		} else {
			res += string(r)
		}

	}
	return res
}

func up_p(phrase []string,nbr int) []string{
	var res []string
	for _,r := range phrase{
		
		if nbr> len(phrase){
			fmt.Println("il y a un problème de mots à mettre en majuscules")
		} else if nbr >=0 {
			if r == phrase[len(phrase)-nbr]{
				res = append(res,up(phrase[len(phrase)-nbr]))
				nbr--
			}else if nbr==0{
				return  res
			} else {
				res = append(res, r)
			}
		}
	}
	return res
}
