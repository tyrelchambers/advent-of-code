package days

import (
	"fmt"
	"os"
	"strings"
)

type Choice struct {
}

func Day2() {
	file, err := os.ReadFile("./day2.txt")
	if err != nil {
		panic(err)
	}

	fileContents := string(file)
	strings := strings.Split(fileContents, "\n")

	var score int

	// A = Rock, B = Paper, C = Scissors, X = Win, Y = Draw, Z = Lose
	// Score: Rock = 1, Paper = 2, Scissors = 3
	// Win = 6, Draw = 3, Lose = 0

	for _, v := range strings {
		if v == "A X" {
			score += 8
		} else if v == "A Y" {
			score += 4
		} else if v == "A Z" {
			score += 3
		} else if v == "B X" {
			score += 9
		} else if v == "B Y" {
			score += 5
		} else if v == "B Z" {
			score += 1
		} else if v == "C X" {
			score += 7
		} else if v == "C Y" {
			score += 6
		} else if v == "C Z" {
			score += 2
		}
	}

	fmt.Println(score)
}
