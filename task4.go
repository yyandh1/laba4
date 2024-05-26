package main

import (
	"fmt"
	"math/rand"
	"time"
)

func unpredictable(roundNumber int, selfChoices []bool, enemyChoices []bool) bool {
	return rand.Intn(2) == 1
}

func reciprocity(roundNumber int, selfChoices []bool, enemyChoices []bool) bool {
	if !enemyChoices[roundNumber-2] {
		return false
	}
	return true
}

func mastery(roundNumber int, selfChoices []bool, enemyChoices []bool) bool {
	if roundNumber == 1 {
		return true
	}

	selfBetrayals := 0
	enemyBetrayals := 0
	for i := 0; i < roundNumber; i++ {
		if !selfChoices[i] {
			selfBetrayals++
		}
		if !enemyChoices[i] {
			enemyBetrayals++
		}
	}

	if selfBetrayals > enemyBetrayals {
		return false
	}

	return true
}

func game(tact1 func(int, []bool, []bool) bool, tact2 func(int, []bool, []bool) bool, pl1Score *int, pl2Score *int) {
	rounds := rand.Intn(101) + 100

	pl1Move := make([]bool, rounds)
	pl2Move := make([]bool, rounds)

	for i := 1; i <= rounds; i++ {
		decision1 := tact1(i, pl1Move, pl2Move)
		decision2 := tact2(i, pl2Move, pl1Move)

		pl1Move[i-1] = decision1
		pl2Move[i-1] = decision2

		if decision1 && decision2 {
			*pl1Score += 24
			*pl2Score += 24
		} else if decision1 && !decision2 {
			*pl1Score += 0
			*pl2Score += 20
		} else if !decision1 && decision2 {
			*pl1Score += 20
			*pl2Score += 0
		} else {
			*pl1Score += 4
			*pl2Score += 4
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	unpredScore := 0
	recipScore := 0
	mastScore := 0

	game(unpredictable, reciprocity, &unpredScore, &recipScore)
	game(unpredictable, mastery, &unpredScore, &mastScore)
	game(reciprocity, mastery, &recipScore, &mastScore)

	fmt.Printf("Счёт игрока с тактикой unpredictable: %d\n", unpredScore)
	fmt.Printf("Счёт игрока с тактикой reciprocity: %d\n", recipScore)
	fmt.Printf("Счёт игрока с тактикой mastery: %d\n", mastScore)
}