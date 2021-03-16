package main

import (
  "testing"
)

func TestPlotReport(t *testing.T) {
  p := AsciiPlot{
    rows: []string{"1", "2"},
    prefixLength: 0,
  }
  got := p.Report()
  expected := "1\n2"
  if got != expected {
    t.Logf("%s != %s", got, expected)
    t.Fail()
  }
}

func TestAttachBar(t *testing.T) {
  rows := make([]string, 1)
  prefixLength := 0
  plot := AsciiPlot{rows, prefixLength}
  plot.AttachBar(1, 1)
  expected := "#|"
  got := plot.Report()
  if got != expected {
    t.Logf("rows: %s", plot.rows)
    t.Fail()
  }
}

func TestAttachBars(t *testing.T) {
}
