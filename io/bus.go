package io

import (
	"github.com/aalquaiti/gbgo/gbgoutil"
	log "github.com/sirupsen/logrus"
)

// Device represents an IO Device
type Device interface {
	Read(uint16) uint8
	Write(uint16, uint8)
	Reset()
}

const (
	VRamSize = 0x2000
	WRamSize = 0x2000
	OamSize  = 0xA0
	IoSize   = 0x80
	HRamSize = 0x7F
	DivAddr  = 0xFF04 // Divider Register Address
	TimaAddr = 0xFF05 // Timer Counter Address
	TmaAddr  = 0xFF06 // Timer Modulo Address
	TacAddr  = 0xFF07 // Time Control Address
)

type Bus struct {
	Cartridge Device
	VRam      [VRamSize]uint8 // Video RAM
	WRam      [WRamSize]uint8 // Work RAM
	Oam       [OamSize]uint8  // Object Attribute Memory
	IO        [IoSize]uint8   // IO Registers
	HRam      [HRamSize]uint8 // High RAM
	IE        uint8           // Interrupt Enable Register
}

// NewBus Creates New Bus
func NewBus(cart Device) Bus {
	return Bus{Cartridge: cart}
}

// Read Returns an 8-bit value from associated device connected to io
// 0x0000 to 0x7FFF		ROM (Handled by Cartridge)
// 0x8000 to 0x9FFF		VRam
// 0xA000 to 0xBFFF		External RAM (Handled by Cartridge)
// 0xC000 to 0xDFFF		Work RAM (WRam)
// 0xE000 to 0xFDFF		Echo. Mirrors 0xC000 to 0xDFFF
// 0xFE00 to 0xFE9F		Object Attribute Table (Oam)
// 0xFEA0 to 0xFEFF		Unusable
// 0xFF00 to 0xFF7F		IO Registers (of which 0xFF40 to 0xFF4B handled by PPU)
// 0xFF80 to 0xFFFE		High RAM (HRam)
// 0xFFFF				Interrupt Enable Register (IE)
func (b *Bus) Read(address uint16) uint8 {
	switch {
	// ROM
	case address <= 0x7FFF:
		return b.Cartridge.Read(address)
	// VRAM
	case address <= 0x9FFF:
		return b.VRam[address&0x7FFF]
	// External RAM
	case address <= 0xBFFF:
		return b.Cartridge.Read(address)
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
		// TODO: remove print
		//fmt.Printf("Reading from IO [%.4X]=%.4X ", address, b.IO[address&0x7F])
		return b.IO[address&0x7F]
	// HRAM
	case address <= 0xFFFE:
		// TODO: remove print
		//fmt.Printf("Reading from HRAM [%.4X]=%.4X ", address, b.HRam[address&0x7F])
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

// Read16 returns an 16-bitutil value from associated device connected to io
func (b *Bus) Read16(pos uint16) uint16 {
	low, high := b.Read16As8(pos)

	return uint16(high)<<8 + uint16(low)
}

// Write an 8-bitutil value to associated device connected to io
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
		b.Cartridge.Write(address, value)
	// VRAM
	case address <= 0x9FFF:
		b.VRam[address&0x7FFF] = value
	// External RAM
	case address <= 0xBFFF:
		b.Cartridge.Write(address, value)
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
			// TODO remove print
			//fmt.Printf("Writing to IO [$FF%.2X]=%.4X ", address, value)
			b.IO[address] = value
		}

	// HRam
	case address <= 0xFFFE:
		// TODO remove print
		//fmt.Printf("Writing to HRAM [%.4X]=%.4X ", address, value)
		b.HRam[address&0x7F] = value
	// IE
	default:
		b.IE = value
	}
}

// Write16 writes a 16-bit value to associated device connected to io
func (b *Bus) Write16(address uint16, value uint16) {
	b.Write(address, uint8(value))
	b.Write(address+1, uint8(value>>8))
}

// region HelperFunctions

// InterruptPending checks if an interrupt is pending, by ANDing the value of Interrupt Enable Register (IE) with the
// value of Interrupt Flag (IF)
func (b *Bus) InterruptPending() bool {
	return (b.IE & b.IO[0x0F]) != 0
}

// IsVBlank determines if VBlank Interrupt is Enabled
func (b *Bus) IsVBlank() bool {
	return gbgoutil.IsBitSet(b.IE, 0)
}

func (b *Bus) SetVBlank(enable bool) {
	b.IE = gbgoutil.SetBit(b.IE, 0, enable)
}

// IsLCDStat determines if LCD Status Interrupt is Enabled
func (b *Bus) IsLCDStat() bool {
	return gbgoutil.IsBitSet(b.IE, 1)
}

func (b *Bus) SetLCDStat(enable bool) {
	b.IE = gbgoutil.SetBit(b.IE, 1, enable)
}

// IsTimerInt determines if Timer Interrupt is Enabled
func (b *Bus) IsTimerInt() bool {
	return gbgoutil.IsBitSet(b.IE, 2)
}

func (b *Bus) SetTimerInt(enable bool) {
	b.IE = gbgoutil.SetBit(b.IE, 2, enable)
}

// IsSerialInt determines if Serial Interrupt is Enabled
func (b *Bus) IsSerialInt() bool {
	return gbgoutil.IsBitSet(b.IE, 3)
}

func (b *Bus) SetSerialInt(enable bool) {
	b.IE = gbgoutil.SetBit(b.IE, 3, enable)
}

// IsJoypadInt determines if Joypad Interrupt is Enabled
func (b *Bus) IsJoypadInt() bool {
	return gbgoutil.IsBitSet(b.IE, 4)
}

func (b *Bus) SetJoypad(enable bool) {
	b.IE = gbgoutil.SetBit(b.IE, 4, enable)
}

// IrqVBlank determines if VBlank Interrupt is Requested
func (b *Bus) IrqVBlank() bool {
	return gbgoutil.IsBitSet(b.IO[0x0F], 0)
}

func (b *Bus) SetIrQVblank(enable bool) {
	b.IO[0x0F] = gbgoutil.SetBit(b.IO[0x0F], 0, enable)
}

// IrqLCDStat determines if LCD Status Interrupt is Requested
func (b *Bus) IrqLCDStat() bool {
	return gbgoutil.IsBitSet(b.IO[0x0F], 1)
}

func (b *Bus) SetIRQLCDStat(enable bool) {
	b.IO[0x0F] = gbgoutil.SetBit(b.IO[0x0F], 1, enable)
}

// IrqTimer determines if Timer Interrupt is Requested
func (b *Bus) IrqTimer() bool {
	return gbgoutil.IsBitSet(b.IO[0x0F], 2)
}

func (b *Bus) SetIRQTimer(enable bool) {
	b.IO[0x0F] = gbgoutil.SetBit(b.IO[0x0F], 2, enable)
}

// IrqSerial determines if Serial Interrupt is Requested
func (b *Bus) IrqSerial() bool {
	return gbgoutil.IsBitSet(b.IO[0x0F], 3)
}

func (b *Bus) SetIrqSerial(enable bool) {
	b.IO[0x0F] = gbgoutil.SetBit(b.IO[0x0F], 3, enable)
}

// IrqJoyPad determines if JoyPad Interrupt is Requested
func (b *Bus) IrqJoyPad() bool {
	return gbgoutil.IsBitSet(b.IO[0x0F], 4)
}

func (b *Bus) SetIrqJoyPad(enable bool) {
	b.IO[0x0F] = gbgoutil.SetBit(b.IO[0x0F], 4, enable)
}

// IncDIV Increment Divider Register by one.
// This is used instead of Write() function as writing to the Divider Register using that function
// resets its value, as an expected behaviour by the game boy io
func (b *Bus) IncDIV() {
	b.IO[DivAddr&0xFF]++
}

// IsTacTimerEnabled determines Timer Control (TAC) bit 2 to determine if Timer is Enabled. When enabled, Timer Counter
// can be incremented. This does not affect Divider Register
func (b *Bus) IsTacTimerEnabled() bool {
	return gbgoutil.IsBitSet(b.IO[TacAddr&0xFF], 2)
}

// GetTacClockSelect Retrieve Timer Control (TAC) bits 0 and 1 that determine the Clock Selected for Timer Counter
func (b *Bus) GetTacClockSelect() uint8 {
	return b.IO[TacAddr&0xFF] & 0b11
}

// endregion
