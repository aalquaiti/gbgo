package io

import "C"
import (
	"fmt"
	"github.com/aalquaiti/gbgo/bit"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	romBankSize          = 0x4000
	titleAddr            = 0x134
	oldTitleSize         = 16
	manufacturerCodeAddr = 0x13F
	manufacturerCodeSize = 4
	cgbFlagAddr          = 0x143
	newLicenseeAddr      = 0x144
	newLicenseeSize      = 2
	sgbFlagAddr          = 0x146
	cartTypeAddr         = 0x147
	romSizeAddr          = 0x148
	ramSizeAddr          = 0x149
	destCodeAddr         = 0x14A
	romVerAddr           = 0x1
	headerChecksumAddr   = 0x14D
	globalChecksumAddr   = 0x14E
)

type (
	CartType    uint8
	DestCode    uint8
	OldLicensee uint8
)

const (
	CartTypeRomOnly    CartType = 0x00
	CartTypeMBC1       CartType = 0x01
	CartTypeMBC1Ram    CartType = 0x02
	CartTypeMBC1RamBat CartType = 0x03
	DestJapanese       DestCode = 00
	DestNonJapanese    DestCode = 01
)

var (
	cartTypeMap = map[CartType]string{
		00: "ROM ONLY", 01: "MBC1", 02: "MBC1+RAM", 03: "MBC1+RAM+BATTERY",
	}
	oldLicenseeMap = map[OldLicensee]string{
		0x00: "none", 0x01: "nintendo", 0x08: "capcom", 0x09: "hot-b", 0x0A: "jaleco", 0x0B: "coconuts",
		0x0C: "elite systems",
	}
)

func (c CartType) String() string {

	if c.IsSupported() {
		return cartTypeMap[c]
	}
	return "Not Supported"
}

func (c CartType) IsSupported() bool {
	if _, ok := cartTypeMap[c]; ok {
		return true
	}

	return false
}

func (d DestCode) String() string {
	switch d {
	case DestJapanese:
		return "Japanese"
	case DestNonJapanese:
		return "Non-Japanese"
	default:
		return "Unknown"
	}
}

func (l OldLicensee) String() string {
	if val, ok := oldLicenseeMap[l]; ok {
		return val
	}

	return "Unknown"
}

// Header Represents Cartridge Header
type Header struct {
	Title            string
	ManufacturerCode string
	CGBFlag          uint8
	NewLicensee      string
	SGBFlag          uint8
	CartType         CartType
	RomSize          uint8
	RamSize          uint8
	DestinationCode  DestCode
	OldLicensee      OldLicensee
	RomVersion       uint8
	HeaderChecksum   uint8
	GlobalChecksum   uint16
}

// OpenRom Reads a ROM file, extract header information and set MBC accordingly
func (h *Cartridge) OpenRom(path string) error {
	var err error
	h.file, err = os.ReadFile(path)

	if err != nil {
		return errors.Wrap(err, "cartridge could not be opened")
	}

	h.Header, err = NewHeader(h.file)

	return err
}

// Cartridge Represents a ROM Device
type Cartridge struct {
	file   []byte
	Rom    [][romBankSize]uint8
	Ram    []uint8
	Header *Header
}

// NewHeader creates Cartridge Header by reading byte slice for information
// Title 				0x134 to 0x143 (16 chars) in old titles.
//						0x134 to 0x13E (11 chars) in CGB Mode.
// Manufacturer Code	0x13F to0x142 in newer Cartridge. This area is part of title in older cartridges
// CGB Flag				0x143
// OldLicensee Code		0x144 - 0x145 as two ASCII chars
// SGB Flag				0x146
// Cartridge CartType	0x147
// ROM Size				0x148
// RAM Size				0x149
// Destination Code 	0x14A
// Old OldLicensee Code 	0x14B
// Rom Version Number 	0x14C
// Header Checksum		0x14D
// Global Checksum		0x14E - 0x14F
// returns error if Header info are not supported, such as in the case of an unsupported MBC
func NewHeader(file []byte) (*Header, error) {
	h := new(Header)

	// TODO implement how CGB handles titles
	// https://gbdev.io/pandocs/The_Cartridge_Header.html
	h.Title = AsciiToStr(file[titleAddr:], oldTitleSize)
	h.ManufacturerCode = AsciiToStr(file[manufacturerCodeAddr:], manufacturerCodeSize)
	h.CGBFlag = file[cgbFlagAddr]
	h.NewLicensee = AsciiToStr(file[newLicenseeAddr:], newLicenseeSize)
	h.SGBFlag = file[sgbFlagAddr]
	h.CartType = CartType(file[cartTypeAddr])
	if !h.CartType.IsSupported() {
		return nil, errors.New("cartridge Type Not supported")
	}
	h.RomSize = file[romSizeAddr]
	h.RamSize = file[ramSizeAddr]
	h.DestinationCode = DestCode(file[destCodeAddr])
	h.RomVersion = file[romVerAddr]
	h.HeaderChecksum = file[headerChecksumAddr]
	h.GlobalChecksum = bit.To16(file[globalChecksumAddr+1], file[globalChecksumAddr])

	return h, nil
}

// GetRomSize Retrieve Size in KB
func (h *Header) GetRomSize() uint16 {
	return romBankSize / 0x400 * (2 << h.RomSize)
}

func (h Header) String() string {
	return fmt.Sprintf("Cartridge {Title: %s, ManufacturerCode: %s, CGBFlag: %d, NewLicense: %s, SGBFlag: %d, "+
		"CartType: %s, RomSize: %d KB, RamSize: %d KB, DestinationCode: %s, OldLicensee: %s, RomVersion: $%.2X, "+
		"HeaderCheckSum: $%.2X, GlobalChecksum: $%.4X}", h.Title, h.ManufacturerCode, h.CGBFlag, h.NewLicensee,
		h.SGBFlag, h.CartType, h.GetRomSize(), h.RamSize, h.DestinationCode, h.OldLicensee, h.RomVersion, h.HeaderChecksum,
		h.GlobalChecksum)
}

func (h *Cartridge) Read(address uint16) uint8 {
	// TODO implement Cartridge external Ram
	if address > MaxRomAddr {
		log.WithField("Address", address).Warn("Cartridge RAM not Supported")
		return 0
	}

	return h.file[address]
}

func (h Cartridge) Write(address uint16, value uint8) {
	// TODO implement Cartridge external Ram
	if address > MaxRomAddr {
		log.WithField("Address", address).Warn("Cartridge RAM not Supported")
	} else {
		log.WithField("Address", address).Warn("Writing to ROM")
	}
}

// AsciiToStr Convert Byte Slice to String
func AsciiToStr(src []byte, length int) string {
	sb := strings.Builder{}
	str := make([]byte, length)
	copy(str, src)

	sb.Write(str)

	return sb.String()
}
