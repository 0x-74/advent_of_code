package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("day_one.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := convertToNum(scanner.Text())

		var first, second string

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				second = string(line[i])
				break
			}
		}

		if first != "" && second != "" {
			numStr := first + second
			numInt, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			total += numInt
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)

	
}
func convertToNum(str string) string {
	ogstr := str
	pattern := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	regexPattern := strings.Join(pattern, "|")
	regex := regexp.MustCompile(regexPattern)

	replacement := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	str = regex.ReplaceAllStringFunc(str, func(match string) string {
		return replacement[match]
	})

	if ogstr != str {
		return convertToNum(str)
	}

	return str
}

