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
	return uint16(b.Ram[pos])<<8 + uint16(b.Ram[pos+1])
}

func (b *Bus) Write(pos uint16, value uint8) {
	b.Ram[pos] = value
}

func (b *Bus) Write16(pos uint16, value uint16) {
	b.Ram[pos] = uint8(value)
	b.Ram[pos+1] = uint8(value >> 8)
}

// TODO add write16 that writes data in small endian format
