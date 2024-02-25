package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a helpMessage",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback:    CommandExit,
		},
	}

}

func StartCli() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("To get available commands enter \"help\"")
	fmt.Println("")
	
	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback

			if err != nil {
				command.callback()
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
