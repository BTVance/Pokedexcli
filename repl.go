package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	format_text := strings.Fields(text)
	return format_text
}
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	initCommands()
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		cleanCommand := cleanInput(command)
		if len(cleanCommand) == 0 {
			continue
		}
		commandName := cleanCommand[0]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
