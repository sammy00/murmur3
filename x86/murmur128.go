package x86

import (
	"encoding/binary"
	"math/bits"

	"github.com/sammyne/murmur3"
)

// Sum128 implements the 128-bit output version of MurmurHash3 in x86 arch
func Sum128(data []byte, seed uint32) [16]byte {
	ell := uint32(len(data))

	h1, h2, h3, h4 := seed, seed, seed, seed

	//fmt.Printf("data=%s***\n", data)

	// fetch decode each 16-byte block into 4 uint32 key in litter endian,
	// then for each key
	//  + mix it
	//  + xor it to the corresponding digest part
	for i := uint32(0); i+16 <= ell; i += 16 {
		// fetch decode each 16-byte block into 4 uint32 key in litter endian,
		k1 := binary.LittleEndian.Uint32(data[i:])
		k2 := binary.LittleEndian.Uint32(data[i+4:])
		k3 := binary.LittleEndian.Uint32(data[i+8:])
		k4 := binary.LittleEndian.Uint32(data[i+12:])
		//fmt.Printf("%x %x %x %x\n", k1, k2, k3, k4)
		//fmt.Println(k1, k2, k3, k4)

		// for each key
		//  + mix it
		//  + xor it to the corresponding digest part
		k1 *= CC1
		//fmt.Println("k1 =", k1)
		k1 = bits.RotateLeft32(k1, RR11)
		//fmt.Println("k1 =", k1, "RR11 =", RR11)
		k1 *= CC2

		h1 ^= k1
		h1 = bits.RotateLeft32(h1, RR21)
		h1 += h2
		h1 = h1*M + NN1
		//fmt.Println("k1 =", k1, "h1 =", h1)

		k2 *= CC2
		k2 = bits.RotateLeft32(k2, RR12)
		k2 *= CC3

		h2 ^= k2
		h2 = bits.RotateLeft32(h2, RR22)
		h2 += h3
		h2 = h2*M + NN2

		k3 *= CC3
		k3 = bits.RotateLeft32(k3, RR13)
		k3 *= CC4

		h3 ^= k3
		h3 = bits.RotateLeft32(h3, RR23)
		h3 += h4
		h3 = h3*M + NN3

		k4 *= CC4
		k4 = bits.RotateLeft32(k4, RR14)
		k4 *= CC1

		h4 ^= k4
		h4 = bits.RotateLeft32(h4, RR24)
		h4 += h1
		h4 = h4*M + NN4
	}

	//fmt.Printf("%x %x %x %x\n", h1, h2, h3, h4)
	//fmt.Println(h1, h2, h3, h4)

	var k1, k2, k3, k4 uint32
	switch tail := data[(ell - ell%16):]; len(tail) {
	case 15:
		k4 ^= uint32(tail[14]) << 16
		fallthrough
	case 14:
		k4 ^= uint32(tail[13]) << 8
		fallthrough
	case 13:
		k4 ^= uint32(tail[12])
		k4 *= CC4
		k4 = bits.RotateLeft32(k4, RR14)
		k4 *= CC1
		fallthrough
	case 12:
		k3 ^= uint32(tail[11]) << 24
		//fmt.Println("k3 =", k3)
		fallthrough
	case 11:
		k3 ^= uint32(tail[10]) << 16
		//fmt.Println("k3 =", k3)
		fallthrough
	case 10:
		k3 ^= uint32(tail[9]) << 8
		//fmt.Println("k3 =", k3)
		fallthrough
	case 9:
		k3 ^= uint32(tail[8])
		k3 *= CC3
		k3 = bits.RotateLeft32(k3, RR13)
		k3 *= CC4
		//fmt.Println("k3 =", k3, "RR13=", RR13)
		fallthrough
	case 8:
		k2 ^= uint32(tail[7]) << 24
		fallthrough
	case 7:
		k2 ^= uint32(tail[6]) << 16
		fallthrough
	case 6:
		k2 ^= uint32(tail[5]) << 8
		fallthrough
	case 5:
		k2 ^= uint32(tail[4])
		k2 *= CC2
		k2 = bits.RotateLeft32(k2, RR12)
		k2 *= CC3
		fallthrough
	case 4:
		k1 ^= uint32(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= CC1
		k1 = bits.RotateLeft32(k1, RR11)
		k1 *= CC2
		fallthrough
	default:
	}
	//fmt.Println(k1, k2, k3, k4)
	h1, h2, h3, h4 = h1^k1, h2^k2, h3^k3, h4^k4

	h1, h2, h3, h4 = h1^ell, h2^ell, h3^ell, h4^ell
	//fmt.Println(h1, h2, h3, h4)

	h1 += h2 + h3 + h4
	h2, h3, h4 = h2+h1, h3+h1, h4+h1
	//fmt.Println(h1, h2, h3, h4)

	h1, h2 = murmur3.FinalizeMix86(h1), murmur3.FinalizeMix86(h2)
	h3, h4 = murmur3.FinalizeMix86(h3), murmur3.FinalizeMix86(h4)
	//fmt.Println(h1, h2, h3, h4)

	h1 += h2 + h3 + h4
	h2, h3, h4 = h2+h1, h3+h1, h4+h1
	//fmt.Println(h1, h2, h3, h4)

	var out [16]byte
	binary.LittleEndian.PutUint32(out[:], h1)
	binary.LittleEndian.PutUint32(out[4:], h2)
	binary.LittleEndian.PutUint32(out[8:], h3)
	binary.LittleEndian.PutUint32(out[12:], h4)

	return out
}
