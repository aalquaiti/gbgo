package gbgo

const (
	WRAM_SIZE = 0x1000
)

type Bus struct {
	ram [WRAM_SIZE]uint8
}

func (b *Bus) read(pos uint16) uint8 {
	return b.ram[pos]
}

func (b *Bus) read16(pos uint16) uint16 {
	return uint16(b.ram[pos])<<8 + uint16(b.ram[pos+1])
}

func (b *Bus) write(pos uint16, value uint8) {
	b.ram[pos] = value
}

func (b *Bus) write16(pos uint16, value uint16) {
	b.ram[pos] = uint8(value)
	b.ram[pos+1] = uint8(value >> 8)
}

// TODO add write16 that writes data in small endian format
