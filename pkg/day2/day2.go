package day2

import "strings"

func createResultMap() map[string]int {
	m := make(map[string]int)
	m["A X"] = 4 //3
	m["A Y"] = 8 //6
	m["A Z"] = 3 //0
	m["B X"] = 1 //0
	m["B Y"] = 5 //3
	m["B Z"] = 9 //6
	m["C X"] = 7 //6
	m["C Y"] = 2 //0
	m["C Z"] = 6 //3
	return m
}

func createResultMap2() map[string]int {
	m := make(map[string]int)
	m["A X"] = 3 //3
	m["A Y"] = 4 //6
	m["A Z"] = 8 //0
	m["B X"] = 1 //0
	m["B Y"] = 5 //3
	m["B Z"] = 9 //6
	m["C X"] = 2 //6
	m["C Y"] = 6 //0
	m["C Z"] = 7 //3
	return m
}

func Task1(input string) int {
	possibleResults := createResultMap()
	splitPerRound := strings.Split(input, "\n")
	total := 0
	for _, round := range splitPerRound {
		total += possibleResults[round]
	}
	return total
}

func Task2(input string) int {
	possibleResults := createResultMap2()
	splitPerRound := strings.Split(input, "\n")
	total := 0
	for _, round := range splitPerRound {
		total += possibleResults[round]
	}
	return total
}
