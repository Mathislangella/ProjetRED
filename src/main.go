package main

import "fmt"

type statistiques struct {
	PVActuels int
	PVMax     int
}

type Level struct {
	Lvl int
	Exp int
}

type Item struct {
	Nom      string
	Quantite int
}

type Character struct {
	Nom        string
	Classe     string
	Niveau     Level
	Stats      statistiques
	Inventaire []Item
}

var C1 Character
var C2 Character
var C3 Character
var equipe [3]Character

func Menu() {
	fmt.Println("Menu Du Jeu")
	fmt.Println("1. Ouvrir Info Personnages")
	fmt.Println("2. Ouvrir Inventaire")
	fmt.Println("3. Quitter Menu")
	fmt.Print("Votre Choix : ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		displayInfo()
	case 2:
		accessInventory()
	case 3:
		return
	default:
		fmt.Print("Mauvais Choix")
	}
}

func initCharacter(nom string, classe string, lvl int, PVMax int, inventaire []Item) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     Level{Lvl: 1, Exp: 0},
		Stats:      statistiques{PVActuels: PVMax, PVMax: PVMax},
		Inventaire: inventaire,
	}
}

func displayInfo() {
	fmt.Println("╔══════════════════╗  ╔══════════════════╗  ╔══════════════════╗")
	fmt.Printf("║ Nom    : %-8s║  ║ Nom    : %-8s║  ║ Nom    : %-8s║\n",
		C1.Nom, C2.Nom, C3.Nom)
	fmt.Printf("║ Classe : %-8s║  ║ Classe : %-8s║  ║ Classe : %-8s║\n",
		C1.Classe, C2.Classe, C3.Classe)
	fmt.Printf("║ Niveau : %-8d║  ║ Niveau : %-8d║  ║ Niveau : %-8d║\n",
		C1.Niveau.Lvl, C2.Niveau.Lvl, C3.Niveau.Lvl)
	fmt.Printf("║ Exp    : %-8d║  ║ Exp    : %-8d║  ║ Exp    : %-8d║\n",
		C1.Niveau.Exp, C2.Niveau.Exp, C2.Niveau.Exp)
	fmt.Printf("║ PV     : %d/%-4d║  ║ PV     : %d/%-4d║  ║ PV     : %d/%-4d║\n",
		C1.Stats.PVActuels, C1.Stats.PVMax, C2.Stats.PVActuels, C2.Stats.PVMax, C3.Stats.PVActuels, C3.Stats.PVMax)
	fmt.Println("╚══════════════════╝  ╚══════════════════╝  ╚══════════════════╝")
}

func accessInventory() {
	fmt.Println("		  	  INVENTAIRES")
	for _, r := range equipe {
		fmt.Printf("║%-8s : ", r.Nom)
		for _, item := range r.Inventaire {
			fmt.Printf("%d %s ", item.Quantite, item.Nom)
		}
		fmt.Println("║")
	}
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Quitter Menu")
	var action int
	fmt.Scan(&action)
	if action == 1 {
		fmt.Println("Choisissez le personnage (1-3): ", equipe[0].Nom, ";", equipe[1].Nom, ";", equipe[2].Nom)
		var choice int
		fmt.Scan(&choice)
		if choice >= 1 && choice <= 3 {
			takePot(&equipe[choice-1])
		} else {
			fmt.Println("Mauvais Choix")
		}
	}
}

func takePot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Vie" && char.Inventaire[i].Quantite > 0 {
			if char.Stats.PVActuels < char.Stats.PVMax {
				char.Inventaire[i].Quantite--
				char.Stats.PVActuels += 50
				if char.Stats.PVActuels > char.Stats.PVMax {
					char.Stats.PVActuels = char.Stats.PVMax
				}
			} else if char.Stats.PVActuels == char.Stats.PVMax {
				fmt.Println(char.Nom, "a maximum points de vie")
			}
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de Vie")
}

func main() {
	C1 = initCharacter("Yanisse", "Golem", 1, 200, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	C2 = initCharacter("Léo", "Sage", 1, 100, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	C3 = initCharacter("Luka", "Assassin", 1, 100, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	equipe = [3]Character{C1, C2, C3}
	Menu()
}
