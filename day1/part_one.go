package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	Solve()
}

func Solve() int {
	file, err := os.Open("numbers.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safecode = 50
	var solution = 0

	for scanner.Scan() {
		direction := scanner.Text()[:1]
		dialTurns, err := strconv.Atoi(scanner.Text()[1:])

		if err != nil {
			log.Fatal(err)
		}

		//overwrote my part one solution as lazy
		if dialTurns > 100 {
			turns := int(dialTurns / 100)
			solution += turns
		}

		if direction == "L" {
			if safecode == 0 {
				safecode = safecode - dialTurns%100
				safecode = 100 + safecode
			} else {
				safecode = safecode - dialTurns%100
				if safecode < 0 {
					solution = solution + 1
					safecode = 100 + safecode
				}
				if safecode == 0 {
					solution = solution + 1
				}
			}

		} else {
			safecode = safecode + dialTurns%100
			if safecode > 99 {
				solution += 1
				safecode -= 100
			}
		}
	}
	fmt.Println(solution)

	return solution
}
