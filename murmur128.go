package murmur3

import (
	"encoding/binary"
	"math/bits"
)

// Sum128 calculates the MurmurHash3 sum of data
func Sum128(data []byte, seed uint32) [16]byte {
	ell := uint64(len(data))

	h1, h2 := uint64(seed), uint64(seed)

	//fmt.Printf("data=%s***\n", data)

	// fetch decode each 16-byte block into 4 uint32 key in litter endian,
	// then for each key
	//  + mix it
	//  + xor it to the corresponding digest part
	for i := uint64(0); i+16 <= ell; i += 16 {
		// fetch decode each 16-byte block into 4 uint32 key in litter endian,
		k1 := binary.LittleEndian.Uint64(data[i:])
		k2 := binary.LittleEndian.Uint64(data[i+8:])
		//fmt.Printf("%x %x %x %x\n", k1, k2, k3, k4)
		//fmt.Println(k1, k2, k3, k4)

		// for each key
		//  + mix it
		//  + xor it to the corresponding digest part
		k1 *= CC1
		//fmt.Println("k1 =", k1)
		k1 = bits.RotateLeft64(k1, RR11)
		//fmt.Println("k1 =", k1, "RR11 =", RR11)
		k1 *= CC2

		h1 ^= k1
		h1 = bits.RotateLeft64(h1, RR21)
		h1 += h2
		h1 = h1*M + NN1
		//fmt.Println("k1 =", k1, "h1 =", h1)

		k2 *= CC2
		k2 = bits.RotateLeft64(k2, RR12)
		k2 *= CC1

		h2 ^= k2
		h2 = bits.RotateLeft64(h2, RR22)
		h2 += h1
		h2 = h2*M + NN2
	}

	//fmt.Printf("%x %x %x %x\n", h1, h2, h3, h4)
	//fmt.Println(h1, h2, h3, h4)

	var k1, k2 uint64
	switch tail := data[(ell - ell%16):]; len(tail) {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		//fmt.Println("k3 =", k3)
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		//fmt.Println("k3 =", k3)
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		//fmt.Println("k3 =", k3)
		fallthrough
	case 9:
		k2 ^= uint64(tail[8])
		k2 *= CC2
		k2 = bits.RotateLeft64(k2, RR12)
		k2 *= CC1
		//fmt.Println("k3 =", k3, "RR13=", RR13)
		fallthrough
	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0])
		k1 *= CC1
		k1 = bits.RotateLeft64(k1, RR11)
		k1 *= CC2
		fallthrough
	default:
	}
	//fmt.Println(k1, k2, k3, k4)
	h1, h2 = h1^k1, h2^k2

	h1, h2 = h1^ell, h2^ell
	//fmt.Println(h1, h2, h3, h4)

	h1 += h2
	h2 += h1
	//fmt.Println(h1, h2, h3, h4)

	h1, h2 = FinalizeMix64(h1), FinalizeMix64(h2)
	//fmt.Println(h1, h2, h3, h4)

	h1 += h2
	h2 += h1
	//fmt.Println(h1, h2, h3, h4)

	var out [16]byte
	binary.LittleEndian.PutUint64(out[:], h1)
	binary.LittleEndian.PutUint64(out[8:], h2)

	return out
}
