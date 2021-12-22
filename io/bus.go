package io

import (
	"fmt"
	"github.com/aalquaiti/gbgo/util/bitutil"
	log "github.com/sirupsen/logrus"
)

// Device represents an IO Device
type Device interface {
	Read(uint16) uint8
	Write(uint16, uint8)
	Reset()
}

const (
	MaxRomAddr = 0xFFF
	VRamSize   = 0x2000
	WRamSize   = 0x2000
	OamSize    = 0xA0
	IoSize     = 0x80
	HRamSize   = 0x7F
	DivAddr    = 0xFF04 // Divider Register Address
	TimaAddr   = 0xFF05 // Timer Counter Address
	TmaAddr    = 0xFF06 // Timer Modulo Address
	TacAddr    = 0xFF07 // Time Control Address
)

type Bus struct {
	Rom  Device
	VRam [VRamSize]uint8 // Video RAM
	WRam [WRamSize]uint8 // Work RAM
	Oam  [OamSize]uint8  // Object Attribute Memory
	IO   [IoSize]uint8   // IO Registers
	HRam [HRamSize]uint8 // High RAM
	IE   uint8           // Interrupt Enable Register
}

// NewBus Creates New Bus
func NewBus(rom Device) Bus {
	return Bus{Rom: rom}
}

// Read Returns an 8-bitutil value from associated device connected to bus
// 0x0000 to 0x7FFF		ROM (Handled by Cartridge)
// 0x8000 to 0x9FFF		Vram
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
		return b.Rom.Read(address)
	// VRAM
	case address <= 0x9FFF:
		return b.VRam[address&0x7FFF]
	// External RAM
	case address <= 0xBFFF:
		return b.Rom.Read(address)
	// WRAM
	case address <= 0xDFFF:
		// TODO add CGB mode with switchable bank (1 to 7)
		return b.WRam[address&0x1FFF]
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
		log.WithFields(log.Fields{
			"Address": address,
		}).Warn("Reading from unusable memory")
		return 0
	// IO
	case address <= 0xFF7F:
		fmt.Printf("Reading from IO [%.4X]=%.4X ", address, b.IO[address&0x7F])
		return b.IO[address&0x7F]
	// HRAM
	case address <= 0xFFFE:
		fmt.Printf("Reading from HRAM [%.4X]=%.4X ", address, b.HRam[address&0x7F])
		return b.HRam[address&0x7F]
	// IE
	default:
		return b.IE
	}
}

// Read16As8 returns an 8-bitutil tuple from memory as address and address + 1
// The first return byte is the least significant byte (low)
// the second return byte is the most significant byte (high)
func (b *Bus) Read16As8(address uint16) (uint8, uint8) {
	return b.Read(address), b.Read(address + 1)
}

// Read16 returns an 16-bitutil value from associated device connected to bus
func (b *Bus) Read16(pos uint16) uint16 {
	low, high := b.Read16As8(pos)

	return uint16(high)<<8 + uint16(low)
}

// Write an 8-bitutil value to associated device connected to bus
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
		b.Rom.Write(address, value)
	// VRAM
	case address <= 0x9FFF:
		b.VRam[address&0x7FFF] = value
	// External RAM
	case address <= 0xBFFF:
		b.Rom.Write(address, value)
	// WRAM
	case address <= 0xDFFF:
		// TODO add CGB mode with switchable bank (1 to 7)
		b.WRam[address&0x1FFF] = value
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
		log.WithFields(log.Fields{
			"Address": address,
			"Value":   value,
		}).Warn("Writing to unusable memory")
	// IO
	case address <= 0xFF7F:

		address &= 0x7F

		// When Divider Register is accessed, it is reset
		if address == 0x04 {
			b.IO[address] = 0
		} else {
			fmt.Printf("Writing to IO [$FF%.2X]=%.4X ", address, value)
			b.IO[address] = value
		}

	// HRAM
	case address <= 0xFFFE:
		fmt.Printf("Writing to HRAM [%.4X]=%.4X ", address, value)
		b.HRam[address&0x7F] = value
	// IE
	default:
		b.IE = value
	}
}

// Write16 writes a 16-bitutil value to associated device connected to bus
func (b *Bus) Write16(address uint16, value uint16) {
	b.Write(address, uint8(value))
	b.Write(address+1, uint8(value>>8))
}

// region HelperFunctions

// ReadIE returns the value of Interrupt Enable Register at address 0xFFFF
func (b *Bus) ReadIE() uint8 {
	return b.IE
}

// ReadIF returns the value of Interrupt Flag at address OxFF0F
func (b *Bus) ReadIF() uint8 {
	return b.IO[0x0F]
}

// InterruptPending checks if an interrupt is pending, by ANDing the value of Interrupt Enable Register (IE) with the
// value of Interrupt Flag (IF)
func (b *Bus) InterruptPending() bool {
	return (b.IE & b.IO[0x0F]) != 0
}

// IsVblank determines if VBlank Interrupt is Enabled
func (b *Bus) IsVblank() bool {
	return b.IE&0x01 == 0x01
}

func (b *Bus) SetVblank(enable bool) {
	if enable {
		b.IE |= 0b00000001
	} else {
		b.IE &= 0b11111110
	}
}

// IsLCDStat determines if LCD Status Interrupt is Enabled
func (b *Bus) IsLCDStat() bool {
	return b.IE&0x02 == 0x02
}

func (b *Bus) SetLCDStat(enable bool) {
	if enable {
		b.IE |= 0b00000010
	} else {
		b.IE &= 0b11111101
	}
}

// IsTimerInt determines if Timer Interrupt is Enabled
func (b *Bus) IsTimerInt() bool {
	return b.IE == 0x04
}

func (b *Bus) SetTimerInt(enable bool) {
	if enable {
		b.IE |= 0b00000100
	} else {
		b.IE &= 0b11111011
	}
}

// IsSerialInt determines if Serial Interrupt is Enabled
func (b *Bus) IsSerialInt() bool {
	return b.IE&0x04 == 0x04
}

func (b *Bus) SetSerialInt(enable bool) {
	if enable {
		b.IE |= 0b00001000
	} else {
		b.IE &= 0b11110111
	}
}

// IsJoypadInt determines if Joypad Interrupt is Enabled
func (b *Bus) IsJoypadInt() bool {
	return b.IE&0x08 == 0x08
}

func (b *Bus) SetJoypad(enable bool) {
	if enable {
		b.IE |= 0b00010000
	} else {
		b.IE &= 0b11101111
	}
}

// IrqVblank determines if VBlank Interrupt is Requested
func (b *Bus) IrqVblank() bool {
	return b.IO[0x0F]&0x01 == 0x01
}

func (b *Bus) SetIrQVblank(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000001
	} else {
		b.IO[0x0F] &= 0b11111110
	}
}

// IrqLCDStat determines if LCD Status Interrupt is Requested
func (b *Bus) IrqLCDStat() bool {
	return b.IO[0x0F]&0x02 == 0x02
}

func (b *Bus) SetIRQLCDStat(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000010
	} else {
		b.IO[0x0F] &= 0b11111101
	}
}

// IrqTimer determines if Timer Interrupt is Requested
func (b *Bus) IrqTimer() bool {
	return b.IO[0x0F]&0x04 == 0x04
}

func (b *Bus) SetIRQTimer(enable bool) {
	if enable {
		b.IO[0x0F] |= 0b00000100
	} else {
		b.IO[0x0F] &= 0b11111011
	}
}

// IrqSerial determines if Serial Interrupt is Requested
func (b *Bus) IrqSerial() bool {
	return b.IO[0x0F]&0x08 == 0x08
}

func (b *Bus) SetIrqSerial(enable bool) {
	b.IO[0x0F] = bitutil.Set(b.IO[0x0F], 3, enable)
}

// IrqJoyPad determines if JoyPad Interrupt is Requested
func (b *Bus) IrqJoyPad() bool {
	return b.IO[0x0F]&0x10 == 0x10
}

func (b *Bus) SetIrqJoypad(enable bool) {
	b.IO[0x0F] = bitutil.Set(b.IO[0x0F], 4, enable)
}

// IncDIV Increment Divider Register by one.
// This is used instead of Write() function as writing to the Divider
// Register using that function to reset its value, as an expected
// behaviour by the game boy bus
func (b *Bus) IncDIV() {
	b.IO[DivAddr&0xFF]++
}

// IsTacTimerEnabled determines Timer Control (TAC) bitutil 2 to determine if Timer is Enabled. When enabled, Timer Counter
// can be incremented. This does not affect Divider Register
func (b *Bus) IsTacTimerEnabled() bool {
	return b.IO[TacAddr&0xFF]&0b100 == 0b100
}

// GetTacClockSelect Retrieve Timer Control (TAC) bits 0 and 1 that determine the Clock Selected for Timer Counter
func (b *Bus) GetTacClockSelect() uint8 {
	return b.IO[TacAddr&0xFF] & 0b11
}

// endregion
