package day3

import "strings"

func Task1(input string) int {
	nummer := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	splitperbag := strings.Split(input, "\n")
	for _, bag := range splitperbag {
		bag1 := bag[0 : len(bag)/2]
		bag2 := bag[len(bag)/2:]
		for _, char := range bag1 {
			if strings.Contains(bag2, string(char)) {
				bag2 = strings.ReplaceAll(bag2, string(char), "")
				total += strings.Index(nummer, string(char)) + 1
			}
		}
	}
	return total

}

func Task2(input string) int {
	nummer := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	splitperbag := strings.Split(input, "\n")

	for i := 0; i < len(splitperbag); i += 3 {
		bag1 := splitperbag[i]
		bag2 := splitperbag[i+1]
		bag3 := splitperbag[i+2]
		for _, char := range bag1 {
			if strings.Contains(bag2, string(char)) && strings.Contains(bag3, string(char)) {
				total += strings.Index(nummer, string(char)) + 1
				break
			}
		}

	}
	return total
}
