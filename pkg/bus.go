package pkg

const (
	WRAM_SIZE = 0x1000
)

type Bus struct {
	Ram [WRAM_SIZE]uint8
}

func (b *Bus) Read(pos uint16) uint8 {
	return b.Ram[pos]
}

func (b *Bus) Read16(pos uint16) uint16 {
	low, high := b.Read16As8(pos)

	return uint16(high)<<8 + uint16(low)
}

// Returns an 8-bit tuple from memory as position and position + 1
// The first return byte is the least significant byte (low)
// the second return byte is the most significant byte (high)
func (b *Bus) Read16As8(pos uint16) (uint8, uint8) {
	return b.Ram[pos], b.Ram[pos+1]
}

func (b *Bus) Write(pos uint16, value uint8) {
	b.Ram[pos] = value
}

func (b *Bus) Write16(pos uint16, value uint16) {
	b.Ram[pos] = uint8(value)
	b.Ram[pos+1] = uint8(value >> 8)
}

// TODO add write16 that writes data in small endian format
