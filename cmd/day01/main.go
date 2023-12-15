package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func parseInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 { break }

		lines = append(lines, line)
	}
	
	return lines
}

type SafeSum struct {
	mu  sync.Mutex
	sum int
}

func (ss *SafeSum) add(addend int) {
	defer ss.mu.Unlock()
	
	ss.mu.Lock()
	ss.sum += addend
}

func getAndSumDigits(wg *sync.WaitGroup, ss *SafeSum, rp *strings.Replacer, line string) {
	line = rp.Replace(rp.Replace(line))

	re := regexp.MustCompile("[0-9]")
	digits := re.FindAllString(line, -1)

	number, _ := strconv.Atoi(digits[0] + digits[len(digits) - 1])
	ss.add(number)

	wg.Done()
}

func part1(input []string) int {
	var wg sync.WaitGroup
	rp := strings.NewReplacer()

	var ss SafeSum
	for _, line := range input {
		wg.Add(1)
		go getAndSumDigits(&wg, &ss, rp, line)
	}

	wg.Wait()

	return ss.sum
}

func part2(input []string) int {
	var wg sync.WaitGroup
	rp := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)

	var ss SafeSum
	for _, line := range input {
		wg.Add(1)
		go getAndSumDigits(&wg, &ss, rp, line)
	}

	wg.Wait()

	return ss.sum
}

func main() {
	input := parseInput()
	
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
