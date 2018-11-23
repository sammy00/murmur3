package murmur3

import (
	"encoding/binary"
	"math/bits"
)

func Sum32(data []byte, seed uint32) [4]byte {
	hash := SumUint32(data, seed)

	var out [4]byte
	binary.LittleEndian.PutUint32(out[:], hash)

	return out
}

func SumUint32(key []byte, seed uint32) uint32 {
	hash := seed

	ell := len(key)
	for i := 0; i+4 <= ell; i += 4 {
		k := binary.LittleEndian.Uint32(key[i:(i + 4)])

		k *= C1
		k = bits.RotateLeft32(k, R1)
		k *= C2

		hash ^= k
		hash = bits.RotateLeft32(hash, R2)
		hash = hash*M + N
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

	return FinalizeMix86(hash)
}
