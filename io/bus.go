package io

import (
	"github.com/aalquaiti/gbgo/gbgoutil"
	"github.com/sirupsen/logrus"
)

// Device represents an IO Device
type Device interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
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
	MaxAddrVRam  uint16 = 0x9FFF
	MaxAddrOam   uint16 = 0xFE9F
	MinAddrLcdIO uint16 = 0xFF40
	MaxAddrLcdIO uint16 = 0xFF4B

	AddrDiv  uint16 = 0xFF04 // Divider Register Address
	AddrTima uint16 = 0xFF05 // Timer Counter Address
	AddrTma  uint16 = 0xFF06 // Timer Modulo Address
	AddrTac  uint16 = 0xFF07 // Time Control Address
	AddrIF   uint16 = 0xFF0F
	AddrLcdc uint16 = 0xFF40
	AddrLcds uint16 = 0xFF41
	AddrScy  uint16 = 0xFF42
	AddrScx  uint16 = 0xFF43
	AddrLy   uint16 = 0xFF44
	AddrLyc  uint16 = 0xFF45
	AddrDma  uint16 = 0xFF46
	AddrBgp  uint16 = 0xFF47
	AddrObp0 uint16 = 0xFF48
	AddrObp1 uint16 = 0xFF49
	AddrWy   uint16 = 0xFF4A
	AddrWx   uint16 = 0xFF4B
)

type Bus struct {
	cart Device
	ppu  Device
	WRam [WRamSize]uint8 // Work RAM

	// IO Registers
	Time TimeReg
	IF   IF

	HRam [HRamSize]uint8 // High RAM
	IE   IE              // Interrupt Enable Register
}

// NewBus Creates New Bus
func NewBus(cart, ppu Device) Bus {
	return Bus{
		cart: cart,
		ppu:  ppu,
	}
}

// Read Returns an 8-bit value from associated device connected to io
// 0x0000 to 0x7FFF		ROM (Handled by cart)
// 0x8000 to 0x9FFF		VRam
// 0xA000 to 0xBFFF		External RAM (Handled by cart)
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
		return b.cart.Read(address)
	// VRAM
	case address <= MaxAddrVRam:
		return b.ppu.Read(address)
	// External RAM
	case address <= 0xBFFF:
		return b.cart.Read(address)
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
	case address <= MaxAddrOam:
		return b.ppu.Read(address)
	// Unusable
	case address <= 0xFEFF:
		// Do nothing
		// TODO add behaviour related to OAM access
		logrus.Warnf("Reading from unusable memory at address $%.4X", address)
		return 0

	// IO
	case address == AddrDiv:
		return b.Time.div
	case address == AddrTima:
		return b.Time.tima
	case address == AddrTma:
		return b.Time.tma
	case address == AddrTac:
		return b.Time.tac
	case address == AddrIF:
		return uint8(b.IF)

	case address >= MinAddrLcdIO && address <= MaxAddrLcdIO:
		return b.ppu.Read(address)

	// HRAM
	case address <= 0xFFFE:
		value := b.HRam[address&0x7F]
		logrus.Debugf("bus: Reading HRAM [%.4X]=%.4X", address, value)
		return value
	// IE
	case address == 0xFFFF:
		return uint8(b.IE)
	}

	logrus.Debugf("bus: Read was not mapped to Device at $%.4X", address)
	return 0
}

// Read16As8 returns an 8-bit tuple from memory as address and address + 1
// The first return byte is the least significant byte (low)
// the second return byte is the most significant byte (high)
func (b *Bus) Read16As8(address uint16) (uint8, uint8) {
	return b.Read(address), b.Read(address + 1)
}

// Read16 returns an 16-bit value from associated device connected to io
func (b *Bus) Read16(pos uint16) uint16 {
	low, high := b.Read16As8(pos)

	return gbgoutil.To16(high, low)
}

// Write an 8-bitutil value to associated device connected to io
// 0x0000 to 0x7FFF		ROM (Handled by cart)
// 0x8000 to 0x9FFF		VRAM
// 0xA000 to 0xBFFF		External RAM (Handled by cart)
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
		b.cart.Write(address, value)
	// VRAM
	case address <= MaxAddrVRam:
		b.ppu.Write(address, value)
	// External RAM
	case address <= 0xBFFF:
		b.cart.Write(address, value)
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
		b.ppu.Write(address, value)
	// Unusable
	case address <= 0xFEFF:
		// Do nothing
		// TODO add behaviour related to OAM access
		logrus.WithFields(logrus.Fields{
			"Address": address,
			"Value":   value,
		}).Warn("bus: Writing to unusable memory")

	// IO

	// When Divider Register is accessed, it is reset
	// Use TimeReg method if change is needed
	case address == AddrDiv:
		b.Time.div = 0
	case address == AddrTima:
		b.Time.tima = value
	case address == AddrTma:
		b.Time.tma = value
	case address == AddrTac:
		b.Time.tac = value
	case address == AddrIF:
		b.IF = IF(value)
	case address >= MinAddrLcdIO && address <= MaxAddrLcdIO:
		b.ppu.Write(address, value)

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
