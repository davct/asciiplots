package main

import (
  "fmt"
  "sort"
  "strings"
  "strconv"
)

// @TODO: min max functions are prob unnessecary & messy + index bugs
func min(data []float64) (min float64) {
  min = data[0]
  for _, d := range data {
    if d < min {
      min = d
    }
  }
  return min
}

func max(data []float64) (max float64) {
  max = data[0]
  for _, d := range data {
    if d > max {
      max = d
    }
  }
  return max
}

func minmaxi(l []int) (min int, max int) {
  min, max = l[0], l[0]
  for _, n := range l {
    if n < min { min = n }
    if n > max { max = n }
  }
  return min, max
}

type BucketInterval struct {
  floor float64;
  ceil float64;
}

type HistData struct {
  intervals []BucketInterval;
  counts []int;
}

func formatLabel(label float64, bytes int) string {
  if bytes == 1 { bytes++ }
  if bytes == 0 { return "" }
  return strconv.FormatFloat(label, 'g', bytes-1, 64)
}

func buildXAxis(labels []string, bw int) string {
  out := labels[0]
  for _, l := range labels[1:] {
    out += strings.Repeat(" ", bw - len(l)) + l
  }
  return out
}

func printHist(hist HistData) {
  w, h := 120, 150
  xLabels, yLabels := getLabels(hist)
  bucketWidth := w / len(xLabels)
  prefixL := bucketWidth / 2
  sxLabels := make([]string, len(xLabels))
  for i, n := range xLabels {
    sxLabels[i] = formatLabel(n, 3)
  }
  xAxisLabels := strings.Repeat(" ", prefixL+1) + buildXAxis(sxLabels, bucketWidth)

  rows := make([]string, int(h / 10.0))
  // @TODO not hardcode prefix length
  cmin, cmax := minmaxi(hist.counts)
  for i, _ := range rows {
    row := ""
    // @TODO make number of y labels variable
    if i == 0 { row += strconv.Itoa(yLabels[2]) }
    if i == (int(len(rows)/2)) { row += strconv.Itoa(yLabels[1]) }
    if i == len(rows)-1 { row += strconv.Itoa(yLabels[0]) }
    whitespace := strings.Repeat(" ", (prefixL - len(row)))
    row += whitespace + "|"
    for x := prefixL+1; x < w; x++ {
      ci := (x - (x % bucketWidth)) / bucketWidth
      if ci > len(hist.counts)-1 { ci = len(hist.counts)-1 }
      c := hist.counts[ci]
      heightp := float64(len(rows) - i) / float64(len(rows))
      countp := float64(cmax - cmin) * heightp
      //fmt.Printf("heightp %f at %d -- countp %f\n", heightp, i, countp)
      if c >= int(countp) {
        row += "#"
      } else {
        row += " "
      }
    }
    rows[i] = row
  }
  rows = append(rows, xAxisLabels)
  for _,r := range rows {
    fmt.Println(r)
  }

  //fmt.Println(strings.Join(rows, "\n"))
}

func getLabels(hist HistData) (xLabels []float64 , yLabels []int) {
  xLabels = make([]float64, len(hist.intervals))
  for i, interval := range hist.intervals {
    xLabels[i] = interval.floor
  }
  xLabels = append(xLabels, hist.intervals[len(hist.intervals)-1].ceil)

  min, max := hist.counts[0],hist.counts[0]
  for _, c := range hist.counts {
    if c > max { max = c }
    if c < min { min = c }
  }
  // @TODO: Y labels
  mid := min + (max - min) / 2
  yLabels = []int{min, mid, max}
  return xLabels, yLabels
}


func getHistData(data []float64, numBuckets int) HistData {
  buckets := make([][]float64, 0, numBuckets)
  min := min(data)
  max := max(data)
  distance := max - min
  bucketSize := distance / float64(numBuckets)
  sort.Float64s(data)
  bucketFloor := min
  bucketCeil := min + bucketSize
  currentBucket := 0

  bucketIntervals := make([]BucketInterval, numBuckets)
  bucketIntervals[currentBucket] = BucketInterval{floor: bucketFloor, ceil: bucketCeil}
  bucket := make([]float64, 0)
  for _, n := range data {
    if n > bucketCeil {
      currentBucket++
      bucketFloor = bucketCeil
      bucketCeil = bucketFloor + bucketSize
      bucketIntervals[currentBucket] = BucketInterval{floor: bucketFloor, ceil: bucketCeil}
      buckets = append(buckets, bucket)
      bucket = make([]float64, 0)
      bucket = append(bucket, n)
    } else {
      bucket = append(bucket, n)
    }
  }
  buckets = append(buckets, bucket)
  bucketSizes := make([]int, numBuckets)
  for i, b := range buckets {
    bucketSizes[i] = len(b)
  }
  return HistData{intervals: bucketIntervals, counts: bucketSizes}
}

func main() {
  data := []float64{1.0, 2.0, 3.0}
  for i := 1; i < 10; i++ {
    for j:= i; j < 10; j++ {
      data = append(data, float64(i))
    }
  }
  hist := getHistData(data, 10)
  printHist(hist)
}
