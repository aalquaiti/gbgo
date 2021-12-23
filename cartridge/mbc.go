package cartridge

import "C"
import (
	"github.com/aalquaiti/gbgo/io"
	"github.com/pkg/errors"
)

const (
	romBankSize         = 0x4000
	romBankMaxAddr      = romBankSize - 1
	bank0MaxAddr        = 0x3FFF
	bank1MaxAddr        = 0x7FFF
	ramBankSize         = 0x2000
	ramEnableRegMaxAddr = 0x1FFF
	romBankRegMaxAddr   = 0x3FFF
	ramBankRegMaxAddr   = 0x5FFF
	bankModeMaxAddr     = 0x7FFF
	externalRamMaxAddr  = 0xBFFF
)

type Mbc struct {
	Header *Header
	Rom    [][romBankSize]byte
	Ram    [][ramBankSize]byte
}

type Mbc0 Mbc

type BankMode uint8

const (
	BankModeSimple  BankMode = 0
	BankModeAdvance BankMode = 1
)

type Mbc1 struct {
	Mbc
	RamEnabled    bool
	RomBank       uint8
	SecondaryBank uint8
	BankMode      BankMode
}

const cartErrorMsg = "cartridge: rom file corrupted"

func (m *Mbc0) Read(address uint16) uint8 {

	if address <= bank1MaxAddr {
		bank := address & romBankSize
		address &= romBankSize - 1
		return m.Rom[bank][address]
	}

	// No support for external ram
	return 0
}

func (m *Mbc0) Write(address uint16, value uint8) {
	// MBC0 has not bank switch or external ram
	// Write nothing
}

func (m *Mbc0) Reset() {

}

func newMbc0(c *Cartridge) (io.Device, error) {
	mbc := &Mbc0{Header: c.Header}

	if err := mbc.validate(); err != nil {
		return nil, err
	}

	mbc.Rom = make([][romBankSize]byte, c.Header.RomCode.GetBankSize())
	for i := 0; i < len(mbc.Rom); i++ {
		// sub slice of 16 KB data to copy from file to each bank
		start := romBankSize * i
		end := romBankSize * (i + 1)
		copy(mbc.Rom[i][:], c.file[start:end])
	}

	return mbc, nil
}

func (m *Mbc0) validate() error {

	// MBC0 should support no bank switch
	if m.Header.RomCode != 0 {
		return errors.New(cartErrorMsg)
	}

	return nil
}

func (m *Mbc1) Read(address uint16) uint8 {

	switch {
	// Bank Zero (or Others in Bank Mode Advance)
	case address <= bank0MaxAddr:
		if m.BankMode == BankModeAdvance {
			return m.Rom[m.GetSelectedRomBank()][address]
		}
		return m.Rom[0][address]
	// Bank One and Up
	case address <= bank1MaxAddr:
		address &= romBankMaxAddr
		// ROM Bank $00 must be accessed from Cartridge[0], so selecting it leads to increment to 1.
		// This is so the selected Cartridge Bank cannot be Cartridge[0].
		// Setting Bank Mode to $1 allows Cartridge[0] to remap the area of the bank zero ($0000 to $3FFF), which
		// allows access to banks in large ROM, such as bank $20, $40 and $60
		selected := m.GetSelectedRomBank()
		if m.RomBank == 0 {
			selected++
		}

		return m.Rom[selected][address]
	case address <= externalRamMaxAddr:
		address &= ramBankSize - 1
		if m.RamEnabled {
			return m.Ram[m.SecondaryBank][address]
		}
	}

	return 0
}

func (m *Mbc1) Write(address uint16, value uint8) {
	switch {
	// RAM Enable Register
	case address <= ramEnableRegMaxAddr:

		value &= 0b1111
		// Any written value in the first four bits equals $A will enable ram
		if value == 0xA {
			m.RamEnabled = true
		} else {
			m.RamEnabled = false
		}

	// ROM Bank Number
	case address <= romBankRegMaxAddr:
		// Reads the first five bits
		value &= 0b11111
		// If value written is higher than number of rom banks, it will be masked to required bits
		// E.g: Cartridge Bank Size is of 256 KB (i.e. 32 rom banks) which needs four bits, and value written is higher,
		// value will be masked to four bits
		mask := m.Header.RomCode.GetBankSize() - 1
		m.RomBank = value & mask

	// RAM bank Number
	// OR
	// Upper Bits of ROM Bank (Secondary ROM Bank)
	case address <= ramBankRegMaxAddr:
		// Reads the first two bits
		value &= 0b11
		m.SecondaryBank = value
		// TODO ignore setting if ROM and RAM are not big enough
		// TODO add support for 1MB Multi-game cartridge

	// Banking Mode
	case address <= bankModeMaxAddr:
		// Reads the first bit
		value &= 0b1
		m.BankMode = BankMode(value)
	// External Ram
	case address <= externalRamMaxAddr:
		address &= ramBankSize - 1
		// TODO check that ram writing is enabled
		if m.RamEnabled {
			m.Ram[m.SecondaryBank][address] = value
		}
	}
}

func (m *Mbc1) GetSelectedRomBank() uint8 {
	// Selected Cartridge comes from Cartridge Bank (for first five bits) and secondary rom bank (for more than 5 bits)
	return m.SecondaryBank<<5 + m.RomBank
}

func (m *Mbc1) Reset() {
	m.RomBank = 1 // Default Cartridge Bank
}

func newMbc1(c *Cartridge) (io.Device, error) {
	mbc := &Mbc1{Mbc: Mbc{Header: c.Header}}

	if err := mbc.validate(); err != nil {
		return nil, err
	}

	mbc.Rom = make([][romBankSize]byte, c.Header.RomCode.GetBankSize())
	for i := 0; i < len(mbc.Rom); i++ {
		// sub slice of 16 KB data to copy from file to each bank
		start := romBankSize * i
		end := romBankSize * (i + 1)
		copy(mbc.Rom[i][:], c.file[start:end])
	}

	mbc.Ram = make([][ramBankSize]byte, c.Header.RamCode.GetBankSize())

	return mbc, nil
}

func (m *Mbc1) validate() error {

	// MBC0 has maximum of 2 MB ROM and 32 KB RAM

	// Cartridge Code up to $6 support 2 MB of ROM
	// Ram Code up to $3 support 32 KB of RAM
	// TODO: Turn RomCode and RamCode to enum equivalent for easier reading of code
	// Use this as reference : https://gbdev.io/pandocs/The_Cartridge_Header.html#0148---rom-size
	if m.Header.RomCode > 6 || m.Header.RamCode > 3 {
		return errors.New(cartErrorMsg)
	}

	return nil
}
