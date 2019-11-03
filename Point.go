package kmeans

import (
	"fmt"
	"math"
)

type Point []float64

func (point Point) print() {

	for i, num := range point {

		if i == len(point)-1 {
			fmt.Printf("%f", num)
		} else {
			fmt.Printf("%f,", num)
		}
	}
	fmt.Printf("\n")
}

func (point Point) pointDist(p2 Point) float64 {
	var sum float64
	for i := 0; i < len(point); i++ {
		if point[i] == p2[i] {
			continue
		}
		sum += math.Pow(point[i]-p2[i], 2.0)
	}
	return math.Sqrt(sum)
}

func (point Point) pointEqual(p2 Point) bool {
	if len(point) != len(p2) {
		return false
	}
	for i := 0; i < len(point); i++ {
		if point[i] != p2[i] {
			return false
		}
	}
	return true
}

func (point Point) subtract(p2 Point) Point {
	for i := range point {
		point[i] -= p2[i]
	}
	return point
}

func (point Point) norm() float64 {
	norm := 0.0
	for _, p := range point {
		norm += math.Pow(p, 2)
	}
	return math.Sqrt(norm)
}

func equals(points1 []Point, points2 []Point) bool {
	if len(points1) != len(points2) {
		return false
	}
	for i := range points1 {
		if !points1[i].pointEqual(points2[i]) {
			return false
		}
	}
	return true
}
