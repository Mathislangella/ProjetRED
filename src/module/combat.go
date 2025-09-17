package module

import "fmt"

func InitGobelin() Monster {
	return Monster{
		Nom:          "Gobelin d'entrainement",
		PVMax:        40,
		PVActuels:    40,
		AttaquePoint: 5,
	}
}

func GobelinPattern(char *Character, _ *Monster) {
	monstre := InitGobelin()
	tour := 1
	fmt.Printf("Un %s apparaît !\n", monstre.Nom)
	for char.Stats.PVActuels > 0 && monstre.PVActuels > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		var degats int
		if tour%3 == 0 {
			degats = monstre.AttaquePoint * 2
		} else {
			degats = monstre.AttaquePoint
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", monstre.Nom, char.Nom, degats)
		char.Stats.PVActuels -= degats
		if char.Stats.PVActuels < 0 {
			char.Stats.PVActuels = 0
		}
		fmt.Printf("%s : %d/%d PV\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
		fmt.Printf("%s : %d/%d PV\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)
		tour++
		if monstre.PVActuels <= 0 {
			fmt.Printf("%s a été vaincu !\n", monstre.Nom)
			char.Argent += 10
			return
		}
		if char.Stats.PVActuels <= 0 {
			char.Stats.PVActuels = 0
			fmt.Printf("%s est KO !\n", char.Nom)
			IsDead(char)
			return
		}
	}
}
