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

func getInput() (string, string) {
	var input string
	fmt.Scanln(&input)
	splitResult := strings.SplitN(input, separator, 2)
	return splitResult[0], splitResult[1]
}

func match(pattern, test string) bool {
	patternLength := len(pattern)
	if len(test) < patternLength {
		return false
	}
	if matchEqual(pattern, test[:patternLength]) {
		return true
	}
	return match(pattern, test[1:])
}

func matchEqual(regex, test string) bool {
	if regex == "" {
		return true
	}
	if test == "" {
		return false
	}
	if matchOne(regex[:1], test[:1]) {
		return matchEqual(regex[1:], test[1:])
	}
	return false
}

func matchOne(regex, test string) bool {
	return regex == "." || regex == test
}
