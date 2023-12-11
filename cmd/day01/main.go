package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func getAndSumDigits(wg *sync.WaitGroup, ss *SafeSum, re *regexp.Regexp, line string) {
	digits := re.FindAllString(line, -1)
	
	digit1, _ := strconv.Atoi(digits[0])
	digit2, _ := strconv.Atoi(digits[len(digits) - 1])
	
	number := digit1 * 10 + digit2
	ss.add(number)

	wg.Done()
}

func part1(input []string) int {
	var wg sync.WaitGroup
	re := regexp.MustCompile("[0-9]")

	var ss SafeSum
	for _, line := range input {
		wg.Add(1)
		go getAndSumDigits(&wg, &ss, re, line)
	}

	wg.Wait()

	return ss.sum
}

func main() {
	input := parseInput()
	
	fmt.Println(part1(input))
}