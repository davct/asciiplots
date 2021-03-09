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
	hist := buildHistogramData(data, numbuckets)

  histogramAsciiPlot := buildHistString(HistChart{
		hist: hist,
    h: h,
    bucketCharWidth: bucketCharWidth,
	})

  fmt.Println(histogramAsciiPlot)
}


func buildHistogramData(inputData []float64, numbuckets int) HistData {
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

func buildHistString(hc HistChart) string {
	rows := make([]string, hc.h)
	xLabels, yLabels := getLabels(hc.hist)
	stringifiedYLabels := intsToStrings(yLabels)

  prefixLength := calculatePrefixLength(stringifiedYLabels)
	rows = attachYAxis(stringifiedYLabels, rows, prefixLength)
  rows = attachBuckets(rows, hc)
	rows = attachXAxis(rows, xLabels, prefixLength, hc.bucketCharWidth+1)
	return strings.Join(rows, "\n")
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


func attachYAxis(ys []string, rows []string, prefixLength int) (newrows []string) {
	for i := 0; i < len(rows); i++ {
		// TODO: Y Label refactor
		if i == 0 {
			rows[i] += ys[2]
		}
		if i == (int(len(rows) / 2)) {
			rows[i] += ys[1]
		}
		if i == len(rows)-1 {
			rows[i] += ys[0]
		}
		ws := strings.Repeat(" ", prefixLength-len(rows[i]))
		rows[i] += ws + "|"
	}
	return rows
}


func attachBuckets(rows []string, hc HistChart) []string {
  cmin, cmax := minmaxi(hc.hist.counts)
  for _, c := range hc.hist.counts {
		rows = attachBucket(rows, hc.bucketCharWidth, c, cmin, cmax)
	}
  return rows
}

func attachBucket(rows []string, width, count, cmin, cmax int) (newrows []string) {
	height := calcNumFilledRows(count, len(rows), cmin, cmax)
	for i, _ := range rows {
		if i < len(rows)-height {
			rows[i] += strings.Repeat(" ", width) + "|"
	} else {
			rows[i] += strings.Repeat("#", width) + "|"
		}
	}
	return rows
}

func calcNumFilledRows(thisBucketCount, maxHeight, minBucketCount, maxBucketCount int) int {
	thisCountDiff := thisBucketCount - minBucketCount
	maxCountDiff := maxBucketCount - minBucketCount
	percentOfMax := float64(thisCountDiff) / float64(maxCountDiff)
	numfilled := float64(maxHeight) * percentOfMax
	return int(math.Round(numfilled))
}

func attachXAxis(rows []string, xLabels []float64, prefixLength, bucketCharWidth int) (newrows []string) {
	labels := make([]string, len(xLabels))
	numChars := 3
	for i, n := range xLabels {
		formatted := formatLabel(n, numChars)
		labels[i] = formatted
	}
	prefix := strings.Repeat(" ", prefixLength)
	xaxis := prefix + buildXAxis(labels, bucketCharWidth)

	return append(rows, xaxis)
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

func buildXAxis(labels []string, bucketCharWidth int) string {
	out := labels[0]
	for _, l := range labels[1:] {
		numfill := bucketCharWidth - len(l)
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
