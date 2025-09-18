package main

import (
	"ProjetRED/src/module"
	"fmt"
)

func main() {
	C1 := module.CharacterCreation()
	var Perso *module.Character = &C1
	fmt.Printf("Nom: %s\n", Perso.Nom)
	var marchandInventory = []module.Item{
		{Nom: "Potion de Vie", Quantite: 10, Prix: 6},
		{Nom: "Potion de poison", Quantite: 10, Prix: 9},
		{Nom: "Livre de Sort: Boule de Feu", Quantite: 1, Prix: 25},
		{Nom: "Fourrure de Loup", Quantite: 10, Prix: 7},
		{Nom: "Peau de Troll", Quantite: 10, Prix: 10},
		{Nom: "Cuir de Sanglier", Quantite: 10, Prix: 5},
		{Nom: "Plume de Corbeau", Quantite: 10, Prix: 3},
		{Nom: "Parchemin d'amélioration d'inventaire", Quantite: 2, Prix: 40},
	}
	var marchandstuf *[]module.Item = &marchandInventory
	module.Menu(Perso, marchandstuf)

}
