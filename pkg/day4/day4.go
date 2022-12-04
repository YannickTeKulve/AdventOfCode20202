package day4

import (
	"strconv"
	"strings"
)

func Task1(input string) int {
	pairs := strings.Split(input, "\n")
	total := 0
	for _, pair := range pairs {
		splitted := strings.Split(pair, ",")
		if contains(splitted[0], splitted[1]) || contains(splitted[1], splitted[0]) {
			total++
		}
	}
	return total
}

func contains(container string, contains string) bool {
	leftbound, leftInsidebound, rightbound, rightInsidebound := splitToBounds(container, contains)

	return leftbound <= leftInsidebound && rightInsidebound <= rightbound
}

func splitToBounds(container string, contains string) (int, int, int, int) {
	containerRange := strings.Split(container, "-")
	containsRange := strings.Split(contains, "-")

	leftbound, _ := strconv.Atoi(containerRange[0])
	leftInsidebound, _ := strconv.Atoi(containsRange[0])
	rightbound, _ := strconv.Atoi(containerRange[1])
	rightInsidebound, _ := strconv.Atoi(containsRange[1])
	return leftbound, leftInsidebound, rightbound, rightInsidebound
}

func Task2(input string) int {
	pairs := strings.Split(input, "\n")
	total := 0
	for _, pair := range pairs {
		splitted := strings.Split(pair, ",")
		if overlap(splitted[0], splitted[1]) || overlap(splitted[1], splitted[0]) {
			total++
		}
	}
	return total
}

func overlap(container string, contains string) bool {
	leftbound, leftInsidebound, rightbound, rightInsidebound := splitToBounds(container, contains)

	left := leftbound <= leftInsidebound && leftInsidebound <= rightbound
	right := leftbound <= rightInsidebound && rightInsidebound <= rightbound
	return left || right
}
