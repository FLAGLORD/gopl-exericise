package countSHA

func countDiffBit(sha1, sha2 [32]byte) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		cnt += countByte(sha1[i], sha2[i])
	}
	return cnt
}

func countByte(byte1, byte2 byte) int {
	cnt := 0
	mask := int(byte1 ^ byte2)
	for mask != 0 {
		cnt++
		mask &= mask - 1
	}
	return cnt
}

/** Github Solution **/
func sha256DiffBitCount(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += diffBitCount(c1[i], c2[i])
	}
	return count
}

func diffBitCount(b1, b2 byte) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		mask := byte(1 << i)
		if b1&mask != b2&mask {
			count++
		}
	}
	return count
}
