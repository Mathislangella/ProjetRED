package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

type Character struct {
	Nom        string
	Classe     string
	Niveau     Level
	Stats      statistiques
	Inventaire []Item
}

// toute les fonction
func initCharacter(nom string, classe string, LvL Level, PVMax int, inventaire []Item) Character {

	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     LvL,
		Stats:      statistiques{PVActuels: PVMax, PVMax: PVMax},
		Inventaire: inventaire,
	}
}

func Menu(char *Character) {
	fmt.Println("Menu Du Jeu")
	fmt.Println("1. Ouvrir Info Personnages")
	fmt.Println("2. Ouvrir Inventaire")
	fmt.Println("0. Quitter Menu")
	fmt.Print("Votre Choix : ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		displayInfo(char)
		Menu(char)
	case "2":
		accessInventory(char)
		Menu(char)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Menu(char)
	}
}

func displayInfo(char *Character) {
	fmt.Println("╔══════════════════╗")
	fmt.Printf("║ Nom    : %s║\n", char.Nom)
	fmt.Printf("║ Classe : %s║ \n", char.Classe)
	fmt.Printf("║ Niveau : %d║\n", char.Niveau.Lvl)
	fmt.Printf("║ Exp    : %d║\n", char.Niveau.Exp)
	fmt.Printf("║ PV     : %d/%d║\n", char.Stats.PVActuels, char.Stats.PVMax)
	fmt.Println("╚══════════════════╝")
}

func accessInventory(char *Character) {
	fmt.Println("		  	  INVENTAIRES")
	fmt.Printf("║%-8s : ", char.Nom)
	for _, item := range char.Inventaire {
		fmt.Printf("%d %s ,", item.Quantite, item.Nom)
	}
	fmt.Println("║")
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Aller voir le Marchand")
	fmt.Println("0. Quitter Menu")
	var action string
	fmt.Scan(&action)
	switch action {
	case "1":
		takePot(char)
		accessInventory(char)
	case "2":
		Marchand(char)
		accessInventory(char)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		accessInventory(char)
	}
}

func addtoInventory(char *Character, item Item, nb int) {
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
		} else {
			char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb})
		}
	}
}

func takePot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Vie" && char.Inventaire[i].Quantite > 0 {
			if char.Stats.PVActuels < char.Stats.PVMax {
				char.Inventaire[i].Quantite--
				addPV(char, 50)
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
	if char.Stats.PVActuels > char.Stats.PVMax {
		char.Stats.PVActuels = char.Stats.PVMax
	}
}

func lessPV(char *Character, nb int) {
	char.Stats.PVActuels -= nb
}
func isDead(char *Character) {
	return
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
	// inventaire du marchand
	var marchandInventaire = []Item{
		{Nom: "Potion de Vie", Quantite: 1, Prix: 0},
	}
	for i, item := range marchandInventaire {
		fmt.Printf("║%d.    ║%d   ║%d       ║%s\n", i+1, item.Prix, item.Quantite, item.Nom)
	}
	fmt.Println("0. Quitter le Marchand")
	fmt.Print("Marchand : Que voulais vous acheter ? ")
	var choix string
	fmt.Scan(&choix)
	intchoix, err := strconv.Atoi(choix)
	if err == nil {
		fmt.Print(err)
	}
	if intchoix > 0 && intchoix <= len(marchandInventaire) {
		if marchandInventaire[intchoix-1].Quantite > 0 {
			addtoInventory(char, marchandInventaire[intchoix-1], 1)
			marchandInventaire[intchoix-1].Quantite--
		} else {
			fmt.Println("Marchand : Vous ne pouvez pas acheter cet objet car soit il n y en a plus en stock soit votre sac a dos est plein")
			Marchand((char))
		}

	} else if intchoix == 0 {
		fmt.Println("Vous quittez la boutique du marchand")
	} else {
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Marchand(char)
	}
}

// Fonction Main
func main() {
	// Initialisation des personnages
	C1 := initCharacter("Yanisse", "Golem", Level{1, 0}, 200, []Item{{Nom: "Potion de Vie", Quantite: 3}})
	var Perso *Character = &C1

	// lancement du jeu
	Menu(Perso)
}
