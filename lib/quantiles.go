package lib

import (
	"strconv"
	"sort"
)
type Quantile struct {
	T float64  // Time
	L float64  // Lower quantile
	M float64  // Median
	U float64  // Upper quantile
}

func quant(ps map[string]int, opts Opts) (float64, float64, float64) {
	var sizes []float64
	
	for n, c := range ps {
		x, e := strconv.ParseFloat(n, 64)
		Check(e)
		for i := 0; i < c; i++ {
			sizes = append(sizes, x)
		}
	}
	if len(sizes) == 0 { return -1, -1, -1 }
	sort.Float64s(sizes)
	i1 := int(float64(len(sizes)) * opts.L)
	i2 := int(float64(len(sizes)) * 0.5)
	i3 := int(float64(len(sizes)) * opts.U)
	
	return sizes[i1], sizes[i2], sizes[i3]
}

func newQuantile(t, l, m, u float64) *Quantile {
	q := new(Quantile)
	q.T = t
	q.L = l
	q.M = m
	q.U = u

	return q
}

func Quantiles(data []Epos, opts Opts) []Quantile {
	var qu []Quantile
	ps := make(map[string]int) // Map of population sizes as strings
	
	for _, ep := range data {
		n := strconv.FormatFloat(ep.N, 'e', -1, 64)
		if ep.T == 0 {
			ps[n]++
		} else {
			l, m, u := quant(ps, opts)
			q := newQuantile(ep.T, l, m, u)
			qu = append(qu, *q)
			if ep.S == true {
				ps[n]++
			} else          {
				ps[n]--
			}
			l, m, u = quant(ps, opts)
			q = newQuantile(ep.T, l, m, u)
			qu = append(qu, *q)
		}
	}

	return qu
}
