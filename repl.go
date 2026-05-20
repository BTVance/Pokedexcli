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
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) > 0 {
			fmt.Printf("Your command was: %s\n", cleanText[0])
		}
	}
}
