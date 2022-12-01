package day1

import (
	"sort"
	"strconv"
	"strings"
)

func task1(input string) int {
	var splitPerElf = strings.Split(input, "\n\n")
	var highest int
	for _, valueE := range splitPerElf {
		var parsedCalories = []int{}
		var caloriesPerElf = strings.Split(valueE, "\n")
		for _, calorie := range caloriesPerElf {
			parsed, _ := strconv.Atoi(calorie)
			parsedCalories = append(parsedCalories, parsed)
		}
		total := sum(parsedCalories)
		if total > highest {
			highest = total
		}
	}
	return highest
}

func task2(input string) int {
	var splitPerElf = strings.Split(input, "\n\n")
	var highest = []int{0, 0, 0}
	for _, valueE := range splitPerElf {
		var parsedCalories = []int{}
		var caloriesPerElf = strings.Split(valueE, "\n")
		for _, calorie := range caloriesPerElf {
			parsed, _ := strconv.Atoi(calorie)
			parsedCalories = append(parsedCalories, parsed)
		}
		total := sum(parsedCalories)
		for i, value := range highest {

			if value < total {
				highest[i] = total
				sort.Ints(highest)
				break
			}
		}
	}
	return sum(highest)
}

func sum(x []int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}
