package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		total += getValidGameId(scanner.Text())
	}

	fmt.Println(total)
}

func getValidGameId(line string) int {
	split := strings.Split(line, ": ")

	// Get GameId from splitted line
	gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])

	if isValidGame(split[1]) {
		return gameId
	}

	return 0
}

// 12 red cubes, 13 green cubes and 14 blue cubes
func isValidGame(gameStr string) bool {
	game := strings.Split(gameStr, "; ")

	for _, sets := range game {
		if !isSetPossible(sets) {
			return false
		}
	}
	return true
}

func isSetPossible(sets string) bool {
	mapOfValidSets := map[string]int{"red": 12, "blue": 14, "green": 13}
	for _, set := range strings.Split(sets, ", ") {
		pick := strings.Split(set, " ")

		// Check if the number of balls is within the valid range
		count, _ := strconv.Atoi(pick[0])
		ball := pick[1]

		if count > mapOfValidSets[ball] {
			return false
		}
	}
	return true
}
