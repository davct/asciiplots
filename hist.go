package main

import (
  "math"
)

type Histogram struct {
	intervals []roomf
	counts    []int
}

type roomf struct {
	floor float64
	ceil  float64
}

func newHistogram(inputData []float64, numbuckets int) Histogram {
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
	return Histogram{intervals, counts}
}

func (h *Histogram) MakePlot(height, width int) string {
  bucketWidth := int(width / len(h.intervals))
	rows := make([]string, height)
	xLabels, yLabels := h.GetDefaultLabels()
	stringifiedYLabels := intsToStrings(yLabels)
  prefixLength := calculatePrefixLength(stringifiedYLabels)
  plot := AsciiPlot{
    rows: rows,
    prefixLength: prefixLength,
  }
	plot.AttachYAxis(stringifiedYLabels)
  heights, widths := h.GetBucketDimensions(height, bucketWidth)
  plot.AttachBars(heights, widths)
	plot.AttachXAxis(xLabels, bucketWidth+1)
  return plot.Report()
}

func (h *Histogram) GetBucketDimensions(maxHeight, defaultWidth int) (heights, widths []int) {
  hardcodedMinHeight := 2
  heights = h.CalculateBucketHeights(hardcodedMinHeight, maxHeight)
  widths = make([]int, len(heights))
  for i, _ := range widths {
    widths[i] = defaultWidth
  }
  return heights, widths
}

func (h *Histogram) CalculateBucketHeights(minHeight, maxHeight int) (heights []int) {
  cmin, cmax := minmaxi(h.counts)
  for _, c := range h.counts {
    thisDiff := c - cmin
    maxDiff := cmax - cmin
    percentOfMax := float64(thisDiff) / float64(maxDiff)
    bucketHeight := int(math.Round(float64(maxHeight) * percentOfMax))
    if bucketHeight < minHeight {
      bucketHeight = minHeight
    }
    heights = append(heights, bucketHeight)
  }
  return heights
}

func (h *Histogram) GetDefaultLabels() (xLabels []float64, yLabels []int) {
	xLabels = make([]float64, len(h.intervals)+1)
	for i, interval := range h.intervals {
		xLabels[i] = interval.floor
	}
	end := len(h.intervals) - 1
	xLabels[end+1] = h.intervals[end].ceil
	min, max := minmaxi(h.counts)

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
