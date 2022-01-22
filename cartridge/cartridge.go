package cartridge

import "C"
import (
	gbio "github.com/aalquaiti/gbgo/io"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

// Cartridge Represents a ROM Device
// Cartridge implements io.ByteReader
type Cartridge struct {
	file   []byte
	Rom    [][]uint8
	Ram    []uint8
	Header *Header
	mbc    gbio.Device

	pos int // Used to point to byte position for ReadByte
}

const (
	CartTypeRomOnly    CartType = 0x00
	CartTypeMBC1       CartType = 0x01
	CartTypeMBC1Ram    CartType = 0x02
	CartTypeMBC1RamBat CartType = 0x03
	DestJapanese       DestCode = 00
	DestNonJapanese    DestCode = 01
)

var mbcFunc = map[CartType]func(*Cartridge) (gbio.Device, error){
	CartTypeRomOnly:    newMbc0,
	CartTypeMBC1:       newMbc1,
	CartTypeMBC1Ram:    newMbc1,
	CartTypeMBC1RamBat: newMbc1,
}

// errors
var (
	ErrorType = errors.New("cartridge: type not supported")
	ErrorMbc  = errors.New("cartridge: mbc not supported")
)

// NewCartridge Reads a ROM file, extract header information and return Cartridge with appropriate MBC accordingly
// returns error if rom file corrupted or not supported
func NewCartridge(path string) (*Cartridge, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.Wrap(err, "cartridge: could not be opened")
	}

	header, err := NewHeader(file)
	if err != nil {
		return nil, errors.Wrap(err, "cartridge: header is corrupted or unsupported")
	}
	cart := new(Cartridge)
	cart.file = file
	cart.Header = header

	if mbc, ok := mbcFunc[header.CartType]; !ok {
		return nil, ErrorMbc
	} else {
		cart.mbc, err = mbc(cart)
	}

	return cart, err
}

func (c *Cartridge) Read(address uint16) uint8 {
	return c.mbc.Read(address)
}

func (c *Cartridge) Write(address uint16, value uint8) {
	c.mbc.Write(address, value)
}

func (c *Cartridge) Reset() {
	c.mbc.Reset()
}

// AsciiToStr Convert Byte Slice to String
func asciiToStr(src []byte, length int) string {
	sb := strings.Builder{}
	str := make([]byte, length)
	copy(str, src)

	sb.Write(str)

	return sb.String()
}

func (c *Cartridge) ReadByte() (byte, error) {
	if c.pos == len(c.file) {
		return 0, io.EOF
	}

	result := c.file[c.pos]
	c.pos++

	return result, nil
}
