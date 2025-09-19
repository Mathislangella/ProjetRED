package module

import (
	"fmt"
	"strings"
)

func spellKeep(char *Character) {
	for i := 0; i < len(char.Inventaire); i++ {
		if strings.HasPrefix(char.Inventaire[i].Nom, "Livre de Sort") {
			learnSpell(char, char.Inventaire[i])
			char.Inventaire[i].Quantite--
			if char.Inventaire[i].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
			}
			break
		}
	}
}

func learnSpell(char *Character, item Item) {
	switch item.Nom {
	case "Livre de Sort: Boule de Feu":
		newSpell := Skill{
			Nom:    "Boule de Feu",
			Mana:   20,
			Degats: 30,
		}
		alreadyKnown := false
		for _, s := range char.Sorts {
			if s.Nom == newSpell.Nom {
				alreadyKnown = true
				break
			}
		}
		if !alreadyKnown {
			char.Sorts = append(char.Sorts, newSpell)
			fmt.Println("Vous avez appris le sort Boule de Feu !")
		} else {
			fmt.Println("Vous connaissez déjà ce sort.")
		}
	default:
		fmt.Println("Ce livre de sort est inconnu.")
	}
}
