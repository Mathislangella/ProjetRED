package module

import "fmt"

func InitGobelin() Monster {
	return Monster{
		Nom:          "Gobelin d'entrainement",
		PVMax:        40,
		PVActuels:    40,
		AttaquePoint: 5,
	}
}

func GobelinPattern(char *Character, monstre *Monster, tour int) {
	fmt.Printf("\n--- Tour %d ---\n", tour)
	var degats int
	if tour%3 == 0 {
		degats = monstre.AttaquePoint * 2
	} else {
		degats = monstre.AttaquePoint
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", monstre.Nom, char.Nom, degats)
	char.Stats.PVActuels -= degats
	if char.Stats.PVActuels < 0 {
		char.Stats.PVActuels = 0
	}

}

func charaterTurn(char *Character, monstre *Monster, tour int) {
	fmt.Printf("\n--- Tour %d ---\n", tour)
	fmt.Println("Votre choix (1-4): ")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Println("4. Fuir le Combat")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		var degats int
		degats = 5
		fmt.Printf("%s inflige à %s %d de dégâts\n", char.Nom, monstre.Nom, degats)
		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
	case "2":
		Clear()
		AccessInventory(char, nil)
		charaterTurn(char, monstre, tour)
	case "3":
		fmt.Println("Vous fuyez le combat !")
		Menu(char, nil)
	default:
		fmt.Println("Mauvais Choix")
		charaterTurn(char, monstre, tour)
	}
}
