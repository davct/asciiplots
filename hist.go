package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	//dev
	"math/rand"
	"time"
)

type roomf struct {
	floor float64
	ceil  float64
}

type HistData struct {
	intervals []roomf
	counts    []int
}

type HistChart struct {
	hist            HistData
	h               int
	bucketCharWidth int
}

func formatLabel(label float64, mlen int) string {
	if mlen == 1 {
		mlen++
	}
	if mlen == 0 {
		return ""
	}
	fmtd := strconv.FormatFloat(label, 'f', mlen-1, 64)
	// TODO Robustness: better number converstion
	if len(fmtd) > mlen {
		fmtd = fmtd[:mlen]
	}
	if fmtd[len(fmtd)-1] == '.' {
		fmtd = fmtd[:len(fmtd)-1]
	}
	return fmtd
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

func attachXAxis(rows []string, xLabels []float64, prefixLength, bucketCharWidth int) (newrows []string) {
	labels := make([]string, len(xLabels))
	numChars := 3
	for i, n := range xLabels {
		fmtd := formatLabel(n, numChars)
		labels[i] = fmtd
	}
	prefix := strings.Repeat(" ", prefixLength)
	xaxis := prefix + buildXAxis(labels, bucketCharWidth)

	return append(rows, xaxis)
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
		numfill := prefixLength - len(rows[i])
		if numfill < 0 {
			fmt.Println("prefixLength is too small, ", prefixLength, "under ", len(rows[i]))
			numfill = 0
		}

		ws := strings.Repeat(" ", numfill)
		rows[i] += ws + "|"
	}
	return rows
}

func calcNumFilledRows(thisBucketCount, maxHeight, minBucketCount, maxBucketCount int) int {
	thisCountDiff := thisBucketCount - minBucketCount
	maxCountDiff := maxBucketCount - minBucketCount
	pctOfMax := float64(thisCountDiff) / float64(maxCountDiff)
	numfilled := float64(maxHeight) * pctOfMax
	return int(math.Round(numfilled))
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

func buildHistString(hc HistChart) string {
	xLabels, yLabels := getLabels(hc.hist)
	ys := stringifyis(yLabels)
	_, maxYLabelLength := minmaxls(ys)
	minLenWhitespaceAfterYLabels := 2
	prefixLength := maxYLabelLength + minLenWhitespaceAfterYLabels

	rows := make([]string, hc.h)
	rows = attachYAxis(ys, rows, prefixLength)
	cmin, cmax := minmaxi(hc.hist.counts)
	for _, c := range hc.hist.counts {
		rows = attachBucket(rows, hc.bucketCharWidth, c, cmin, cmax)
	}
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

func getHistData(data []float64, numbuckets int) HistData {
	sort.Float64s(data)
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
			fmt.Println(n, ceil)
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

func printHist(data []float64, numbuckets, h, bucketCharWidth int) {
	hist := getHistData(data, numbuckets)
	fmt.Println(buildHistString(HistChart{
		hist: hist, h: h, bucketCharWidth: bucketCharWidth,
	}))
}

func main() {
	data := []float64{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100000; i++ {
		data = append(data, rand.Float64()*500)
	}
	sort.Float64s(data)
	printHist(data, 10, 5, 4)
}
