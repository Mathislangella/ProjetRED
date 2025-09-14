package main

import "fmt"

type Item struct {
	Nom      string
	Quantite int
}

type Character struct {
	Nom        string
	Classe     string
	Lvl        int
	PVMax      int
	PVActuels  int
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
		Lvl:        lvl,
		PVMax:      PVMax,
		PVActuels:  PVMax,
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
		C1.Lvl, C2.Lvl, C3.Lvl)
	fmt.Printf("║ PV     : %d/%-4d║  ║ PV     : %d/%-4d║  ║ PV     : %d/%-4d║\n",
		C1.PVActuels, C1.PVMax, C2.PVActuels, C2.PVMax, C3.PVActuels, C3.PVMax)
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
}

func main() {
	C1 = initCharacter("Yanisse", "Golem", 1, 200, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	C2 = initCharacter("Léo", "Sage", 1, 100, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	C3 = initCharacter("Luka", "Assassin", 1, 100, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	equipe = [3]Character{C1, C2, C3}
	Menu()
}
