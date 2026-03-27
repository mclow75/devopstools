package main

import (
	"fmt"
	"os"
)

func saveStringToFile(fileName string) {
	newFile, _ := os.Create(fileName)
	defer newFile.Close()

	bytes, _ := newFile.WriteString("New content for new file")
	fmt.Printf("%d byte(s) saved to file\n", bytes)
}

func readStringFromFile(fileName string) {
	fileContent, _ := os.ReadFile(fileName)
	fmt.Print(string(fileContent))
}

func main() {
	fileName := "string.txt"
	saveStringToFile(fileName)
	readStringFromFile(fileName)
}
