package plot

import (
	"fmt"

	"github.com/Nimsaja/PortfolioPerformance/store"
)

//Create create Graph for values and diffs
func Create(data []store.Data) {
	x, y, d := getXYData(data)

	CreatePlot(x, y, "Depot Values", "Date", "Values")
	CreatePlot(x, d, "Diff Values", "Date", "Diff")
}

func getXYData(data []store.Data) (x []int, y []float32, ydiff []float32) {
	x = make([]int, len(data))
	y = make([]float32, len(data))
	ydiff = make([]float32, len(data))

	for i, d := range data {
		x[i] = d.Time
		y[i] = d.Value
		ydiff[i] = d.Diff
	}

	fmt.Println(x)

	return
}
