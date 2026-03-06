package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.Russian)
	p.Printf("You have %d unread messages.\n", 5)
	// Output: Vous avez 5 messages non lus.
}
