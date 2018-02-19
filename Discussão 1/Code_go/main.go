package main

import (
	"fmt"
)

// pontos da coordenada
var (
	xPoints = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	yPoints = []float64{1, 2, 3, 1, 4, 5, 4, 6, 7, 10, 15}
)

func main() {

	learn := 0.001
	m := 0.0
	b := 0.0
	y := func(x float64) float64 {
		return m*x + b
	}
	plotLine(y, xPoints)
	fmt.Println("-------------------------")
	for i := 0; i < 80; i++ {
		s1, s2 := summation(y, xPoints, yPoints)
		m = m - learn*s2
		b = b - learn*s1
		plotLine(y, xPoints)
		fmt.Println("-------------------------")
	}

	fmt.Println("m:", m)
	fmt.Println("b:", b)

}

func listPoints(xValues, yValues []float64) {
	for i := 0; i < len(xValues); i++ {
		fmt.Println("x: ", xValues[i], "y: ", yValues[i])
	}
}

func max(data []float64) (ret float64) {
	for _, i := range data {
		if i > ret {
			ret = i
		}
	}
	return
}

func min(data []float64) (ret float64) {
	ret = 0xFFFFFFFF
	for _, i := range data {
		if i < ret {
			ret = i
		}
	}
	return
}

func plotLine(y func(float64) float64, data []float64) {
	var xValues []float64
	var yValues []float64
	for i := min(data) - 1; i < max(data)+2; i++ {
		xValues = append(xValues, i)
		yValues = append(yValues, y(i))

		fmt.Println("x:", xValues[int(i)], "y:", yValues[int(i)])
	}
}

func summation(y func(float64) float64, xValues, yValues []float64) (float64, float64) {
	total1 := 0.0
	total2 := 0.0
	for i := 1; i < len(xValues); i++ {
		total1 += y(xValues[i]) - y(yValues[i])
		total2 += (y(yValues[i]) - yValues[i]) * xValues[i]
	}
	return total1 / float64(len(xValues)), total2 / float64(len(xValues))
}
