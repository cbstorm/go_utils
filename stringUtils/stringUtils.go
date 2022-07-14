package stringUtils

func CharCodeAt(str string, idx int) int {
	code := -1
	for i, e := range str {
		if i == idx {
			code = int(e)
		}
	}
	return code
}

func CharCode(str string) int {
	c := str[0:1]
	code := -1
	for _, e := range c {
		code = int(e)
	}
	return code
}

func CharCodeToString(c int) string {
	return string(rune(c))
}
