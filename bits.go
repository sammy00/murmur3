package murmur3

// FinalizeMix64 is a bit finalizer obfuscating a 64-bit integer
func FinalizeMix64(h uint64) uint64 {
	h ^= h >> 33
	h *= DD1
	h ^= h >> 33
	h *= DD2
	h ^= h >> 33

	return h
}

// FinalizeMix86 is a bit finalizer obfuscating a 32-bit integer
func FinalizeMix86(h uint32) uint32 {
	h ^= h >> 16
	h *= D1
	h ^= h >> 13
	h *= D2
	h ^= h >> 16

	return h
}
