package popcount

func PopCountV2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountV3(x uint64) int {
	result := 0
	for i := 0; i < 64; i++ {
		result += int((x >> i) & 1)
	}
	return result
}

func PopCountV4(x uint64) int {
	result := 0
	for x != 0 {
		result++
		x &= x - 1
	}
	return result
}
