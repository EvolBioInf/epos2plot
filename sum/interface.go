package sum

import (
	"flag"
	"fmt"
	"os"
)
var VERSION, DATE string
type Opts struct {
	V bool
	E bool
	Files []string
}

const (
	progStr = "sumPlot"
)

func version(date, ver string) {
	fmt.Printf("%s %s, %s\n", progStr, ver, date)
	fmt.Printf("Written by Bernhard Haubold\n")
	fmt.Printf("Distributed under the GNU General Public License\n")
	fmt.Printf("Please send bug reports to haubold@evolbio.mpg.de.\n")
	os.Exit(2)
}

func usage() {
	fmt.Printf("Usage: %s [options] [files]\n", progStr)
	fmt.Printf("Summarize several sets of epos2plot output.\n")
	fmt.Printf("Example: epos2plot *.epos | %s\n", progStr)
	fmt.Printf("Options:\n")
	fmt.Printf("\t[-e standard error of the mean; default: standard deviation]\n")
	fmt.Printf("\t[-h help]\n")
	fmt.Printf("\t[-v version]\n")
	os.Exit(2)
}

func ParseCL(date, ver string) Opts {
	var o Opts
	
	flag.BoolVar(&o.V, "v", false, "")
	flag.BoolVar(&o.E, "e", false, "")
	flag.Usage = usage
	flag.Parse()
	if o.V == true { version(date, ver) }
	o.Files = flag.Args()
	
	return o
}
