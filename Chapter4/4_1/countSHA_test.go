package countSHA

import (
	"crypto/sha256"
	"testing"
)

func TestCountSHA(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	ans1 := countDiffBit(c1, c2)
	ans2 := sha256DiffBitCount(c1, c2)
	if ans1 != ans2 {
		t.Error("error!")
	}
}
