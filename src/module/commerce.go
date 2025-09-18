package module

import (
	"fmt"
	"strconv"
)

// Marchand regroupe la logique du marchand
func Marchand(char *Character, shop *[]Item) {
	Clear()
	fmt.Printf("Marchand : Bienvenu dans ma boutique.\n\n")
	fmt.Printf("╔══════╦════╦════════╦══════════════════════════════════════╗\n")
	fmt.Printf("║index ║" + Jaune + "Prix" + reset + "║Quantite║Object                                ║\n")
	fmt.Printf("╠══════╬════╬════════╬══════════════════════════════════════╣\n")
	for i, item := range *shop {
		fmt.Printf("║%-6d║"+Jaune+"%-4d"+reset+"║%-8d║%-38s║\n", i+1, item.Prix, item.Quantite, item.Nom)
	}
	fmt.Printf("╚══════╩════╩════════╩══════════════════════════════════════╝\n")
	fmt.Printf("Il vous reste"+Jaune+" %-4d pièce d'or"+reset+"\n", char.Argent)
	fmt.Println("0. Quitter le Marchand")
	fmt.Print("Marchand : Que voulais vous acheter ?   ")
	var choix string
	fmt.Scan(&choix)
	intchoix, err := strconv.Atoi(choix)
	if err != nil {
		fmt.Print(err)
	}
	if intchoix > 0 && intchoix <= len(*shop) {
		if char.Argent >= (*shop)[intchoix-1].Prix {
			if (*shop)[intchoix-1].Quantite > 0 {
				if AddToInventory(char, (*shop)[intchoix-1], 1) {
					(*shop)[intchoix-1].Quantite -= 1
					fmt.Println("Vous venez d'acheter ", (*shop)[intchoix-1].Nom)
					fmt.Println("Marchand : Merci de votre achat")
					char.Argent -= (*shop)[intchoix-1].Prix
				} else {
					fmt.Println("Marchand : Votre sac a dos est plein ou vous avez annulé le remplacement.")
				}
			} else {
				fmt.Println("Marchand : Vous ne pouvez pas acheter cet objet car il n y en a plus en stock ")
			}
		} else {
			fmt.Println("Marchand : Vous n'avez pas assez d'or ")
		}
		Marchand(char, shop)
	} else if intchoix == 0 {
		fmt.Println("Vous quittez la boutique du marchand")
	} else {
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Marchand(char, shop)
	}
}

// Forgeron regroupe la logique du forgeron
func ForgeEquipment(char *Character, nom string) {
	if char.Argent < 5 {
		fmt.Println(char.Nom, "N'as pas assez pieces d'or")
		return
	}
	char.Argent -= 5
	AddToInventory(char, Item{Nom: nom, Quantite: 1}, 1)
	fmt.Printf("%s a fabrique %v, il vous reste %d pieces d'or \n", char.Nom, nom, char.Argent)
}

func Forgeron(char *Character) {
	fmt.Println("Bienvenue chez le Forgeron !")
	fmt.Println("Voici la liste des equipements a fabriquer :")
	fmt.Println("1. Chapeau de l'aventurier (5 pieces d'or)")
	fmt.Println("2. Tunique de l'aventurier (5 pieces d'or)")
	fmt.Println("3. Bottes de l'aventurier (5 pieces d'or)")
	fmt.Println("0. Quitter le Forgeron")
	fmt.Print("Votre choix : ")
	var choix string
	fmt.Scan(&choix)
	switch choix {
	case "1":
		material_1 := "Plume de Corbeau"
		material_2 := "Cuir de Sanglier"
		var T1 string
		var T2 string
		x := 0
		for x < len(char.Inventaire) {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 1 {
				T1 = material_1
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			x++
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (o/n) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "o" {
				ForgeEquipment(char, "Chapeau de l'aventurier")
			} else if choice == "n" {
				Forgeron(char)
				AddToInventory(char, Item{Nom: material_1, Quantite: 1}, 1)
				AddToInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "2":
		material_2 := "Peau de Troll"
		material_1 := "Fourrure de Loup"
		var T1 string
		var T2 string
		x := 0
		for x < len(char.Inventaire) {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 2 {
				T1 = material_1
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			x++
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (o/n) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "o" {
				ForgeEquipment(char, "Tunique de l'aventurier")
			} else if choice == "n" {
				Forgeron(char)
				AddToInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
				AddToInventory(char, Item{Nom: material_1, Quantite: 2}, 1)
			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "3":
		material_1 := "Cuir de Sanglier"
		material_2 := "Fourrure de Loup"
		var T1 string
		var T2 string
		x := 0
		for x < len(char.Inventaire) {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 1 {
				T1 = material_1
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
				if char.Inventaire[x].Quantite <= 0 {
					char.Inventaire = append(char.Inventaire[:x], char.Inventaire[x+1:]...)
					continue
				}
			}
			x++
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (o/n) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "o" {
				ForgeEquipment(char, "Bottes de l'aventurier")
			} else if choice == "n" {
				Forgeron(char)
				AddToInventory(char, Item{Nom: material_1, Quantite: 1}, 1)
				AddToInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "0":
		return
	default:
		Forgeron(char)
	}
}
