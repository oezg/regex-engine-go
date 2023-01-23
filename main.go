package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const separator = "|"

var (
	regex              []Rex
	test               string
	mustBegin, mustEnd bool
)

type Rex struct {
	char                 string
	must, repeat, escape bool
}

func main() {
	getInput()
	fmt.Println(matchInput())
}

func getInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if strings.Count(input, separator) != 1 {
		log.Fatal("Input must contain one '|' as separator")
	}
	splitResult := strings.SplitN(input, separator, 2)
	makeRegex(splitResult[0])
	test = splitResult[1]
}

func makeRegex(pattern string) {
	regex = make([]Rex, len(pattern))
	for i := range pattern {
		character := string(pattern[i])
		if regex[i].escape {
			regex[i].set(character)
		} else {
			switch character {
			case "^":
				if i == 0 {
					mustBegin = true
				} else {
					regex[i].set(character)
				}
			case "$":
				if i == len(pattern)-1 {
					mustEnd = true
				} else {
					regex[i].set(character)
				}
			case "?":
				if i == 0 || regex[i-1].char == "" {
					invalid(character)
				}
				regex[i-1].must = false
			case "+":
				if i == 0 || regex[i-1].char == "" {
					invalid(character)
				}
				regex[i-1].repeat = true
			case "*":
				if i == 0 || regex[i-1].char == "" {
					invalid(character)
				}
				regex[i-1].must = false
				regex[i-1].repeat = true
			case "\\":
				if i == len(pattern)-1 {
					invalid(character)
				}
				regex[i+1].escape = true
			default:
				regex[i].set(character)
			}
		}
	}
	k := 0
	for _, rex := range regex {
		if rex.char == "" {
			continue
		}
		regex[k] = rex
		k++
	}
	regex = regex[:k]
}

func (rex *Rex) set(character string) {
	rex.char = character
	rex.must = true
}

func invalid(metacharacter string) {
	log.Fatalf("Regular expression contains an invalid '%s'", metacharacter)
}

func matchInput() bool {
	switch {
	case mustBegin && mustEnd:
		return match(0, 0, false) && match(len(regex)-1, len(test)-1, true)
	case mustBegin:
		return match(0, 0, false)
	case mustEnd:
		return match(len(regex)-1, len(test)-1, true)
	default:
		return matchFlex(0)
	}
}

func match(irx, itx int, back bool) bool {
	if mustBegin && mustEnd && (irx < 0 || irx >= len(regex)) {
		return itx < 0 || itx >= len(test)
	} else if irx < 0 || irx >= len(regex) {
		return true
	} else if restIsNotMust(irx, back) {
		return true
	} else if itx < 0 || itx >= len(test) {
		return false
	}
	rex := regex[irx]
	change := 1
	if back {
		change = -1
	}
	if !rex.must {
		if match(irx+change, itx, back) {
			return true
		}
	}
	if matchChar(rex.char, string(test[itx])) {
		if rex.repeat {
			if match(irx, itx+change, back) {
				return true
			}
		}
		if match(irx+change, itx+change, back) {
			return true
		}
	}
	return false
}

func matchFlex(itx int) bool {
	if itx >= len(test) {
		return false
	}
	if match(0, itx, false) {
		return true
	}
	return matchFlex(itx + 1)
}

func matchChar(patternChar, testChar string) bool {
	return patternChar == "." || patternChar == testChar
}

func restIsNotMust(irx int, back bool) bool {
	if back {
		return allNotMust(regex[:irx+1])
	}
	return allNotMust(regex[irx:])
}

func allNotMust(regexEnd []Rex) bool {
	for _, rex := range regexEnd {
		if rex.must {
			return false
		}
	}
	return true
}
