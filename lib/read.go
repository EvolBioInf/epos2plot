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
func (p EposSlice) Less(i, j int) bool { return p[i].T < p[j].T }
func (p EposSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Read(f *os.File) []Epos {
	var ep []Epos
	var st float64  // start time for interval

	input := bufio.NewScanner(f)
	st = 0
	for input.Scan() {
		s := input.Text()
		if s[0:1] != "#" {
			arr := strings.Split(input.Text(), "\t")
			l, er := strconv.Atoi(arr[0])
			Check(er)
			t, er := strconv.ParseFloat(arr[1], 64)
			Check(er)
			n, er := strconv.ParseFloat(arr[2], 64)
			Check(er)
			if t > 0 && n > 0 {            // Skip negative population sizes
				e1 := newEpos(l, st, n, true)
				e2 := newEpos(l,  t, n, false)
				ep = append(ep, *e1)
				ep = append(ep, *e2)
				st = t
			}
		} else {
			st = 0
		}
	}
	sort.Sort(EposSlice(ep))
	
	return ep
}
