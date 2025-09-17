package module

// Statistiques du personnage
type Statistiques struct {
	PVActuels int
	PVMax     int
}

type Level struct {
	Lvl int
	Exp int
}

type Monster struct {
	Nom          string
	PVActuels    int
	PVMax        int
	AttaquePoint int
}

type Item struct {
	Nom      string
	Quantite int
	Prix     int
}

type Skill struct {
	Nom    string
	Degats int
	Mana   int
}

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Nom          string
	Classe       string
	Niveau       Level
	Stats        Statistiques
	Nb_vie       int
	Sorts        []Skill
	Inventaire   []Item
	InventoryMax int
	Argent       int
	Equipement   Equipment
}
