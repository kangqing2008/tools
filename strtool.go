package tools

// 截取指定位置的字符串
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func SubstrToEnd(s string, pos int) string {
	runes := []rune(s)
	return string(runes[pos:])
}