package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
	Nb_vie     int
	Sorts      []Skill
	Inventaire []Item
	Argent     int
}

// toute les fonction

func initCharacter(nom string, classe string, LvL Level, PVMax int, Nb_vie int, sorts []Skill, inventaire []Item) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     LvL,
		Stats:      statistiques{PVActuels: PVMax / 2, PVMax: PVMax},
		Nb_vie:     Nb_vie,
		Sorts:      sorts,
		Inventaire: inventaire,
		Argent:     100,
	}
}

func characterCreation() Character {
	clear()
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
		clear()
		characterCreation()
	}
	clear()
	return initCharacter(nom, classe, Level{1, 0}, pv, 1, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
}

func Menu(char *Character, marchand *[]Item) {
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
		Menu(char, marchand)
	case "2":
		clear()
		accessInventory(char, marchand)
		clear()
		Menu(char, marchand)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		clear()
		Menu(char, marchand)
	}
}

func displayInfo(char *Character) {
	clear()
	fmt.Println("╔══════════════════╗")
	fmt.Printf("║ Nom    : %-8s║\n", char.Nom)
	fmt.Printf("║ Classe : %-8s║ \n", char.Classe)
	fmt.Printf("║ Niveau : %-8d║\n", char.Niveau.Lvl)
	fmt.Printf("║ Exp    : %-8d║\n", char.Niveau.Exp)
	fmt.Printf("║ PV     : %d/%-5d║\n", char.Stats.PVActuels, char.Stats.PVMax)
	fmt.Printf("║ OR     : %-8d║\n", char.Argent)
	fmt.Println("╚══════════════════╝")
}

func accessInventory(char *Character, marchand *[]Item) {
	fmt.Println("		  	  INVENTAIRES")
	fmt.Printf("Vous avez %d Pieces d'or\n", char.Argent)
	for i, item := range char.Inventaire {
		fmt.Printf("║%-4d : ", i)
		fmt.Printf("%d %-10s ,", item.Quantite, item.Nom)
		fmt.Println("║")
	}
	fmt.Println("")
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Utiliser Potion de Poison")
	fmt.Println("3. Aller voir le Marchand")
	fmt.Println("0. Quitter Menu")
	var action string
	fmt.Scan(&action)
	switch action {
	case "1":
		takePot(char)
		accessInventory(char, marchand)
	case "2":
		poisonPot(char)
		accessInventory(char, marchand)
	case "3":
		Marchand(char, marchand)
		clear()
		accessInventory(char, marchand)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		accessInventory(char, marchand)
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

func poisonPot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Poison" && char.Inventaire[i].Quantite > 0 {
			char.Inventaire[i].Quantite--
			for range 3 {
				lessPV(char, 10)
				fmt.Printf("%s subit du poison ! PV restants : %d/%d\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
				time.Sleep(1 * time.Second)
			}
			if char.Stats.PVActuels <= 0 {
				char.Stats.PVActuels = 0
				fmt.Println(char.Nom, "est KO à cause du poison !")
			}
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de Poison")
}

func isInventoryFull(char *Character) bool {
	return len(char.Inventaire) >= 10
}

func InventoryFull(char *Character, newItem Item) bool {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Voulez-vous remplacer un objet ? (o/n)")
		var rep string
		fmt.Scan(&rep)
		if strings.ToLower(rep) == "o" {
			fmt.Println("Voici vos objets :")
			for i, item := range char.Inventaire {
				fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
			}
			fmt.Print("Selectionnez le numéro de l'objet à remplacer : ")
			var choix string
			fmt.Scan(&choix)
			idx, err := strconv.Atoi(choix)
			if err != nil || idx < 1 || idx > len(char.Inventaire) {
				fmt.Println("Choix invalide.")
				return true
			}
			char.Inventaire[idx-1] = newItem
			fmt.Printf("L'objet a été remplacé par %s.\n", newItem.Nom)
		}
		return true
	}
	return false
}

func addtoInventory(char *Character, item Item, nb int) {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Impossible d'ajouter un nouvel objet.")
		return
	}
	etat := false
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
			etat = true
		}
	}
	if !etat {
		char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb})
	}
}

func addPV(char *Character, nb int) {
	char.Stats.PVActuels += nb
	if char.Stats.PVActuels > char.Stats.PVMax {
		char.Stats.PVActuels = char.Stats.PVMax
	}
}

func lessPV(char *Character, nb int) {
	char.Stats.PVActuels -= nb
	isDead(char)
}

func isDead(char *Character) {
	if char.Stats.PVActuels == 0 {
		if char.Nb_vie > 0 {
			fmt.Println("1. Ressurection")
			fmt.Println("2. Quitter Le Jeu")
			var choice string
			fmt.Scan(&choice)
			switch choice {
			case "1":
				addPV(char, char.Stats.PVMax/2)
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

// ╔══╦══╗
// ║  ║  ║
// ╠══╬══╣
// ║  ║  ║
// ╚══╩══╝
func Marchand(char *Character, shop *[]Item) {
	clear()
	fmt.Println("Marchand : Bienvenu dans ma boutique.")
	fmt.Println("")
	fmt.Println("╔══════╦════╦════════╦══════════════════════════════════════╗")
	fmt.Printf("║index ║Prix║Quantite║Object                                ║\n")
	fmt.Println("╠══════╬════╬════════╬══════════════════════════════════════╣")
	for i, item := range *shop {
		fmt.Printf("║%d.    ║%-4d║%-8d║%-38s║\n", i+1, item.Prix, item.Quantite, item.Nom)
	}
	fmt.Println("╚══════╩════╩════════╩══════════════════════════════════════╝")
	fmt.Println("0. Quitter le Marchand")
	fmt.Print("Marchand : Que voulais vous acheter ? ")
	var choix string
	fmt.Scan(&choix)
	intchoix, err := strconv.Atoi(choix)
	if err == nil {
		fmt.Print(err)
	}
	if intchoix > 0 && intchoix <= len(*shop) {
		if (*shop)[intchoix-1].Quantite > 0 {
			addtoInventory(char, (*shop)[intchoix-1], 1)
			(*shop)[intchoix-1].Quantite -= 1
			fmt.Println("Marchand : Merci de votre achat")
			Marchand(char, shop)
		} else {
			fmt.Println("Marchand : Vous ne pouvez pas acheter cet objet car soit il n y en a plus en stock soit votre sac a dos est plein")
			Marchand(char, shop)
		}
	} else if intchoix == 0 {
		fmt.Println("Vous quittez la boutique du marchand")
	} else {
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Marchand(char, shop)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func upgradeInventorySlot(char *Character) {

// Fonction Main
func main() {
	// Initialisation des personnages
	C1 := characterCreation()
	var Perso *Character = &C1
	var marchandInventory = []Item{
		{Nom: "Potion de Vie", Quantite: 1, Prix: 3},
		{Nom: "Potion de poison", Quantite: 10, Prix: 6},
		{Nom: "Livre de Sort", Quantite: 1, Prix: 25},
		{Nom: "Fourrure de Loup", Quantite: 1, Prix: 4},
		{Nom: "Peau de Troll", Quantite: 1, Prix: 7},
		{Nom: "Cuir de Sanglier", Quantite: 1, Prix: 3},
		{Nom: "Plume de Corbeau", Quantite: 1, Prix: 1},
	}
	var marchandstuf *[]Item = &marchandInventory
	// lancement du jeu
	Menu(Perso, marchandstuf)
}
