package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Tout les struct
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
	Prix     int
}

type Skill struct {
	Nom    string
	Degats int
	Mana   int
}

type Character struct {
	Nom        string
	Classe     string
	Niveau     Level
	Stats      statistiques
	Sorts      []Skill
	Inventaire []Item
}

// toute les fonction
func initCharacter(nom string, classe string, LvL Level, PVMax int, sorts []Skill, inventaire []Item) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     LvL,
		Stats:      statistiques{PVActuels: PVMax, PVMax: PVMax},
		Sorts:      sorts,
		Inventaire: inventaire,
	}
}

func Menu(char *Character) {
	fmt.Println("Menu Du Jeu")
	fmt.Println("1. Ouvrir Info Personnages")
	fmt.Println("2. Ouvrir Inventaire")
	fmt.Println("0. Quitter Menu")
	fmt.Print("Votre Choix : ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		displayInfo(char)
	case 2:
		accessInventory(char)
	case 0:
		return
	default:
		fmt.Print("Mauvais Choix")
	}
}

func displayInfo(char *Character) {
	fmt.Println("╔══════════════════╗")
	fmt.Printf("║ Nom    : %-8s║\n", char.Nom)
	fmt.Printf("║ Classe : %-8s║ \n", char.Classe)
	fmt.Printf("║ Niveau : %-8d║\n", char.Niveau.Lvl)
	fmt.Printf("║ Exp    : %-8d║\n", char.Niveau.Exp)
	fmt.Printf("║ PV     : %d/%-4d║\n", char.Stats.PVActuels, char.Stats.PVMax)
	fmt.Println("╚══════════════════╝")
}

func accessInventory(char *Character) {
	fmt.Println("		  	  INVENTAIRES")
	fmt.Printf("║%-8s : ", char.Nom)
	for _, item := range char.Inventaire {
		fmt.Printf("%d %s ", item.Quantite, item.Nom)
	}
	fmt.Println("║")
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Aller voir le Marchand")
	fmt.Println("0. Quitter Menu")
	var action int
	fmt.Scan(&action)
	switch action {
	case 1:
		takePot(char)
	case 2:
		Marchand(char)
	case 0:
		return
	default:
		fmt.Print("Mauvais Choix")
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

func addPV(char *Character, nb int) {
	char.Stats.PVActuels += nb
}

func lessPV(char *Character, nb int) {
	char.Stats.PVActuels -= nb
}
func isDead(char *Character) {
}

// ╔══╦══╗
// ║  ║  ║
// ╠══╬══╣
// ║  ║  ║
// ╚══╩══╝
func Marchand(char *Character) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Marchand : Bienvenu dans ma boutique.")
	fmt.Println("")
	fmt.Println("╔══════╦════╦════════╦══════════════════════════════════════")
	fmt.Printf("║index ║Prix║Quantite║Object\n")
	fmt.Println("╠══════╬════╬════════╬══════════════════════════════════════")
	var marchandInventaire = []Item{
		{Nom: "Potion de Vie", Quantite: 1, Prix: 0},
	}
	for i, item := range marchandInventaire {
		fmt.Printf("║%d.    ║%d   ║%d       ║%s\n", i+1, item.Prix, item.Quantite, item.Nom)
	}
	fmt.Println("0. Quitter le Marchand")
	var choix int
	fmt.Print("Marchand : Que voulais vous acheter ? ")
	fmt.Scan(&choix)
}

// Fonction Main
func main() {
	// Initialisation des personnages
	C1 := initCharacter("Yanisse", "Golem", Level{1, 0}, 200, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	var Perso *Character = &C1
	// lancement du jeu
	Menu(Perso)
}
