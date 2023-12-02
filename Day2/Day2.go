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
		total += powerOfMinCubes(scanner.Text())
	}

	fmt.Println(total)
}

func powerOfMinCubes(line string) int {
	game := strings.Split(line, ": ")[1]

	minBallsCount := make(map[string]int)

	fillMinBallsCount(game, &minBallsCount)

	return minBallsCount["red"] * minBallsCount["green"] * minBallsCount["blue"]
}

func fillMinBallsCount(game string, minBallsCount *map[string]int) {

	for _, set := range strings.Split(game, "; ") {
		for _, pick := range strings.Split(set, ", ") {
			ballCount := strings.Split(pick, " ")
			num, _ := strconv.Atoi(ballCount[0])
			ball := ballCount[1]

			// Store the maximum number of balls got at any picks
			if (*minBallsCount)[ball] < num {
				(*minBallsCount)[ball] = num
			}
		}
	}
}
