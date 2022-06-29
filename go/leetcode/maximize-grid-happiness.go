package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxGridHappiness(2, 3, 1, 2))
	// fmt.Println(getMaxGridHappiness(1, 2, 2, 2))
	// fmt.Println(getMaxGridHappiness(5, 4, 6, 3)) //920
	// fmt.Println(getMaxGridHappiness(2, 3, 2, 3)) //420
	// fmt.Println(getMaxGridHappiness(2, 3, 1, 2)) //240
	// fmt.Println(getMaxGridHappiness(3, 1, 2, 1)) //260
	// fmt.Println(getMaxGridHappiness(2, 2, 3, 2)) //280
}

func getMaxGridHappiness(m int, n int, nx int, wx int) int {
	calcMap := [3][3]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}}
	n3 := 1
	nn := make([]int, n+1)
	nn[0] = 1
	for i := 0; i < n; i++ {
		n3 *= 3
		nn[i+1] = n3
	}
	ML := n * m
	type IdxStruct struct {
		Idx int
		N   int
		W   int
	}
	preMax := make([][]int, n3)
	for i := range preMax {
		preMax[i] = make([]int, (nx+1)*(wx+1))
	}
	maxs := make([][]int, n3)
	for i := range maxs {
		maxs[i] = make([]int, (nx+1)*(wx+1))
	}
	for P := ML - 1; P >= 0; P-- {
		nidx := n
		if P < nidx {
			nidx = P
		}
		N := nx
		if (ML - P) < N {
			N = ML - P
		}
		W := wx
		if (ML - P) < W {
			W = ML - P
		}
		for idx := 0; idx < nn[nidx]; idx++ {
			lf, top, calcIdx := idx%3, (idx*3)/n3, (idx*3)%n3
			if P%n == 0 {
				lf = 0
			}
			for tn := N; tn >= 0; tn-- {
				for tw := W; tw >= 0; tw-- {
					max := preMax[calcIdx][tn*(wx+1)+tw]
					if tn > 0 {
						if tm := preMax[calcIdx+1][(tn-1)*(wx+1)+tw] + 120 + calcMap[1][lf] + calcMap[1][top]; tm > max {
							max = tm
						}
					}
					if tw > 0 {
						if tm := preMax[calcIdx+2][tn*(wx+1)+tw-1] + 40 + calcMap[2][lf] + calcMap[2][top]; tm > max {
							max = tm
						}
					}
					maxs[idx][tn*(wx+1)+tw] = max
				}
			}
		}
		preMax, maxs = maxs, preMax
	}
	max := 0
	for _, m := range preMax[0] {
		if m > max {
			max = m
		}
	}
	fmt.Println(preMax[0])
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
func calc(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a == 1 && b == 1 {
		return -60
	}
	if a == 2 && b == 2 {
		return 40
	}
	return -10
}
func getMaxGridHappiness1(m int, n int, nx int, wx int) int {
	n3 := 1
	for i := 0; i < n; i++ {
		n3 *= 3
	}
	mask_span := make([][]int, n3)
	highest := n3 / 3
	truncate := make([][3]int, n3)
	for mask := 0; mask < n3; mask++ {
		line := make([]int, n)
		for mask_tmp, i := mask, 0; i < n; i++ {
			line[n-i-1] = mask_tmp % 3
			mask_tmp /= 3
		}
		mask_span[mask] = line
		truncate[mask][0] = mask % highest * 3
		truncate[mask][1] = mask%highest*3 + 1
		truncate[mask][2] = mask%highest*3 + 2
	}
	dp := make([][][][]int, m*n)
	for i := range dp {
		dp[i] = make([][][]int, n3)
		for j := range dp[i] {
			dp[i][j] = make([][]int, nx+1)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, wx+1)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	// dfs(位置，轮廓线上的 mask，剩余的内向人数，剩余的外向人数)
	var dfs func(pos, borderline, nx, wx int) int
	dfs = func(pos, borderline, nx, wx int) int {
		// 边界条件：如果已经处理完，或者没有人了
		if pos == m*n || nx+wx == 0 {
			return 0
		}
		// 记忆化
		if dp[pos][borderline][nx][wx] != -1 {
			return dp[pos][borderline][nx][wx]
		}

		// 什么都不做
		best := dfs(pos+1, truncate[borderline][0], nx, wx)
		top := 0
		// 放一个内向的人
		if nx > 0 {
			if pos%n != 0 {
				top = calc(1, mask_span[borderline][n-1])
			}
			best = max(best, 120+calc(1, mask_span[borderline][0])+
				top+
				dfs(pos+1, truncate[borderline][1], nx-1, wx))
		}
		// 放一个外向的人
		if wx > 0 {
			if pos%n != 0 {
				top = calc(2, mask_span[borderline][n-1])
			}
			best = max(best, 40+calc(2, mask_span[borderline][0])+
				top+
				dfs(pos+1, truncate[borderline][2], nx, wx-1))
		}
		dp[pos][borderline][nx][wx] = best
		return best
	}
	return dfs(0, 0, nx, wx)
}
