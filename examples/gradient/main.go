package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	// a smooth wave whose height sweeps the full value range
	data := make([]float64, 0, 80)
	for x := 0; x < 80; x++ {
		data = append(data, math.Sin(float64(x)/8)*10)
	}

	// SeriesColorGradient colors each point by its value: low values use the
	// first stop, high values the last. HeatmapSpectrum is a built-in
	// cool-to-warm palette (blue → cyan → green → yellow → red).
	graph := asciigraph.Plot(data,
		asciigraph.Height(15),
		asciigraph.Precision(0),
		asciigraph.Caption("Heatmap gradient: low values cool, high values warm"),
		asciigraph.SeriesColorGradient(asciigraph.HeatmapSpectrum...),
	)

	fmt.Println(graph)
	// Output:
	//   10 ┤                                                              ╭╮
	//    9 ┤        ╭───────╮                                         ╭───╯╰───╮
	//    7 ┤      ╭─╯       ╰─╮                                     ╭─╯        ╰╮
	//    6 ┤     ╭╯           ╰╮                                   ╭╯           ╰─╮
	//    5 ┤   ╭─╯             ╰─╮                                ╭╯              ╰╮
	//    3 ┤  ╭╯                 ╰╮                             ╭─╯                ╰╮
	//    2 ┤ ╭╯                   ╰╮                           ╭╯                   ╰╮
	//    1 ┤╭╯                     ╰╮                         ╭╯                     ╰╮
	//   -1 ┼╯                       ╰╮                       ╭╯                       ╰╮
	//   -2 ┤                         ╰╮                     ╭╯                         ╰─╮
	//   -3 ┤                          ╰╮                   ╭╯                            ╰╮
	//   -5 ┤                           ╰─╮                ╭╯                              ╰
	//   -6 ┤                             ╰╮              ╭╯
	//   -7 ┤                              ╰╮           ╭─╯
	//   -9 ┤                               ╰─╮       ╭─╯
	//  -10 ┤                                 ╰───────╯
	//                     Heatmap gradient: low values cool, high values warm
}
