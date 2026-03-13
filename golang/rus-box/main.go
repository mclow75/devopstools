package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.Russian)
	p.Printf("You have %d archiced messages.\n", 112345)
}
