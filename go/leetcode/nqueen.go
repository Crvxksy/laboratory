package main

import "fmt"

func main() {
	queens := solveNQueens2(8)
	fmt.Println(len(queens), queens)
}

func solveNQueens2(n int) (res [][]string) {
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	halfN := (n - 1) / 2
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			hplan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
				hplan[i] = placements[n-1-prev[i]]
			}
			res = append(res, plan)
			if prev[0] != (n - 1 - prev[0]) {
				res = append(res, hplan)
			}
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit<<dist | bit>>dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied>>i)&1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			if prev[0] <= halfN {
				construct(prev)
			}
		}
	}
	construct(make([]int, 0, n))
	return
}
