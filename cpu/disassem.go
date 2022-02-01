package cpu

import (
	"fmt"
	"github.com/aalquaiti/gbgo/gbgoutil"
	"github.com/pkg/errors"
	"io"
)

type Options struct {
}

type dsmFunc [OPCodeSize]func() string

type Disassembler struct {
	op      OpCode
	arg1    byte
	arg2    byte
	funcs   dsmFunc
	cpFuncs dsmFunc
}

func NewDisassembler() *Disassembler {
	dsm := &Disassembler{}
	dsm.init(opCodes, &dsm.funcs)
	dsm.init(cpOpCodes, &dsm.cpFuncs)

	return dsm
}

// init Initialise Disassembler functions
func (d *Disassembler) init(codes [OPCodeSize]OpCode, funcs *dsmFunc) {

	for _, op := range codes {
		switch {
		// No Operands
		case op.oprs[0] == nil && op.oprs[1] == nil:
			funcs[op.code] = d.mnc

		// One Operand
		case op.oprs[0] != nil && op.oprs[1] == nil:
			switch op.oprs[0].(type) {
			case OprReg:
				funcs[op.code] = d.mncReg
			case OprConst8:
				funcs[op.code] = d.mncConst8
			case OprConst16:
				funcs[op.code] = d.mncConst16
			case OprFlag:
				funcs[op.code] = d.mncFlag
			case OprVec:
				funcs[op.code] = d.mncVec
			case OprBit:
				funcs[op.code] = d.mncBitReg
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
					funcs[op.code] = d.mncConst8Reg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprConst16:
				switch op.oprs[1].(type) {
				case OprReg:
					funcs[op.code] = d.mncConst16Reg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprReg:
				switch op.oprs[1].(type) {
				case OprConst8:
					funcs[op.code] = d.mncRegConst8
				case OprConst16:
					funcs[op.code] = d.mncRegConst16
				case OprReg:
					funcs[op.code] = d.mncRegReg
				default:
					// This should not happen
					panic("op codes was not initialised appropriately")
				}
			case OprFlag:
				switch op.oprs[1].(type) {
				case OprConst8:
					funcs[op.code] = d.mncFlagConst8
				case OprConst16:
					funcs[op.code] = d.mncFlagConst16
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

func (d *Disassembler) DisassembleAll(r io.ByteReader) ([]string, error) {
	var result []string

	var err error
	var line string
	for {
		line, err = d.Disassemble(r)
		if err != nil {
			break
		}
		result = append(result, line)
	}

	if err != nil {
		if errors.Is(err, io.EOF) {
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

		// io.EOF to be handled by caller for first read byte, as it indicates no other instructions left in
		// io.Reader. If EOF is received after reading an opcode that has a length > 2, that would indicate
		// not having all info needed for an opcode disassembling, and an error is returned
		//
		if errors.Is(err, io.EOF) {
			return "", err
		}

		return "", errors.Wrap(err, "cpu/disassembler: Could not read byte for disassembling")
	}
	// TODO add support for cpOpCode
	d.op = opCodes[value]

	// Reading the first operand (if any)
	if d.op.length > 1 {
		value, err = r.ReadByte()
		if err != nil {
			return "", errors.Wrapf(err,
				"cpu/disassembler: Could not read first byte for opcode %.2X", d.op.code)
		}
		d.arg1 = value

	}

	// Reading second operand (if any)
	if d.op.length > 2 {
		value, err = r.ReadByte()
		if err != nil {
			return "", errors.Wrapf(err,
				"cpu/disassembler: Could not read second byte for opcode %.2X", d.op.code)
		}
		d.arg2 = value
	}

	// Checking if the operand is Prefix CF
	if d.op.mnc == PrefixCB {
		value, err = r.ReadByte()
		d.op = cpOpCodes[value]
		if err != nil {
			return "", errors.Wrapf(err,
				"cpu/disassembler: Could not read first byte for opcode %.2X", d.op.code)
		}

		return d.cpFuncs[d.op.code](), nil
	}

	return d.funcs[d.op.code](), nil
}

// mnc Mnemonic (without operand)
func (d *Disassembler) mnc() string {
	return d.op.mnc.String()
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
		d.op.mnc, OprConst8Text(d.op.oprs[0], d.arg1), OprRegText(d.op.oprs[1]))
}

func (d *Disassembler) mncConst16Reg() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprConst16Text(d.op.oprs[0], gbgoutil.To16(d.arg2, d.arg1)), OprRegText(d.op.oprs[1]))
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

func (d *Disassembler) mncBitReg() string {
	return fmt.Sprintf("%s %s, %s",
		d.op.mnc, OprBitText(d.op.oprs[0]), OprRegText(d.op.oprs[1]))
}
