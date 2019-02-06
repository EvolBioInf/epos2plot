package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func PrintQuantiles(qu []Quantile, t float64) {
	fmt.Printf("#Time\tLowerQ\tMedian\tUpperQ\tUncoalescedSamples\n")
	mi := int(math.Round(float64(qu[0].N) * t)) // minimun number of measurements for consideration in output
	for i, q := range qu {
		if q.N >= mi && (i == 0 || qu[i].T != qu[i-1].T) {
			fmt.Printf("%g\t%g\t%g\t%g\t%d\n", q.T, q.L, q.M, q.U, q.N)
		} 
	}
}

func PrintRaw(f *os.File) {
	var st float64  // Start time for interval.
	var first bool
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if s[0:1] == "#" {    // Reset start time.
			st = 0
			first = true
		} else {
			arr := strings.Split(input.Text(), "\t")
			t, er := strconv.ParseFloat(arr[1], 64)
			Check(er)
			n, er := strconv.ParseFloat(arr[2], 64)
			Check(er)
			if st > 0 && t > 0 && n > 0 {                      // Skip negative population sizes.
				if first {
					first = false
					fmt.Printf("\n")
				}
				fmt.Printf("%g\t%g\n", st, n)
				fmt.Printf("%g\t%g\n", t,  n)
			}
			st = t
		} 
	}
}
