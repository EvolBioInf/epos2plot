package main

import (
	"os"
	"strings"

	"github.com/evolbioinf/epos2plot/plot"
)

func run(f *os.File, opts plot.Opts) {
	if opts.R {
		plot.PrintRaw(f)
	} else {
		data := plot.Read(f)
		quant := plot.Quantiles(data, opts)
		plot.PrintQuantiles(quant)
	}
}
var VERSION, DATE string
func main() {
	date := strings.Replace(DATE, "_", " ", -1)
	opts := plot.ParseCL(date, VERSION)
	if len(opts.Files) == 0 {
		run(os.Stdin, opts)
	} else {
		for _, f := range opts.Files {
			fi, err := os.Open(f)
			plot.Check(err)
			run(fi, opts)
			err = fi.Close()
			plot.Check(err)
		}
	}
}
