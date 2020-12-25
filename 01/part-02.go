package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input-part-02.txt")

	if err != nil {
		log.Fatal("reading input file", err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var input []int
	for reader.Scan() {
		n, err := strconv.ParseInt(reader.Text(), 10, 64)

		if err != nil {
			log.Fatal("parsing number line on file", err)
		}

		input = append(input, int(n))
	}

	for i1, val1 := range input {
		for _, val2 := range input[i1:] {
			if val1+val2 >= 2020 {
				continue
			}

			for _, val3 := range input {
				if val1+val2+val3 == 2020 {
					fmt.Printf("found: %d + %d + %d = %d \n", val1, val2, val3, val1+val2+val3)
					fmt.Printf("multiplication is %d \n", val1*val2*val3)
					return
				}
			}

		}
	}
}
