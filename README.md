# Go-reloaded

## Présentation
Go-reloaded est un outil en ligne de commande écrit en Go qui corrige et transforme un texte selon un ensemble de règles. Il lit un fichier texte, applique chaque règle dans l'ordre, puis écrit le résultat dans `sortie.txt`.

Projet réalisé à Ynov Bordeaux.

## Fonctionnalités

| Règle | Entrée | Sortie |
|---|---|---|
| `(hex)` : convertit le mot précédent de l'hexadécimal vers le décimal | `1E (hex) files were added` | `30 files were added` |
| `(bin)` : convertit le mot précédent du binaire vers le décimal | `It has been 10 (bin) years` | `It has been 2 years` |
| `(up)` : mot précédent en majuscules | `Ready, set, go (up) !` | `Ready, set, GO!` |
| `(low)` : mot précédent en minuscules | `I should stop SHOUTING (low)` | `I should stop shouting` |
| `(cap)` : mot précédent avec une majuscule initiale | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| `(up, N)` : les N mots précédents en majuscules | `This is so exciting (up, 2)` | `This is SO EXCITING` |
| Ponctuation `. , ! ? : ;` et groupes (`...`, `!?`) collés au mot précédent | `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` |
| Apostrophes `'` collées aux mots qu'elles entourent | `describe me: ' awesome '` | `describe me: 'awesome'` |
| `a` devient `an` devant une voyelle ou un `h` | `There it was. A amazing rock!` | `There it was. An amazing rock!` |

## Technologies
- Go
- Bibliothèque standard uniquement : `os`, `fmt`, `strings`, `strconv`
- Package `testing` pour les tests

## Prérequis
- Go, version indiquée dans `go.mod` (1.27.1)
- Git, pour récupérer le projet

## Installation
```bash
git clone <url-du-depot>
cd <dossier-du-projet>
```

## Configuration
Aucune variable d'environnement n'est nécessaire. Le seul paramètre est le chemin du fichier à corriger, passé en argument. Le résultat est toujours écrit dans `sortie.txt`, dans le dossier courant.

## Lancement
```bash
go run . sample.txt
```
Le texte corrigé est écrit dans `sortie.txt`.

## Tests
```bash
go test ./...
```

## Arborescence du projet
```text
Parsing_go/
├── README.md
├── go.mod
├── main1.go                     # point d'entrée : lit le fichier, applique les règles et écrit sortie.txt
├── sample.txt                   # exemple de fichier d'entrée
├── sortie.txt                   # résultat généré lors de l'exécution
└── folder_func/                 # package contenant une règle par fichier
    ├── a_to_an.go
    ├── bin_func.go
    ├── cap_func.go
    ├── hex_func.go
    ├── lower_func.go
    ├── ponctuation.go
    ├── up_func.go
    └── up_precedent.go
```

## Auteur
Andy Abbas - Ynov Bordeaux