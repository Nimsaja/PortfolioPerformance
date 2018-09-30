package plot

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PortfolioPerformance/store"
)

//Create create Graph for values and diffs
func Create(data []store.Data, value, buySum float32) {
	x, y, d := getXYData(data, value, buySum)

	CreatePlot(x, y, "Depot Values", "Date", "Values")
	CreatePlot(x, d, "Diff Values", "Date", "Diff")
}

func getXYData(data []store.Data, valueToday, buySum float32) (x []int, y []float32, ydiff []float32) {
	x = make([]int, len(data)+1)
	y = make([]float32, len(data)+1)
	ydiff = make([]float32, len(data)+1)

	for i, d := range data {
		fmt.Println("add data time ", d.Time)
		x[i] = d.Time
		y[i] = d.Value
		ydiff[i] = d.Value - buySum
	}

	fmt.Println("add cur time ", int(time.Now().Unix()))
	x[len(data)] = int(time.Now().Unix())
	y[len(data)] = valueToday
	ydiff[len(data)] = valueToday - buySum

	fmt.Println(x)

	return
}
