package src

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"main"
)



func takePot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Vie" && char.Inventaire[i].Quantite > 0 {
			if char.Stats.PVActuels < char.Stats.PVMax {
				char.Inventaire[i].Quantite--
				addPV(char, 50)
				fmt.Println(char.Nom, "gagne 50 points de vie ! (PV:", char.Stats.PVActuels, "/", char.Stats.PVMax, ")")
			} else if char.Stats.PVActuels == char.Stats.PVMax {
				fmt.Println(char.Nom, "est déja en très bonne santé ! (PV:", char.Stats.PVMax, "/", char.Stats.PVMax, ")")
				fmt.Println("Vous ne pouvez pas utiliser cette potion.")
			}
			return
		}
	}
	fmt.Println(char.Nom, "n'a pas de Potion de Vie.")
}

func poisonPot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de poison" && char.Inventaire[i].Quantite > 0 {
			char.Inventaire[i].Quantite--
			for range 3 {
				char.Stats.PVActuels -= 10
				fmt.Printf("%s subit du poison ! PV restants : %d/%d.\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
				time.Sleep(1 * time.Second)
			}
			if char.Stats.PVActuels <= 0 {
				char.Stats.PVActuels = 0
				fmt.Println(char.Nom, "est KO à cause du poison !")
			}
			isDead(char)
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de poison")
}
// augmente la limite d'inventaire si le joueur a un parchemin
func upgradeInventorySlot(char *Character) {
	found := false
	for _, item := range char.Inventaire {
		if item.Nom == "Parchemin d'amélioration d'inventaire" && item.Quantite > 0 {
			found = true
			break
		}
	}
	if found {
		removeParchemin(char, "Parchemin d'amélioration d'inventaire", 1)
		char.InventoryMax += 5
		fmt.Printf("Inventaire agrandi à %d emplacements.\n", char.InventoryMax)
	} else {
		fmt.Println("Vous n'avez pas assez de parchemins pour agrandir votre inventaire.")
		fmt.Println("Vous pouvez en acheter chez le marchand.")
	}
}