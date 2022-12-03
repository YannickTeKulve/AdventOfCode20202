package day3

import (
	"os"
	"testing"
)

func TestTask1(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	//data := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	total := Task1(string(data))
	if total != 157 {
		t.Errorf("incorrect, got: %d", total)
	}
}

func TestTask2(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	//data := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	total := Task2(string(data))
	if total != 70 {
		t.Errorf("incorrect, got: %d", total)
	}
}
