package main

func main() {
	print(dieSimulator(4, []int{2, 1, 1, 3, 3, 2}), "\n")
	print(dieSimulator(500, []int{7, 5, 4, 1, 2, 7}), "\n")
}

func dieSimulator(n int, rollMax []int) int {
	// 用于记录当前数字的前rollMax[i]次的次数和
	var history [6][]int
	// 用于记录当前数字的前rollMax[i]次的隔代累加和
	var accum [6][]int
	// 用于记录当前数字在上面两个数组的N-1次的索引
	var maxIdx [6]int
	// 准备工作：初始化数据
	for idx, rmax := range rollMax {
		history[idx] = make([]int, rmax+1)
		accum[idx] = make([]int, rmax)
		maxIdx[idx] = 0
		history[idx][maxIdx[idx]] = 1
	}
	const mod = 1000000007
	rollNum := len(rollMax)
	// 从第一次开始推导，到N-1次结束
	for i := 1; i < n; i++ {
		for idx := 0; idx < rollNum; idx++ {
			rmax := rollMax[idx]
			rmaxIdx := maxIdx[idx]
			num := 0
			// 获取其他数字第i-1次的次数和
			for idx2 := 0; idx2 < rollNum; idx2++ {
				if idx2 != idx {
					num += history[idx2][maxIdx[idx2]]
				}
			}
			// 获取当前数字的隔代差累加和
			if rmax > 1 {
				distance := history[idx][rmaxIdx] - history[idx][(rmaxIdx+1)%rmax] + accum[idx][rmaxIdx]
				// 因为取模的原因会出现隔代差累加和为负数的情况，把模加回去就为正了
				if distance < 0 {
					distance += mod
				}
				accum[idx][rmaxIdx] = distance
				num += distance
			}
			num = num % mod
			// 因为当前回合其他数字还需要使用到i-1次的次数和，所以先临时把最新的数字和存起来
			history[idx][rmax] = num
		}
		// 更新下一轮使用到的i-1次索引
		for idx := 0; idx < rollNum; idx++ {
			rmax := rollMax[idx]
			rmaxIdx := (maxIdx[idx] + 1) % rmax
			maxIdx[idx] = rmaxIdx
			history[idx][rmaxIdx] = history[idx][rmax]
		}
	}
	// 推导到N-1次后，把每个数字最新的历史次数和加起来就是第N次的次数
	total := 0
	for idx := 0; idx < rollNum; idx++ {
		total += history[idx][rollMax[idx]]
		total = total % mod
	}
	return total
}
