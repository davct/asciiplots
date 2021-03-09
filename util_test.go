package main

import (
	"testing"
  "sort"
)

func TestMinMax(t *testing.T) {
	mini, maxi := minmaxi([]int{1, 2, 3, 4, 5, 4})
	msg := "Error in testing utils on min/max"
	if mini != 1 {
		t.Errorf("Error %s:\n", msg+"i")
	}
	if maxi != 5 {
		t.Errorf("Error %s:\n", msg+"i")
	}

	minf, maxf := minmaxf([]float64{1.0, 1.1, 1.3, 0.2})
	if minf != float64(0.2) {
		t.Errorf("Error %s:\n", msg+"f")
	}
	if maxf != float64(1.3) {
		t.Errorf("Error %s:\n", msg+"f")
	}

	minl, maxl := minmaxls([]string{"1", "22", "333"})
	if minl != 1 {
		t.Errorf("Error %s:\n", msg+"stringl")
	}
	if maxl != 3 {
		t.Errorf("Error %s:\n", msg+"stringl")
	}
}

func TestConversions(t *testing.T) {
	inp := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}
	got := intsToStrings(inp)
	for i, _ := range got {
		if got[i] != expected[i] {
			t.Errorf("Error changing int slice to string slice...")
		}
	}
}

func TestTruncatef(t *testing.T) {
	input := []float64{1.12345, 1.0 / 3.0, 5}
	truncto := []int{3, 3, 10}
	expected := []float64{1.12, 0.33, 5}
	for i, _ := range input {
		out := truncatef(input[i], truncto[i])
		if out != expected[i] {
			t.Errorf("Error in truncatef on test case %d\n", i)
			t.Errorf("Expected %f but got %f", expected[i], out)
		}
	}
}


func TestGetSortedCopy(t *testing.T) {
  input := []float64{1.0, 0, 2.0, 5.0, 3.0}
  got := getSortedCopy(input)
  sort.Float64s(input)
  epsilon := 0.0001
  equal, _ := areFloatSlicesEqual(got, input, epsilon)
  if equal != true {
    t.Errorf("TestGetSortedCopy 1: got a bad copy")
  }
}
