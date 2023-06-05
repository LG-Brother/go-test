package main

import "fmt"

func main() {
	ints := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	floats := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
		"d": 4.4,
		"e": 5.5,
	}
	ints["f"] = 1
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}
