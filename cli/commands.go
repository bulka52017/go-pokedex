package cli

import (
	"fmt"
	"os"
)

func CommandHelp() error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func CommandExit() error {
	os.Exit(0)
	return nil
}
