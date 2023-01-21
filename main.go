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
	fmt.Println(match())
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
		if regex[i].escape {
			regex[i].char = string(pattern[i])
			regex[i].must = true
		} else {
			switch string(pattern[i]) {
			case "^":
				if i == 0 {
					mustBegin = true
				} else {
					regex[i].char = string(pattern[i])
					regex[i].must = true
				}
			case "$":
				if i == len(pattern)-1 {
					mustEnd = true
				} else {
					regex[i].char = string(pattern[i])
					regex[i].must = true
				}
			case "?":
				regex[i-1].must = false
			case "+":
				regex[i-1].repeat = true
			case "*":
				regex[i-1].must = false
				regex[i-1].repeat = true
			case "\\":
				regex[i+1].escape = true
			default:
				regex[i].char = string(pattern[i])
				regex[i].must = true
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

func match() bool {
	switch {
	case mustBegin && mustEnd:
		return matchStrict(0, 0)
	case mustBegin:
		return matchBeginning(0, 0)
	case mustEnd:
		return matchEnd(len(regex)-1, len(test)-1)
	default:
		return matchFlex(0)
	}
}

func matchStrict(irx, itx int) bool {
	if irx >= len(regex) && itx >= len(test) {
		return true
	}
	if itx >= len(test) || irx >= len(regex) {
		return false
	}
	rex := regex[irx]
	if !rex.must {
		if matchStrict(irx+1, itx) {
			return true
		}
	}
	if matchChar(rex.char, string(test[itx])) {
		if rex.repeat {
			if matchStrict(irx, itx+1) {
				return true
			}
		}
		if matchStrict(irx+1, itx+1) {
			return true
		}
	}
	return false
}

func matchBeginning(irx, itx int) bool {
	if irx >= len(regex) {
		return true
	}
	if itx >= len(test) {
		return false
	}
	rex := regex[irx]
	if !rex.must {
		if matchBeginning(irx+1, itx) {
			return true
		}
	}
	if matchChar(rex.char, string(test[itx])) {
		if rex.repeat {
			if matchBeginning(irx, itx+1) {
				return true
			}
		}
		if matchBeginning(irx+1, itx+1) {
			return true
		}
	}
	return false
}

func matchEnd(irx, itx int) bool {
	if irx < 0 {
		return true
	}
	if itx < 0 {
		return false
	}
	rex := regex[irx]
	if !rex.must {
		if matchEnd(irx-1, itx) {
			return true
		}
	}
	if matchChar(rex.char, string(test[itx])) {
		if rex.repeat {
			if matchEnd(irx, itx-1) {
				return true
			}
		}
		if matchEnd(irx-1, itx-1) {
			return true
		}
	}
	return false
}

func matchFlex(itx int) bool {
	if itx >= len(test) {
		return false
	}
	if matchBeginning(0, itx) {
		return true
	}
	return matchFlex(itx + 1)
}

func matchChar(patternChar, testChar string) bool {
	return patternChar == "." || patternChar == testChar
}
