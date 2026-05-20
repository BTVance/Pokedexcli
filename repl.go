package main

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	format_text := strings.Fields(text)
	return format_text
}
