package main

import (
	"fmt"
	"math"
)

var Empty struct{}

var pentas = []float64{}
var pentamap = map[float64]struct{}{}

func main() {
	for i := 1; i < 9000000; i++ {
		if isPentagonal(float64(i)) {
			pentas = append(pentas, float64(i))
			pentamap[float64(i)] = Empty
		}
	}
	lowestDiff()
	fmt.Println("length of map is", len(pentamap))
}

func isPentagonal(x float64) bool {
	r := math.Sqrt(1+24*x) + 1
	return (math.Mod(r, 6.0) == 0)
}

func abs(i float64) float64 {
	if i < 0 {
		return -i
	}
	return i
}

func ptest(i, n float64) bool {
	if _, ok := pentamap[i+n]; ok {
		if _, ok := pentamap[abs(i-n)]; ok {
			return true
		}
	}
	return false
}

func lowestDiff() {
	for i := 0; i < len(pentas); i++ {
		for j := 1; j < len(pentas); j++ {
			pi, pj := pentas[i], pentas[j]
			if ptest(pi, pj) {
				fmt.Println(abs(pi - pj))
				return
			}
		}
	}
}
