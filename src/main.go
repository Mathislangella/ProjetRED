struct Character {
	Nom string,
	Classe string,
	LVL int,
	PVMax int,
	PVActuels int,
	Inventaire []string,
}

func initCharacter(nom string, classe string) Character {
		Nom := nom
		Classe := classe
		LVL := 1
		PVMax := 100
		PVActuels := PVMax
		Inventaire := []string
}
	