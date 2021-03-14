package main

//import "fmt"

func areFloatsEqual(first float64, second float64, epsilon float64) bool {
  if epsilon < 0 {
    epsilon = -epsilon
  }

  var diff float64
  if first > second {
    diff = first - second
  } else {
    diff = second - first
  }
  return diff < epsilon
}

func areFloatSlicesEqual(first []float64, second []float64, epsilon float64) (equal bool, firstDifference int)  {
  equal = true
  firstDifference = -1
  for i, _ := range first {
    if !areFloatsEqual(first[i], second[i], epsilon) {
      equal = false
      firstDifference = i
      break
    }
  }
  return equal, firstDifference
}
