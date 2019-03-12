package murmur3_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/sammyne/murmur3"
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
			expect: "213163d23b7f8a73e516c07e727345f9",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: "6c1b07bc7bbc4be347939ac4a93c437a",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: "ff9d0cd2ee401fac26f5efde525c9338",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: "94618270b057cdb83cf873585b456f55",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: "9a2685ff70a98c653e5c8ea6eae3fe43",
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: "8c935c5b843839f964b24f7ad58bbccd",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: "6c8ad027e39089788ad2cb67789ca4cf",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x00000000,
			expect: "c9fb0a32011820a64b3c7a60b06c3982",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: "69e483f3d9f149f2f480ce3a0527fe34",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: "dd904d9a52d2fb5e71ead0546624bea0",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x00000000,
			expect: "24b0e694c86c766a6c8fd44492bb010b",
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: "d81bb41cba1c1f083e81a8624ee4f16f",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: "a8fa6851bd2c21cdf33e80c8968b74d0",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: "b386ade2fee9e4bc7f4b6e4074e3e20a",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: "02f78a1b0296ec885cc5692ec8430864",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: "714c9a5add16aa271f90a72183fd2be0",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: "3222507256fe092f24d124bb1e8d7586",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: "48517672b6b419924084009f9f6d7381",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x9747b28c,
			expect: "76bc171b8f86b798b31a53645dc52e03",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0x00000000,
			expect: "02bad780310b2ff6166f516d34dfdb56",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog zzz"),
			seed:   0xc58f1a7b,
			expect: "862e5694107703e8be8935037efeded7",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x9747b28c,
			expect: "c503d778947729ff4798445c23e32b14",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0x00000000,
			expect: "00ca110863217b6749c0809e630e9154",
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog here"),
			seed:   0xc58f1a7b,
			expect: "0740903fd031fc1aba2cf790967b1f2c",
		},
	}

	for i, c := range testCases {
		expect, err := hex.DecodeString(c.expect)
		if nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		if got := murmur3.Sum128(c.data, c.seed); !bytes.Equal(got[:], expect) {
			t.Fatalf("#%d failed: got %x, expect %x", i, got, expect)
		}
	}
}
