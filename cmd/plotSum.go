package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/evolbioinf/epos2plot/sum"
	"github.com/evolbioinf/epos2plot/plot"
)

func run(f *os.File, opts sum.Opts) {
	var i int
	rd := sum.Read(f)
	fd := sum.Analysis(rd, opts)
	fmt.Printf("#Pos\tMean\t")
	if opts.E  {
		fmt.Printf("SEM\n")
	} else {
		fmt.Printf("SD\n")
	}
	fmt.Printf("%.0f\t%.2f\t%.2f\n", fd.T[0], fd.M[0], fd.S[0])
	for i = 1; i < len(fd.T) - 1; i++ {
		if fd.M[i] != fd.M[i-1] {
			fmt.Printf("%.0f\t%.2f\t%.2f\n", fd.T[i], fd.M[i], fd.S[i])
		}
	}
	fmt.Printf("%.0f\t%.2f\t%.2f\n", fd.T[i], fd.M[i], fd.S[i])
}
var VERSION, DATE string
func main() {
	date := strings.Replace(DATE, "_", " ", -1)
	opts := sum.ParseCL(date, VERSION)
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
