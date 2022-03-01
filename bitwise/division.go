package bitwise

func Divide(dividend, divisor int64) int64 {
	sign := 1
	if !((dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0)) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	quotient := int64(0)
	temp := int64(0)
	for i := int64(63); i >= 0; i-- {
		if temp+(divisor<<i) <= dividend && temp+(divisor<<i) > 0 {
			temp += divisor << i
			quotient |= int64(1) << i
		}
	}

	if sign == -1 {
		quotient = -quotient
	}

	return quotient
}

func DivideInt8(dividend, divisor int8) int8 {
	sign := 1
	if !((dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0)) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	quotient := int8(0)
	temp := int8(0)
	for i := int8(7); i >= 0; i-- {
		c := temp + (divisor << i)
		if c <= dividend && c > 0 {
			temp += divisor << i
			quotient |= int8(1) << i
		}
	}

	if sign == -1 {
		quotient = -quotient
	}

	return quotient
}
