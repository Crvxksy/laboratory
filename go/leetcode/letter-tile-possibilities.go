package main

import (
	"fmt"
)

func main() {
	// fmt.Println(numTilePossibilities("AAB"))        // 8
	// fmt.Println(numTilePossibilities("AAABBC"))     // 188
	// fmt.Println(numTilePossibilities("AAAABBBCCD")) // 38847
	fmt.Println(getMaxGridHappiness(2, 3, 1, 2))

}

func numTilePossibilities(tiles string) int {
	m := make(map[byte]int)
	for i := range tiles {
		m[tiles[i]]++
	}
	arr := make([]int, 0)
	for _, v := range m {
		arr = append(arr, v)
	}
	fmt.Println(arr)
	return numTileDFS(arr)
}

func numTileDFS(arr []int) (r int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		r++
		arr[i]--
		r += numTileDFS(arr)
		arr[i]++
	}
	return
}
