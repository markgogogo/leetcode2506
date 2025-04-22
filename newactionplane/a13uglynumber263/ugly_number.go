package a13uglynumber263

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for {
		if n%2 == 0 {
			n /= 2
			continue
		}
		if n%3 == 0 {
			n /= 3
			continue
		}
		if n%5 == 0 {
			n /= 5
			continue
		}
		break
	}

	return n == 1
}
