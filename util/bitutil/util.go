package bitutil

// To16 Combine two 8-bit values to one 16-bit value
func To16(rHigh, rLow uint8) uint16 {
	return uint16(rHigh)<<8 + uint16(rLow)
}

// From16 Convert 16-bit value to an 8-bit tuple.
// The first return byte is the most significant byte (high)
// the second return byte is the least significant byte (low)
func From16(value uint16) (uint8, uint8) {
	high := uint8(value >> 8)
	low := uint8(value)
	return high, low
}
