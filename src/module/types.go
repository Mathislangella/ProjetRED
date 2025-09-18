package module

// Statistiques du personnage
type Statistiques struct {
	Initiative   int //vitesse d'attaque(qui joue en premier)
	Force        int //degats au corps a corps
	Intelligence int //degats magique
	Agilite      int //chance d'esquive
	Chance       int //chance de coup critique
}

type Ressources struct {
	PVActuels   int //pv actuel
	PVMax       int //pv maximum
	ManaActuels int //mana actuel
	ManaMax     int //mana maximum
	Nb_vie      int //nombre de vie
}

type Level struct {
	Lvl     int //lv actuel
	Exp_Cap int //exp necessaire pour lv up
	Exp     int //exp actuel
}

type Monster struct {
	Nom          string
	statistiques Statistiques
	ressources   Ressources
}

type Item struct {
	Nom      string
	Quantite int
	Prix     int
	slot     string
	bonus    int
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
	Ressources   Ressources
	Sorts        []Skill
	Inventaire   []Item
	InventoryMax int
	Argent       int
	Equipement   Equipment
}
