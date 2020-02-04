package main

import "fmt"

const team = "Green Team"

func getValue(scores map[string]int, key string) (int, error) {
	var value int
	var ok bool

	if value, ok = scores[key]; !ok {
		return -1, fmt.Errorf("value not in map")
	}

	return value, nil
}

func main() {
	scores := make(map[string]int)

	value, err := getValue(scores, team)
	if err != nil {
		fmt.Printf("failed to get value: %s\n", err)
		return
	}

	fmt.Println(team, value)
}
