package main

import (
	"fmt"
	"os"
	"strings"
	"main.go/folder_func"
)

func main() { // Il lit le fichier, applique chaque règle de modification dans l'ordre.
	entre := os.Args[1]
	var phrase string

	txt, err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	} else {

	}
		res := strings.Fields(string(txt))
	
		res= regles(res)
		fmt.Println(res)
		
		for _,r := range res{
		phrase += r + " "
		}
		resultat_final:= []byte(phrase)
		os.WriteFile("sortie.txt",resultat_final,0644)
	

}

func regles(matrice []string) []string {
	
	matrice = folderfunc.Bin(matrice)
	matrice = folderfunc.Bin(matrice)
	matrice= folderfunc.Cap(matrice)
	matrice= folderfunc.Hex(matrice)
	matrice= folderfunc.Lower(matrice)
	matrice= folderfunc.Up(matrice)
	matrice=folderfunc.Up_p(matrice)
	matrice=folderfunc.Ponctuation_simple(matrice)
	matrice=folderfunc.Indefinite_article_rule(matrice)

	return matrice
}


