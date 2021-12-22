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

// Set the value of bit at position to true (1) or false (0)
func Set(value, position uint8, set bool) uint8 {
	// TODO: Optimise by using a pre-calculated array
	var mask uint8 = 1 << position

	if set {
		return value | mask
	}

	return value & ^mask
}

// IsSet determines if bit at position is true (1) or false (0)
func IsSet(value, position uint8) bool {
	var mask uint8 = 1 << position

	return value&mask == mask
}
