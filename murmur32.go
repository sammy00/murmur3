package murmur3

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

func SumUint32(key []byte, seed uint32) uint32 {

	hash := seed

	ell := len(key)
	for i := 0; i+4 <= ell; i += 4 {
		k := binary.LittleEndian.Uint32(key[i:(i + 4)])
		fmt.Println("k =", k)

		k *= C1
		//fmt.Println(k)
		k = bits.RotateLeft32(k, R1)
		//fmt.Println(k)
		k *= C2
		//fmt.Println(k)

		hash ^= k
		hash = bits.RotateLeft32(hash, R2)
		hash = hash*M + N

		//fmt.Println(hash)
	}

	var remaining uint32
	switch j := ell - ell%4; j {
	case ell - 3:
		remaining = uint32(key[j]) | (uint32(key[j+1]) << 8) |
			(uint32(key[j+2]) << 16)
	case ell - 2:
		remaining = uint32(key[j]) | (uint32(key[j+1]) << 8)
	case ell - 1:
		remaining = uint32(key[j])
	default:
	}
	remaining *= C1
	remaining = bits.RotateLeft32(remaining, R1)
	remaining *= C2
	hash ^= remaining

	hash ^= uint32(ell)

	hash ^= (hash >> 16)
	hash *= C3
	hash ^= (hash >> 13)
	hash *= C4
	hash ^= (hash >> 16)

	return hash
}
