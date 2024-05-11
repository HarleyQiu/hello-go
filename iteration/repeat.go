package iteration

// Repeat 返回一个新字符串，由指定的字符重复指定次数组成。
func Repeat(character string, count int) string {
	var repeatd string
	for i := 0; i < count; i++ {
		repeatd = repeatd + character
	}
	return repeatd
}
