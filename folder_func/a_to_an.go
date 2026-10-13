// Package folderfunc regroupe les règles de transformation du texte.
package folderfunc

import "strings"

// Indefinite_article_rule remplace "a" par "an" devant une voyelle ou le son h.
func Indefinite_article_rule(phrase []string) []string {
	var res []string

	for i := 0; i < len(phrase); i++ {
		switch phrase[i] {
		case "a":
			if i+1 < len(phrase) {
				if strings.ToLower(string(phrase[i+1][0]))== "a" ||strings.ToLower(string(phrase[i+1][0]))== "e"   || strings.ToLower(string(phrase[i+1][0]))== "i"   || strings.ToLower(string(phrase[i+1][0]))== "o" || strings.ToLower(string(phrase[i+1][0]))== "u" || strings.ToLower(string(phrase[i+1][0]))== "h"{
					res = append(res, "an")
				} else {
					res = append(res, "a")
				}
			} else {
				res = append(res, "a")
			}
		case "A":
			if i+1 < len(phrase) {
				if strings.ToLower(string(phrase[i+1][0]))== "a" ||strings.ToLower(string(phrase[i+1][0]))== "e"   || strings.ToLower(string(phrase[i+1][0]))== "i"   || strings.ToLower(string(phrase[i+1][0]))== "o" || strings.ToLower(string(phrase[i+1][0]))== "u" || strings.ToLower(string(phrase[i+1][0]))== "h"{
					res = append(res, "An")
				} else {
					res = append(res, "A")
				}
			} else {
				res = append(res, "A")
			}

		default:
			res = append(res, phrase[i])

		}
	}
	return res
}
