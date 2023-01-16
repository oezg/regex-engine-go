package main

import (
	"fmt"
	"strings"
)

const separator = "|"

func main() {
	pattern, testString := getInput()
	fmt.Println(match(pattern, testString))
}

func match(regex, test string) bool {
	if regex == "" {
		return true
	}
	if test == "" {
		return false
	}
	if matchOne(regex[:1], test[:1]) {
		return match(regex[1:], test[1:])
	}
	return false
}

func getInput() (string, string) {
	var input string
	fmt.Scanln(&input)
	splitResult := strings.SplitN(input, separator, 2)
	return splitResult[0], splitResult[1]
}

func matchOne(regex, test string) bool {
	return regex == "." || regex == test
}
