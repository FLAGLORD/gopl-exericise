package popcount

var pc [256]byte

func init() {
	for i := range pc {
		//偶数 pc[i] = pc[i/2] ,奇数 pc[i] = pc[i/2] + Chapter1
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
