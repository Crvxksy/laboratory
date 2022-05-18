package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxGridHappiness(2, 3, 1, 2)) //240
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func getMaxGridHappiness(m int, n int, nx int, wx int) int {
	n3 := n * n * n
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
			dp[i][j] = make([][]int, nx)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, wx)
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
	fmt.Println(mask_span)
	fmt.Println(truncate)
	fmt.Println(dp)

	return dfs(0, 0, nx, wx)
}
