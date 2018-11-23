package murmur3

func FinalizeMix64(h uint64) uint64 {
	h ^= h >> 33
	h *= DD1
	h ^= h >> 33
	h *= DD2
	h ^= h >> 33

	return h
}
