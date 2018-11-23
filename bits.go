package murmur3

func FinalizeMix64(h uint64) uint64 {
	h ^= h >> 33
	h *= DD1
	h ^= h >> 33
	h *= DD2
	h ^= h >> 33

	return h
}

func FinalizeMix86(h uint32) uint32 {
	h ^= h >> 16
	h *= D1
	h ^= h >> 13
	h *= D2
	h ^= h >> 16

	return h
}
