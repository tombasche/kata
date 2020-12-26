package divide

func Divide(dividend, divisor int) int {
	var counter int
	sign := 1
	if divisor < 0 {
		sign = -1
	}
	if dividend == divisor {
		return 1
	}
	return divide(dividend, divisor*sign, counter) * sign
}

func divide(dividend, divisor int, counter int) int {
	if dividend-divisor <= 0 {
		return counter
	}
	counter++
	return divide(dividend-divisor, divisor, counter)
}
