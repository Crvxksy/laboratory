package main

import "fmt"

func main() {
	pieces := [][]int{{3, 4, 0}, {5, 3, 1}, {5, 8, 0}, {5, 5, 0}, {1, 8, 0}, {7, 8, 0}, {5, 7, 0}, {9, 3, 1}, {8, 3, 1}, {3, 0, 1}, {9, 9, 1}, {9, 0, 1}, {7, 1, 0}, {4, 3, 1}, {6, 6, 1}, {3, 2, 1}, {3, 3, 0}, {3, 6, 1}, {0, 9, 0}, {7, 5, 1}, {0, 0, 0}, {3, 5, 1}, {3, 8, 1}, {8, 1, 1}, {2, 4, 0}, {0, 7, 1}, {1, 7, 1}, {5, 6, 1}, {2, 9, 0}, {8, 0, 0}, {8, 5, 0}, {1, 4, 0}, {0, 1, 1}, {9, 6, 1}, {1, 0, 0}, {5, 2, 1}, {2, 1, 1}, {5, 4, 1}, {3, 9, 0}, {1, 5, 0}, {8, 6, 0}, {5, 1, 1}, {7, 2, 1}, {9, 2, 0}, {2, 3, 1}, {9, 7, 0}, {2, 2, 1}, {6, 5, 0}, {0, 8, 0}, {7, 0, 1}, {7, 7, 1}, {8, 4, 0}, {3, 1, 0}}
	print(gobang(pieces), " ") // W
	// printChess(pieces)

	pieces = [][]int{{0, 0, 1}, {0, 1, 0}, {0, 3, 0}, {0, 4, 0}, {0, 7, 0}, {0, 8, 1}}
	print(gobang(pieces), " ") // B

	pieces = [][]int{{813382118, -172692001, 0}, {813382118, -172692000, 0}, {813382118, -172691997, 0}, {813382118, -172691996, 0},
		{813382118, -172691993, 0}, {813382118, -172691992, 0}, {813382118, -172691989, 0}, {813382118, -172691988, 0}, {813382118, -172691985, 0},
	}
	print(gobang(pieces), " ") // N

	pieces = [][]int{{1, 9, 1}, {3, 2, 1}, {6, 9, 0}, {1, 5, 0}, {4, 7, 1}, {5, 7, 1}, {7, 5, 0}, {6, 4, 1}, {2, 7, 1},
		{0, 3, 0}, {3, 4, 0}, {8, 4, 0}, {0, 6, 0}, {8, 5, 1}, {3, 3, 1}, {2, 8, 1}, {4, 4, 0}, {5, 9, 0}, {0, 1, 1},
		{5, 1, 1}, {6, 3, 0}, {9, 4, 1}, {2, 6, 0}, {0, 8, 1}, {7, 3, 0}, {9, 5, 0}, {9, 2, 0}, {8, 3, 0}, {9, 0, 0},
		{8, 9, 0}, {4, 3, 1}, {2, 3, 1}, {7, 4, 1}, {2, 1, 0}, {6, 5, 1}, {8, 8, 0}, {6, 1, 1}, {6, 6, 1}, {1, 3, 1},
		{4, 6, 1}, {7, 1, 1}, {5, 5, 0}, {2, 5, 0}}
	print(gobang(pieces), " ") // N

	pieces = [][]int{{2, 4, 1}, {0, 9, 1}, {1, 7, 0}, {7, 1, 1}, {8, 7, 1}, {0, 1, 1},
		{1, 2, 1}, {3, 3, 1}, {3, 2, 1}, {5, 3, 0}, {3, 7, 0}, {7, 0, 0}, {9, 9, 0},
		{2, 5, 1}, {1, 3, 1}, {6, 2, 1}, {7, 3, 0}, {1, 6, 1}, {4, 4, 1}, {2, 7, 0},
		{4, 5, 1}, {7, 9, 1}, {2, 6, 0}, {8, 4, 1}, {3, 8, 0}, {8, 9, 0}, {9, 5, 1},
		{7, 2, 0}, {4, 1, 1}, {0, 3, 0}, {7, 7, 0}, {1, 4, 1}, {7, 6, 0}, {0, 6, 0},
		{4, 9, 0}, {8, 3, 1}, {5, 0, 0}, {4, 8, 0}, {2, 2, 0}, {6, 0, 1}}
	print(gobang(pieces), " ") // B

	pieces = [][]int{{0, 2, 1}, {4, 3, 0}, {5, 5, 1}, {7, 9, 0}, {3, 9, 1}, {3, 8, 0},
		{8, 8, 0}, {4, 6, 0}, {0, 8, 1}, {8, 6, 1}, {2, 1, 0}, {3, 6, 1}, {1, 0, 1},
		{5, 3, 0}, {7, 6, 1}, {7, 7, 1}, {4, 4, 1}, {6, 9, 1}, {4, 8, 1}, {2, 2, 1},
		{8, 0, 1}, {6, 6, 1}, {0, 0, 1}, {4, 7, 1}, {3, 7, 1}, {4, 9, 1}, {6, 2, 1},
		{5, 9, 0}, {1, 7, 0}, {9, 1, 0}, {5, 8, 0}}

	print(gobang(pieces), " ") // N

	pieces = [][]int{{4, 1, 1}, {0, 2, 0}, {9, 2, 1}, {7, 0, 0}, {0, 5, 0}, {4, 6, 1}, {9, 9, 0},
		{2, 5, 0}, {4, 3, 1}, {0, 6, 1}, {9, 6, 1}, {9, 1, 0}, {2, 1, 0}, {5, 3, 0}, {5, 4, 1}, {1, 3, 0},
		{3, 3, 0}, {5, 5, 1}, {4, 7, 1}, {5, 2, 1}, {2, 9, 1}, {3, 1, 1}, {1, 8, 0}, {7, 6, 1}, {9, 7, 0},
		{8, 8, 1}, {0, 7, 1}, {1, 5, 1}, {4, 2, 0}, {6, 0, 1}, {5, 9, 0}, {6, 2, 0}, {4, 0, 0}, {4, 5, 0},
		{3, 2, 0}, {8, 6, 1}, {2, 8, 0}, {6, 6, 1}, {0, 9, 0}, {6, 3, 0}, {5, 0, 0}, {7, 8, 0}, {0, 1, 1},
		{5, 6, 0}, {7, 4, 0}, {8, 1, 1}, {2, 7, 1}, {2, 4, 0}, {3, 4, 1}, {8, 0, 0}, {8, 3, 0}, {6, 8, 0},
		{4, 8, 0}, {1, 6, 0}, {8, 2, 0}}
	print(gobang(pieces), " ") // N

	pieces = [][]int{
		{0, 2, 1}, {0, 3, 0}, {0, 4, 1}, {0, 5, 1}, {0, 9, 0},
		{1, 2, 1}, {1, 3, 0},
		{2, 0, 0}, {2, 1, 1}, {2, 2, 1}, {2, 3, 0}, {2, 4, 1}, {2, 8, 0},
		{3, 0, 0}, {3, 6, 1}, {3, 7, 1}, {3, 9, 1},
		{4, 0, 1}, {4, 1, 0}, {4, 3, 1}, {4, 4, 1}, {4, 6, 0}, {4, 8, 1},
		{5, 0, 1}, {5, 1, 1}, {5, 2, 1}, {5, 3, 0}, {5, 5, 1}, {5, 6, 1}, {5, 8, 0},
		{6, 1, 1}, {6, 2, 0}, {6, 4, 1}, {6, 6, 0}, {6, 8, 1},
		{7, 2, 1}, {7, 3, 1}, {7, 5, 1},
		{8, 0, 1}, {8, 1, 0}, {8, 2, 0}, {8, 4, 1}, {8, 5, 0}, {8, 9, 1},
		{9, 3, 1}, {9, 6, 1}, {9, 7, 0},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{0, 3, 1}, {0, 6, 0}, {0, 7, 0}, {0, 9, 0},
		{1, 0, 0}, {1, 7, 1},
		{2, 0, 1}, {2, 1, 1}, {2, 2, 0}, {2, 3, 0}, {2, 5, 1}, {2, 6, 0}, {2, 7, 0},
		{3, 0, 0}, {3, 2, 1}, {3, 3, 0}, {3, 4, 1}, {3, 5, 1}, {3, 7, 1}, {3, 8, 0},
		{4, 2, 1}, {4, 3, 0}, {4, 5, 0},
		{5, 0, 0}, {5, 5, 0}, {5, 6, 1}, {5, 8, 0}, {5, 9, 1},
		{6, 0, 1}, {6, 1, 1}, {6, 2, 1}, {6, 5, 0}, {6, 6, 1}, {6, 7, 0}, {6, 9, 1},
		{7, 1, 1}, {7, 2, 1}, {7, 4, 0}, {7, 9, 1},
		{8, 1, 1}, {8, 3, 1}, {8, 4, 0}, {8, 5, 1}, {8, 8, 0}, {8, 9, 0},
		{9, 0, 1}, {9, 7, 1}}
	print(gobang(pieces), " ") // N

	pieces = [][]int{
		{0, 2, 0},
		{0, 3, 0},
		{0, 4, 0},
		{0, 5, 1},
		{1, 1, 1},
		{2, 0, 0},
		{2, 2, 1},
		{3, 0, 0},
		{3, 3, 1},
		{4, 0, 0},
		{4, 4, 1},
		{5, 0, 1},
		{5, 5, 0},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{0, 2, 0},
		{0, 3, 0},
		{0, 7, 0},
		{0, 8, 0},
		{0, 9, 0},
		{0, 10, 1},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{0, -2, 1},
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{0, 4, 1},
		{1, 4, 0},
		{2, 4, 0},
		{4, 4, 0},
	}
	print(gobang(pieces), " ") // N

	pieces = [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{0, 2, 0},
		{0, 3, 0},
		{0, 7, 0},
		{0, 8, 0},
		{0, 9, 0},
		{0, 10, 1},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{0, 0, 0},
		{0, 6, 0},
		{1, 1, 0},
		{1, 5, 0},
		{3, 0, 1},
		{3, 1, 1},
		{3, 2, 1},
		{3, 4, 1},
		{4, 2, 0},
		{4, 4, 0},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{10, 10, 1},
		{8, 8, 1},
		{7, 7, 1},
	}
	print(gobang(pieces), " ") // N

	pieces = [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{0, 3, 0},
		{0, 4, 0},
		{0, 7, 0},
		{0, 8, 0},
	}
	print(gobang(pieces), " ") // B

	pieces = [][]int{
		{0, -2, 1},
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{0, 4, 1},
		{1, 4, 0},
		{2, 4, 0},
		{4, 4, 0},
	}
	print(gobang(pieces), " ") // N

	pieces = [][]int{
		{1, 2, 1},
		{1, 4, 1},
		{1, 5, 1},
		{2, 1, 0},
		{2, 3, 0},
		{2, 4, 0},
		{3, 2, 1},
		{3, 4, 0},
		{4, 2, 1},
		{5, 2, 1},
	}
	print(gobang(pieces), " ") // B
}

func gobang(pieces [][]int) string {
	posMap := make(map[[2]int]uint8, len(pieces))
	black := make([]*[2]int, 0)
	white := make([]*[2]int, 0)
	for _, p := range pieces {
		pos := [2]int{p[1], p[0]}
		posMap[pos] = uint8(p[2])
		if p[2] == 0 {
			black = append(black, &pos)
		} else {
			white = append(white, &pos)
		}
	}
	var muti = [][2]int{{1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	var mutiNodeMap = make(map[[2]int]uint8, 0)
	var willWinMap = make(map[[2]int]struct{}, 0)
	var whiteWinMap = make(map[[2]int]struct{}, 0)
	var canWin = func(iPos [2]int) int {
		pos := posMap[iPos]
		x, y := iPos[0], iPos[1]
		colr := pos & 1
		one := colr == 1
		lose := pos>>1 == 15
		if lose {
			return 2
		}
		var rm uint8 = 1
		for idx, mp := range muti {
			idx++
			if pos>>idx&1 != 0 {
				continue
			}
			lose = false
			rmDir := rm << idx
			nones := make([][2]int, 0, 4)
			mx, my := mp[0], mp[1]
			exists := [][2]int{}
			endNone := false
			for i := 1; i <= 4; i++ {
				p := [2]int{x + mx*i, y + my*i}
				c, ok := posMap[p]
				if ok {
					if c&1 == colr {
						exists = append(exists, p)
					} else {
						lose = true
					}
				} else {
					if i == 4 {
						endNone = true
					}
					nones = append(nones, p)
					lose = len(nones) > 2
				}
				if lose {
					for _, ep := range exists {
						if epos, ok := posMap[ep]; ok {
							posMap[ep] = rmDir | epos
						}
					}
					break
				}
			}
			if !lose || len(exists) >= 2 {
				existNum := len(exists)
				noneNum := len(nones)
				p1 := [2]int{x + mx*-1, y + my*-1}
				c, ok1 := posMap[p1]
				if existNum == 3 {
					if one {
						if !ok1 && endNone {
							return 0
						}
						if noneNum == 1 {
							whiteWinMap[nones[0]] = struct{}{}
						}
						if !ok1 {
							if noneNum == 0 || endNone {
								whiteWinMap[p1] = struct{}{}
							}
						}
						if len(whiteWinMap) > 1 {
							return 0
						}
						continue
					} else if !ok1 || noneNum > 0 {
						return 0
					}
				}
				if !one {
					if ok1 {
						if lose && c&1 != colr {
							continue
						}
					} else {
						mnp1, ok1 := mutiNodeMap[p1]
						p2 := [2]int{x + mx*-2, y + my*-2}
						c2, ok2 := posMap[p2]
						if !ok2 {
							if noneNum == 1 || ok1 && mnp1^rmDir != 0 {
								willWinMap[p1] = struct{}{}
								ok1 = false
							}
							if _, ok := mutiNodeMap[p2]; ok && !ok1 {
								willWinMap[p2] = struct{}{}
							}
							mutiNodeMap[p2] = mutiNodeMap[p2] | rmDir
						} else if noneNum == 2 && c2&1 == colr {
							_, ok1 := mutiNodeMap[nones[0]]
							_, ok2 := mutiNodeMap[nones[1]]
							if ok1 && !ok2 {
								willWinMap[nones[0]] = struct{}{}
							}
						}
						if noneNum > 0 {
							if noneNum == 2 {
								if endNone {
									if !ok2 {
										willWinMap[p1] = struct{}{}
									}
									willWinMap[nones[0]] = struct{}{}
								}
							}
							nones = append(nones, p1)
						}
					}
					for _, n := range nones {
						if mnp, ok := mutiNodeMap[n]; ok && mnp^rmDir != 0 {
							willWinMap[n] = struct{}{}
						}
						mutiNodeMap[n] = mutiNodeMap[n] | rmDir
					}
				}
			}
		}
		return 2
	}
	for _, pos := range black {
		if canWin(*pos) == 0 {
			return "Black"
		}
	}
	for _, pos := range white {
		if canWin(*pos) == 0 {
			return "White"
		}
	}
	if len(willWinMap) > 0 {
		// fmt.Println("willWinMap:", willWinMap)
		if len(whiteWinMap) > 0 {
			for pos := range whiteWinMap {
				if _, ok := willWinMap[pos]; ok {
					return "Black"
				}
			}
			// fmt.Println("whiteWinMap:", whiteWinMap)
			return "None"
		}
		return "Black"
	}
	return "None"
}

func printChess(pieces [][]int) {
	maxX, minX, maxY, minY := pieces[0][0], pieces[0][0], pieces[0][1], pieces[0][1]
	chessMap := make(map[[2]int]int)
	for _, p := range pieces {
		x, y := p[0], p[1]
		chessMap[[2]int{x, y}] = p[2]
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}
	fmt.Println(maxX, minX, maxY, minY)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if p, ok := chessMap[[2]int{x, y}]; ok {
				fmt.Print(p, " ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}
