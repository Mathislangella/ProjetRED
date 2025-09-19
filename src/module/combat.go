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

func castSpell(char *Character, monstre *Monster, nom string) bool {
	if char.Ressources.ManaActuels <= 0 {
		fmt.Println("Vous avez pas de mana")
		return false
	}
	if len(char.Sorts) == 0 {
		fmt.Println("Vous avez aucune sort a utiliser")
		return false
	}
	for i := range char.Sorts {
		if char.Sorts[i].Nom == nom {
			sort := char.Sorts[i]
			if char.Ressources.ManaActuels < sort.Mana {
				fmt.Println("Vous n'avez pas de mana")
				return false
			}
			char.Ressources.ManaActuels -= sort.Mana
			damage := sort.Degats + char.Stats.Intelligence/8
			if damage < 0 {
				// negative damage = healing
				heal := -damage
				char.Ressources.PVActuels += heal
				if char.Ressources.PVActuels > char.Ressources.PVMax {
					char.Ressources.PVActuels = char.Ressources.PVMax
				}
				fmt.Printf("%s utilise %s et soigne %d PV\n", char.Nom, sort.Nom, heal)
				fmt.Printf("Mana restante : %d/%d\n", char.Ressources.ManaActuels, char.Ressources.ManaMax)
				return true
			}
			fmt.Printf("%s lance %s et inflicte %d de dégâts\n", char.Nom, sort.Nom, damage)
			fmt.Printf("Mana restante : %d/%d\n", char.Ressources.ManaActuels, char.Ressources.ManaMax)
			monstre.ressources.PVActuels -= damage
			if monstre.ressources.PVActuels < 0 {
				monstre.ressources.PVActuels = 0
			}
			return true
		}
	}
	fmt.Println("Vous avez pas ce sort :", nom)
	return false
}

func characterTurn(char *Character, monstre *Monster, tour int) {
	fmt.Println("Votre choix (1-4): ")
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
		characterTurn(char, monstre, tour)
	case "3":
		Clear()
		if len(char.Sorts) == 0 {
			fmt.Println("Vous n'avez aucun sort")
			characterTurn(char, monstre, tour)
			return
		}
		fmt.Println("Choisissez un sort :")
		for i, sort := range char.Sorts {
			fmt.Printf("%d. %s (Mana necessaire : %d)", i+1, sort.Nom, sort.Mana)
			if char.Ressources.ManaActuels < sort.Mana {
				fmt.Printf(" Mana insuffisant")
			}
			fmt.Println()
		}
		var spellChoice int
		fmt.Scan(&spellChoice)
		if spellChoice < 1 || spellChoice > len(char.Sorts) {
			fmt.Println("Mauvais Choix")
			characterTurn(char, monstre, tour)
			return
		}
		ok := castSpell(char, monstre, char.Sorts[spellChoice-1].Nom)

		if !ok {
			characterTurn(char, monstre, tour)
			return
		}
	case "4":
		fmt.Println("Vous fuyez le combat !")
		Menu(char, nil)
	default:
		fmt.Println("Mauvais Choix")
		characterTurn(char, monstre, tour)
	}
}

func trainingFight(char *Character, monstre *Monster) {
	fmt.Printf("Un %s apparaît !\n", monstre.Nom)
	tour := 1
	for monstre.ressources.PVActuels > 0 && char.Ressources.PVActuels > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		fmt.Printf("%s : %d/%d PV\n", char.Nom, char.Ressources.PVActuels, char.Ressources.PVMax)
		fmt.Printf("%s : %d/%d PV\n", monstre.Nom, monstre.ressources.PVActuels, monstre.ressources.PVMax)
		characterTurn(char, monstre, tour)
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
