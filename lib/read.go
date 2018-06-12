package lib

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Epos struct {
	L int     // Level
	T float64 // Time
	N float64 // Population size (Ne)
	S bool    // Start
}

func newEpos(l int, t, n float64, s bool) *Epos {
	e := new(Epos)
	e.L = l
	e.T = t
	e.N = n
	e.S = s

	return e
}

type EposSlice []Epos
func (p EposSlice) Len() int           { return len(p) }
func (p EposSlice) Less(i, j int) bool { // When tied, let end < start
	var t1, t2 float64
	t1 = p[i].T
	t2 = p[j].T
	if p[i].S == true { t1++ }
	if p[j].S == true { t2++ }
	return t1 < t2
}
func (p EposSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Read(f *os.File) []Epos {
	var ep []Epos
	var coal []Epos
	var st float64  // Start time for interval.

	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if s[0:1] == "#" {    // Reset start time.
			st = 0
			for _, c := range coal {
				ep = append(ep, c)
			}
			coal = make([]Epos, 0)
		} else {
			arr := strings.Split(input.Text(), "\t")
			l, er := strconv.Atoi(arr[0])
			Check(er)
			t, er := strconv.ParseFloat(arr[1], 64)
			Check(er)
			n, er := strconv.ParseFloat(arr[2], 64)
			Check(er)
			if t > 0 && n > 0  && coal != nil {      // Skip entire coalescent with negative pop sizes
				e1 := newEpos(l, st, n, true)    // Start at previous time, st.
				e2 := newEpos(l,  t, n, false)   // End at current time, t.
				coal = append(coal, *e1)
				coal = append(coal, *e2)
				st = t
			} else {
				coal = nil
			}
		} 
	}
	for _, c := range coal {
		ep = append(ep, c)
	}
	sort.Sort(EposSlice(ep))
	
	return ep
}
