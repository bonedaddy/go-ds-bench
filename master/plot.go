package master

import (
	"fmt"
	"github.com/ipfs/go-ds-bench/options"
	"strings"

	"golang.org/x/tools/benchmark/parse"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func genplots(plotName string, pathPrefix string, bopts []options.BenchOptions, results map[string]map[int]*parse.Benchmark, x *xsel, y *ysel, yscale plot.Normalizer, ymarker plot.Ticker, suffix string) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = plotName
	p.Y.Label.Text = y.name
	p.X.Label.Text = x.name
	p.X.Scale = ZeroLogScale{}
	p.Y.Scale = yscale
	p.Legend.Top = true
	p.X.Tick.Marker = Log2Ticks{}
	p.Y.Tick.Marker = ymarker

	p.Add(plotter.NewGrid())

	var lp []interface{}
	for dsname, p := range results {
		pts := make(plotter.XYs, len(p))

		for n, bench := range p {
			pts[n].X = x.sel(bopts[n])
			pts[n].Y = y.sel(bench)
		}
		lp = append(lp, dsname, pts)
	}

	if err := plotutil.AddLinePoints(p, lp...); err != nil {
		return err
	}

	plotName = fmt.Sprintf("%s-%s-%s%s.png", plotName, x.name, y.name, suffix)
	plotName = strings.Replace(plotName, "/", "", -1)
	return p.Save(8*vg.Inch, 6*vg.Inch, pathPrefix+plotName)
}
