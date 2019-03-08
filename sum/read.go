package sum

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"github.com/evolbioinf/epos2plot/plot"
)

type Data struct {
	T float64 // position
	M float64 // median
	R int     // row
	C int     // column
}

func newData(t, m float64, r, c int) *Data {
	d := new(Data)
	d.T = t
	d.M = m
	d.R = r
	d.C = c
	
	return d
}

type DataSlice []Data
func (p DataSlice) Len() int { return len(p) }
func (p DataSlice) Less(i, j int) bool { return p[i].T < p[j].T }
func (p DataSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Read(f *os.File) DataSlice {
	var n int

	input := bufio.NewScanner(f)
	ds := make(DataSlice, 0)
	for input.Scan() {
		s := input.Text()
		if s[0:1] == "#" { // New data set
			n++
		} else {
			arr := strings.Split(input.Text(), "\t")
			t, er := strconv.ParseFloat(arr[0], 64) // Time
			plot.Check(er)
			m, er := strconv.ParseFloat(arr[2], 64) // Median
			plot.Check(er)
			d := newData(t, m, 0, n - 1)
			ds = append(ds, *d)
		}
	}
	sort.Sort(DataSlice(ds))

	return ds
}
