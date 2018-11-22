package x86_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/sammy00/murmur3/x86"
)

func TestSum128(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect string
	}{
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: "5ed5d48a7161b84c9c3aa78e3e79b6cd",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: "0e8829ccf91081107c6f5ebf3433e188",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: "85c392715a1701c682dfbf4a14508190",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: "98bb1b3c95194f656e79de6d66243f09",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: "6ec157f6790f9bee9e2be436c509f88a",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: "ba7799e33bd19f1c3d49a292432788fe",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x0,
			expect: "c383152f672ceeec6cf67b5d2c1de9e5",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x0,
			expect: "8843d60e79e79a3e934503972ff3b349",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x0,
			expect: "5cb9e26451952b2fcc493cb9fed0b4c9",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x0,
			expect: "484f7f9a1b2a6202c3ab64c4f4ddbbab",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x0,
			expect: "fdb47daf02170d403f6093539b388af2",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x0,
			expect: "fbf39633637c9e70efce1df157359e42",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: "4f3ba94a618f373a2b66976aa38ff8e8",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: "086525191fee0a51bcd01caaf5b988fd",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: "2b9a5cc8c8ec392164ae9207e0428b18",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: "1fc7b642b0ac04d7bb65238419cbc824",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: "a0a355f5e26816115ccf78a1925d1b4e",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: "6c22b9aa4bba167c809d3e087b1b9d46",
		},
	}

	for i, c := range testCases {
		expect, err := hex.DecodeString(c.expect)
		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if got := x86.Sum128(c.data, c.seed); !bytes.Equal(got[:], expect) {
			t.Fatalf("#%d failed: got %x, expect %x", i, got, expect)
		}
	}
}
