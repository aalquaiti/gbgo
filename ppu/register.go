package ppu

// Reg Represents LCD and PPU Registers
type Reg struct {
	// Holds values for Registers as follows:
	// Register    	Original Address    Mapped Address
	// LCDC        	0xFF40			   	0x00
	// STAT			0xFF41				0x01
	// SCY			0xFF42				0X02
	// SCX			0xFF43				0x03
	// LY			0xFF44				0x04
	// LYC			0xFF45				0x05
	// DMA			0xFF46				0x06
	// BGP			0xFF47				0x07
	// OBP0			0xFF48				0x08
	// OBP1			0xFF49				0x09
	// WY			0xFF4A				0x0A
	// WX			0xFF4B				0x0B
	val [0x0C]uint8
}

// IncLY Increment LY Register. Register Value is always within range 0 and 153
func (r *Reg) IncLY() uint8 {
	r.val[0x04] = (r.val[0x04] + 1) % 154
	return r.val[0x04]
}
