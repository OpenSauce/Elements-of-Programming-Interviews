package chapter_4

import (
	"math"
	"math/bits"
)

var precomputedArray = make([]uint64, math.MaxUint16)

func ReverseBits(x uint64) uint64 {
	for i := uint16(0); i < math.MaxUint16; i++ {
		precomputedArray[i] = uint64(bits.Reverse16(i))
	}
	maskSize := uint64(16)
	bitMask := uint64(0xFFFF)
	return precomputedArray[x&bitMask]<<(maskSize*3) | // 16 LSB reversed and moved to MSB.
		precomputedArray[(x>>maskSize)&bitMask]<<(2*maskSize) | // 2nd 16 LSB moved to 16 LSB, reversed and moved to 2nd MSB partition.
		precomputedArray[(x>>(2*maskSize))&bitMask]<<maskSize | // 3rd 16 LSB moved to 16 LSB, reversed and moved to the 3rd MSB partition.
		precomputedArray[(x>>(3*maskSize))&bitMask] // 16 MSB moved to the LSB and reversed.
}
