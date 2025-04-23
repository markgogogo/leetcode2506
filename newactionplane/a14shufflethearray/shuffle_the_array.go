package a14shufflethearray

import "fmt"

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
	}
	fmt.Println(res)
	// for i, j := 0, n; j < 2*n; i, j = i+1, j+1 {
	// 	res[2*i+1] = nums[j]
	// }

	return res
}
