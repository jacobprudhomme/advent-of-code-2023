package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type GameRound struct {
	r int
	g int
	b int
}

func parseInput() [][]GameRound {
	scanner := bufio.NewScanner(os.Stdin)

	var games [][]GameRound
	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 { break }
		
		gameStr := strings.Split(line, ": ")[1]
		roundsStr := strings.Split(gameStr, "; ")
		var game []GameRound
		for _, roundStr := range roundsStr {
			cubes := strings.Split(roundStr, ", ")
			
			var round GameRound
			for _, cube := range cubes {
				cubeData := strings.Split(cube, " ")
				count, _ := strconv.Atoi(cubeData[0])
				colour := cubeData[1]

				switch colour {
				case "red":
					round.r = count
				case "green":
					round.g = count
				case "blue":
					round.b = count
				}
			}
			
			game = append(game, round)
		}

		games = append(games, game)
	}

	return games
}

func isRoundPossible(round GameRound) bool {
	return round.r <= 12 && round.g <= 13 && round.b <= 14
}

func isGamePossible(game []GameRound, gameId int) bool {
	var wg sync.WaitGroup
	wg.Add(len(game))
	
	possibleGame := true
	ch := make(chan bool, len(game))
	for _, round := range game {
		go func(round GameRound) {
			ch <- isRoundPossible(round)
			
			wg.Done()
		}(round)
	}
	
	wg.Wait()
	close(ch)
	
	for res := range ch {
		possibleGame = possibleGame && res
	}
	
	return possibleGame
}

func part1(input [][]GameRound) int {
	var wg sync.WaitGroup
	wg.Add(len(input))
	
	var gameIdsSum atomic.Int32
	for idx, game := range input {
		go func(game []GameRound, gameId int) {
			isPossible := isGamePossible(game, gameId)
			
			if isPossible {
				gameIdsSum.Add(int32(gameId))
			}
			
			wg.Done()
		}(game, idx + 1)
	}
	
	wg.Wait()
	
	return int(gameIdsSum.Load())
}

func main() {
	input := parseInput()
	
	fmt.Println(part1(input))
}