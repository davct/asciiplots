package main

import (
  "testing"
)

func errt(msg string, t *testing.T) {
  t.Errorf("Error %s:\n", msg)
}
func byNewlines(e string, g string) string {
    return "\n expected: \n" + e + "\n got: \n" + g
}

func TestUtils(t *testing.T) {
  min, max := minmaxi([]int{1, 2, 3, 4, 5, 4})
  if min != 1 {
    errt("error min wrong", t)
  }
  if max != 5 {
    errt("error max wrong", t)
  }
}

func expectEqualBool(a bool, b bool, t *testing.T, env string) {
  if a != b {
    t.Errorf("Error in %s: bools not equal", env)
  }
}

func compareStrings(e string, g string, t *testing.T, env string) {
  if e != g {
    t.Errorf("Error in %s: %s", env, byNewlines(e, g))
  }
}

func TestFormatLabel(t *testing.T) {
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

func TestShouldColor(t *testing.T) {
  hist := HistData{
    intervals: []BucketInterval{
      BucketInterval{
        floor: 1.0,
        ceil: 1.8,
      },
      BucketInterval{
        floor: 1.8,
        ceil: 2.8,
      },
    },
    counts: []int{4, 2},
  }

  expectEqualBool(
    shouldColor(3, 3, 6, 10, hist), true, t, "should color 1")
}
