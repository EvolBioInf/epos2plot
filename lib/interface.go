package lib

import (
	"flag"
	"fmt"
	"os"
)

type Opts struct {
	L float64
	U float64
	V bool
	R bool
	T float64
	F float64
	Files []string
}

const (
	progStr = "epos2plot"
	verStr  = "0.2"
	defL    = 0.025
	defU    = 0.975
	defT    = 0
	defF    = 0.5
)

func version() {
	fmt.Printf("%s %s\n", progStr, verStr)
	fmt.Printf("Written by Bernhard Haubold\n")
	fmt.Printf("Distributed under the GNU General Public License\n")
	fmt.Printf("Please send bug reports to haubold@evolbio.mpg.de.\n")
	os.Exit(2)
}

func usage() {
	fmt.Printf("Usage: %s [options] [files]\n", progStr)
	fmt.Printf("Convert epos output to quantiles ready for plotting.\n")
	fmt.Printf("Example: %s foo.epos\n", progStr)
	fmt.Printf("Options:\n")
	fmt.Printf("\t[-l NUM lower quantile; default: %.3f]\n", defL)
	fmt.Printf("\t[-u NUM upper quantile; default: %.3f]\n", defU)
	fmt.Printf("\t[-t NUM minimum time step; default: %d, i.e. use all steps]\n", defT)
	fmt.Printf("\t[-f NUM fraction of measurements for inclusion in output; default %.2f]\n", defF)
	fmt.Printf("\t[-r print raw output; default: quantiles]\n")
	fmt.Printf("\t[-h help]\n")
	fmt.Printf("\t[-v version]\n")
	os.Exit(2)
}

func ParseCL() Opts {
	var o Opts
	
	flag.Float64Var(&o.L, "l", defL,  "")
	flag.Float64Var(&o.U, "u", defU,  "")
	flag.Float64Var(&o.T, "t", defT,  "")
	flag.Float64Var(&o.F, "f", defF,  "")
	flag.BoolVar(   &o.V, "v", false, "")
	flag.BoolVar(   &o.R, "r", false, "")
	flag.Usage = usage
	flag.Parse()
	if o.V == true { version() }
	o.Files = flag.Args()
	
	return o
}
