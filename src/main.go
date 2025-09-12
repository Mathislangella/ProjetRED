package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Lvl        int
	PVMax      int
	PVActuels  int
	Inventaire [][]string
}

// Persos globaux
var C1 Character
var C2 Character
var C3 Character

func initCharacter(nom string, classe string, lvl int, PVMax int, inventaire [][]string) Character {
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
	fmt.Println("")
	fmt.Println("		  	  INVENTAIRES")
	fmt.Println("╔══════════════════╗  ╔══════════════════╗  ╔══════════════════╗")
	fmt.Printf("║     %-8s     ║  ║	      %-8s   ║  ║	    %-8s   ║\n",
		C1.Nom, C2.Nom, C3.Nom)
	fmt.Println("╚══════════════════╝  ╚══════════════════╝  ╚══════════════════╝")
}

func main() {
	C1 = initCharacter("Yanisse", "Golem", 1, 200, [][]string{{"3", "Potion de Vie"}})
	C2 = initCharacter("Léo", "Sage", 1, 100, [][]string{{"3", "Potion de Vie"}})
	C3 = initCharacter("Luka", "Assassin", 1, 100, [][]string{{"3", "Potion de Vie"}})

	displayInfo()
}
