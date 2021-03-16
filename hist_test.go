package main

import (
  "testing"
)

func TestNewHistogram(t *testing.T) {
  d := []float64{1, 2, 3, 4, 5}
  hist := newHistogram(d, 1)
  if len(hist.counts) != 1 {
    t.Logf("Wrong number of buckets")
    t.Fail()
  }
  expected := roomf{floor: 1.0, ceil: 5.0}
  got := hist.intervals[0]
  if got != expected {
    t.Logf("Interval is wrong in new Histogram")
    t.Fail()
  }
}
