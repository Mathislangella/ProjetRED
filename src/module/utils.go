package module

import (
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func zebi(c *Character) {
	fmt.Printf("%s: Zebi !!!!", c.Nom)
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
