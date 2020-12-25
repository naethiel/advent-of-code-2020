package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input-part-01.txt")

	if err != nil {
		log.Fatal("reading input file", err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var input []int64

	for reader.Scan() {
		number, err := strconv.ParseInt(reader.Text(), 10, 64)
		if err != nil {
			log.Fatal("converting line to int", err)
		}

		input = append(input, number)
	}

	for i, n1 := range input {
		for _, n2 := range input[i:] {
			if n1+n2 == 2020 {
				fmt.Printf("found, n1 is: %d, n2 is: %d, multiplication is %d \n", n1, n2, n1*n2)
				log.Print("found result, exiting")
				return
			}
		}
	}
}
