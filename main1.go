
// Package main est le point d'entrée de l'outil Go-reloaded.
// Il lit un fichier texte, applique les règles de transformation
// et écrit le résultat corrigé dans sortie.txt.
package main

import (
	"fmt"
	"os"
	"strings"

	folderfunc "main.go/folder_func"
)

// main lit le chemin fourni en argument, applique les règles et
// enregistre le texte transformé dans le fichier de sortie.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Veuillez fournir exactement un argument : le chemin du fichier à corriger.")
		return
	}

	entre := os.Args[1]
	txt, err := os.ReadFile(entre)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	lignes := strings.Split(string(txt), "\n")
	lignes_finales := make([]string, 0, len(lignes))

	for _, ligne := range lignes {
		mots := strings.Fields(ligne)
		if len(mots) == 0 {
			lignes_finales = append(lignes_finales, "")
			continue
		}

		mots = regles(mots)
		ligne_finale := strings.Join(mots, " ")
		lignes_finales = append(lignes_finales, ligne_finale)
	}

	phrase := strings.Join(lignes_finales, "\n")

	err = os.WriteFile("sortie.txt", []byte(phrase), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}

// regles applique la chaîne de transformations sur un tableau de mots.
func regles(matrice []string) []string {
	if len(matrice) == 0 {
		return matrice
	}

	matrice = folderfunc.Bin(matrice)
	matrice = folderfunc.Cap(matrice)
	matrice = folderfunc.Hex(matrice)
	matrice = folderfunc.Lower(matrice)
	matrice = folderfunc.Up(matrice)
	matrice = folderfunc.Up_precedent(matrice)
	matrice = folderfunc.Ponctuation_simple(matrice)
	matrice = folderfunc.Indefinite_article_rule(matrice)

	return matrice
}