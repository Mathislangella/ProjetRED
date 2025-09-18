package module

import (
	"fmt"
)

// Menu principal du jeu
// Affiche les options et redirige vers les écrans dédiés
func Menu(char *Character, marchand *[]Item) {
	fmt.Printf("Menu Du Jeu\n")
	fmt.Printf("1. Ouvrir Info Personnages\n")
	fmt.Printf("2. Ouvrir Inventaire\n")
	fmt.Printf("3. Forgeron\n")
	fmt.Printf("4. Marchand\n")
	fmt.Printf("5. Combat d'entrainement\n")
	fmt.Printf("0. Quitter Menu\n")
	fmt.Printf("Votre choix (0-5): ")
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
	case "5":
		// Lancer un combat d'entrainement contre un Gobelin
		gobelin := InitGobelin()
		trainingFight(char, &gobelin)
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
