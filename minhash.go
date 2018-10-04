package minhash

import (
	"math/big"

	"github.com/spaolacci/murmur3"
)

const (
	k = 128 // Number of hash functions.
)

func Sketch(words []string) []byte {
	sketch := big.NewInt(0)
	for i := 0; i < k; i++ {
		min := make([]uint64, 2)
		min[0], min[1] = murmur3.Sum128WithSeed([]byte(words[0]), uint32(i))
		for j := 1; j < len(words); j++ {
			h1, h2 := murmur3.Sum128WithSeed([]byte(words[j]), uint32(i))
			if h1 < min[0] {
				min[0] = h1
				min[1] = h2
			} else if h1 == min[0] && h2 < min[1] {
				min[0] = h1
				min[1] = h2
			}
		}
		sketch.SetBit(sketch, i, uint(min[1]&0x1))
	}
	return sketch.Bytes()
}

func Jaccard(s1, s2 []byte) float64 {
	sb1 := big.NewInt(0).SetBytes(s1)
	sb2 := big.NewInt(0).SetBytes(s2)
	popCnt := 0
	for _, byte := range sb1.Xor(sb1, sb2).Bytes() {
		popCnt += onesCount8(uint8(byte))
	}
	return 2.0 * (float64(k-popCnt)/float64(k) - 0.5)
}
