package main

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13, 25, 26, 27, 28, 29, 30, 31, 32}, []int{11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 33, 34, 35, 36}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 18, 21, 22, 23, 24, 25}, []int{11, 12, 13, 14, 15, 16, 17, 19, 20, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}))

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		nums1Mid = (high + low) >> 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}
