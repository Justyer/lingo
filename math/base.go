package math

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
