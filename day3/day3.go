package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	Solve()
}

func Solve() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var partOne int
	var partTwo int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bank := scanner.Text()
		var ints []int
		for _, char := range bank {
			digit := int(char - '0')
			ints = append(ints, digit)

		}
		partOne += MaxNSize(2, ints, 0)
		partTwo += MaxNSize(12, ints, 0)
	}
	fmt.Printf("Joltage for part one: %d \n", partOne)
	fmt.Printf("Joltage for part two: %d \n", partTwo)

}

func MaxNSize(n int, pBank []int, joltage int) int {

	pow := int(math.Pow(10, float64(n-1)))

	if n == 1 {
		slices.Sort(pBank)
		joltage += pBank[len(pBank)-1]

	} else {

		cutLastN := slices.Clone(pBank[:len(pBank)-n+1])
		slices.Sort(cutLastN)

		largestDigit := cutLastN[len(cutLastN)-1]

		indexOfLargest := slices.Index(pBank, largestDigit)
		remainder := pBank[indexOfLargest+1:]

		joltage += largestDigit * pow
		return MaxNSize(n-1, remainder, joltage)

	}

	return joltage

}
