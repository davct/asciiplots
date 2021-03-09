package main

import (
  "strconv"
  "sort"
)


func minmaxi(l []int) (min int, max int) {
	min, max = l[0], l[0]
	for _, n := range l {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func minmaxf(l []float64) (min float64, max float64) {
	min, max = l[0], l[0]
	for _, n := range l {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func minmaxls(strings []string) (minl, maxl int) {
	if len(strings) == 0 {
		return 0, 0
	}
	l := len(strings[0])
	maxl = l
	minl = l
	for _, s := range strings {
		l = len(s)
		if l < minl {
			minl = l
		}
		if l > maxl {
			maxl = l
		}
	}
	return minl, maxl
}

func intsToStrings(is []int) (rs []string) {
	for _, i := range is {
		rs = append(rs, strconv.Itoa(i))
	}
	return rs
}

func truncatef(f float64, length int) float64 {
	transferpow := 1.0
	for i := 1; i < length; i++ {
		transferpow *= 10
	}
	return float64(int(f*transferpow)) / transferpow
}

func getSortedCopy(data []float64) (copied []float64) {
  for _, e := range data {
    copied = append(copied, e)
  }
  sort.Float64s(copied)
  return copied
}
