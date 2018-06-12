package main

import (
	"os"

	"github.com/evolbioinf/epos2plot/lib"
)

func run(f *os.File, opts lib.Opts) {
	data := lib.Read(f)
	quant := lib.Quantiles(data, opts)
	lib.Print(quant)
}

func main() {
	opts := lib.ParseCL()

	if len(opts.Files) == 0 {
		run(os.Stdin, opts)
	} else {
		for _, f := range opts.Files {
			fi, err := os.Open(f)
			lib.Check(err)
			run(fi, opts)
			err = fi.Close()
			lib.Check(err)
		}
	}
}
