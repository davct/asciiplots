package main

import (
  "strings"
	"strconv"
)

type AsciiPlot struct {
  rows []string
  prefixLength int
}

func (plot *AsciiPlot) Report() string {
  return strings.Join(plot.rows, "\n")
}

func (plot *AsciiPlot) attachYAxis(ys []string) {
  yAxisParts := buildYAxisParts(ys, plot.prefixLength, len(plot.rows))
  for i, _ := range plot.rows {
    plot.rows[i] += yAxisParts[i]
  }
}

func (plot *AsciiPlot) attachBars(heights, widths []int) {
  for i, _ := range heights {
		plot.attachBar(heights[i], widths[i])
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

func (plot *AsciiPlot) attachBar(height, width int) {
	for i, _ := range plot.rows {
		if (len(plot.rows) > height) && (i < len(plot.rows)-height) {
			plot.rows[i] += strings.Repeat(" ", width) + "|"
    } else {
      plot.rows[i] += strings.Repeat("#", width) + "|"
    }
  }
}

func (plot *AsciiPlot) attachXAxis(xLabels []float64, bucketCharWidth int) {
	labels := make([]string, len(xLabels))
	numChars := 3
	for i, n := range xLabels {
		formatted := formatLabel(n, numChars)
		labels[i] = formatted
	}
	prefix := strings.Repeat(" ", plot.prefixLength)
	xaxis := prefix + buildXAxis(labels, bucketCharWidth)
	plot.rows = append(plot.rows, xaxis)
}

func formatLabel(label float64, mlen int) string {
	if mlen == 1 {
		mlen++
	}
	if mlen == 0 {
		return ""
	}
	formatted := strconv.FormatFloat(label, 'f', mlen-1, 64)
	// TODO Robustness: better number converstion
	if len(formatted) > mlen {
		formatted = formatted[:mlen]
	}
	if formatted[len(formatted)-1] == '.' {
		formatted = formatted[:len(formatted)-1]
	}
	return formatted
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
