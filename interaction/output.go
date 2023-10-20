package interaction

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action              string
	PlayerAttackDamage  int
	PlayerHealValue     int
	MonsterAttackDamage int
	PlayerHealth        int
	MonsterHealth       int
}

func PrintGreetings() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action:")
	fmt.Println("--------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage\n", roundData.PlayerAttackDamage)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage\n", roundData.PlayerAttackDamage)
	} else {
		// HEAL
		fmt.Printf("Player healed for %v\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage\n", roundData.MonsterAttackDamage)

	fmt.Println("--------------------------")
	fmt.Printf("Player health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v\n", roundData.MonsterHealth)
	fmt.Println("--------------------------")
}

func DeclareWinner(winner string) {
	fmt.Println("--------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("--------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(roundData *[]RoundData) {
	fmt.Println("Writing log file...")
	file, err := os.Create("gamelog.txt")

	if err != nil {
		fmt.Println("Error creating log file")
		return
	}

	for index, round := range *roundData {
		logEntry := map[string]string{
			"Round":                 fmt.Sprintf("%v", index+1),
			"Action":                round.Action,
			"Player Attack Damage":  fmt.Sprintf("%v", round.PlayerAttackDamage),
			"Player Heal Value":     fmt.Sprintf("%v", round.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprintf("%v", round.MonsterAttackDamage),
			"Player Health":         fmt.Sprintf("%v", round.PlayerHealth),
			"Monster Health":        fmt.Sprintf("%v", round.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Error writing log file")
			continue
		}
	}

	file.Close()
	fmt.Println("Log file written successfully")
}
