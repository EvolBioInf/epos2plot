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
	N int      // Sample size
}

func quant(ps map[string]int, opts Opts) (float64, float64, float64, int) {
	var sizes []float64
	
	for n, c := range ps {
		x, e := strconv.ParseFloat(n, 64)
		Check(e)
		for i := 0; i < c; i++ {
			sizes = append(sizes, x)
		}
	}
	if len(sizes) == 0 { return -1, -1, -1, -1}
	sort.Float64s(sizes)
	i1 := int(float64(len(sizes)) * opts.L)
	i2 := int(float64(len(sizes)) * 0.5)
	i3 := int(float64(len(sizes)) * opts.U)
	
	return sizes[i1], sizes[i2], sizes[i3], len(sizes)
}

func newQuantile(t, l, m, u float64, n int) *Quantile {
	q := new(Quantile)
	q.T = t
	q.L = l
	q.M = m
	q.U = u
	q.N = n

	return q
}

func Quantiles(data []Epos, opts Opts) []Quantile {
	var qu []Quantile
	var i, ii int
	var n string
	ps := make(map[string]int) // Map of population sizes as strings

	// Collect measurements at time zero
	for i = 0; i < len(data) && data[i].T == 0; i++ {
		n = strconv.FormatFloat(data[i].N, 'e', -1, 64)
		ps[n]++
	}
	l, m, u, s := quant(ps, opts)
	q := newQuantile(0, l, m, u, s)
	qu = append(qu, *q)
	for ii = i; ii < len(data); ii++ {
		l, m, u, s = quant(ps, opts)
		q = newQuantile(data[ii].T, l, m, u, s)
		qu = append(qu, *q)
		n = strconv.FormatFloat(data[ii].N, 'e', -1, 64)
		if data[ii].S == true {
			ps[n]++
		} else          {
			ps[n]--
		}
	}

	return qu
}
