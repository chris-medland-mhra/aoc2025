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
	Solve()
}

func Solve() int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ids := strings.SplitSeq(scanner.Text(), ",")
		var result int
		var resultPartTwo int

		//lordamercy gl understanding because I cant
		for idRange := range ids {
			splitIdRange := strings.Split(idRange, "-")
			start, _ := strconv.Atoi(splitIdRange[0])
			end, _ := strconv.Atoi(splitIdRange[1])
			for i := start; i <= end; i++ {
				targ := strconv.Itoa(i)
				if len(targ)%2 != 0 {
					continue
				}
				mid := len(targ) / 2
				firstHalf := targ[:mid]
				secondHalf := targ[mid:]
				if firstHalf == secondHalf {
					result = result + i
				}
			}

			for i := start; i <= end; i++ {
				targ := strconv.Itoa(i)
				mid := len(targ) / 2
				for j := 1; j <= mid; j++ {
					subStrings := SplitStrByN(targ, j)
					if len(subStrings[0]) != len(subStrings[len(subStrings)-1]) {
						continue
					}
					comparator := subStrings[0]
					var count int
					for _, val := range subStrings[1:] {
						if comparator == val {
							count = count + 1
						}
					}
					if count == len(subStrings)-1 {
						fmt.Println(i)
						resultPartTwo = resultPartTwo + i
						break
					}
				}
			}
		}

		fmt.Println(result)
		fmt.Println(resultPartTwo)
	}
	return 1
}

func SplitStrByN(s string, n int) []string {
	var strs []string
	for i := 0; i < len(s); i += n {
		endpoint := i + n
		if endpoint > len(s) {
			endpoint = len(s)
		}
		strs = append(strs, s[i:endpoint])
	}
	return strs
}
