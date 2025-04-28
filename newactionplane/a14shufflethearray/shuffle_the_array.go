package a14shufflethearray

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
	}
	for i, j := 0, n; j < 2*n; i, j = i+1, j+1 {
		res[2*i+1] = nums[j]
	}
	return res
}

func shuffle1(nums []int, n int) []int {
	return []int{1}
}

func shuffle2(nums []int, n int) []int {
	return []int{1}
}

func shuffle3(nums []int, n int) []int {
	return []int{1}
}

func shuffle4(nums []int, n int) []int {
	return []int{1}
}
