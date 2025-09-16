package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"strconv"
	"strings"
	"time"
)

// Tout les struct
type statistiques struct {
	PVActuels int
	PVMax     int
}

type Level struct {
	Lvl int
	Exp int
}
type monster struct {
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

// toute les fonction


func initCharacter(nom string, classe string, LvL Level, PVMax int, Nb_vie int, sorts []Skill, inventaire []Item) Character {
	return Character{
		Nom:          nom,
		Classe:       classe,
		Niveau:       LvL,
		Stats:        statistiques{PVActuels: PVMax / 2, PVMax: PVMax},
		Nb_vie:       Nb_vie,
		Sorts:        sorts,
		Inventaire:   inventaire,
		InventoryMax: 10,
		Argent:       100,
	}
}

func characterCreation() Character {
	clear()
	fmt.Println("Quel est votre nom :")
	var nom string
	fmt.Scan(&nom)
	nom = strings.ToLower(nom)
	nom0 := strings.Split(nom, "")
	nom0[0] = strings.ToUpper((string(nom[0])))
	nom = strings.Join(nom0, "")
	fmt.Println("Quel est votre Classe choisissez |1. Humain|2. Elf|3. Nain")
	var choi string
	fmt.Scan(&choi)
	var classe string
	var pv int
	switch choi {
	case "1":
		classe = "Humain"
		pv = 100
	case "2":
		classe = "Elf"
		pv = 80
	case "3":
		classe = "Nain"
		pv = 120
	default:
		fmt.Printf("Pour choisir votre choisisez 1 ,2 ou 3")
		clear()
		characterCreation()
	}
	clear()
	return initCharacter(nom, classe, Level{1, 0}, pv, 1, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
}

func Menu(char *Character, marchand *[]Item) {
func characterCreation() Character {
	clear()
	fmt.Println("Quel est votre nom :")
	var nom string
	fmt.Scan(&nom)
	nom = strings.ToLower(nom)
	nom0 := strings.Split(nom, "")
	nom0[0] = strings.ToUpper((string(nom[0])))
	nom = strings.Join(nom0, "")
	fmt.Println("Quel est votre Classe choisissez |1. Humain|2. Elf|3. Nain")
	var choi string
	fmt.Scan(&choi)
	var classe string
	var pv int
	switch choi {
	case "1":
		classe = "Humain"
		pv = 100
	case "2":
		classe = "Elf"
		pv = 80
	case "3":
		classe = "Nain"
		pv = 120
	default:
		fmt.Printf("Pour choisir votre choisisez 1 ,2 ou 3")
		clear()
		characterCreation()
	}
	clear()
	return initCharacter(nom, classe, Level{1, 0}, pv, 1, []Skill{{Nom: "Coup de Poign", Degats: 5, Mana: 0}}, []Item{{Nom: "Potion de Vie", Quantite: 3}})
}

func Menu(char *Character, marchand *[]Item) {
	fmt.Println("Menu Du Jeu")
	fmt.Println("1. Ouvrir Info Personnages")
	fmt.Println("2. Ouvrir Inventaire")
	fmt.Println("3. Forgeron")
	fmt.Println("0. Quitter Menu")
	fmt.Print("Votre Choix : ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		displayInfo(char)
		Menu(char, marchand)
	case "2":
		clear()
		accessInventory(char, marchand)
		clear()
		Menu(char, marchand)
	case "3":
		clear()
		Forgeron(char)
		Menu(char, marchand)
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		clear()
		Menu(char, marchand)
	}
}

func displayInfo(char *Character) {
	clear()
	fmt.Println("╔══════════════════╗")
	fmt.Printf("║ Nom    : %-8s║\n", char.Nom)
	fmt.Printf("║ Classe : %-8s║ \n", char.Classe)
	fmt.Printf("║ Niveau : %-8d║\n", char.Niveau.Lvl)
	fmt.Printf("║ Exp    : %-8d║\n", char.Niveau.Exp)
	fmt.Printf("║ PV     :  %d/%-4d║\n", char.Stats.PVActuels, char.Stats.PVMax)
	fmt.Printf("║ OR     : %-8d║\n", char.Argent)
	fmt.Println("╚══════════════════╝")
}

func accessInventory(char *Character, marchand *[]Item) {
	fmt.Println("		  	  INVENTAIRES")
	fmt.Printf("Vous avez %d Pieces d'or\n", char.Argent)
	for i, item := range char.Inventaire {
		fmt.Printf("║%-4d : ", i)
		fmt.Printf("%d %-10s ,", item.Quantite, item.Nom)
		fmt.Println("║")
	}
	fmt.Println("")
	fmt.Println("1. Utiliser Potion de Vie")
	fmt.Println("2. Utiliser Potion de poison")
	fmt.Println("3. Utiliser Parchemin d'amélioration d'inventaire")
	fmt.Println("4. Aller voir le Marchand")
	fmt.Println("9. Qui sont-ils ?")
	fmt.Println("0. Quitter Menu")
	var action string
	fmt.Scan(&action)
	switch action {
	case "1":
		takePot(char)
		accessInventory(char, marchand)
	case "2":
		poisonPot(char)
		accessInventory(char, marchand)
	case "3":
		upgradeInventorySlot(char)
		accessInventory(char, marchand)
	case "4":
		Marchand(char, marchand)
		clear()
		accessInventory(char, marchand)
	case "9":
		fmt.Println("Abba & La Chose")
	case "0":
		return
	default:
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		accessInventory(char, marchand)
	}
}

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

func InventoryFull(char *Character, newItem Item) bool {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Voulez-vous remplacer un objet ? (o/n)")
		var rep string
		fmt.Scan(&rep)
		if strings.ToLower(rep) == "o" {
			fmt.Println("Voici vos objets :")
			for i, item := range char.Inventaire {
				fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
			}
			fmt.Print("Selectionnez le numéro de l'objet à remplacer : ")
			var choix string
			fmt.Scan(&choix)
			idx, err := strconv.Atoi(choix)
			if err != nil || idx < 1 || idx > len(char.Inventaire) {
				fmt.Println("Choix invalide.")
				return false
			}
			char.Inventaire[idx-1].Quantite -= 1 // On enlève un de l'item remplacé
			if char.Inventaire[idx-1].Quantite <= 0 {
				char.Inventaire = append(char.Inventaire[:idx-1], char.Inventaire[idx:]...)
			}
			addtoInventory(char, newItem, 1)
			fmt.Printf("L'objet a été remplacé par %s.\n", newItem.Nom)
			return true
		}
		return false
	}
	return false
}

func isInventoryFull(char *Character) bool {
	temp := 0
	for _, r := range char.Inventaire {
		temp += r.Quantite
	}
	if temp < char.InventoryMax {
		return false
	}
	return true
}

func addtoInventory(char *Character, item Item, nb int) bool {
	if isInventoryFull(char) {
		// Si l'utilisateur remplace un objet, on retourne true, sinon false
		return InventoryFull(char, item)
	}
	etat := false
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
			etat = true
		}
	}
	if !etat {
		char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb, Prix: item.Prix})
	}
	return true
}

func poisonPot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Poison" && char.Inventaire[i].Quantite > 0 {
			char.Inventaire[i].Quantite--
			for range 3 {
				lessPV(char, 10)
				fmt.Printf("%s subit du poison ! PV restants : %d/%d\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
				time.Sleep(1 * time.Second)
			}
			if char.Stats.PVActuels <= 0 {
				char.Stats.PVActuels = 0
				fmt.Println(char.Nom, "est KO à cause du poison !")
			}
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de Poison")
}

func isInventoryFull(char *Character) bool {
	return len(char.Inventaire) >= 10
}

func InventoryFull(char *Character, newItem Item) bool {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Voulez-vous remplacer un objet ? (o/n)")
		var rep string
		fmt.Scan(&rep)
		if strings.ToLower(rep) == "o" {
			fmt.Println("Voici vos objets :")
			for i, item := range char.Inventaire {
				fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
			}
			fmt.Print("Selectionnez le numéro de l'objet à remplacer : ")
			var choix string
			fmt.Scan(&choix)
			idx, err := strconv.Atoi(choix)
			if err != nil || idx < 1 || idx > len(char.Inventaire) {
				fmt.Println("Choix invalide.")
				return true
			}
			char.Inventaire[idx-1] = newItem
			fmt.Printf("L'objet a été remplacé par %s.\n", newItem.Nom)
		}
		return true
	}
	return false
}

func addtoInventory(char *Character, item Item, nb int) {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Impossible d'ajouter un nouvel objet.")
		return
	}
	etat := false
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
			etat = true
		}
	}
	if !etat {
		char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb})
	}

}

func poisonPot(char *Character) {
	for i := range char.Inventaire {
		if char.Inventaire[i].Nom == "Potion de Poison" && char.Inventaire[i].Quantite > 0 {
			char.Inventaire[i].Quantite--
			for range 3 {
				lessPV(char, 10)
				fmt.Printf("%s subit du poison ! PV restants : %d/%d\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
				time.Sleep(1 * time.Second)
			}
			if char.Stats.PVActuels <= 0 {
				char.Stats.PVActuels = 0
				fmt.Println(char.Nom, "est KO à cause du poison !")
			}
			return
		}
	}
	fmt.Println(char.Nom, "N'as pas de Potion de Poison")
}

func isInventoryFull(char *Character) bool {
	return len(char.Inventaire) >= 10
}

func InventoryFull(char *Character, newItem Item) bool {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Voulez-vous remplacer un objet ? (o/n)")
		var rep string
		fmt.Scan(&rep)
		if strings.ToLower(rep) == "o" {
			fmt.Println("Voici vos objets :")
			for i, item := range char.Inventaire {
				fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
			}
			fmt.Print("Selectionnez le numéro de l'objet à remplacer : ")
			var choix string
			fmt.Scan(&choix)
			idx, err := strconv.Atoi(choix)
			if err != nil || idx < 1 || idx > len(char.Inventaire) {
				fmt.Println("Choix invalide.")
				return true
			}
			char.Inventaire[idx-1] = newItem
			fmt.Printf("L'objet a été remplacé par %s.\n", newItem.Nom)
		}
		return true
	}
	return false
}

func addtoInventory(char *Character, item Item, nb int) {
	if isInventoryFull(char) {
		fmt.Println("Inventaire plein ! Impossible d'ajouter un nouvel objet.")
		return
	}
	etat := false
	for i, r := range char.Inventaire {
		if r.Nom == item.Nom {
			char.Inventaire[i].Quantite += nb
			etat = true
		}
	}
	if !etat {
		char.Inventaire = append(char.Inventaire, Item{Nom: item.Nom, Quantite: nb})
	}

}

func addPV(char *Character, nb int) {
	char.Stats.PVActuels += nb
	if char.Stats.PVActuels > char.Stats.PVMax {
		char.Stats.PVActuels = char.Stats.PVMax
	}
}

func lessPV(char *Character, nb int) {
	char.Stats.PVActuels -= nb
	isDead(char)
}

func isDead(char *Character) {
	if char.Stats.PVActuels == 0 {
		if char.Nb_vie > 0 {
			fmt.Println("1. Ressurection")
			fmt.Println("2. Quitter Le Jeu")
			var choice string
			fmt.Scan(&choice)
			switch choice {
			case "1":
				addPV(char, char.Stats.PVMax/2)
				fmt.Printf("%s gagne %d points de vie", char.Nom, char.Stats.PVMax/2)
				fmt.Printf(": %d/%d", char.Stats.PVActuels, char.Stats.PVMax)
				char.Nb_vie--
				return
			case "2":
				os.Exit(0)
			}
		}
		if char.Nb_vie == 0 {
			fmt.Println(char.Nom, "est mort")
		}
	}
	if char.Stats.PVActuels == 0 {
		if char.Nb_vie > 0 {
			fmt.Println("1. Ressurection")
			fmt.Println("2. Quitter Le Jeu")
			var choice string
			fmt.Scan(&choice)
			switch choice {
			case "1":
				addPV(char, char.Stats.PVMax/2)
				fmt.Printf("%s gagne %d points de vie", char.Nom, char.Stats.PVMax/2)
				fmt.Printf(": %d/%d", char.Stats.PVActuels, char.Stats.PVMax)
				char.Nb_vie--
				return
			case "2":
				os.Exit(0)
			}
		}
		if char.Nb_vie == 0 {
			fmt.Println(char.Nom, "est mort")
		}
	}
}

// ╔══╦══╗
// ║  ║  ║
// ╠══╬══╣
// ║  ║  ║
// ╚══╩══╝
func Marchand(char *Character, shop *[]Item) {
	clear()
	fmt.Println("Marchand : Bienvenu dans ma boutique.")
	fmt.Println("")
	fmt.Println("╔══════╦════╦════════╦══════════════════════════════════════╗")
	fmt.Printf("║index ║Prix║Quantite║Object                                ║\n")
	fmt.Println("╠══════╬════╬════════╬══════════════════════════════════════╣")
	for i, item := range *shop {
		fmt.Printf("║%d.    ║%-4d║%-8d║%-38s║\n", i+1, item.Prix, item.Quantite, item.Nom)
	}
	fmt.Println("╚══════╩════╩════════╩══════════════════════════════════════╝")
	fmt.Printf("Il vous reste %-4d pièce d'or\n", char.Argent)
	fmt.Println("0. Quitter le Marchand")
	fmt.Print("Marchand : Que voulais vous acheter ?   ")
	var choix string
	fmt.Scan(&choix)
	intchoix, err := strconv.Atoi(choix)
	if err != nil {
		fmt.Print(err)
	}
	if intchoix > 0 && intchoix <= len(*shop) {
		if char.Argent >= (*shop)[intchoix-1].Prix {
			if (*shop)[intchoix-1].Quantite > 0 {
				// On ne débite l'or et ne retire du stock seulement si l'ajout a réussi
				if addtoInventory(char, (*shop)[intchoix-1], 1) {
					(*shop)[intchoix-1].Quantite -= 1
					fmt.Println("Vous venez d'acheter ", (*shop)[intchoix-1].Nom)
					fmt.Println("Marchand : Merci de votre achat")
					char.Argent -= (*shop)[intchoix-1].Prix
				} else {
					fmt.Println("Marchand : Votre sac a dos est plein ou vous avez annulé le remplacement.")
				}
			} else {
				fmt.Println("Marchand : Vous ne pouvez pas acheter cet objet car il n y en a plus en stock ")
			}
		} else {
			fmt.Println("Marchand : Vous n'avez pas assez d'or ")
		}
		Marchand(char, shop)
	} else if intchoix == 0 {
		fmt.Println("Vous quittez la boutique du marchand")
	} else {
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Marchand(char, shop)
	}
}

// réinitialise l'affichage de l'écran

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func removeParchemin(char *Character, nom string, quantite int) {
	for i := 0; i < len(char.Inventaire); i++ {
		// Vérifie si l'objet est à l'indice i
		if char.Inventaire[i].Nom == nom {
			// Réduit la quantité de l'objet
			char.Inventaire[i].Quantite -= quantite
			if char.Inventaire[i].Quantite <= 0 {
				// Supprime l'objet de l'inventaire si la quantité est nulle ou négative
				char.Inventaire = append(char.Inventaire[:i], char.Inventaire[i+1:]...)
				return
			}
			// Arret de la boucle après avoir trouvé et modifié l'objet
			break
		}
	}
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

func initGobelin() monster {
	return monster{
		Nom:          "Gobelin d'entrainement",
		PVMax:        40,
		PVActuels:    40,
		AttaquePoint: 5,
	}
}

func gobelinPattern(char *Character, _ *monster) {
	// Création d'un gobelin d'entraînement avec ses stats de base
	monstre := initGobelin()
	tour := 1
	fmt.Printf("Un %s apparaît !\n", monstre.Nom)
	// Boucle principale du combat : continue tant que les deux sont vivants
	for char.Stats.PVActuels > 0 && monstre.PVActuels > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		var degats int

		// Tous les 3 tours, le gobelin se voit doubler ses dégâts ( soit 200% de dégâts )
		if tour%3 == 0 {
			degats = monstre.AttaquePoint * 2
		} else {
			degats = monstre.AttaquePoint
		}

		// Affichage de l'attaque et application des dégâts au joueur
		fmt.Printf("%s inflige à %s %d de dégâts\n", monstre.Nom, char.Nom, degats)
		char.Stats.PVActuels -= degats

		// Empêche les PV du joueur de devenir négatifs
		if char.Stats.PVActuels < 0 {
			char.Stats.PVActuels = 0
		}

		// Affichage des PV restants pour le joueur et le gobelin
		fmt.Printf("%s : %d/%d PV\n", char.Nom, char.Stats.PVActuels, char.Stats.PVMax)
		fmt.Printf("%s : %d/%d PV\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)

		tour++ // Passage au tour suivant

		// vérification de la fin du combat
		// Si le gobelin est vaincu, le joueur gagne 10 pièces d'or
		// Si le joueur est KO, il perd une vie
		if monstre.PVActuels <= 0 {
			fmt.Printf("%s a été vaincu !\n", monstre.Nom)
			char.Argent += 10
			return
		}
		if char.Stats.PVActuels <= 0 {
			char.Stats.PVActuels = 0
			fmt.Printf("%s est KO !\n", char.Nom)
			isDead(char)
			return
		}
	}
	intchoix, err := strconv.Atoi(choix)
	if err == nil {
		fmt.Print(err)
	}
	if intchoix > 0 && intchoix <= len(*shop) {
		if (*shop)[intchoix-1].Quantite > 0 {
			addtoInventory(char, (*shop)[intchoix-1], 1)
			(*shop)[intchoix-1].Quantite -= 1
			fmt.Println("Marchand : Merci de votre achat")
			Marchand(char, shop)
		} else {
			fmt.Println("Marchand : Vous ne pouvez pas acheter cet objet car soit il n y en a plus en stock soit votre sac a dos est plein")
			Marchand(char, shop)
		}
	} else if intchoix == 0 {
		fmt.Println("Vous quittez la boutique du marchand")
	} else {
		fmt.Println("Mauvais Choix")
		fmt.Println("")
		Marchand(char, shop)
	}
}

func ForgeEquipment(char *Character, nom string) {
	if char.Argent < 5 {
		fmt.Println(char.Nom, "N'as pas assez pieces d'or")
		return
	}
	char.Argent -= 5
	addtoInventory(char, Item{Nom: nom, Quantite: 1}, 1)
	fmt.Printf("%s a fabrique %v, il vous reste %d pieces d'or \n", char.Nom, nom, char.Argent)

}

func Forgeron(char *Character) {
	fmt.Println("Bienvenue chez le Forgeron !")
	fmt.Println("Voici la liste des equipements a fabriquer :")
	fmt.Println("1. Chapeau de l'aventurier (5 pieces d'or)")
	fmt.Println("2. Tunique de l'aventurier (5 pieces d'or)")
	fmt.Println("3. Bottes de l'aventurier (5 pieces d'or)")
	fmt.Println("0. Quitter le Forgeron")
	fmt.Print("Votre choix : ")
	var choix string
	fmt.Scan(&choix)
	switch choix {
	case "1":
		material_1 := "Plume de Corbeau"
		material_2 := "Cuir de Sanglier"
		var T1 string
		var T2 string
		for x := range char.Inventaire {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 1 {
				T1 = material_1
				char.Inventaire[x].Quantite--
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
			}
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (oui/non) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "oui" {
				ForgeEquipment(char, "Chapeau de l'aventurier")
			} else if choice == "non" {
				Forgeron(char)
				addtoInventory(char, Item{Nom: material_1, Quantite: 1}, 1)
				addtoInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "2":
		material_2 := "Peau de Troll"
		material_1 := "Fourrure de Loup"
		var T1 string
		var T2 string
		for x := range char.Inventaire {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 2 {
				T1 = material_1
				char.Inventaire[x].Quantite--
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
			}
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (oui/non) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "oui" {
				ForgeEquipment(char, "Tunique de l'aventurier")
			} else if choice == "non" {
				Forgeron(char)
				addtoInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
				addtoInventory(char, Item{Nom: material_1, Quantite: 2}, 1)

			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "3":
		material_1 := "Cuir de Sanglier"
		material_2 := "Fourrure de Loup"
		var T1 string
		var T2 string
		for x := range char.Inventaire {
			if char.Inventaire[x].Nom == material_1 && char.Inventaire[x].Quantite >= 1 {
				T1 = material_1
				char.Inventaire[x].Quantite--
			}
			if char.Inventaire[x].Nom == material_2 && char.Inventaire[x].Quantite >= 1 {
				T2 = material_2
				char.Inventaire[x].Quantite--
			}
		}
		if T1 == material_1 && T2 == material_2 {
			fmt.Println("Est-ce que vous-voulez fabriquer un Chapeau de l'aventurier ? (oui/non) ")
			var choice string
			fmt.Scan(&choice)
			if choice == "oui" {
				ForgeEquipment(char, "Bottes de l'aventurier")
			} else if choice == "non" {
				Forgeron(char)
				addtoInventory(char, Item{Nom: material_1, Quantite: 1}, 1)
				addtoInventory(char, Item{Nom: material_2, Quantite: 1}, 1)
			}
		} else {
			fmt.Println("Vous avez pas assez de materiaux")
		}
	case "0":
		return
	default:
		Forgeron(char)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fonction Main
func main() {
	// Initialisation des personnages
	C1 := characterCreation()
	C1 := characterCreation()
	var Perso *Character = &C1
	displayInfo(Perso)
	var marchandInventory = []Item{
		{Nom: "Potion de Vie", Quantite: 10, Prix: 6},
		{Nom: "Potion de poison", Quantite: 10, Prix: 9},
		{Nom: "Livre de Sort", Quantite: 1, Prix: 25},
		{Nom: "Fourrure de Loup", Quantite: 10, Prix: 7},
		{Nom: "Peau de Troll", Quantite: 10, Prix: 10},
		{Nom: "Cuir de Sanglier", Quantite: 10, Prix: 5},
		{Nom: "Plume de Corbeau", Quantite: 10, Prix: 3},
		{Nom: "Parchemin d'amélioration d'inventaire", Quantite: 2, Prix: 40},
	}
	var marchandstuf *[]Item = &marchandInventory
	// lancement du jeu
	Menu(Perso, marchandstuf)
	isDead(Perso)
	gobelinPattern(Perso, nil)
}
