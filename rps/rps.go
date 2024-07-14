package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Piedra. Vence a las tijeras . (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel. Vence a la piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijeras. Vence al papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Ganaste!",
	"Buen trabajo",
	"¡Sigue así!",
	"Deverias comprar un boleto de lotería",
	"¡Eres un genio!",
}

var loseMessages = []string{
	"¡Perdiste!",
	"¡Intentalo de nuevo!",
	"¡No te rindas!",
	"¡Sigue intentando!",
	"¡Que lastima!",
	"Hoy simplemente no es tu día",
}

var drawMessages = []string{
	"¡Empate!",
	"¡Vaya!",
	"¡Que coincidencia!",
	"Las grandes mentes piensan igual",
	"Oh oh, parece que tenemos un empate",
	"Nadie gana, nadie pierde",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERAS"
	}
	// Genera un numero aleatorio entre 0 y 2, que usamos para elegir un mensaje aleatorio
	var messageInt int = 0

	// Declara una variable para contener el mensaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		messageInt = len(drawMessages)
		message = drawMessages[rand.Intn(messageInt)]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "Ganaste"
		messageInt = len(winMessages)
		message = winMessages[rand.Intn(messageInt)]
	} else {
		ComputerScore++
		roundResult = "Perdiste"
		messageInt = len(loseMessages)
		message = loseMessages[rand.Intn(messageInt)]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
