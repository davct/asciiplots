package main

import "testing"

func TestAreFloatsEqual(t *testing.T) {
  first, second, epsilon := 1.01, 1.029, 0.02
  if areFloatsEqual(first, second, epsilon) != true {
    t.Errorf("Expected 'areFloatsEqual' to say that %f == %f with epsilon=%f", first, second, epsilon)
  }

  first, second, epsilon = 1.01, 1.01, 0.000001
  if areFloatsEqual(first, second, epsilon) != true {
    t.Errorf("Expected 'areFloatsEqual' to say that %f == %f with epsilon=%f", first, second, epsilon)
  }

  first, second, epsilon = 1.00001, 1.00001, 0.000001
  if areFloatsEqual(first, second, epsilon) != true {
    t.Errorf("Expected 'areFloatsEqual' to say that %f == %f with epsilon=%f", first, second, epsilon)
  }
}

func TestAreFloatSlicesEqual(t *testing.T) {
  first := []float64{1.0, 2.0, 3.01}
  second := []float64{1.0091, 1.991, 3.01}
  epsilon := 0.01
  equal, i := areFloatSlicesEqual(first, second, epsilon)
  if equal != true {
    t.Errorf("Failed on float slices compare test 1\n")
    t.Errorf("Thought that lists differ at %d, i.e. %f != %f with epsilon %f", i, first[i], second[i], epsilon)
  }

  first = []float64{1.0, 2.0, 3.01}
  second = []float64{1.1, 2.1, 3.01}
  epsilon = 0.01
  equal, i = areFloatSlicesEqual(first, second, epsilon)
  if equal != false || i != 0 {
    t.Errorf("Failed on float slices compare test 2\n")
  }
}
