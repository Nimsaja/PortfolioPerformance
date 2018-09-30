package plot

import (
	"fmt"
	"log"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//CreatePlot create x y plot - x is time, y is float32
func CreatePlot(xdata []int, ydata []float32, title, xLabel, yLabel string) {
	if len(xdata) != len(ydata) {
		log.Println("Length of x data and y data should be the same")
		return
	}

	xticks := plot.TimeTicks{Format: "2006-01-02\n15:04"}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel
	p.X.Tick.Marker = xticks

	pts := make(plotter.XYs, len(xdata))
	for i, v := range ydata {
		pts[i].X = float64(xdata[i])
		pts[i].Y = float64(v)
	}

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	path := strings.Replace(title, " ", "", -1) + ".png"
	if err := p.Save(10*vg.Inch, 4*vg.Inch, path); err != nil {
		panic(err)
	}

	fmt.Println("\n*****Please open ", path, " ********")
}
