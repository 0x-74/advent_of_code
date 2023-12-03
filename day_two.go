package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	power, total := 0, 0
	f, err := os.Open("day_two.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	games := make(map[string][]map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		key := strings.TrimSpace(line[0])

		colorBallsPairs := strings.Split(line[1], ";")
		var colorBallsMaps []map[string]int

		for _, pair := range colorBallsPairs {
			values := strings.Split(pair, ",")
			colorBallsMap := make(map[string]int)

			for _, value := range values {
				parts := strings.Split(strings.TrimSpace(value), " ")
				color := parts[1]
				numBalls, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatal(err)
				}

				colorBallsMap[color] = numBalls
			}

			colorBallsMaps = append(colorBallsMaps, colorBallsMap)
		}

		games[key] = colorBallsMaps
	}

	for gameno, subgames := range games {
		isvalid := true
		for _, colors := range subgames {
			if colors["red"] > 12 || colors["green"] > 13 || colors["blue"] > 14 {
				isvalid = false
				break
			}
		}
		if isvalid {
			gameNumber, err := strconv.Atoi(strings.Split(gameno, " ")[1])
			if err != nil {
				log.Fatal(err)
			}
			total += gameNumber
		}
	}
	fmt.Println("total : ", total)

	// Part 2
	for _, subgames := range games {
		r, g, b := 0, 0, 0 
		for _, colors := range subgames {
			if colors["red"] > r {
				r = colors["red"]
			}
			if colors["green"] > g {
				g = colors["green"]
			}
			if colors["blue"] > b {
				b = colors["blue"]
			}
		}
		power += r * g * b
	}
	fmt.Println("total : ", power)
}
