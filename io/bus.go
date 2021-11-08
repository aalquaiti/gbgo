package io

import (
	"github.com/golang/glog"
)

const (
	VRAM_SIZE = 0x2000
	WRAM_SIZE = 0x2000
	OAM_SIZE  = 0xA0
	IO_SIZE   = 0x80
	HRAM_SIZE = 0x7F
	DIV_ADDR  = 0xFF04 // Divider Register Address
	TIMA_ADDR = 0xFF05 // Timer Counter Address
	TMA_ADDR  = 0xFF06 // Timer Modulo Address
	TAC_ADDR  = 0xFF07 // Time Control Address
)

type Bus struct {
	VRam [VRAM_SIZE]uint8 // Video RAM
	WRam [WRAM_SIZE]uint8 // Work RAM
	Oam  [OAM_SIZE]uint8  // Object Attribute Memory
	IO   [IO_SIZE]uint8   // IO Registers
	HRam [HRAM_SIZE]uint8 // High RAM
	IE   uint8            // Interrupt Enable Register
}

// Returns an 8-bit value from associated device connected to bus
// 0x0000 to 0x7FFF		ROM (Handled by Cartridge)
// 0x8000 to 0x9FFF		VRAM
// 0xA000 to 0xBFFF		External RAM (Handled by Cartridge)
// 0xC000 to 0xDFFF		Work RAM (WRAM)
// 0xE000 to 0xFDFF		Echo. Mirrors 0xC000 to 0xDFFF
// 0xFE00 to 0xFE9F		Object Attribute Table (OAM)
// 0xFEA0 to 0xFEFF		Unusable
// 0xFF00 to 0xFF7F		IO Registers
// 0xFF80 to 0xFFFE		High RAM (HRAM)
// 0xFFFF				Interrupt Enable Register
func (b *Bus) Read(address uint16) uint8 {
	switch {
	// ROM
	case address <= 0x7FFF:
		// TODO Implement
		// TODO add bank switch functionality
		return 0
	// VRAM
	case address <= 0x9FFF:
		return b.VRam[address&0x7FFF]
	// External RAM
	case address <= 0xBFFF:
		// TODO Implement
		return 0
	// WRAM
	case address <= 0xDFFF:
		// TODO add CGB mode with switchable bank (1 to 7)
		return b.WRam[0x1FFF]
	// WRAM Echo
	case address <= 0xFDFF:
		// address is AND with 0x1DFF as the echo does not mirror
		// the whole WRAM address
		return b.WRam[address&0x1DFF]
	// OAM
	case address <= 0xFE9F:
		return b.Oam[address&0x9F]
	// Unusable
	case address <= 0xFEFF:
		// Do nothing
		// TODO add behaviour related to OAM access
		glog.Warningf("Unusable address accessed at $%x", address)
		return 0
	// IO
	case address <= 0xFF7F:
		return b.IO[address&0x7F]
	// HRAM
	case address <= 0xFFFE:
		return b.HRam[address&0x7F]
	// IE
	default:
		return b.IE
	}
}

// Returns an 8-bit tuple from memory as address and address + 1
// The first return byte is the least significant byte (low)
// the second return byte is the most significant byte (high)
func (b *Bus) Read16As8(address uint16) (uint8, uint8) {
	return b.Read(address), b.Read(address)
}

// Returns an 16-bit value from associated device connected to bus
func (b *Bus) Read16(pos uint16) uint16 {
	low, high := b.Read16As8(pos)

	return uint16(high)<<8 + uint16(low)
}

// Writes an 8-bit value to associated device connected to bus
// 0x0000 to 0x7FFF		ROM (Handled by Cartridge)
// 0x8000 to 0x9FFF		VRAM
// 0xA000 to 0xBFFF		External RAM (Handled by Cartridge)
// 0xC000 to 0xDFFF		Work RAM (WRAM)
// 0xE000 to 0xFDFF		Echo. Mirrors 0xC000 to 0xDFFF
// 0xFE00 to 0xFE9F		Object Attribute Table (OAM)
// 0xFEA0 to 0xFEFF		Unusable
// 0xFF00 to 0xFF7F		IO Registers
// 0xFF80 to 0xFFFE		High RAM (HRAM)
// 0xFFFF				Interrupt Enable Register
func (b *Bus) Write(address uint16, value uint8) {
	switch {
	// ROM
	case address <= 0x7FFF:
		// TODO Implement
		// TODO add bank switch functionality
	// VRAM
	case address <= 0x9FFF:
		b.VRam[address&0x7FFF] = value
	// External RAM
	case address <= 0xBFFF:
		// TODO Implement
	// WRAM
	case address <= 0xDFFF:
		// TODO add CGB mode with switchable bank (1 to 7)
		b.WRam[0x1FFF] = value
	// WRAM Echo
	case address <= 0xFDFF:
		// address is AND with 0x1DFF as the echo does not mirror
		// the whole WRAM address
		b.WRam[address&0x1DFF] = value
	// OAM
	case address <= 0xFE9F:
		b.Oam[address&0x9F] = value
	// Unusable
	case address <= 0xFEFF:
		// Do nothing
		// TODO add behaviour related to OAM access
		glog.Warningf("Value: $%X tried to write to unusable address: $%x",
			value, address)
	// IO
	case address <= 0xFF7F:
		address &= 0x7F

		// When Divider Register is accessed, it is reset
		if address == 0x04 {
			b.IO[address] = 0
		} else {
			b.IO[address] = value
		}

	// HRAM
	case address <= 0xFFFE:
		b.HRam[address&0x7F] = value
	// IE
	default:
		b.IE = value
	}
}

// Writes a 16-bit value to associated device connected to bus
func (b *Bus) Write16(address uint16, value uint16) {
	b.Write(address, uint8(value))
	b.Write(address+1, uint8(value>>8))
}

// HelperFunctions

// Checks if VBlank Interrupt is Enabled
func (b *Bus) IsVblank() bool {
	return b.IE&0x01 == 0x01
}

// Set VBlank Interrupt
func (b *Bus) SetVblank(enable bool) {
	if enable {
		b.IE |= 0b00000001
	} else {
		b.IE &= 0b11111110
	}
}

// Checks if LCD Status Interrupt is Enabled
func (b *Bus) IsLCDStat() bool {
	return b.IE&0x02 == 0x02
}

// Set LCD Status Interrupt
func (b *Bus) SetLCDStat(enable bool) {
	if enable {
		b.IE |= 0b00000010
	} else {
		b.IE &= 0b11111101
	}
}

// Checks if Timer Interrupt is Enabled
func (b *Bus) IsTimerInt() bool {
	return b.IE == 0x04
}

// Set Timer Interrupt
func (b *Bus) SetTimerInt(enable bool) {
	if enable {
		b.IE |= 0b00000100
	} else {
		b.IE &= 0b11111011
	}
}

// Checks if Serial Interrupt is Enabled
func (b *Bus) IsSerialInt() bool {
	return b.IE&0x04 == 0x04
}

// Set Serial Interrupt
func (b *Bus) SetSerialInt(enable bool) {
	if enable {
		b.IE |= 0b00001000
	} else {
		b.IE &= 0b11110111
	}
}

// Checks if Joypad Interrupt is Enabled
func (b *Bus) IsJoypadInt() bool {
	return b.IE&0x08 == 0x08
}

// Set Joypad Interrupt
func (b *Bus) SetJoypad(enable bool) {
	if enable {
		b.IE |= 0b00010000
	} else {
		b.IE &= 0b11101111
	}
}

// Checks if VBlank Interrupt is Requested
func (b *Bus) IrqVblank() bool {
	return b.IO[0x0F]&0x01 == 0x01
}

// Set VBlank Interrupt Request
func (b *Bus) SetIrQVblank(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000001
	} else {
		b.IO[0x0F] &= 0b11111110
	}
}

// Checks if LCD Status Interrupt is Requested
func (b *Bus) IrqLCDStat() bool {
	return b.IO[0x0F]&0x02 == 0x02
}

// Set LCD Status Interrupt Request
func (b *Bus) SetIRQLCDStat(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000010
	} else {
		b.IO[0x0F] &= 0b11111101
	}
}

// Checks if Timer Interrupt is Requested
func (b *Bus) IrqTimer() bool {
	return b.IO[0x0F]&0x04 == 0x04
}

// Set Timer Interrupt Request
func (b *Bus) SetIRQTimer(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000100
	} else {
		b.IO[0x0F] &= 0b11111011
	}
}

// Checks if Serial Interrupt is Requested
func (b *Bus) IrqSerial() bool {
	return b.IO[0x0F]&0x08 == 0x08
}

// Set Serial Interrupt Request
func (b *Bus) SetIrqSerial(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00001000
	} else {
		b.IO[0x0F] &= 0b11110111
	}
}

// Checks if Joypad Interrupt is Requested
func (b *Bus) IrqJoypad() bool {
	return b.IO[0x0F]&0x10 == 0x10
}

// Set Joypad Interrupt Request
func (b *Bus) SetIrqJoypad(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00010000
	} else {
		b.IO[0x0F] &= 0b11101111
	}
}

// Increment Divider Register by one.
// This is used instead of Write() function as writing to the Divider
// Register using that function to reset its value, as an expected
// behaviour by the game boy bus
func (b *Bus) IncDIV() {
	b.IO[DIV_ADDR&0xFF]++
}

// Checks Timer Control (TAC) bit 2 to determine if Timer is Enabled.
// when enabled, Timer Counter can be incremented. This does not
// affect Divider Register
func (b *Bus) IsTacTimerEnabled() bool {
	return b.IO[TAC_ADDR&0xFF]&0b100 == 0b100
}

// Retrieve Timer Contrl (TAC) bits 0 and 1 that determine the Clock
// Selected for Timer Counter
func (b *Bus) GetTacClockSelect() uint8 {
	return b.IO[TAC_ADDR&0xFF] & 0b11
}