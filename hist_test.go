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
        ceil: 1.8,
      },
      roomf{
        floor: 1.8,
        ceil: 2.8,
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

func expectEqualBool(a bool, b bool, t *testing.T, env string) {
  if a != b {
    t.Errorf("Error in %s: bools not equal", env)
  }
}

func TestUtils(t *testing.T) {
  mini, maxi := minmaxi([]int{1, 2, 3, 4, 5, 4})
  msg := "Error in testing utils on min/max"
  if mini != 1 {
    t.Errorf("Error %s:\n", msg)
  }
  if maxi != 5 {
    t.Errorf("Error %s:\n", msg)
  }

  minf, maxf := minmaxf([]float64{1.0, 1.1, 1.3, 0.2})
  if minf != float64(0.2) {
    t.Errorf("Error %s:\n", msg)
  }
  if maxf != float64(1.3) {
    t.Errorf("Error %s:\n", msg)
  }
}

func TestFormatLabel(t *testing.T) {
  // trim trailing dot
  compareStrings("50", formatLabel(50.002, 3), t, "formatLabel")

  compareStrings("0.0", formatLabel(26.0/10000.0, 3), t, "formatLabel")
  compareStrings("3.3", formatLabel(3 + (1.0/3.0), 3), t, "formatLabel")
  compareStrings("3.141", formatLabel(3.141, 5), t, "formatLabel")
  compareStrings("3.14", formatLabel(3.141, 4), t, "formatLabel")
  compareStrings("3.1", formatLabel(3.141, 3), t, "formatLabel")
  compareStrings("3", formatLabel(3.141, 1), t, "formatLabel")
  compareStrings("", formatLabel(3.141, 0), t, "formatLabel")
}

func TestBuildXAxis(t *testing.T) {
  compareStrings(
    "1.0    3.0    5.0",
    buildXAxis([]string{"1.0","3.0","5.0"}, 7), t, "buildXAxis")

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

func TestGetIntervals(t *testing.T) {
  data := []float64{1.0, 2.0, 3.0, 4.0}
  numbuckets := 3
  got := getIntervals(data, numbuckets)
  expected := []roomf{
    roomf{1.0, 2.0},
    roomf{2.0, 3.0},
    roomf{3.0, 4.0},
  }
  for i, room := range got {
    if room != expected[i] {
      t.Errorf("Get interval test 1 failed.\n")
    }
  }
}





