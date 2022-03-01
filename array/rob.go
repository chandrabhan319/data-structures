package array

func RobHouse(values []int) int {
	len := len(values)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return values[0]
	}

	robbing := make([]int, len)
	robbed := make([]int, len)

	robbing[0] = 0
	robbed[0] = values[0]

	for i := 1; i < len; i++ {
		robbing[i] = max(robbing[i-1], robbed[i-1])
		robbed[i] = robbing[i-1] + values[i]
	}

	if robbed[len-1] > robbed[len-2] {
		return robbed[len-1]
	}

	return robbed[len-2]
}

func RobHouseII(values []int) int {
	len := len(values)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return values[0]
	}

	r := make([]int, len)
	r[0] = values[0]
	r[1] = max(values[0], values[1])
	for i := 2; i < len; i++ {
		r[i] = max(r[i-2]+values[i], r[i-1])
	}

	return r[len-1]
}

func RobHouseIII(values []int) int {
	len := len(values)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return values[0]
	}

	r0 := values[0]
	r1 := max(values[0], values[1])
	for i := 2; i < len; i++ {
		r0 = max(r0+values[i], r1)
		r1 = max(r1, values[i-2]+values[i])
	}

	return r1
}
