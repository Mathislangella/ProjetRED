package module

import (
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
