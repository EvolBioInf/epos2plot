package sum

import (
	"fmt"
	"math"
	"os"
)

type Row []float64

type Result struct {
	T []float64 // time
	M []float64 // means
	S []float64 // standard error
}

func newResult() *Result {
	r := new(Result)
	r.T = make([]float64, 0)
	r.M = make([]float64, 0)
	r.S = make([]float64, 0)
	return r
}

// ms returns mean and standard error of numbers in d
func ms(d []float64, opts Opts) (float64, float64) {
	n := len(d)
	var sx, sxx float64
	for i := 0; i < n; i++ {
		x := d[i]
		sx += x
		sxx += x * x
		
	}
	v := (sxx - sx * sx / float64(n)) / float64(n - 1)
	m := sx / float64(n)
	e := math.Sqrt(float64(v))
	if opts.E {
		e /= math.Sqrt(float64(n))
	}
	return m, e
}

func printMat(data []Row) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("%.2f", data[i][0])
		for j := 1; j < len(data[i]); j++ {
			fmt.Printf("\t%.2f", data[i][j])
		}
		fmt.Printf("\n")
	}
}

func Analysis(ds DataSlice, opts Opts) Result {
	var n, m int

	// construct m x n data matrix
	m = 1
	n = 1
	r := newResult()
	r.T = append(r.T, ds[0].T)
	ds[0].R = 0
	// determine m & n
	for i := 1; i < len(ds); i++ {
		if ds[i].T != ds[i-1].T {
			r.T = append(r.T, ds[i].T)
			m++
		}
		ds[i].R = m - 1
		if n - 1 < ds[i].C {
			n = ds[i].C + 1
		}
	}
	
	dm := make([]Row, m)
	for i := 0; i < m; i++ {
		dm[i] = make([]float64, n)
	}
	for i := 0; i < len(ds); i++ {
		dm[ds[i].R][ds[i].C] = ds[i].M
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if dm[j][i] == 0 {
				dm[j][i] = dm[j - 1][i]
			}
		}
	}
	if opts.M {
		printMat(dm)
		os.Exit(0)
	}
	// calculate means & standard errors
	for i := 0; i < m; i++ {
		x, y := ms(dm[i], opts)
		r.M = append(r.M, x)
		r.S = append(r.S, y)
	}

	return *r
}
