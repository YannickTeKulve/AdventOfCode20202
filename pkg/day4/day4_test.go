package day4

import (
	"os"
	"testing"
)

func TestTask1(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	//data := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	total := Task1(string(data))
	if total != 2 {
		t.Errorf("incorrect, got: %d", total)
	}
}

func TestTask2(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	//data := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	total := Task2(string(data))
	if total != 4 {
		t.Errorf("incorrect, got: %d", total)
	}
}
