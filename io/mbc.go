package io

import "github.com/pkg/errors"

type Mbc struct {
	Rom [][romBankSize]byte
	Ram []byte
}

type Mbc0 Mbc

func (m *Mbc0) Read(address uint16) uint8 {

	if address <= bank1MaxAddr {
		bank := address & romBankSize
		address &= romBankSize - 1
		return m.Rom[bank][address]
	}

	//switch {
	//case address <= bank0MaxAddr:
	//	return m.Rom[0][address]
	//case address < bank1MaxAddr:
	//	return m.Rom[1][address]
	//}

	// No support for external ram
	return 0
}

func (m *Mbc0) Write(address uint16, value uint8) {

	// MBC0 has not bank switch or external ram
	// Write nothing
}

func newMbc0(c *Cartridge) error {
	m := Mbc0{}

	if err := m.validate(c.Header); err != nil {
		return err
	}

	m.Rom = make([][romBankSize]byte, c.Header.RomCode.GetBankSize())
	for i := 0; i < int(c.Header.RomCode.GetBankSize()); i++ {
		//m.Rom[i] = make([]byte, romBankSize)
		copy(m.Rom[i][:], c.file[romBankSize*i:romBankSize*(i+1)])
		//m.Rom[i] = ([romBankSize]byte)(file[romBankSize*i : romBankSize*i+1])
	}
	c.mbc = &m

	return nil
}

func (m *Mbc0) validate(header *Header) error {

	// MBC0 should support no bank switch
	if header.RomCode != 0 {
		return errors.New("rom file corrupted")
	}

	return nil
}
