package io

import "os"

// Represents a Cartridge ROM IO
type Cartridge struct {
	rom [0x8000]byte
}

// Open Rom's File
func (c *Cartridge) OpenRom(path string) error {
	// var err error
	output, err := os.ReadFile(path)
	copy(c.rom[:], output)

	return err
}

func (c Cartridge) Read(address uint16) uint8 {
	return c.rom[address]
}

func (c Cartridge) Write(address uint16, value uint8) {
	c.rom[address] = value
}
