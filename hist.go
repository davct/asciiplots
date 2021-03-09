package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	//dev
	"math/rand"
	"time"
)

type HistChart struct {
	hist            HistData
	h               int
	bucketCharWidth int
}

type HistData struct {
	intervals []roomf
	counts    []int
}

type roomf struct {
	floor float64
	ceil  float64
}

func printHist(data []float64, numbuckets, h, bucketCharWidth int) {
	hist := getHistogramData(data, numbuckets)

  histogramAsciiPlot := buildHistogramPlot(HistChart{
		hist: hist,
    h: h,
    bucketCharWidth: bucketCharWidth,
	})

  fmt.Println(histogramAsciiPlot)
}


func getHistogramData(inputData []float64, numbuckets int) HistData {
  data := getSortedCopy(inputData)
	intervals := []roomf{}
	min, max := minmaxf(data)
	distance := max - min
	bw := distance / float64(numbuckets)
	floor := min
	ceil := min + bw
	counts := make([]int, numbuckets)
	bucketIndex := 0
	for i, n := range data {
		if bucketIndex+1 == numbuckets {
			counts[bucketIndex] += len(data) - i
			break
		}
		if n > ceil {
			intervals = append(intervals, roomf{floor, ceil})
			floor = ceil
			ceil += bw
			bucketIndex++
		}
		counts[bucketIndex]++
	}
	intervals = append(intervals, roomf{floor, ceil})
	return HistData{intervals, counts}
}

func buildHistogramPlot(hc HistChart) string {
	rows := make([]string, hc.h)
	xLabels, yLabels := getLabels(hc.hist)
	stringifiedYLabels := intsToStrings(yLabels)
  prefixLength := calculatePrefixLength(stringifiedYLabels)
  plot := AsciiPlot{
    rows: rows,
    prefixLength: prefixLength,
  }
	plot.attachYAxis(stringifiedYLabels)
  plot.attachBuckets(hc)
	plot.attachXAxis(xLabels, hc.bucketCharWidth+1)
	return strings.Join(plot.rows, "\n")
}

type AsciiPlot struct {
  rows []string
  prefixLength int
}

func getLabels(hist HistData) (xLabels []float64, yLabels []int) {
	xLabels = make([]float64, len(hist.intervals)+1)
	for i, interval := range hist.intervals {
		xLabels[i] = interval.floor
	}
	end := len(hist.intervals) - 1
	xLabels[end+1] = hist.intervals[end].ceil
	min, max := minmaxi(hist.counts)

	// TODO: Y labels
	mid := min + (max-min)/2
	yLabels = []int{min, mid, max}
	return xLabels, yLabels
}

func calculatePrefixLength(yLabels []string) (prefixLength int) {
	_, maxYLabelLength := minmaxls(yLabels)
	minLenWhitespaceAfterYLabels := 2
	prefixLength = maxYLabelLength + minLenWhitespaceAfterYLabels
  return prefixLength
}

func (plot *AsciiPlot) attachYAxis(ys []string) {
  yAxisParts := buildYAxisParts(ys, plot.prefixLength, len(plot.rows))
  for i, _ := range plot.rows {
    plot.rows[i] += yAxisParts[i]
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

func (plot *AsciiPlot) attachBuckets(hc HistChart) {
  cmin, cmax := minmaxi(hc.hist.counts)
  for _, c := range hc.hist.counts {
    thisDiff := c - cmin
    maxDiff := cmax - cmin
    percentOfMax := float64(thisDiff) / float64(maxDiff)
    height := int(math.Round(float64(hc.h) * percentOfMax))
    bp := BucketParams{height: height, width: hc.bucketCharWidth}
		plot.attachBucket(bp)
	}
}

type BucketParams struct {
  height int
  width int
}

func (plot *AsciiPlot) attachBucket(bp BucketParams) {
	for i, _ := range plot.rows {
		if i < len(plot.rows)-bp.height {
			plot.rows[i] += strings.Repeat(" ", bp.width) + "|"
    } else {
      plot.rows[i] += strings.Repeat("#", bp.width) + "|"
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

func main() {
	data := []float64{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100000; i++ {
		data = append(data, rand.Float64()*500)
	}
	printHist(data, 10, 10, 8)
}
