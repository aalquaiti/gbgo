package ppu

import (
	"github.com/aalquaiti/gbgo/io"
	log "github.com/sirupsen/logrus"
)

// ppuRegMask masks address to make it within Range of PPU Register address
const ppuRegMask = 0x0F

type PPU struct {
	vram [io.VRamSize]uint8
	oam  [io.OamSize]uint8
	reg  Reg

	lx uint8
}

func NewPPU() PPU {
	return PPU{}
}

func (p PPU) Read(address uint16) uint8 {

	switch {
	case address <= io.MaxAddrVRam:
		return p.vram[address&(io.VRamSize-1)]
	case address <= io.MaxAddrOam:
		return p.oam[address&(io.OamSize-1)]
	case address >= io.MinAddrLcdIO && address <= io.MaxAddrLcdIO:
		return p.reg.val[address&ppuRegMask]
	default:
		// Should never reach this line
		log.Errorf("ppu: read in unreachable address $%.4X", address)
		return 0
	}
}

func (p PPU) Write(address uint16, value uint8) {
	switch {
	case address <= io.MaxAddrVRam:
		p.vram[address&(io.VRamSize-1)] = value
	case address <= io.MaxAddrOam:
		p.oam[address&(io.OamSize-1)] = value
	// Registers
	case address >= io.MinAddrLcdIO && address <= io.MaxAddrLcdIO:
		p.reg.val[address&ppuRegMask] = value
	default:
		// Should never reach this line
		log.Errorf("ppu: write in unreachable address $%.4X", address)
	}
}

func (p PPU) Reset() {
	//TODO implement me
	panic("implement me")
}

func (p *PPU) Tick() {
	p.lx++

	if p.lx == 154 {
		p.lx = 0
		p.reg.IncLY()
	}
}
