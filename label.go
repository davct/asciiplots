package main

import (
	"strconv"
)

type AxisLabel string

func MakeAxisLabel(label float64, mlen int) AxisLabel {
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
	return AxisLabel(formatted)
}
