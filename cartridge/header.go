package cartridge

import (
	"fmt"
	"github.com/aalquaiti/gbgo/gbgoutil"
	log "github.com/sirupsen/logrus"
)

const (
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
	NewLicensee string
	CartType    uint8
	DestCode    uint8
	OldLicensee uint8
	RomCode     uint8
	RamCode     uint8
)

var (
	newLicenseeMap = map[NewLicensee]string{
		"00": "None", "01": "Nintendo R&D1", "08": "Capcom", "12": "Electronic Arts", "18": "Hudson Soft",
		"19": "b-ai", "20": "kss",
	}
	cartTypeMap = map[CartType]string{
		00: "ROM ONLY", 01: "MBC1", 02: "MBC1+RAM", 03: "MBC1+RAM+BATTERY",
	}
	oldLicenseeMap = map[OldLicensee]string{
		0x00: "none", 0x01: "nintendo", 0x08: "capcom", 0x09: "hot-b", 0x0A: "jaleco", 0x0B: "coconuts",
		0x0C: "elite systems", 0x13: "electronic arts", 0x18: "hudsonsoft", 0x19: "itc entertainment", 0x1A: "yanoman",
		0x1D: "clary", 0x1F: "virgin", 0x24: "pcm complete", 0x25: "san-x", 0x28: "kotobuki systems", 0x29: "seta",
		0x30: "infogrames", 0x31: "nintendo", 0x32: "bandai", 0x33: "OTHER", 0x34: "konami", 0x35: "hector",
		0x38: "capcom", 0x39: "banpresto",
	}
)

// Header Represents Cartridge Header
type Header struct {
	Title            string
	ManufacturerCode string
	CGBFlag          uint8
	NewLicensee      NewLicensee
	SGBFlag          uint8
	CartType         CartType
	RomCode          RomCode
	RamCode          RamCode
	DestinationCode  DestCode
	OldLicensee      OldLicensee
	RomVersion       uint8
	HeaderChecksum   uint8
	GlobalChecksum   uint16
}

func (n NewLicensee) String() string {
	if val, ok := newLicenseeMap[n]; ok {
		return val
	}

	return "N/A"
}

func (c CartType) String() string {
	if c.IsSupported() {
		return cartTypeMap[c]
	}
	log.WithField("CartType Header", string(c)).Warn("Header value not supported")

	return "Not Supported"
}

// IsSupported determines if Cartridge type MBC is supported
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

// IsSupported determines if Cartridge Size header within supported values
func (r RomCode) IsSupported() bool {
	// Currently only up to code $08 is supported
	// TODO: Optional: Look into supporting Codes $52, $53 and $54
	return r <= 8
}

// GetBankSize retrieve no. of banks of Cartridge Needed.
func (r RomCode) GetBankSize() uint8 {

	// Assume not supported in this case
	if r > 8 {
		log.Panicf("RomCode Header has unsupported value of %d", r)
	}

	return 2 << r
}

func (r RomCode) String() string {

	return fmt.Sprintf("%d KB", 16*r.GetBankSize())
}

// IsSupported Determine if external Ram supported
func (r RamCode) IsSupported() bool {
	return r <= 5
}

// GetBankSize retrieve no. of banks of Ram Needed. It assumes Ram is supported
func (r RamCode) GetBankSize() uint8 {
	switch r {
	case 0:
		return 0
	case 2:
		return 1
	case 3:
		return 4
	case 4:
		return 16
	case 5:
		return 8
	}

	// It is assumed this line would be reached as external ram is supported
	return 0
}

func (r RamCode) String() string {
	if !r.IsSupported() {
		return "N//a"
	}

	return fmt.Sprintf("%d KB", 8*r.GetBankSize())
}

// NewHeader creates Cartridge Header by reading byte slice for information
// Title 				0x134 to 0x143 (16 chars) in old titles.
//						0x134 to 0x13E (11 chars) in CGB Mode.
// Manufacturer Code	0x13F to0x142 in newer Cartridge. This area is part of title in older cartridges
// CGB Flag				0x143
// New Licensee Code	0x144 - 0x145 as two ASCII chars
// SGB Flag				0x146
// Cartridge CartType	0x147
// ROM Size	Code		0x148
// RAM Size	Code		0x149
// Destination Code 	0x14A
// Old Licensee Code 	0x14B
// Cartridge Version Number 	0x14C
// Header Checksum		0x14D
// Global Checksum		0x14E - 0x14F
// returns error if Header info are not supported, such as in the case of an unsupported MBC
func NewHeader(file []byte) (*Header, error) {
	h := new(Header)

	// TODO implement how CGB handles titles
	// https://gbdev.io/pandocs/The_Cartridge_Header.html
	h.Title = asciiToStr(file[titleAddr:], oldTitleSize)
	h.ManufacturerCode = asciiToStr(file[manufacturerCodeAddr:], manufacturerCodeSize)
	h.CGBFlag = file[cgbFlagAddr]
	h.NewLicensee = NewLicensee(asciiToStr(file[newLicenseeAddr:], newLicenseeSize))
	h.SGBFlag = file[sgbFlagAddr]
	h.CartType = CartType(file[cartTypeAddr])
	if !h.CartType.IsSupported() {
		return nil, ErrorType
	}
	h.RomCode = RomCode(file[romSizeAddr])
	h.RamCode = RamCode(file[ramSizeAddr])
	h.DestinationCode = DestCode(file[destCodeAddr])
	h.RomVersion = file[romVerAddr]
	h.HeaderChecksum = file[headerChecksumAddr]
	h.GlobalChecksum = gbgoutil.To16(file[globalChecksumAddr+1], file[globalChecksumAddr])

	return h, nil
}

func (h Header) String() string {
	return fmt.Sprintf("cart Header {Title: %s, ManufacturerCode: %s, CGBFlag: %d, NewLicense: %s, SGBFlag: %d, "+
		"CartType: %s, RomCode: %s, RamCode: %s, DestinationCode: %s, OldLicensee: %s, RomVersion: $%.2X, "+
		"HeaderCheckSum: $%.2X, GlobalChecksum: $%.4X}", h.Title, h.ManufacturerCode, h.CGBFlag, h.NewLicensee,
		h.SGBFlag, h.CartType, h.RomCode, h.RamCode, h.DestinationCode, h.OldLicensee, h.RomVersion, h.HeaderChecksum,
		h.GlobalChecksum)
}
