package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		total += getNumber(scanner.Text())
	}

	fmt.Println(total)
}

func getNumber(line string) int {

	numsInWord := mapOfNums()
	m := make(map[int]int)
	for key := range numsInWord {
		i := strings.Index(line, key)

		if i != -1 {
			m[i] = numsInWord[key]
		}
	}

	// to track the lowest and highest index
	lowest := math.MaxInt
	highest := math.MinInt

	if len(m) == 0 {
		return 0
	}

	for key := range m {
		if key > highest {
			highest = key
		}

		if key < lowest {
			lowest = key
		}
	}

	return m[lowest]*10 + m[highest]
}

func mapOfNums() map[string]int {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9

	return m
}
