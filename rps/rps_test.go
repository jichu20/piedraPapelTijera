package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case ROCK:
			fmt.Println("El jugador eligió PIEDRA")
		case PAPER:
			fmt.Println("El jugador eligió PAPEL")
		case SCISSORS:
			fmt.Println("El jugador eligió TIJERAS")
		}

		fmt.Printf("Computer choice: %s\n", round.ComputerChoice)
		fmt.Printf("Round result: %s\n", round.RoundResult)
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("Computer choice int: %d\n", round.ComputerChoiceInt)
		fmt.Printf("Computer score: %s\n", round.ComputerScore)
		fmt.Printf("Player score: %s\n", round.PlayerScore)

		fmt.Println("\n=========================================================")
	}

}
