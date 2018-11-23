package murmur3

// configuration contants for hashing
const (
	C1 = 0xcc9e2d51
	C2 = 0x1b873593
	R1 = 15
	R2 = 13
	M  = 5
	N  = 0xe6546b64

	// extra constants
	D1 = 0x85ebca6b
	D2 = 0xc2b2ae35
)

// constants enumeration for x64-128bit
const (
	CC1 uint64 = 0x87c37b91114253d5
	CC2 uint64 = 0x4cf5ad432745937f

	NN1 = 0x52dce729
	NN2 = 0x38495ab5

	DD1 uint64 = 0xff51afd7ed558ccd
	DD2 uint64 = 0xc4ceb9fe1a85ec53
)

const (
	RR11 = 31
	RR12 = 33
)

const (
	RR21 = 27
	RR22 = 31
)
