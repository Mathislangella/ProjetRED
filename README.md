# ProjetRED

Mini jeu en ligne de commande (Go) avec un personnage, inventaire, marchand, forgeron et combat de gobelin.

## Structure

```
src/
	main.go               # Entrée du programme
	module/
		types.go            # Types de base (Character, Item, Monster, ...)
		character.go        # Création/affichage personnage, vie
		inventory.go        # Inventaire + menu d'inventaire
		inventory_extra.go  # Consommables, équipement
		utils.go            # Utilitaires (Clear)
		commerce.go         # Marchand et Forgeron (achat/fabrication)
		menu.go             # Menu principal (déplacé depuis main)
		combat.go           # Gobelin et pattern de combat
```

## Prérequis

- Go 1.22+
- Module initialisé au repo root (`go.mod` avec `module ProjetRED`)

## Lancer le jeu

```
go run ./src
```

## Notes

- Le marchand permet d'acheter des objets s'il y a assez d'or et de place.
- Le forgeron fabrique des équipements basiques pour 5 pièces d'or chacun.
- L'inventaire peut être agrandi via un parchemin (si possédé).
- Pour nettoyer l'écran, la fonction `Clear()` est utilisée.
