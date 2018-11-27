package murmur3_test

import (
	"testing"

	"github.com/sammy00/murmur3"
)

func TestSum32(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect [4]byte
	}{
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: [4]byte{0xcd, 0x26, 0xa8, 0x2f},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: [4]byte{0x23, 0xf7, 0x4f, 0x2e},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xd6, 0x22, 0xb4, 0x75},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: [4]byte{0xd4, 0x9e, 0x9c, 0x8d},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: [4]byte{0xfc, 0x00, 0x82, 0xf0},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0x27, 0x75, 0x75, 0x27},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: [4]byte{0xf5, 0x8c, 0x7b, 0x63},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x00000000,
			expect: [4]byte{0xae, 0xb0, 0x1b, 0x2c},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xe0, 0x60, 0x49, 0xf2},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: [4]byte{0x36, 0xf2, 0xb8, 0x21},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x00000000,
			expect: [4]byte{0x85, 0x54, 0xb4, 0x21},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xfc, 0x70, 0xa3, 0x1e},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: [4]byte{0x1c, 0x2f, 0xc0, 0xed},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: [4]byte{0xff, 0x62, 0xde, 0x02},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0x60, 0x14, 0x56, 0x14},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: [4]byte{0x2b, 0x57, 0x81, 0x97},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: [4]byte{0x41, 0x33, 0x61, 0xbd},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xff, 0x92, 0xd2, 0xb6},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x9747b28c,
			expect: [4]byte{0xb8, 0xc8, 0xb4, 0x8b},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x00000000,
			expect: [4]byte{0x15, 0x5c, 0x9d, 0x22},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xd3, 0x02, 0x45, 0xa9},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x9747b28c,
			expect: [4]byte{0xaf, 0x87, 0x64, 0xba},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x00000000,
			expect: [4]byte{0x10, 0xd1, 0xb0, 0xf4},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0x62, 0xf4, 0x89, 0x28},
		},
	}

	for i, c := range testCases {
		if got := murmur3.Sum32(c.data, c.seed); got != c.expect {
			t.Fatalf("#%d invalid hash: got 0x%x, expect 0x%x", i, got, c.expect)
		}
	}
}

func TestSumUint32(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect uint32
	}{
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: 0x2fa826cd,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: 0x2e4ff723,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: 0x75b422d6,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: 0x8d9c9ed4,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: 0xf08200fc,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: 0x27757527,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: 0x637b8cf5,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x00000000,
			expect: 0x2c1bb0ae,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: 0xf24960e0,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: 0x21b8f236,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x00000000,
			expect: 0x21b45485,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: 0x1ea370fc,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: 0xedc02f1c,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: 0x02de62ff,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: 0x14561460,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: 0x9781572b,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: 0xbd613341,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: 0xb6d292ff,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x9747b28c,
			expect: 0x8bb4c8b8,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x00000000,
			expect: 0x229d5c15,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0xc58f1a7b,
			expect: 0xa94502d3,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x9747b28c,
			expect: 0xba6487af,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x00000000,
			expect: 0xf4b0d110,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0xc58f1a7b,
			expect: 0x2889f462,
		},
		// more test cases from btcsuite/btcutil/bloom.MurmurHash3
		{
			data:   []byte{},
			seed:   0x00000000,
			expect: 0x00000000,
		},
		{
			data:   []byte{},
			seed:   0xfba4c795,
			expect: 0x6a396f08,
		},
		{
			data:   []byte{},
			seed:   0xffffffff,
			expect: 0x81f16f39,
		},
		{
			data:   []byte{0x00},
			seed:   0x00000000,
			expect: 0x514e28b7,
		},
		{
			data:   []byte{0x00},
			seed:   0xfba4c795,
			expect: 0xea3f0b17,
		},
		{
			data:   []byte{0xff},
			seed:   0x00000000,
			expect: 0xfd6cf10d,
		},
		{
			data:   []byte{0x00, 0x11},
			seed:   0x00000000,
			expect: 0x16c6b7ab,
		},
		{
			data:   []byte{0x00, 0x11, 0x22},
			seed:   0x00000000,
			expect: 0x8eb51c3d,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33},
			seed:   0x00000000,
			expect: 0xb4471bf8,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33, 0x44},
			seed:   0x00000000,
			expect: 0xe2301fa8,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
			seed:   0x00000000,
			expect: 0xfc2e4a15,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66},
			seed:   0x00000000,
			expect: 0xb074502c,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
			seed:   0x00000000,
			expect: 0x8034d2a0,
		},
		{
			data:   []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88},
			seed:   0x00000000,
			expect: 0xb4698def,
		},
	}

	/*
	   	format := `
	   {
	   	data: []byte("%s"),
	   	seed: 0x%08x,
	   	expect: [4]byte{0x%02x, 0x%02x, 0x%02x, 0x%02x},
	   },`*/
	for i, c := range testCases[24:] {
		/* var out [4]byte
		binary.LittleEndian.PutUint32(out[:], c.expect)
		fmt.Printf(format, c.data, c.seed, out[0], out[1], out[2], out[3])
		*/
		if got := murmur3.SumUint32(c.data, c.seed); got != c.expect {
			t.Fatalf("#%d invalid hash: got 0x%x, expect 0x%x", i, got, c.expect)
		}
	}
}
