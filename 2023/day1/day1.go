package day1

import (
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Digit struct {
	index int
	value string
}

func find_digit_indices(hash_map map[string]string, key string, str string) []Digit {
	var indices []Digit
	pattern := key

	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringIndex(str, -1)

	if len(matches) == 0 {
		return nil
	}

	for _, match := range matches {
		indices = append(indices, Digit{index: match[0], value: hash_map[key]})
	}

	fmt.Println(indices)

	return indices
}

func concat_num(a, b Digit) string {
	return a.value + b.value
}

func get_last_digit(digits []Digit) Digit {
	if len(digits) == 1 {
		return digits[0]
	}

	return digits[len(digits)-1]
}

func Day1() {
	num_match := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
	}

	data_file, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(data_file), "\n")

	var values []string

	for _, v := range data {
		var digits []Digit
		var first Digit
		var last Digit

		for key := range num_match {
			matches := find_digit_indices(num_match, key, v)

			if len(matches) != 0 {
				digits = append(digits, matches...)
			}

		}

		slices.SortFunc(digits, func(a, b Digit) int {
			return cmp.Compare(a.index, b.index)
		})

		first = digits[0]
		last = get_last_digit(digits)

		values = append(values, concat_num(first, last))
	}

	final := 0

	for _, v := range values {
		f, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println(err)
		}

		final += f
	}
	fmt.Println(final)
}
