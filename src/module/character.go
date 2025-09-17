package module

import (
	"fmt"
	"os"
	"strings"
)

func InitCharacter(nom string, classe string, LvL Level, stat Statistiques, ressource Ressources, sorts []Skill, inventaire []Item) Character {
	return Character{
		Nom:          nom,
		Classe:       classe,
		Niveau:       LvL,
		Stats:        stat,
		Ressources:   ressource,
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
	var mana int
	var nb_vie int = 3
	var initiative int = 10
	var force int = 10
	var intelligence int = 10
	var agilite int = 10
	var chance int = 10
	switch choi {
	case "1":
		classe = "Humain"
		pv = 100
		mana = 50
		initiative = 10
		force = 10
		intelligence = 10
		agilite = 10
		chance = 10
	case "2":
		classe = "Elf"
		pv = 80
		mana = 80
		initiative = 15
		force = 8
		intelligence = 12
		agilite = 12
		chance = 13
	case "3":
		classe = "Nain"
		pv = 120
		mana = 30
		initiative = 8
		force = 12
		intelligence = 8
		agilite = 8
		chance = 7
	default:
		fmt.Printf("Pour choisir votre classe choisisez 1 ,2 ou 3 \n")
		CharacterCreation()
	}
	return InitCharacter(nom, classe, Level{1, 30, 0}, Statistiques{initiative, force, intelligence, agilite, chance}, Ressources{pv / 2, pv, mana, mana, nb_vie}, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
}

func IsDead(char *Character) {
	if char.Ressources.PVActuels == 0 {
		if char.Ressources.Nb_vie > 0 {
			fmt.Println("1. Ressurection")
			fmt.Println("2. Quitter Le Jeu")
			var choice string
			fmt.Scan(&choice)
			switch choice {
			case "1":
				AddPV(char, char.Ressources.PVMax/2)
				fmt.Printf("%s gagne %d points de vie", char.Nom, char.Ressources.PVMax/2)
				fmt.Printf(": %d/%d", char.Ressources.PVActuels, char.Ressources.PVMax)
				char.Ressources.Nb_vie--
				return
			case "2":
				os.Exit(0)
			}
		}
		if char.Ressources.Nb_vie == 0 {
			fmt.Println(char.Nom, "est mort")
		}
	}
}

func AddPV(char *Character, nb int) {
	char.Ressources.PVActuels += nb
	if char.Ressources.PVActuels > char.Ressources.PVMax {
		char.Ressources.PVActuels = char.Ressources.PVMax
	}
}

func DisplayInfo(char *Character) {
	Clear()
	fmt.Printf("               ╔══════════════════╗\n")
	fmt.Printf("               ║ Nom    : %-8s║\n", char.Nom)
	fmt.Printf("               ║ Classe : %-8s║ \n", char.Classe)
	fmt.Printf("               ║ Niveau : %-8d║\n", char.Niveau.Lvl)
	fmt.Printf("               ║ Exp    : %-8d║\n", char.Niveau.Exp)
	fmt.Printf("               ║ PV     :  %d/%-4d║\n", char.Ressources.PVActuels, char.Ressources.PVMax)
	fmt.Printf("               ║ OR     : %-8d║\n", char.Argent)
	fmt.Printf("╔══════════════╩══════════════════╩══════════════╗\n")
	fmt.Printf("║------------------ Equipement ------------------║\n")
	fmt.Printf("║ Tete   : %-38s║\n", char.Equipement.Tete)
	fmt.Printf("║ Torse  : %-38s║\n", char.Equipement.Torse)
	fmt.Printf("║ Pieds  : %-38s║\n", char.Equipement.Pieds)
	fmt.Printf("╚════════════════════════════════════════════════╝\n")
}
