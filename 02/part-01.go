package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("reading input file", err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var validPasswords []string

	for reader.Scan() {
		line := reader.Text()

		r, pw := Split(line, ":")

		rule, err := parseRule(r)

		if err != nil {
			log.Fatal("parsing rule", err)
		}

		count := strings.Count(pw, string(rule.char))

		if count >= rule.min && count <= rule.max {
			validPasswords = append(validPasswords, pw)
		}
	}

	log.Println("valid passwords", len(validPasswords))
}

type Rule struct {
	min  int
	max  int
	char rune
}

func parseRule(rule string) (Rule, error) {

	var err error
	var r Rule

	length, letter := Split(rule, " ")

	minStr, maxStr := Split(length, "-")

	min, err := strconv.Atoi(minStr)

	max, err := strconv.Atoi(maxStr)

	if err != nil {
		return r, errors.New("could not parse max value")
	}

	r.max = max
	r.min = min
	r.char = rune(letter[0])

	return r, nil
}

func Split(s, sep string) (string, string) {
	// Empty string should just return empty
	if len(s) == 0 {
		return s, s
	}

	slice := strings.SplitN(s, sep, 2)

	// Incase no separator was present
	if len(slice) == 1 {
		return slice[0], ""
	}

	return slice[0], slice[1]
}
