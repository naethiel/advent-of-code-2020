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
	data, err := loadFile("input.txt")

	if err != nil {
		log.Fatal("loading data from input file", err)
	}

	var validPasswords []string

	for _, line := range data {
		l, err := parseLine(line)

		if err != nil {
			log.Fatal("parsing line", line, err)
		}

		var isAtMin bool = false
		var isAtMax bool = false

		// offsetting by -1 because min and max values in the rule are 1-based
		// instead of 0-based
		if len(l.password) >= l.min-1 {
			isAtMin = rune(l.password[l.min-1]) == l.char
		}

		if len(l.password) >= l.max-1 {
			isAtMax = rune(l.password[l.max-1]) == l.char
		}

		if (isAtMin && !isAtMax) || (!isAtMin && isAtMax) {
			validPasswords = append(validPasswords, l.password)
		}
	}

	log.Println("valid password count", len(validPasswords))
}

func loadFile(path string) ([]string, error) {

	// open file, as usual
	file, err := os.Open(path)

	if err != nil {
		log.Print("reading input file", "path", path, "error", err)
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var data []string
	for reader.Scan() {
		data = append(data, reader.Text())
	}

	return data, nil
}

type Rule struct {
	min      int
	max      int
	char     rune
	password string
}

func parseLine(line string) (Rule, error) {

	var err error
	var r Rule

	rule, password := split(line, ":")

	length, letter := split(rule, " ")

	minStr, maxStr := split(length, "-")

	min, err := strconv.Atoi(minStr)

	max, err := strconv.Atoi(maxStr)

	if err != nil {
		return r, errors.New("could not parse max value")
	}

	r.max = max
	r.min = min
	r.char = rune(letter[0])
	r.password = strings.TrimSpace(password)

	return r, nil
}

func split(s, sep string) (string, string) {
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
