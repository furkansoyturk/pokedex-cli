package main

import (
	"bufio"
	"fmt"
	"github.com/furkansoyturk/pokedex-cli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	nextLocURL *string
	prevLocURL *string
	pokeClient pokeapi.Client
}

func StartRepl(ccc *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Type your command : ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(ccc)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not available")

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
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Show next 20 Pokeman locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 Pokeman locations",
			callback:    commandMapB,
		},
	}
}
