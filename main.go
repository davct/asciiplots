package main

import (
  "math/rand"
  "time"
  "fmt"
)

func main() {
	data := []float64{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100000; i++ {
		data = append(data, rand.Float64()*500)
	}
  hist := newHistogram(data, 10)
  plot := hist.MakePlot(14, 100)
  fmt.Println(plot)
}
