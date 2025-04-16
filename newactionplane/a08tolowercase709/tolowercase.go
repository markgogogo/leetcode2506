package a08tolowercase709

func toLowerCase(s string) string {
	// slice的定义方式：
	// 方式1.
	// res := make([]byte, len(s))
	// 赋值方式：res[i] = xx
	// 方式2.
	// var res []byte
	// 赋值方式：res = append(res, xx)
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res[i] = s[i] + 32
		} else {
			res[i] = s[i]
		}
	}
	return string(res)
}

func toLowerCase1(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res = append(res, s[i]+32)
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
