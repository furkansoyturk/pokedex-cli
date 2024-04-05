package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Type your command : ")
		scanner.Scan()
		words := scanner.Text()

		command, exists := getCommands()[words]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not available")

		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
