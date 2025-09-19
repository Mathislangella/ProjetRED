package module

import "fmt"

func InitGobelin() Monster {
	return Monster{
		Nom:          "Gobelin d'entrainement",
		statistiques: Statistiques{Force: 8, Intelligence: 5, Agilite: 12, Chance: 7, Initiative: 10},
		ressources:   Ressources{PVMax: 30, PVActuels: 30, ManaMax: 0, ManaActuels: 0, Nb_vie: 1},
	}
}

func GobelinPattern(char *Character, monstre *Monster, tour int) {
	fmt.Printf("\n--- Tour %d ---\n", tour)
	var degats int
	if tour%3 == 0 {
		degats = monstre.statistiques.Force * 2
	} else {
		degats = monstre.statistiques.Force
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", monstre.Nom, char.Nom, degats)
	char.Ressources.PVActuels -= degats
	if char.Ressources.PVActuels < 0 {
		char.Ressources.PVActuels = 0
	}

}

func charaterTurn(char *Character, monstre *Monster, tour int) {
	fmt.Printf("\n--- Tour %d ---\n", tour)
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Println("4. Fuir le Combat")
	fmt.Println("Votre choix (1-4): ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		var degats int
		degats = 5
		fmt.Printf("%s inflige à %s %d de dégâts\n", char.Nom, monstre.Nom, degats)
		monstre.ressources.PVActuels -= degats
		if monstre.ressources.PVActuels < (monstre.ressources.PVMax/2) && monstre.ressources.PVActuels > 0 {
			fmt.Printf("%s Zebi !!! Je suis blesser !\n", monstre.Nom)
		}
		if monstre.ressources.PVActuels < 0 {
			monstre.ressources.PVActuels = 0
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

func trainingFight(char *Character, monstre *Monster) {
	fmt.Printf("Un %s apparaît !\n", monstre.Nom)
	tour := 1
	for monstre.ressources.PVActuels > 0 && char.Ressources.PVActuels > 0 {
		fmt.Printf("%s : %d/%d PV\n", char.Nom, char.Ressources.PVActuels, char.Ressources.PVMax)
		fmt.Printf("%s : %d/%d PV\n", monstre.Nom, monstre.ressources.PVActuels, monstre.ressources.PVMax)
		charaterTurn(char, monstre, tour)
		if monstre.ressources.PVActuels <= 0 {
			fmt.Printf("%s a été vaincu !\n", monstre.Nom)
			char.Argent += 10
			fmt.Printf("%s gagne 10 pièces d'or et 10 xp!\n", char.Nom)
			addXP(char, 10)
			return
		}
		if char.Ressources.PVActuels <= 0 {
			char.Ressources.PVActuels = 0
			fmt.Printf("%s est KO !\n", char.Nom)
			IsDead(char)
			return
		}
		GobelinPattern(char, monstre, tour)
		tour++
	}
}
