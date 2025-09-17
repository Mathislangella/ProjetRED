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
