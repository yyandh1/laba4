package main

import (
	"fmt"
	"math"
	"math"
)

func inputTable(LeftPoint, RightPoint float64, iteration int) {
	if iteration == 1 {
		fmt.Printf("N\t     a \t\t\t     b \t\t\t     b-a \n")
		fmt.Println("---------------------------------------------------------------------")
	}
	fmt.Printf("%d |\t%10.6f\t\t%10.6f\t\t%12.6f\n", iteration, LeftPoint, RightPoint, RightPoint-LeftPoint)
}

func function(x float64) float64 {
	return 5*math.Sin(x) - x
}

func df(x float64) float64 {
	return 5*math.Cos(x) - 1
}

func findGraficalSolution(left, right, x0 *float64) {
	for x := -1.0; x < 5.0; x += 0.01 {
		if math.Ceil(math.Sin(x)-(-2*x-1)) == 0 {
			*left = x - 1.0
			*right = x + 1.0
			*x0 = x
		}
	}
}

func chordMethod(a, b, epsilon float64) float64 {
	for i := 1; math.Abs(b-a) >= epsilon; i++ {
		a = a - (b-a)*function(a)/(function(b)-function(a))
		b = b - (a-b)*function(b)/(function(a)-function(b))
		inputTable(a, b, i)
	}
	return b
}

func NewtownMethod(x0, epsilon float64) float64 {
	var x float64
	for i := 1; math.Abs(function(x0)) >= epsilon && i < 10; i++ {
		x = x0 - function(x0)/df(x0)
		inputTable(x, x0, i)
		x0 = x
	}
	return x
}

func halfDivisionMethod(LeftPoint, RightPoint, epsilon float64) float64 {
	iteration := 1
	var midPoint float64

	if function(LeftPoint)*function(RightPoint) < 0 {
		for math.Abs(RightPoint-LeftPoint) > epsilon {
			midPoint = (RightPoint + LeftPoint) / 2
			inputTable(LeftPoint, RightPoint, iteration)

			if function(LeftPoint)*function(midPoint) < 0 {
				RightPoint = midPoint
			} else {
				LeftPoint = midPoint
			}
			iteration++
		}
	}
	return midPoint
}

func main() {
	var left, right, x0 float64
	findGraficalSolution(&left, &right, &x0)

	eps := 0.0001

	fmt.Printf("Root through the chord method:\n%f\n\n", chordMethod(left, right, eps))
	fmt.Printf("The root through the method of half division:\n%f\n\n", halfDivisionMethod(left, right, eps))
	fmt.Printf("The root through the Newtown method:\n%f\n", NewtownMethod(x0, eps))

	fmt.Println("\nBy the rate of convergence")
	fmt.Println("Newtown method > Chord method > Half divisions")
}


