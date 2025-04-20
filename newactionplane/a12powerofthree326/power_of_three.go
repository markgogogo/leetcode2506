package a12powerofthree326

func isPowerOfThree(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		m := n % 3
		if m != 0 {
			return false
		}
		n /= 3
	}
	return false
}
