package module

import (
	"fmt"
)

// Menu principal du jeu
// Affiche les options et redirige vers les écrans dédiés
func Menu(char *Character, marchand *[]Item) {
	fmt.Println("Menu Du Jeu")
	fmt.Println("1. Ouvrir Info Personnages")
	fmt.Println("2. Ouvrir Inventaire")
	fmt.Println("3. Forgeron")
	fmt.Println("4. Marchand")
	fmt.Println("0. Quitter Menu")
	fmt.Print("Votre Choix : ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		DisplayInfo(char)
		Menu(char, marchand)
	case "2":
		Clear()
		AccessInventory(char, marchand)
		Clear()
		Menu(char, marchand)
	case "3":
		Clear()
		Forgeron(char)
		Menu(char, marchand)
	case "4":
		Clear()
		Marchand(char, marchand)
		Menu(char, marchand)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Clear()
		Menu(char, marchand)
	}
}
