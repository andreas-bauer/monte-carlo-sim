package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	numWeeks := 100000
	numIslands := 10

	visited := simulate(numWeeks, numIslands)
	fmt.Println(visited)

	plotSim(visited)
}

func simulate(numWeeks int, numIslands int) []int {
	visited := make([]int, numIslands)

	current := 0

	for i := 0; i < numWeeks; i++ {
		// record current position
		visited[current]++

		// flip a coin and generate a proposal
		proposal := current + flipCoin()

		if proposal < 0 {
			proposal = numIslands - 1
		}
		if proposal >= numIslands {
			proposal = 0
		}

		// move if propability is higher then random number int the interval [0.0,1.0)
		propabilityMove := float32(proposal+1) / float32(current)
		if rand.Float32() < propabilityMove {
			current = proposal
		}
	}
	return visited
}

// flipCoin returns either 1 or -1
func flipCoin() int {
	return -1 + 2*rand.Intn(2)
}

func plotSim(visited []int) {
	p := plot.New()
	p.Title.Text = "Simulation of 1e5 weeks"
	p.Y.Label.Text = "number of weeks spent on island"
	p.X.Label.Text = "island"
	p.NominalX("1", " 2", "3", "4", "5", "6", "7", "8", "9", "10")

	ps := plotter.Values(asFloatArray(visited))
	width := vg.Points(20)

	bar, err := plotter.NewBarChart(ps, width)
	if err != nil {
		panic(err)
	}
	bar.Color = color.RGBA{127, 188, 165, 1}

	p.Add(bar)
	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}

func asFloatArray(a []int) []float64 {
	var retVal []float64 = make([]float64, len(a))
	for i, v := range a {
		retVal[i] = float64(v)
	}
	return retVal
}
