type Character struct{
	Nom string,
	Classe string,
	LVL int,
	PVMax int,
	PVActuels int,
	Inventaire [][]string,
}

func initCharacter(nom string, classe string) Character {
	return var Character = Character{nom, Classe, 1, 100, 100, [[3,"Potion de Vie"]]}
}
	