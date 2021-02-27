package main

import (
	"testing"
)

// @TODO: more cases

// @TODO: don't do this
func mkhist() HistData {
	hist := HistData{
		intervals: []roomf{
			roomf{
				floor: 1.0,
				ceil:  1.8,
			},
			roomf{
				floor: 1.8,
				ceil:  2.8,
			},
		},
		counts: []int{4, 2},
	}
	return hist
}

func byNewlines(e string, g string) string {
	return "\n expected: \n" + e + "\n got: \n" + g
}

func compareStrings(e string, g string, t *testing.T, env string) {
	if e != g {
		t.Errorf("Error in %s: %s", env, byNewlines(e, g))
	}
}

func TestFormatLabel(t *testing.T) {
	// trim trailing dot
	compareStrings("50", formatLabel(50.002, 3), t, "formatLabel")

	compareStrings("0.0", formatLabel(26.0/10000.0, 3), t, "formatLabel")
	compareStrings("3.3", formatLabel(3+(1.0/3.0), 3), t, "formatLabel")
	compareStrings("3.141", formatLabel(3.141, 5), t, "formatLabel")
	compareStrings("3.14", formatLabel(3.141, 4), t, "formatLabel")
	compareStrings("3.1", formatLabel(3.141, 3), t, "formatLabel")
	compareStrings("3", formatLabel(3.141, 1), t, "formatLabel")
	compareStrings("", formatLabel(3.141, 0), t, "formatLabel")
}

func TestBuildXAxis(t *testing.T) {
	compareStrings(
		"1.0    3.0    5.0",
		buildXAxis([]string{"1.0", "3.0", "5.0"}, 7), t, "buildXAxis")

	compareStrings(
		"1.0",
		buildXAxis([]string{"1.0"}, 25), t, "buildXAxis")
}

func TestGetHistData(t *testing.T) {
	data := []float64{1, 1, 1, 2, 3}
	counts := []int{3, 1, 1}
	outh := getHistData(data, 3)

	for i := 0; i < len(counts); i++ {
		if counts[i] != outh.counts[i] {
			t.Errorf("getHistData: Count on histogram is wrong for bucket %d", i)
			t.Errorf("Expected %d got %d", counts[i], outh.counts[i])
		}
	}
}

func TestGetLabels(t *testing.T) {
	h := mkhist()
	xs, ys := getLabels(h)
	expxs := []float64{1.0, 1.8, 2.8}
	expys := []int{2, 3, 4}
	for i := 0; i < 3; i++ {
		if xs[i] != expxs[i] || ys[i] != expys[i] {
			t.Errorf("Error in get labels: ")
		}
	}
}
