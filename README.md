# Regular Expression Engine

## How to use
Enter your regular expression and the string you want to match in one line separated by a pipe character `|`

## About
Regular expressions are a fundamental part of computer science and natural language processing. 
In this project, I wrote an extendable regular expression engine that can handle 
- basic regex syntax, 
- including literals (a, b, c, etc.), 
- wild-cards (`.`), and 
- metacharacters (`?`, `*`, `+`, `^`, `$`).

## Learning Outcomes
- Learn about the syntax of regular expressions, 
- practice working with the string type, parsing, and slicing, and 
- get more familiar with boolean algebra and recursion.

## Stages
1. Implement a program that compares two single character strings (including the wildcard) and determines if there's a match.
2. Extend your engine to compare two equal length strings using recursion.
3. Add the ability to compare a regex to strings that vary in length.
4. Extend the engine to handle the operators `^` and `$` that control the position of the regex within a string.
5. Support the additional operators `?`, `*`, and `+` that control the repetition of a character within a string.
6. Finally, implement the backslash `\` as an escape symbol that allows to use metacharacters as literals.

