package days

import (
	"os"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.ReadFile("./day1.txt")
	if err != nil {
		panic(err)
	}

	fileContents := string(file)
	strings := strings.Split(fileContents, "\n")
	formattedStrings := make([]int, len(strings))

	for i := 0; i < len(strings); i++ {
		val, _ := strconv.Atoi(strings[i])
		formattedStrings[i] = val
	}

	addedList := []int{}

	var val = 0
	for _, v := range formattedStrings {

		if v == 0 {
			addedList = append(addedList, val)
			val = 0
		} else {
			val += v
		}
	}

	for i := 0; i < len(addedList); i++ {
		for j := 0; j < len(addedList)-1-i; j++ {
			if addedList[j] > addedList[j+1] {
				addedList[j], addedList[j+1] = addedList[j+1], addedList[j]
			}
		}
	}

}
