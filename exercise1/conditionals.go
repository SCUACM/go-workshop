package main

import "fmt"

func main() {
	scores := make(map[string]int)
	scores["Blue Team"] = 26

	redScore := 14

	if blueScore, ok := scores["Blue Team"]; ok {
		if blueScore < redScore {
			fmt.Println("Red wins!")
		} else if blueScore > redScore {
			fmt.Println("Blue wins!")
		} else {
			fmt.Println("Tie!")
		}
	}
}
