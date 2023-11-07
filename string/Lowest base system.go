package string

// https://www.codewars.com/kata/58bc16e271b1e4c5d3000151/train/go
import "math"

func check(n, b uint64) bool {
	for n > 0 {
		if n%b != 1 {
			return false
		}
		n /= b
	}
	return true
}

func GetMinBase(n uint64) uint64 {
	x, m := uint64(0), n
	for m > 0 {
		m /= 2
		x++
	}

	for i := x; i > 0; i-- {
		b := uint64(math.Pow(float64(n), 1/float64(i-1)))
		if check(n, b) {
			return b
		}
	}
	return n - 1
}
