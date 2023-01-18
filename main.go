package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const separator = "|"

func main() {
	pattern, test := getInput()
	fmt.Println(match(pattern, test))
}

func getInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if strings.Count(input, separator) != 1 {
		log.Fatal("Input must contain one '|' as separator")
	}
	splitResult := strings.SplitN(input, separator, 2)
	return splitResult[0], splitResult[1]
}

func match(pattern, test string) bool {
	pLen, tLen := len(pattern), len(test)
	if pLen == 0 || (pattern == "^$" && tLen == 0) || pattern == "^" || pattern == "$" {
		return true
	} else if tLen == 0 || pattern == "^$" {
		return false
	} else if strings.HasPrefix(pattern, "^") && strings.HasSuffix(pattern, "$") {
		return matchExact(pattern[1:pLen-1], test)
	} else if strings.HasPrefix(pattern, "^") {
		return matchExact(pattern[1:], test[:pLen-1])
	} else if strings.HasSuffix(pattern, "$") {
		return matchExact(pattern[:pLen-1], test[tLen-pLen+1:])
	} else {
		return matchFlex(pattern, test)
	}
}

func matchFlex(pattern, test string) bool {
	pLen, tLen := len(pattern), len(test)
	return pLen <= tLen && (matchExact(pattern, test[:pLen]) || matchFlex(pattern, test[1:]))
}

func matchExact(pattern, test string) bool {
	pLen, tLen := len(pattern), len(test)
	return pLen == 0 || pLen == tLen && matchOne(pattern[:1], test[:1]) && matchExact(pattern[1:], test[1:])
}

func matchOne(pattern, test string) bool {
	return pattern == "." || pattern == test
}
