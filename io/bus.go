package io

import (
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
	HRamSize = 0x7F
)

// Memory Addresses
const (
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

	// IO Registers
	Time TimeReg
	IF   IF

	//IO        [IoSize]uint8   // IO Registers
	HRam [HRamSize]uint8 // High RAM
	IE   IE              // Interrupt Enable Register
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
// 0xFF00 to 0xFF7F		IO Registers
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
		switch address {
		case DivAddr:
			return b.Time.Div
		case TimaAddr:
			return b.Time.Tima
		case TmaAddr:
			return b.Time.Tma
		case TacAddr:
			return b.Time.Tac
		}
		// TODO handle other IO registers
		log.Debugf("bus-read: Unsupported IO Register $%.4X", address)
		return 0
	// HRAM
	case address <= 0xFFFE:
		// TODO: remove print
		//fmt.Printf("Reading from HRAM [%.4X]=%.4X ", address, b.HRam[address&0x7F])
		return b.HRam[address&0x7F]
	// IE
	default:
		return uint8(b.IE)
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
		// the whole WRam address
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

		log.Debugf("Writing: IO [$%.4X]=$%.4X", address, value)
		switch address {
		case DivAddr:
			// When Divider Register is accessed, it is reset
			b.Time.Div = 0
		case TimaAddr:
			b.Time.Tima = value
		case TmaAddr:
			b.Time.Tma = value
		case TacAddr:
			b.Time.Tac = value
		default:
			// TODO handle other IO Registers
			log.Debugf("bus-write: Unsupported IO Register $%.4X", address)
		}
	// HRam
	case address <= 0xFFFE:
		// TODO remove print
		//fmt.Printf("Writing to HRAM [%.4X]=%.4X ", address, value)
		b.HRam[address&0x7F] = value
	// IE
	default:
		b.IE = IE(value)
	}
}

// Write16 writes a 16-bit value to associated device connected to io
func (b *Bus) Write16(address uint16, value uint16) {
	b.Write(address, uint8(value))
	b.Write(address+1, uint8(value>>8))
}

// InterruptPending checks if an interrupt is pending, by ANDing the value of Interrupt Enable Register (IE) with the
// value of Interrupt Flag (IF)
func (b *Bus) InterruptPending() bool {
	return (uint8(b.IE) & uint8(b.IF)) != 0
}
