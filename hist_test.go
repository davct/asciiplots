package main

import (
  "testing"
)

func DefaultHist() Histogram {
  d := []float64{1, 2, 3, 4, 5}
  h := newHistogram(d, 2)
  return h
}

func TestNewHistogram(t *testing.T) {
  h := DefaultHist()
  if len(h.counts) != 2 {
    t.Logf("Wrong number of buckets")
    t.Fail()
  }
  expected := roomf{floor: 1.0, ceil: 3.0}
  got := h.intervals[0]
  if got != expected {
    t.Logf("Interval is wrong in new hogram")
    t.Logf("%+v\n", got)
    t.Fail()
  }
}

func TestCalculateBucketHeights(t *testing.T) {
  h := DefaultHist()
  hs := h.CalculateBucketHeights(2, 15)
  if hs[0] != 15 {
    t.Fail()
  }
}

func TestGetBucketDimensions(t *testing.T) {
  h := DefaultHist()
  _, ws := h.GetBucketDimensions(15, 10)
  for i, _ := range ws {
    if ws[i] != 10 {
      t.Fail()
    }
  }
}

func TestGetDefaultLabels(t *testing.T) {
  h := DefaultHist()
  xls, yls := h.GetDefaultLabels()
  if yls[0] != 2 || yls[2] != 3 {
    t.Logf("y labels: %v\n", yls)
    t.Fail()
  }
  if len(xls) != len(h.intervals)+1 {
    t.Logf("x labels: %v\n", xls)
    t.Fail()
  }
}

func TestCalculatePrefixLength(t *testing.T) {
  yls := []string{"label", "two"}
  got := calculatePrefixLength(yls)
  if got != 7 {
    t.Fail()
  }
}
