package module

import (
	"fmt"
	"time"
)

// Exported version of TakePot
func TakePot(char *Character) {
	for i := 0; i < len(char.Inventaire); i++ {
		if char.Inventaire[i].Nom == "Potion de Vie" && char.Inventaire[i].Quantite > 0 {
			if char.Ressources.PVActuels < char.Ressources.PVMax {
				char.Inventaire[i].Quantite--
				AddPV(char, 50)
				fmt.Println(char.Nom, "gagne 50 points de vie ! (PV:", char.Ressources.PVActuels, "/", char.Ressources.PVMax, ")")
			} else if char.Ressources.PVActuels == char.Ressources.PVMax {
				fmt.Println(char.Nom, "est déja en très bonne santé ! (PV:", char.Ressources.PVMax, "/", char.Ressources.PVMax, ")")
				fmt.Println("Vous ne pouvez pas utiliser cette potion.")
			}
			if char.Inventaire[i].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
			}
			return
		}
	}
	fmt.Println(char.Nom, "n'a pas de Potion de Vie.")
}

// Exported version of PoisonPot
func PoisonPot(char *Character) {
	for i := 0; i < len(char.Inventaire); i++ {
		if char.Inventaire[i].Nom == "Potion de poison" && char.Inventaire[i].Quantite > 0 {
			char.Inventaire[i].Quantite--
			for range 3 {
				char.Ressources.PVActuels -= 10
				fmt.Printf("%s subit du poison ! PV restants : %d/%d.\n", char.Nom, char.Ressources.PVActuels, char.Ressources.PVMax)
				time.Sleep(1 * time.Second)
			}
			if char.Ressources.PVActuels <= 0 {
				char.Ressources.PVActuels = 0
				fmt.Println(char.Nom, "est KO à cause du poison !")
			}
			if char.Inventaire[i].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
			}
			IsDead(char)
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de poison")
}

// Exported version of removeParchemin
func RemoveParchemin(char *Character, nom string, quantite int) {
	for i := 0; i < len(char.Inventaire); i++ {
		if char.Inventaire[i].Nom == nom {
			char.Inventaire[i].Quantite -= quantite
			if char.Inventaire[i].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
				return
			}
			break
		}
	}
}

// Exported version of UpgradeInventorySlot
func UpgradeInventorySlot(char *Character) {
	found := false
	for _, item := range char.Inventaire {
		if item.Nom == "Parchemin d'amélioration d'inventaire" && item.Quantite > 0 {
			found = true
			break
		}
	}
	if found {
		RemoveParchemin(char, "Parchemin d'amélioration d'inventaire", 1)
		char.InventoryMax += 5
		fmt.Printf("Inventaire agrandi à %d emplacements.\n", char.InventoryMax)
	} else {
		fmt.Println("Vous n'avez pas assez de parchemins pour agrandir votre inventaire.")
		fmt.Println("Vous pouvez en acheter chez le marchand.")
	}
}

// Exported version of EquiperObjet
func EquiperObjet(char *Character, objet Item) {
	var oldEquip string
	var oldBonus int

	switch objet.Nom {
	case "Chapeau de l'aventurier":
		objet.slot = "Tete"
		objet.bonus = 10
	case "Tunique de l'aventurier":
		objet.slot = "Torse"
		objet.bonus = 25
	case "Bottes de l'aventurier":
		objet.slot = "Pieds"
		objet.bonus = 15
	default:
		fmt.Println("Vous pouvez pas equiper cet objet")
		return
	}

	switch objet.slot {
	case "Tete":
		oldEquip = char.Equipement.Tete
		if oldEquip == "Chapeau de l'aventurier" {
			oldBonus = 10
		}
		char.Equipement.Tete = objet.Nom
	case "Torse":
		oldEquip = char.Equipement.Torse
		if oldEquip == "Tunique de l'aventurier" {
			oldBonus = 25
		}
		char.Equipement.Torse = objet.Nom
	case "Pieds":
		oldEquip = char.Equipement.Pieds
		if oldEquip == "Bottes de l'aventurier" {
			oldBonus = 15
		}
		char.Equipement.Pieds = objet.Nom
	}
	char.Ressources.PVMax -= oldBonus

	for i, item := range char.Inventaire {
		if item.Nom == objet.Nom && item.Quantite > 0 {
			char.Inventaire[i].Quantite--
			if char.Inventaire[i].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
			}
			return
		}
	}

	if oldEquip != "" {
		found := false
		for i, item := range char.Inventaire {
			if item.Nom == oldEquip {
				char.Inventaire[i].Quantite++
				found = true
				break
			}
		}
		if !found {
			char.Inventaire = append(char.Inventaire, Item{Nom: oldEquip, Quantite: 1})
		}
	}
	char.Ressources.PVMax += objet.bonus
	fmt.Printf("%s equipe ! PV Max : %d\n", objet.Nom, char.Ressources.PVMax)
}
