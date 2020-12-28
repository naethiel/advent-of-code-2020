package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var inputFilePath string
	// flags declaration using flag package
	flag.StringVar(&inputFilePath, "file", "./input.txt", "Specify input file path. Default is ./input.txt")

	flag.Parse() // after declaring flags we need to call it

	data, err := loadFile(inputFilePath)

	if err != nil {
		log.Fatal("loading input file", err)
	}

	treeCount := getTreesInPath(1, 1, &data)
	treeCount *= getTreesInPath(1, 3, &data)
	treeCount *= getTreesInPath(1, 5, &data)
	treeCount *= getTreesInPath(1, 7, &data)
	treeCount *= getTreesInPath(2, 1, &data)

	log.Println("trees", treeCount)
}

func getTreesInPath(vslope, hslope int, topography *[]string) int {
	var h, v = 0, 0
	var treeCount = 0

	for v < len(*topography) {
		if (*topography)[v][h%len((*topography)[v])] == '#' {
			treeCount++
		}

		h += hslope
		v += vslope
	}

	return treeCount
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
