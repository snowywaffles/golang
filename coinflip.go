package main

import (
	"fmt"
	"math/rand"
)

func flipNCoins(n int) []int {
	allFlips := []int{}
	for i := 1; i <= n; i++ {
		flip := rand.Intn(2) // will psuedorandomly generate either a 0 or a 1. 0 = tails, and 1 = heads
		allFlips = append(allFlips, flip)
	}
	return allFlips
}

func calculateAdjustment(allFlips []int) int {
	adjustment := 0
	for _, flip := range allFlips {
		if flip == 0 {
			adjustment--
		} else if flip == 1 {
			adjustment++
		}
	}
	return adjustment
}

func main() {
	currentTotal := 0
	moneyEarned := 0
	numberOfCoins := 3
	for true {

		outOfRange := (currentTotal < -2) || (currentTotal > 2)
		if outOfRange == true {
			moneyEarned = 0 // lose all your money!
			break
		} else {
			moneyEarned += (1 * numberOfCoins)
		}

		allFlips := flipNCoins(numberOfCoins)
		adjustment := calculateAdjustment(allFlips)
		currentTotal += adjustment
		fmt.Println("allFlips: ", allFlips, "adj: ", adjustment, "currentTotal: ", currentTotal)

	}
}
