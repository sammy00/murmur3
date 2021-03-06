package murmur3

import (
	"encoding/binary"
	"math/bits"
)

// Sum32 serves an alias as SumUint32 but with the hashed data encoded
// as a 4-byte little-endian byte sequences
func Sum32(data []byte, seed uint32) [4]byte {
	hash := SumUint32(data, seed)

	var out [4]byte
	binary.LittleEndian.PutUint32(out[:], hash)

	return out
}

// SumUint32 hashes the given 32-bit seed with the key block into a new
// 32-bit integer, based on the snippet as
//  key *= C1
//  key = rotl(key, R1)
//  key *= C2
//
//  hash ^= k
//
//  hash = rotl(hash, R1)
//  hash = hash*M+N
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

	var k uint32
	switch tail := key[ell-ell%4:]; len(tail) {
	case 3:
		k ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k ^= uint32(tail[0])
		k *= C1
		k = bits.RotateLeft32(k, R1)
		k *= C2
		fallthrough
	default:
	}
	hash ^= k

	hash ^= uint32(ell)

	return FinalizeMix86(hash)
}
