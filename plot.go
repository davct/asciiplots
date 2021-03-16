package main

import (
  "strings"
)

type AsciiPlot struct {
  rows []string
  prefixLength int
}

func (plot *AsciiPlot) Report() string {
  return strings.Join(plot.rows, "\n")
}

func (plot *AsciiPlot) AttachYAxis(ys []string) {
  yAxisParts := buildYAxisParts(ys, plot.prefixLength, len(plot.rows))
  for i, _ := range plot.rows {
    plot.rows[i] += yAxisParts[i]
  }
}

func (plot *AsciiPlot) AttachBars(heights, widths []int) {
  for i, _ := range heights {
		plot.AttachBar(heights[i], widths[i])
	}
}

func buildYAxisParts(ys []string, prefixLength, numRows int) []string {
  parts := make([]string, numRows)
	for i := 0; i < numRows; i++ {
		// TODO: Y Label refactor
		if i == 0 {
			parts[i] += ys[2]
		}
		if i == (int(numRows / 2)) {
			parts[i] += ys[1]
		}
		if i == numRows-1 {
			parts[i] += ys[0]
		}
		ws := strings.Repeat(" ", prefixLength-len(parts[i]))
		parts[i] += ws + "|"
  }
  return parts
}

func (plot *AsciiPlot) AttachBar(height, width int) {
	for i, _ := range plot.rows {
		if (len(plot.rows) > height) && (i < len(plot.rows)-height) {
			plot.rows[i] += strings.Repeat(" ", width) + "|"
    } else {
      plot.rows[i] += strings.Repeat("#", width) + "|"
    }
  }
}

func (plot *AsciiPlot) AttachXAxis(xLabels []float64, bucketCharWidth int) {
	labels := make([]string, len(xLabels))
	numChars := 3
	for i, n := range xLabels {
		formatted := MakeAxisLabel(n, numChars)
		labels[i] = string(formatted)
	}
	prefix := strings.Repeat(" ", plot.prefixLength)
	xaxis := prefix + buildXAxis(labels, bucketCharWidth)
	plot.rows = append(plot.rows, xaxis)
}

func buildXAxis(labels []string, labelDistance int) string {
	out := labels[0]
	for _, l := range labels[1:] {
		numfill := labelDistance - len(l)
		if numfill < 0 {
			numfill = 0
		}
		out += strings.Repeat(" ", numfill) + l
	}
	return out
}
