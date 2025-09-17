package module

import (
	"fmt"
	"os"
	"strings"
)

func InitCharacter(nom string, classe string, LvL Level, PVMax int, Nb_vie int, sorts []Skill, inventaire []Item) Character {
	return Character{
		Nom:          nom,
		Classe:       classe,
		Niveau:       LvL,
		Stats:        Statistiques{PVActuels: PVMax / 2, PVMax: PVMax},
		Nb_vie:       Nb_vie,
		Sorts:        sorts,
		Inventaire:   inventaire,
		InventoryMax: 10,
		Argent:       100,
	}
}

func CharacterCreation() Character {
	fmt.Println("Quel est votre nom :")
	var nom string
	fmt.Scan(&nom)
	nom = strings.ToLower(nom)
	nom0 := strings.Split(nom, "")
	nom0[0] = strings.ToUpper((string(nom[0])))
	nom = strings.Join(nom0, "")
	fmt.Println("Quel est votre Classe choisissez |1. Humain|2. Elf|3. Nain")
	var choi string
	fmt.Scan(&choi)
	var classe string
	var pv int
	switch choi {
	case "1":
		classe = "Humain"
		pv = 100
	case "2":
		classe = "Elf"
		pv = 80
	case "3":
		classe = "Nain"
		pv = 120
	default:
		fmt.Printf("Pour choisir votre choisisez 1 ,2 ou 3")
		CharacterCreation()
	}
	return InitCharacter(nom, classe, Level{1, 0}, pv, 1, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
}

func IsDead(char *Character) {
	if char.Stats.PVActuels == 0 {
		if char.Nb_vie > 0 {
			fmt.Println("1. Ressurection")
			fmt.Println("2. Quitter Le Jeu")
			var choice string
			fmt.Scan(&choice)
			switch choice {
			case "1":
				AddPV(char, char.Stats.PVMax/2)
				fmt.Printf("%s gagne %d points de vie", char.Nom, char.Stats.PVMax/2)
				fmt.Printf(": %d/%d", char.Stats.PVActuels, char.Stats.PVMax)
				char.Nb_vie--
				return
			case "2":
				os.Exit(0)
			}
		}
		if char.Nb_vie == 0 {
			fmt.Println(char.Nom, "est mort")
		}
	}
}

func AddPV(char *Character, nb int) {
	char.Stats.PVActuels += nb
	if char.Stats.PVActuels > char.Stats.PVMax {
		char.Stats.PVActuels = char.Stats.PVMax
	}
}

func DisplayInfo(char *Character) {
	Clear()
	fmt.Println("╔══════════════════╗")
	fmt.Printf("║ Nom    : %-8s║\n", char.Nom)
	fmt.Printf("║ Classe : %-8s║ \n", char.Classe)
	fmt.Printf("║ Niveau : %-8d║\n", char.Niveau.Lvl)
	fmt.Printf("║ Exp    : %-8d║\n", char.Niveau.Exp)
	fmt.Printf("║ PV     :  %d/%-4d║\n", char.Stats.PVActuels, char.Stats.PVMax)
	fmt.Printf("║ OR     : %-8d║\n", char.Argent)
	fmt.Println("║------------------ Equipement ------------------║")
	fmt.Printf("║ Tete  : %-25s\n", char.Equipement.Tete)
	fmt.Printf("║ Torse : %-25s\n", char.Equipement.Torse)
	fmt.Printf("║ Pieds : %-25s\n", char.Equipement.Pieds)
	fmt.Println("╚════════════════════════════════════════════════╝")
}
