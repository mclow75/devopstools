package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("version: [ %s ] commit hash LONG:[ %s ] SHORT[ %s ]\n", Version(), LongCommitHash(), ShortCommitHash())
}
