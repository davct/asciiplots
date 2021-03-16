package main

import (
  "testing"
)

func TestMakeAxisLabel(t *testing.T) {
  inputs := []float64{1.01, 1.01, 1.123, 1.0/3}
  lens := []int{1, 3, 4, 3}
  expected := []AxisLabel{"1", "1.0", "1.12", "0.3"}
  for i, inp := range inputs {
    label := MakeAxisLabel(inp, lens[i])
    if label != expected[i] {
      t.Logf("%s != %s", label, expected[i])
      t.Fail()
    }
  }
}
