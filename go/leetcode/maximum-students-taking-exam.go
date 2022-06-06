package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxStudents([][]byte{{'#', '.', '#', '#', '.', '#'},
		{'.', '.', '.', '.', '.', '.'},
		{'#', '.', '#', '.', '#', '#'}}))
	// fmt.Println(maxStudents([][]byte{{'#', '.', '#', '#', '.', '#'},
	// 	{'.', '#', '#', '#', '#', '.'},
	// 	{'#', '.', '#', '#', '.', '#'}}))

}

func maxStudents(seats [][]byte) int {
	type st struct {
		N int
		M int
		B uint
	}
	seatBits := make([]st, 0)
	preSeat := make([]st, 0)
	m, max := 0, 0
	for _, seat := range seats {
		seatBits = append(seatBits[:0], st{0, max, 0})
		for i, b := range seat {
			if b == '.' {
				for _, s := range seatBits {
					sb := s.B | 1<<i
					if sb<<1&sb == 0 && sb>>1&sb == 0 {
						m = 0
						for _, ps := range preSeat {
							if ps.B<<1&sb == 0 && ps.B>>1&sb == 0 {
								if ps.M+ps.N > m {
									m = ps.M + ps.N
								}
							}
						}
						if s.N+1+m > max {
							max = s.N + 1 + m
						}
						seatBits = append(seatBits, st{s.N + 1, m, sb})
					}
				}
			}
		}
		preSeat = append(preSeat[:0], seatBits...)
	}
	return max
}
