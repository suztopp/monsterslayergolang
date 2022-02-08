package main

import (
	"example.com/monsterslayer/actions"
	"example.com/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{} //new slice of type roundData

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}

	interaction.PrintRoundStatistics(&roundData) //reference the struct here with pointer

	gameRounds = append(gameRounds, roundData) //over ride gameRounds with brand new slice

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
