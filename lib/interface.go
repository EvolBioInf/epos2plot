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

	Files []string
}

const (
	progStr = "epos2plot"
	verStr  = "0.1"
	defL    = 0.025
	defU    = 0.975
)

func version() {
	fmt.Printf("%s %s\n", progStr, verStr)
	fmt.Printf("Written by Bernhard Haubold\n")
	fmt.Printf("Distributed under the GNU General Public License\n")
	fmt.Printf("Please send bug reports to haubold@evolbio.mpg.de.\n")
	os.Exit(2)
}

func usage() {
	fmt.Printf("Usage: %s [options]\n", progStr)
	fmt.Printf("Options:\n")
	fmt.Printf("\t[-l NUM lower quantile; default: %.3f]\n", defL)
	fmt.Printf("\t[-u NUM upper quantile; default: %.3f]\n", defU)
	fmt.Printf("\t[-h help]\n")
	fmt.Printf("\t[-v version]\n")
	os.Exit(2)
}

func ParseCL() Opts {
	var o Opts
	
	flag.Float64Var(&o.L, "l", defL,  "")
	flag.Float64Var(&o.U, "u", defU,  "")
	flag.BoolVar(   &o.V, "v", false, "")
	flag.Usage = usage
	flag.Parse()
	if o.V == true { version() }
	o.Files = flag.Args()
	
	return o
}
