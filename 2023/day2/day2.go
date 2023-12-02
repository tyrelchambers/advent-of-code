package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Red   int
	Blue  int
	Green int
}

func getGameID(str string) int {
	reg := regexp.MustCompile(`(\d+)`)
	match := reg.FindStringSubmatch(str)

	if len(match) == 0 {
		return -1
	}

	toId, err := strconv.Atoi(match[0])

	if err != nil {
		return -1
	}

	return toId
}

func getCubes(str, cube string) int {
	reg := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, cube))
	match := reg.FindStringSubmatch(str)

	if len(match) == 0 {
		return 0
	}

	toId, err := strconv.Atoi(match[1])

	if err != nil {
		return -1
	}

	return toId
}

func Main() {

	data_file, err := os.ReadFile("./day2/data.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(data_file), "\n")

	total := 0

	for _, str := range data {
		redMax := 0
		blueMax := 0
		greenMax := 0
		sets := strings.Split(str, ";")

		game_id := getGameID(str)
		fmt.Println("Game: ", game_id)

		for _, set := range sets {
			redCubes := getCubes(set, "red")
			blueCubes := getCubes(set, "blue")
			greenCubes := getCubes(set, "green")

			if redCubes > redMax {
				redMax = redCubes
			}

			if blueCubes > blueMax {
				blueMax = blueCubes
			}

			if greenCubes > greenMax {
				greenMax = greenCubes
			}
		}

		total += redMax * blueMax * greenMax

	}

	fmt.Println("Total: ", total)
}
