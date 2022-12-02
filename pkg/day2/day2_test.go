package day2

import (
	"os"
	"testing"
)

func TestTask1(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	//input := "A Y\nB X\nC Z"
	total := Task1(string(data))
	if total != 15 {
		t.Errorf("incorrect, got: %d", total)
	}
}

func TestTask2(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	//input := "A Y\nB X\nC Z"
	total := Task2(string(input))
	if total != 12 {
		t.Errorf("incorrect, got: %d", total)
	}
}
