package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var stake, goal, trials int
	fmt.Print("Enter stake: ")
	fmt.Scanln(&stake)
	fmt.Print("Enter goal: ")
	fmt.Scanln(&goal)
	fmt.Print("Enter number of trials: ")
	fmt.Scanln(&trials)

	rand.Seed(time.Now().UnixNano())

	totalBets := 0
	totalWins := 0

	for i := 0; i < trials; i++ {
		bets := 0
		wins := 0

		for stake > 0 && stake < goal {
			bets++
			if rand.Intn(2) == 0 {
				stake++
				wins++
			} else {
				stake--
			}
		}
		totalBets += bets
		totalWins += wins
	}

	winPercent := 100.0 * float64(totalWins) / float64(totalBets)
	avgBets := float64(totalBets) / float64(trials)
	fmt.Printf("Number of times won: %d\n", totalWins)
	fmt.Printf("Win percent: %.2f%%\n", winPercent)
	fmt.Printf("Average number of bets: %.2f\n", avgBets)
}
