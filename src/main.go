struct Character {
	nom string,
	classe string,
	niveau int,
	pointsDeVieMax int,
	pointsDeVieActuels int,
	inventaire []string,
}

func initCharacter(nom string, classe string) Character {
	