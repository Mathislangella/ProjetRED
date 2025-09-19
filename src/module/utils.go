package module

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func addXP(c *Character, xp int) {
	c.Niveau.Exp += xp
	fmt.Printf("%s gagne %d XP !\n", c.Nom, xp)
	c.Niveau.Exp_Cap = int(math.Pow(float64(c.Niveau.Lvl)*10, 1.5))
	if c.Niveau.Exp >= c.Niveau.Exp_Cap {
		c.Niveau.Lvl++
		c.Niveau.Exp -= c.Niveau.Exp_Cap
		c.Stats.Force += 2
		c.Stats.Intelligence += 2
		c.Stats.Agilite += 2
		c.Stats.Chance += 1
		c.Stats.Initiative += 1
		c.Ressources.PVMax += 10
		c.Ressources.ManaMax += 5
		c.Ressources.PVActuels = c.Ressources.PVMax
		c.Ressources.ManaActuels = c.Ressources.ManaMax
		fmt.Printf("%s passe au niveau %d !\n", c.Nom, c.Niveau)
	}
}

const (
	reset    = "\033[0m"
	Noir     = "\033[30m"
	Rouge    = "\033[31m"
	Vert     = "\033[32m"
	Jaune    = "\033[33m"
	Bleu     = "\033[34m"
	Magenta  = "\033[35m"
	Cyan     = "\033[36m"
	Gris     = "\033[37m"
	Blanc    = "\033[97m"
	Gras     = "\033[1m"
	Italique = "\033[3m"
	Souligné = "\033[4m"
	Inverser = "\033[7m"
)
