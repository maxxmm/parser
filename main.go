package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() 

	var fileText []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        fileText = append(fileText, strings.Split(scanner.Text(), "")...)
		//add blank after a line for readability
		fileText = append(fileText, " ")
    }

	for _, char := range(fileText) {
		fmt.Println(char)
	}
}