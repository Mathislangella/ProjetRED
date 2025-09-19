package module

import (
	"fmt"
	"strconv"
	"strings"
)

func IsInventoryFull(char *Character) bool {
	temp := 0
	for _, r := range char.Inventaire {
		temp += r.Quantite
	}
	return temp >= char.InventoryMax
}

func AddToInventory(char *Character, item Item, nb int) bool {
	if IsInventoryFull(char) {
		return InventoryFull(char, item)
	}
	etat := false
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
			etat = true
		}
	}
	if !etat {
		char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb, Prix: item.Prix})
	}
	return true
}

func InventoryFull(char *Character, newItem Item) bool {
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
		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(char.Inventaire) {
			fmt.Println("Choix invalide.")
			return false
		}
		char.Inventaire[index-1].Quantite -= 1
		if char.Inventaire[index-1].Quantite <= 0 {
			char.Inventaire = append(char.Inventaire[:index-1], char.Inventaire[index:]...)
		}
		AddToInventory(char, newItem, 1)
		fmt.Printf("L'objet a été remplacé par %s.\n", newItem.Nom)
		return true
	}
	return false
}

func AccessInventory(char *Character, marchand *[]Item) {
	fmt.Println("\t\t  \t  INVENTAIRES")
	fmt.Printf("Vous avez %d Pieces d'or\n", char.Argent)
	fmt.Printf("╔═════╦════════╦══════════════════════════════════════╗\n")
	fmt.Printf("║index║Quantite║Object                                ║\n")
	fmt.Printf("╠═════╬════════╬══════════════════════════════════════╣\n")
	for i, item := range char.Inventaire {
		fmt.Printf("║%-5d║%-8d║%-38s║\n", i, item.Quantite, item.Nom)
	}
	fmt.Printf("╚═════╩════════╩══════════════════════════════════════╝\n")
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Utiliser Potion de poison")
	fmt.Println("3. Utiliser Parchemin d'amélioration d'inventaire")
	fmt.Println("4. Equiper un equipement")
	fmt.Println("9. Qui sont-ils ?")
	fmt.Println("0. Quitter Menu")
	var action string
	fmt.Scan(&action)
	switch action {
	case "1":
		TakePot(char)
		AccessInventory(char, marchand)
	case "2":
		PoisonPot(char)
		AccessInventory(char, marchand)
	case "3":
		UpgradeInventorySlot(char)
		AccessInventory(char, marchand)
	case "4":
		equippable_Names := []string{"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier"}
		var equippable []int
		fmt.Println("Equipments que vous pouvez equiper :")
		for i, item := range char.Inventaire {
			for _, eq := range equippable_Names {
				if item.Nom == eq && item.Quantite > 0 {
					fmt.Printf("%d. %s (x%d)\n", len(equippable)+1, item.Nom, item.Quantite)
					equippable = append(equippable, i)
				}
			}
		}
		if len(equippable) == 0 {
			fmt.Println("Vous avez Aucune equipment dans l'inventaire")
			AccessInventory(char, marchand)
			return
		}
		fmt.Print("Choisissez le numero d'equipment : ")
		var choix int
		fmt.Scan(&choix)
		if choix < 1 || choix > len(equippable) {
			fmt.Println("Mauvais choix")
			AccessInventory(char, marchand)
			return
		}
		index := equippable[choix-1]
		EquiperObjet(char, char.Inventaire[index])
		AccessInventory(char, marchand)
	case "9":
		fmt.Println("Abba & La Chose")
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		AccessInventory(char, marchand)
	}
}
