package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure" //need to run go mod tidy
	//updates the go mod file
)

type RoundData struct {
	Action           string //capital so available outside of package
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {

	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print() //prints ascii art to screen

	// fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func DeclareWinner(winner string) {

	asciiFigure := figure.NewFigure("GAME OVER", "", true)

	fmt.Println("-------------------------")
	asciiFigure.Print() //prints ascii art to screen
	// fmt.Println("GAME OVER!")
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player Attacked Monsters for %v Damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v Damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v Damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func WriteLogFile(rounds *[]RoundData) { //type of data is slice, list of roundData, dynamic array so to say

	exPath, err := os.Executable() //returns a path to the executable file with the full path and file name

	if err != nil {
		fmt.Println("Writing log file failed. Exiting")
	}

	//filepath package
	exPath = filepath.Dir(exPath) //gives us directory that holds executable - cleans exPath

	file, err := os.Create(exPath + "/gamelog.txt") //tells it to create the file in the folder that contains the executable
	//file, err := os.Create("gamelog.txt") - this is for 'go run'

	if err != nil {
		fmt.Println("Saving a log - file failed. Exiting")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine) //write string version of this map into file

		if err != nil {
			fmt.Println("Writing into Log File Failed, Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote Data to Log!")
}
