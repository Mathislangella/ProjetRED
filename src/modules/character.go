package modules

type Character struct {
	Nom          string
	Classe       string
	Niveau       Level
	Stats        statistiques
	Nb_vie       int
	Sorts        []Skill
	Inventaire   []Item
	InventoryMax int
	Argent       int
}
