package main

import (
	"github.com/abrl91/monsterslayer/actions"
	"github.com/abrl91/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreetings()
}

// return the winner
func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDamage int

	if userChoice == "ATTACK" {
		playerAttackDamage = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		// SPECIAL_ATTACK
		playerAttackDamage = actions.AttackMonster(true)
	}

	monsterAttackDamage = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthState()

	roundData := interaction.RoundData{
		Action:              userChoice,
		PlayerHealth:        playerHealth,
		MonsterHealth:       monsterHealth,
		PlayerAttackDamage:  playerAttackDamage,
		PlayerHealValue:     playerHealValue,
		MonsterAttackDamage: monsterAttackDamage,
	}

	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "MONSTER"
	} else if monsterHealth <= 0 {
		return "PLAYER"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
