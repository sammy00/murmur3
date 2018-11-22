package x86

// configuration constants for 128-bit version
// it's named so to avoid conflict with those for 32-bit output
const (
	CC1 = 0x239b961b
	CC2 = 0xab0e9789
	CC3 = 0x38b34ae5
	CC4 = 0xa1e38b93

	M = 5

	NN1 = 0x561ccd1b
	NN2 = 0x0bcaa747
	NN3 = 0x96cd1c35
	NN4 = 0x32ac3b17
)

const (
	RR11 = iota + 15
	RR12
	RR13
	RR14
)

const (
	RR21 = 19 - 2*iota
	RR22
	RR23
	RR24
)
