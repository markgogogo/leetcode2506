package a08tolowercase709

func toLowerCase(s string) string {
	res := make([]byte, len(s)) // TODO：必须make初始化？
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res[i] = s[i] + 32
		} else {
			res[i] = s[i]
		}
	}
	return string(res)
}
