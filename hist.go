package main

import (
  "fmt"
  "sort"
  "strings"
  "strconv"
  "math"
  //dev
  "math/rand"
  "time"
)

func minmaxi(l []int) (min int, max int) {
  min, max = l[0], l[0]
  for _, n := range l {
    if n < min { min = n }
    if n > max { max = n }
  }
  return min, max
}

func minmaxf(l []float64) (min float64, max float64) {
  min, max = l[0], l[0]
  for _, n := range l {
    if n < min { min = n }
    if n > max { max = n }
  }
  return min, max
}

type roomf struct {
  floor float64;
  ceil float64;
}

type HistData struct {
  intervals []roomf;
  counts []int;
}

type HistChart struct {
  hist HistData;
  h int; // height in terms of bytes
  bsize int; // bucket size in bytes
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

func attachXAxis(rows []string, xLabels []float64, prefixl, bw int) (newrows []string) {
  labels := make([]string, len(xLabels))
  numChars := 3
  for i, n := range xLabels {
    labels[i] = formatLabel(n, numChars)
  }
  prefix := strings.Repeat(" ", prefixl)
  xaxis := prefix + buildXAxis(labels, bw)

  return append(rows, xaxis)
}

func attachYAxis(yLabels []int, rows []string, prefixl int) (newrows []string) {
  for i := 0; i < len(rows); i++ {
    // @TODO: Y Label refactor
    if i == 0 { rows[i] += strconv.Itoa(yLabels[2]) }
    if i == (int(len(rows)/2)) { rows[i] += strconv.Itoa(yLabels[1]) }
    if i == len(rows)-1 { rows[i] += strconv.Itoa(yLabels[0]) }
    ws := strings.Repeat(" ", prefixl - len(rows[i]))
    rows[i] += ws + "|"
  }
  return rows
}

func checkCell(currentCount, mheight, cheight, min, max int) bool {
  countRatio := float64(currentCount - min) / float64(max - min)
  heightRatio := float64(cheight) / float64(mheight)
  return heightRatio <= countRatio
}

func calcNumFilledRows(bc, mheight, bmin, bmax int) int {
  countp := float64(bc - bmin) / float64(bmax - bmin)
  numfilled := float64(mheight) * countp
  return int(math.Round(numfilled))
}

func attachBucket(rows []string, width, count, cmin, cmax int) (newrows []string) {
  height := calcNumFilledRows(count, len(rows), cmin, cmax)
  fmt.Println("height:, ", height)
  for i, k := len(rows)-1, 0; k < height; k++ {
    rows[i-k] += strings.Repeat("#", width)
  }
  for i, _ := range rows[:len(rows)-height] {
    rows[i] += strings.Repeat(" ", width)
  }
  for i, _ := range rows { rows[i] += "|" }
  return rows
}

func buildHistString(hc HistChart) string {
  prefixl := int(hc.bsize / 2)
  xLabels, yLabels := getLabels(hc.hist)

  rows := make([]string, hc.h)
  rows = attachYAxis(yLabels, rows, prefixl)
  cmin, cmax := minmaxi(hc.hist.counts)
  for _, c := range hc.hist.counts {
    // @TODO compute bucket height -> simpler fill
    rows = attachBucket(rows, hc.bsize, c, cmin, cmax)
  }
  rows = attachXAxis(rows, xLabels, prefixl, hc.bsize)
  return strings.Join(rows, "\n")
}

func getLabels(hist HistData) (xLabels []float64 , yLabels []int) {
  xLabels = make([]float64, len(hist.intervals)+1)
  for i, interval := range hist.intervals {
    xLabels[i] = interval.floor
  }
  end := len(hist.intervals)-1
  xLabels[end+1] = hist.intervals[end].ceil
  min, max := minmaxi(hist.counts)

  // @TODO: Y labels
  mid := min + (max - min) / 2
  yLabels = []int{min, mid, max}
  return xLabels, yLabels
}

func getIntervals(data []float64, numbuckets int) []roomf {
  sort.Float64s(data)
  res := []roomf{}
  min, max := minmaxf(data)
  distance := max - min
  bsize := distance / float64(numbuckets)
  floor := min
  ceil := min + bsize
  for _, n := range data {
    if n > ceil {
      res = append(res, roomf{floor, ceil})
      floor = ceil
      ceil += bsize
    }
  }
  res = append(res, roomf{floor, ceil})
  return res
}

func getHistData(data []float64, numbuckets int) HistData {
  if len(data) == 0 {
    return HistData{
      intervals: []roomf{},
      counts: []int{},
    }
  }
  sort.Float64s(data)
  intervals := getIntervals(data, numbuckets)
  counts := make([]int, numbuckets, numbuckets)
  index := 0
  for _, d := range data {
    croom := intervals[index]
    if d > croom.ceil && index+1 < numbuckets {
      index++
    }
    counts[index]++
  }
  return HistData{intervals, counts}
}

func main() {
  data := []float64{}
  rand.Seed(time.Now().UTC().UnixNano())
  for i := 0; i < 500; i++ {
    data = append(data, rand.Float64())
  }
  hist := getHistData(data, 10)
  fmt.Println(buildHistString(HistChart{
    hist: hist, h: 10, bsize: 10,
  }))
}
