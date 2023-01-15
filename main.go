package main

import (
	"fmt"
	"log"
	"strings"
)

const separator = "|"

var (
	input      string
	regex      string
	testString string
	result     bool
)

func main() {
	getInput()
	splitInput()
	match()
	showResult()
}

func getInput() {
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
}

func splitInput() {
	splitResult := strings.SplitN(input, separator, 2)
	regex = splitResult[0]
	testString = splitResult[1]
}

func match() {
	switch {
	case regex == "":
		result = true
	case testString == "":
		result = false
	case regex == testString:
		result = true
	case regex == ".":
		result = true
	}
}

func showResult() {
	fmt.Println(result)
}