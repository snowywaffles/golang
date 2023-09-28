package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	myScore                     int
	numbersIHaveChosenInThePast []int
	numbersChosenLastRound      []int
	myCurrentlyChosenNumber     int
	behavior                    string
}

func (p *Player) GainOnePoint() {
	p.myScore++
}

func (p *Player) ChooseNumberRandomly() {
	p.myCurrentlyChosenNumber = rand.Intn(10) + 1 // generates an integer from 1-10
}

func (p *Player) ChooseNumber() {
	switch p.behavior {
	case "rock":
		if p.myCurrentlyChosenNumber == 0 {
			fmt.Println("WHOOPS rock forgot to choose an initial number with ChooseInitialNumberRandomly")
		}
		// else, do nothing. rock just keeps his number for every single round.
	case "opportunist":
		// do stuff
	case "cherry":
		// do stuff
	case "completelyRandom":
		p.ChooseNumberRandomly()
	}
	p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber)
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func SimulateRound(numbersChosenLastRound []int) {
	numbersChosenLastRound = []int{} // empty the slice
}

func main() {
	numbersChosenLastRound := []int{}
	allPlayers := []Player{
		{0, []int{}, []int{}, 0, "completelyRandom"},
		{0, []int{}, []int{}, 0, "rock"},
	}
	for player := range allPlayers {
		allPlayers[player].ChooseNumberRandomly()
		allPlayers[player].GainOnePoint()
		allPlayers[player].ChooseNumber()
		allPlayers[player].ChooseNumber()
		allPlayers[player].ChooseNumber()
		fmt.Println(allPlayers[player])
	}
}
