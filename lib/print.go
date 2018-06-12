package lib

import "fmt"

func Print(qu []Quantile) {
	fmt.Printf("#t\tl\tm\tu\n")
	for _, q := range qu {
		if q.L > -1 {
			fmt.Printf("%f\t%f\t%f\t%f\n", q.T, q.L, q.M, q.U)
		}
	}
}
