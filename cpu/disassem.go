package cpu

import (
	"fmt"
	"github.com/aalquaiti/gbgo/gbgoutil"
	"github.com/pkg/errors"
	"io"
)

type Options struct {
}

type Disassembler struct {
	op    OpCode
	arg1  byte
	arg2  byte
	funcs [OPCodeSize]func() string
}

func NewDisassembler() *Disassembler {
	dsm := &Disassembler{}
	dsm.init()

	return dsm
}

func (d *Disassembler) DisassembleAll(r io.ByteReader) ([]string, error) {
	var result []string

	var err error
	var line string
	for line, err = d.Disassemble(r); err == nil; {
		result = append(result, line)
	}

	if err != nil {
		if errors.As(err, io.EOF) {
			return result, nil
		} else {
			return nil, errors.Wrap(err, "cpu/disassembler: Failed while disassembling")
		}
	}

	return result, nil
}

func (d *Disassembler) Disassemble(r io.ByteReader) (string, error) {
	value, err := r.ReadByte()
	if err != nil {
		// io.EOF to be handled by caller for first read byte
		// This is as EOF has been reached before reading any opcode
		if errors.As(err, io.EOF) {
			return "", err
		}

		return "", errors.Wrap(err, "cpu/disassembler: Could not read byte for disassembling")
	}
	// TODO add support for cpOpCode
	d.op = opCodes[value]
	if d.op.oprs[0] != nil {
		value, err = r.ReadByte()
		if err != nil {
			return "", errors.Wrap(err, "cpu/disassembler: Could not read byte for first operand")
		}
		d.arg1 = value

	}
	if d.op.oprs[1] != nil {
		value, err = r.ReadByte()
		if err != nil {
			return "", errors.Wrap(err, "cpu/disassembler: Could not read byte for second operand")
		}
		d.arg2 = value
	}

	return d.funcs[d.op.code](), nil
}

// init Initialise Disassembler functions
func (d *Disassembler) init() {

	for _, op := range opCodes {
		switch {
		// No Operands
		case op.oprs[0] == nil && op.oprs[1] == nil:
			d.funcs[op.code] = d.mnc

		// One Operand
		case op.oprs[0] != nil && op.oprs[1] == nil:
			switch op.oprs[0].(type) {
			case OprReg:
				d.funcs[op.code] = d.mncReg
			case OprConst8:
				d.funcs[op.code] = d.mncConst8
			case OprConst16:
				d.funcs[op.code] = d.mncConst16
			case OprFlag:
				d.funcs[op.code] = d.mncFlag
			case OprVec:
				d.funcs[op.code] = d.mncVec
			default:
				// This should not happen
				panic("op codes was not initialised appropriately")
			}

		// Two Operands
		case op.oprs[0] != nil && op.oprs[1] != nil:
			switch op.oprs[0].(type) {

			case OprConst8:
				switch op.oprs[1].(type) {
				case OprReg:
					d.funcs[op.code] = d.mncConst8Reg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprConst16:
				switch op.oprs[1].(type) {
				case OprReg:
					d.funcs[op.code] = d.mncConst16Reg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprReg:
				switch op.oprs[1].(type) {
				case OprConst8:
					d.funcs[op.code] = d.mncRegConst8
				case OprConst16:
					d.funcs[op.code] = d.mncRegConst16
				case OprReg:
					d.funcs[op.code] = d.mncRegReg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprFlag:
				switch op.oprs[1].(type) {
				case OprConst8:
					d.funcs[op.code] = d.mncFlagConst8
				case OprConst16:
					d.funcs[op.code] = d.mncFlagConst16
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			}

		default:
			// This should not happen
			panic("op codes was not initialised appropriately")
		}
	}
}

// mnc Mnemonic (without operand)
func (d *Disassembler) mnc() string {
	return string(d.op.mnc)
}

func (d *Disassembler) mncConst8() string {
	return fmt.Sprintf("%s %s",
		d.op.mnc, OprConst8Text(d.op.oprs[0], d.arg1))
}

func (d *Disassembler) mncConst16() string {
	return fmt.Sprintf("%s %s", d.op.mnc, OprConst16Text(d.op.oprs[0], gbgoutil.To16(d.arg2, d.arg1)))
}

func (d *Disassembler) mncReg() string {
	return fmt.Sprintf("%s %s", d.op.mnc, OprRegText(d.op.oprs[0]))
}

func (d *Disassembler) mncFlag() string {
	return fmt.Sprintf("%s %s", d.op.mnc, OprFlagText(d.op.oprs[0]))
}

func (d *Disassembler) mncVec() string {
	return fmt.Sprintf("%s %s", d.op.mnc, OprVecText(d.op.oprs[0]))
}

func (d *Disassembler) mncConst8Reg() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprConst8Text(d.op.oprs[1], d.arg1), OprRegText(d.op.oprs[0]))
}

func (d *Disassembler) mncConst16Reg() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprConst16Text(d.op.oprs[1], gbgoutil.To16(d.arg2, d.arg1)), OprRegText(d.op.oprs[0]))
}

func (d *Disassembler) mncRegConst8() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprRegText(d.op.oprs[0]), OprConst8Text(d.op.oprs[1], d.arg1))
}

func (d *Disassembler) mncRegConst16() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprRegText(d.op.oprs[0]), OprConst16Text(d.op.oprs[1], gbgoutil.To16(d.arg2, d.arg1)))
}

func (d *Disassembler) mncRegReg() string {
	return fmt.Sprintf("%s %s, %s", d.op.mnc, OprRegText(d.op.oprs[0]), OprRegText(d.op.oprs[1]))
}

func (d *Disassembler) mncFlagConst8() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprFlagText(d.op.oprs[0]), OprConst8Text(d.op.oprs[1], d.arg1))
}

func (d *Disassembler) mncFlagConst16() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprFlagText(d.op.oprs[0]), OprConst16Text(d.op.oprs[1], gbgoutil.To16(d.arg2, d.arg1)))
}
