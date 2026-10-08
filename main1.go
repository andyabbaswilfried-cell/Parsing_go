package main

import (
	"fmt"
	"os"
	"strings"
	"main.go/folder_func"
)

func main() { // Il lit le fichier, applique chaque règle de modification dans l'ordre.
	entre := os.Args[1]

	txt, err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	} else {

	}
		res := strings.Fields(string(txt))
		res = folderfunc.Bin(res)
		res= folderfunc.Cap(res)
		res= folderfunc.Hex(res)
		res= folderfunc.Lower(res)
		res= folderfunc.Up(res)
		res=folderfunc.Up_p(res)
		fmt.Println(res)
}

func regles(matrice []string) []string {
	
	matrice = folderfunc.Bin(matrice)
	
	return matrice
}


