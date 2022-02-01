package cpu

import "fmt"

type Mnemonic int

const (
	ADC Mnemonic = iota + 1
	ADD
	AND
	BIT
	CALL
	CCF
	CP
	CPL
	DAA
	DEC
	DI
	EI
	HALT
	INC
	JP
	JR
	LD
	LDH
	NOP
	OR
	POP
	PUSH
	RES
	RET
	RETI
	RL
	RLA
	RLC
	RLCA
	RR
	RRA
	RRC
	RRCA
	RST
	SBC
	SCF
	SET
	SLA
	SRA
	SRL
	STOP
	SUB
	SWAP
	XOR

	PrefixCB
	IllegalOp
)

var mncText = map[Mnemonic]string{
	ADC:  "ADC",
	ADD:  "ADD",
	AND:  "AND",
	BIT:  "BIT",
	CALL: "CALL",
	CCF:  "CCF",
	CP:   "CP",
	CPL:  "CPL",
	DAA:  "DAA",
	DEC:  "DEC",
	DI:   "DI",
	EI:   "EI",
	HALT: "HALT",
	INC:  "INC",
	JP:   "JP",
	JR:   "JR",
	LD:   "LD",
	LDH:  "LDH",
	NOP:  "NOP",
	OR:   "OR",
	POP:  "POP",
	PUSH: "PUSH",
	RES:  "RES",
	RET:  "RET",
	RETI: "RETI",
	RL:   "RL",
	RLA:  "RLA",
	RLC:  "RLC",
	RLCA: "RLCA",
	RR:   "RR",
	RRA:  "RRA",
	RRC:  "RRC",
	RRCA: "RRCA",
	RST:  "RST",
	SBC:  "SBC",
	SCF:  "SCF",
	SET:  "SET",
	SLA:  "SLA",
	SRA:  "SRA",
	SRL:  "SRL",
	STOP: "STOP",
	SUB:  "SUB",
	SWAP: "SWAP",
	XOR:  "XOR",

	// TODO handle PrefixCB and IllegalOp
}

func MncText(mnc Mnemonic) string {
	return mncText[mnc]
}

func (mnc Mnemonic) String() string {
	return mncText[mnc]
}

// Operand represents an OP Code Operand
type Operand interface {
	Operand() // Implementation to be ignored
}

type (
	OprConst8  int
	OprConst16 int
	OprReg     int
	OprFlag    int // Flag Condition Operand
	OprVec     int
	OprBit     int
)

func (opr OprConst8) Operand()  {}
func (opr OprConst16) Operand() {}
func (opr OprReg) Operand()     {}
func (opr OprFlag) Operand()    {}
func (opr OprVec) Operand()     {}
func (opr OprBit) Operand()     {}

const (
	OprU8   OprConst8 = iota + 1 // Immediate Unsigned 8-bit
	OprI8                        // Immediate Signed value (i.e. Relative value)
	OprFFU8                      // Indirect address of (FF00 + U8)
	OprSPI8                      // Value of SP +I8
)

const (
	OprU16 OprConst16 = iota + 1 // Immediate Unsigned 16-bit
	OprInd                       // Indirect Unsigned 16-bit
)

const (
	OprRegA OprReg = iota + 1
	OprRegB
	OprRegC
	OprIRegC // Indirect address of (FF00 + RegC)
	OprRegD
	OprRegE
	OprRegH
	OprRegL

	OprRegAF
	OprRegBC
	OprIRegBC
	OprRegDE
	OprIRegDE
	OprRegHL
	OprIRegHL
	OprIRegHLI // HL+
	OprIRegHLD // HL-
	OprRegSP
)

const (
	OprFlagZ OprFlag = iota + 1
	OprFlagNZ
	OprFlagC
	OprFlagNC
)

const (
	OprVec00 OprVec = iota + 1
	OprVec08
	OprVec10
	OprVec18
	OprVec20
	OprVec28
	OprVec30
	OprVec38
)

const (
	OprBit0 OprBit = iota
	OprBit1
	OprBit2
	OprBit3
	OprBit4
	OprBit5
	OprBit6
	OprBit7
)

var oprConst8Text = map[Operand]string{
	OprU8: "%.02X",
	OprI8: "%.02X",

	OprFFU8: "(FF00+%.02X)",
	OprSPI8: "SP + %.02X",
}

var oprConst16Text = map[Operand]string{
	OprU16: "%.04X",
	OprInd: "(%.04X)",
}

var oprRegText = map[Operand]string{
	OprRegA:    "A",
	OprRegB:    "B",
	OprRegC:    "C",
	OprRegD:    "D",
	OprRegE:    "E",
	OprRegH:    "H",
	OprRegL:    "L",
	OprRegAF:   "AF",
	OprRegBC:   "BC",
	OprIRegBC:  "(BC)",
	OprRegDE:   "DE",
	OprIRegDE:  "(DE)",
	OprRegHL:   "HL",
	OprIRegHL:  "(HL)",
	OprIRegHLI: "(HL+)",
	OprIRegHLD: "(HL-)",
	OprRegSP:   "SP",

	OprIRegC: "(FF00 + C)",
}

var oprFlagText = map[Operand]string{
	OprFlagZ:  "Z",
	OprFlagNZ: "NZ",
	OprFlagC:  "C",
	OprFlagNC: "NC",
}

var oprVecText = map[Operand]string{
	OprVec00: "00",
	OprVec08: "09",
	OprVec10: "10",
	OprVec18: "18",
	OprVec20: "20",
	OprVec28: "28",
	OprVec30: "30",
	OprVec38: "38",
}

var oprBitText = map[Operand]string{
	OprBit0: "0",
	OprBit1: "1",
	OprBit2: "2",
	OprBit3: "3",
	OprBit4: "4",
	OprBit5: "5",
	OprBit6: "6",
	OprBit7: "7",
}

// OperandText returns a text for OpCode Operand
func OperandText(code Operand) string {
	switch code.(type) {
	case OprConst8:
		return oprConst8Text[code]
	case OprConst16:
		return oprConst16Text[code]
	case OprReg:
		return oprRegText[code]
	case OprFlag:
		return oprFlagText[code]
	case OprVec:
		return oprVecText[code]
	}

	panic("cpu: Operand Not supported")
}

// OprConst8Text returns a text for OpCode Operand including an 8-bit value
func OprConst8Text(opr Operand, value uint8) string {
	return fmt.Sprintf(oprConst8Text[opr], value)
}

// OprConst16Text returns a text for OpCode Operand including an 16-bit value
func OprConst16Text(opr Operand, value uint16) string {
	return fmt.Sprintf(oprConst16Text[opr], value)
}

func OprRegText(opr Operand) string {
	return oprRegText[opr]
}

func OprFlagText(opr Operand) string {
	return oprFlagText[opr]
}

func OprVecText(opr Operand) string {
	return oprVecText[opr]
}

func OprBitText(opr Operand) string {
	return oprBitText[opr]
}

//
//// OperandVal16Text returns a text for OpCode Operand including an 16-bit value
//func OperandVal16Text(code Operand, value uint16) string {
//	return fmt.Sprintf(operandValText[code], value)
//}

const OPCodeSize = 0x100 // OpCode Size

// Operands represents an Operand array of size 2
type Operands [2]Operand

// OpCode defines a CPU instruction
type OpCode struct {
	_      struct{}
	code   uint8
	ticks  uint8 // how many m-ticks required
	mnc    Mnemonic
	oprs   Operands
	length uint8
}

var opCodes = [OPCodeSize]OpCode{
	// NOP
	{code: 0x00, ticks: 1, mnc: NOP, oprs: Operands{}, length: 1},
	// LD BC, $FFFF
	{code: 0x01, ticks: 1, mnc: LD, oprs: Operands{OprRegBC, OprU16}, length: 2},
	// LD (BC), A
	{code: 0x02, ticks: 2, mnc: LD, oprs: Operands{OprIRegBC, OprRegA}, length: 1},
	// INC BC
	{code: 0x03, ticks: 2, mnc: INC, oprs: Operands{OprRegBC}, length: 1},
	// INC B
	{code: 0x04, ticks: 1, mnc: INC, oprs: Operands{OprRegB}, length: 1},
	// DEC B
	{code: 0x05, ticks: 1, mnc: DEC, oprs: Operands{OprRegB}, length: 1},
	// LD B, $FF
	{code: 0x06, ticks: 2, mnc: LD, oprs: Operands{OprRegB, OprU8}, length: 2},
	// RLCA
	{code: 0x07, ticks: 1, mnc: RLCA, oprs: Operands{}, length: 1},
	// LD ($FFFF), SP
	{code: 0x08, ticks: 5, mnc: LD, oprs: Operands{OprInd, OprRegSP}, length: 3},
	// ADD HL, BC
	{code: 0x09, ticks: 2, mnc: ADD, oprs: Operands{OprRegHL, OprRegBC}, length: 1},
	//// LD A, (BC)
	{code: 0x0A, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprIRegBC}, length: 1},
	//// DEC BC
	{code: 0x0B, ticks: 2, mnc: DEC, oprs: Operands{OprRegBC}, length: 1},
	//// INC C
	{code: 0x0C, ticks: 1, mnc: INC, oprs: Operands{OprRegC}, length: 1},
	// DEC C
	{code: 0x0D, ticks: 1, mnc: DEC, oprs: Operands{OprRegC}, length: 1},
	// LD C, $FF
	{code: 0x0E, ticks: 2, mnc: LD, oprs: Operands{OprRegC, OprU8}, length: 2},
	// RRCA
	{code: 0x0F, ticks: 1, mnc: RRCA, oprs: Operands{}, length: 1},

	//// STOP
	{code: 0x10, ticks: 1, mnc: STOP, oprs: Operands{}, length: 1},
	// LD DE, $FFFF
	{code: 0x11, ticks: 3, mnc: LD, oprs: Operands{OprRegDE, OprU16}, length: 3},
	// LD (DE), A
	{code: 0x12, ticks: 2, mnc: LD, oprs: Operands{OprIRegDE, OprRegA}, length: 1},
	// INC DE
	{code: 0x13, ticks: 2, mnc: INC, oprs: Operands{OprRegDE}, length: 1},
	// INC D
	{code: 0x14, ticks: 1, mnc: INC, oprs: Operands{OprRegD}, length: 1},
	// DEC D
	{code: 0x15, ticks: 1, mnc: DEC, oprs: Operands{OprRegD}, length: 1},
	// LD D, $FF
	{code: 0x16, ticks: 2, mnc: LD, oprs: Operands{OprRegD, OprU8}, length: 2},
	// RLA
	{code: 0x17, ticks: 1, mnc: RLA, oprs: Operands{}, length: 1},
	// JR $FF
	{code: 0x18, ticks: 1, mnc: JR, oprs: Operands{OprI8}, length: 2},
	// ADD HL, DE
	{code: 0x19, ticks: 2, mnc: ADD, oprs: Operands{OprRegHL, OprRegDE}, length: 1},
	//// LD A, (DE)
	{code: 0x1A, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprIRegDE}, length: 1},
	// DEC DE
	{code: 0x1B, ticks: 2, mnc: DEC, oprs: Operands{OprRegDE}, length: 1},
	// INC E
	{code: 0x1C, ticks: 1, mnc: INC, oprs: Operands{OprRegE}, length: 1},
	// DEC E
	{code: 0x1D, ticks: 1, mnc: DEC, oprs: Operands{OprRegE}, length: 1},
	// LD E, $FF
	{code: 0x1E, ticks: 2, mnc: LD, oprs: Operands{OprRegE, OprU8}, length: 2},
	// RRA
	{code: 0x1F, ticks: 1, mnc: RRA, oprs: Operands{}, length: 1},

	// JR NZ, $FF
	{code: 0x20, ticks: 2, mnc: JR, oprs: Operands{OprFlagNZ, OprI8}, length: 2},
	// LD HL, $FFFF
	{code: 0x21, ticks: 3, mnc: LD, oprs: Operands{OprRegHL, OprU16}, length: 3},
	// LD (HLI), A
	{code: 0x22, ticks: 2, mnc: LD, oprs: Operands{OprIRegHLI}, length: 1},
	// INC HL
	{code: 0x23, ticks: 2, mnc: INC, oprs: Operands{OprRegHL}, length: 1},
	// INC H
	{code: 0x24, ticks: 1, mnc: INC, oprs: Operands{OprRegH}, length: 1},
	// DEC H
	{code: 0x25, ticks: 1, mnc: DEC, oprs: Operands{OprRegH}, length: 1},
	// LD H, $FF
	{code: 0x26, ticks: 2, mnc: LD, oprs: Operands{OprRegH, OprU8}, length: 2},
	// DAA
	{code: 0x27, ticks: 1, mnc: DAA, oprs: Operands{}, length: 1},
	// JR Z, $FF
	{code: 0x28, ticks: 2, mnc: JR, oprs: Operands{OprFlagZ, OprI8}, length: 2},
	// ADD HL, HL
	{code: 0x29, ticks: 2, mnc: ADD, oprs: Operands{OprRegHL, OprRegHL}, length: 1},
	// LD A, (HLI)
	{code: 0x2A, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprIRegHLI}, length: 1},
	// DEC HL
	{code: 0x2B, ticks: 2, mnc: DEC, oprs: Operands{OprRegHL}, length: 1},
	//// INC L
	{code: 0x2C, ticks: 1, mnc: INC, oprs: Operands{OprRegL}, length: 1},
	// DEC L
	{code: 0x2D, ticks: 1, mnc: DEC, oprs: Operands{OprRegL}, length: 1},
	// LD L, $FF
	{code: 0x2E, ticks: 2, mnc: LD, oprs: Operands{OprRegL, OprU8}, length: 2},
	// CPL
	{code: 0x2F, ticks: 1, mnc: CPL, oprs: Operands{}, length: 1},

	// JR NC, $FF
	{code: 0x30, ticks: 2, mnc: JR, oprs: Operands{OprFlagNC, OprI8}, length: 2},
	// LD SP, $FFFF
	{code: 0x31, ticks: 3, mnc: LD, oprs: Operands{OprRegSP, OprU16}, length: 3},
	// LD (HLD), A
	{code: 0x32, ticks: 2, mnc: LD, oprs: Operands{OprIRegHLD, OprRegA}, length: 1},
	// INC SP
	{code: 0x33, ticks: 2, mnc: INC, oprs: Operands{OprRegSP}, length: 1},
	// INC (HL)
	{code: 0x34, ticks: 3, mnc: INC, oprs: Operands{OprIRegHL}, length: 1},
	// DEC (HL)
	{code: 0x35, ticks: 3, mnc: DEC, oprs: Operands{OprIRegHL}, length: 1},
	// LD (HL), $FF
	{code: 0x36, ticks: 3, mnc: LD, oprs: Operands{OprIRegHL, OprU8}, length: 2},
	// SCF
	{code: 0x37, ticks: 1, mnc: SCF, oprs: Operands{}, length: 1},
	// JR C, $FF
	{code: 0x38, ticks: 2, mnc: JR, oprs: Operands{OprRegC, OprI8}, length: 2},
	// ADD HL, SP
	{code: 0x39, ticks: 2, mnc: ADD, oprs: Operands{OprRegHL, OprRegSP}, length: 1},
	// LD A, (HLD)
	{code: 0x3A, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprIRegHLD}, length: 1},
	// DEC SP
	{code: 0x3B, ticks: 2, mnc: DEC, oprs: Operands{OprRegSP}, length: 1},
	// INC A
	{code: 0x3C, ticks: 1, mnc: INC, oprs: Operands{OprRegA}, length: 1},
	// DEC A
	{code: 0x3D, ticks: 1, mnc: DEC, oprs: Operands{OprRegA}, length: 1},
	// LD A, $FF
	{code: 0x3E, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprU8}, length: 2},
	// CCF
	{code: 0x3F, ticks: 1, mnc: CCF, oprs: Operands{}, length: 1},

	// LD B, B
	{code: 0x40, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegB}, length: 1},
	// LD B, C
	{code: 0x41, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegC}, length: 1},
	// LD B, D
	{code: 0x42, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegD}, length: 1},
	// LD B, E
	{code: 0x43, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegE}, length: 1},
	// LD B, H
	{code: 0x44, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegH}, length: 1},
	// LD B, L
	{code: 0x45, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegL}, length: 1},
	// LD B, (HL)
	{code: 0x46, ticks: 2, mnc: LD, oprs: Operands{OprRegB, OprIRegHL}, length: 1},
	// LD B, A
	{code: 0x47, ticks: 1, mnc: LD, oprs: Operands{OprRegB, OprRegA}, length: 1},
	// LD C, B
	{code: 0x48, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegB}, length: 1},
	// LD C, C
	{code: 0x49, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegC}, length: 1},
	// LD C, D
	{code: 0x4A, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegD}, length: 1},
	// LD C, E
	{code: 0x4B, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegE}, length: 1},
	// LD C, H
	{code: 0x4C, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegH}, length: 1},
	// LD C, L
	{code: 0x4D, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegL}, length: 1},
	// LD C, (HL)
	{code: 0x4E, ticks: 2, mnc: LD, oprs: Operands{OprRegC, OprIRegHL}, length: 1},
	// LD C, A
	{code: 0x4F, ticks: 1, mnc: LD, oprs: Operands{OprRegC, OprRegA}, length: 1},

	// LD D, B
	{code: 0x50, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegB}, length: 1},
	// LD D, C
	{code: 0x51, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegC}, length: 1},
	// LD D, D
	{code: 0x52, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegD}, length: 1},
	// LD D, E
	{code: 0x53, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegE}, length: 1},
	// LD D, H
	{code: 0x54, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegH}, length: 1},
	// LD D, L
	{code: 0x55, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegL}, length: 1},
	// LD D, (HL)
	{code: 0x56, ticks: 2, mnc: LD, oprs: Operands{OprRegD, OprIRegHL}, length: 1},
	// LD D, A
	{code: 0x57, ticks: 1, mnc: LD, oprs: Operands{OprRegD, OprRegA}, length: 1},
	// LD E, B
	{code: 0x58, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegB}, length: 1},
	// LD E, C
	{code: 0x59, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegC}, length: 1},
	// LD E, D
	{code: 0x5A, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegD}, length: 1},
	// LD E, E
	{code: 0x5B, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegE}, length: 1},
	// LD E, H
	{code: 0x5C, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegH}, length: 1},
	// LD E, L
	{code: 0x5D, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegB}, length: 1},
	// LD E, (HL)
	{code: 0x5E, ticks: 2, mnc: LD, oprs: Operands{OprRegE, OprIRegHL}, length: 1},
	// LD E, A
	{code: 0x5F, ticks: 1, mnc: LD, oprs: Operands{OprRegE, OprRegA}, length: 1},

	// LD H, B
	{code: 0x60, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegB}, length: 1},
	// LD H, C
	{code: 0x61, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegC}, length: 1},
	// LD H, D
	{code: 0x62, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegD}, length: 1},
	// LD H, E
	{code: 0x63, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegE}, length: 1},
	// LD H, H
	{code: 0x64, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegH}, length: 1},
	// LD H, L
	{code: 0x65, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegL}, length: 1},
	// LD H, (HL)
	{code: 0x66, ticks: 2, mnc: LD, oprs: Operands{OprRegH, OprIRegHL}, length: 1},
	// LD H, A
	{code: 0x67, ticks: 1, mnc: LD, oprs: Operands{OprRegH, OprRegA}, length: 1},
	// LD L, B
	{code: 0x68, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegB}, length: 1},
	// LD L, C
	{code: 0x69, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegC}, length: 1},
	// LD L, D
	{code: 0x6A, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegD}, length: 1},
	// LD L, E
	{code: 0x6B, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegE}, length: 1},
	// LD L, H
	{code: 0x6C, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegH}, length: 1},
	// LD L, L
	{code: 0x6D, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegL}, length: 1},
	// LD L, (HL)
	{code: 0x6E, ticks: 2, mnc: LD, oprs: Operands{OprRegL, OprIRegHL}, length: 1},
	// LD L, A
	{code: 0x6F, ticks: 1, mnc: LD, oprs: Operands{OprRegL, OprRegA}, length: 1},

	// LD (HL), B
	{code: 0x70, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegB}, length: 1},
	// LD (HL), C
	{code: 0x71, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegC}, length: 1},
	// LD (HL), D
	{code: 0x72, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegD}, length: 1},
	// LD (HL), E
	{code: 0x73, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegE}, length: 1},
	// LD (HL), H
	{code: 0x74, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegH}, length: 1},
	// LD (HL), L
	{code: 0x75, ticks: 1, mnc: LD, oprs: Operands{OprIRegHL, OprRegL}, length: 1},
	// HALT
	{code: 0x76, ticks: 1, mnc: HALT, oprs: Operands{}, length: 1},
	// LD (HL), A
	{code: 0x77, ticks: 2, mnc: LD, oprs: Operands{OprIRegHL, OprRegA}, length: 1},
	//// LD A, B
	{code: 0x78, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// LD A, C
	{code: 0x79, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// LD A, D
	{code: 0x7A, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// LD A, E
	{code: 0x7B, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// LD A, H
	{code: 0x7C, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// LD A, L
	{code: 0x7D, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// LD A, (HL)
	{code: 0x7E, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// LD A, A
	{code: 0x7F, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprRegA}, length: 1},

	//// ADD A, B
	{code: 0x80, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// ADD A, C
	{code: 0x81, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// ADD A, D
	{code: 0x82, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// ADD A, E
	{code: 0x83, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// ADD A, H
	{code: 0x84, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// ADD A, L
	{code: 0x85, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// ADD A, (HL)
	{code: 0x86, ticks: 2, mnc: ADD, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// ADD A, A
	{code: 0x87, ticks: 1, mnc: ADD, oprs: Operands{OprRegA, OprRegA}, length: 1},
	// ADC A, B
	{code: 0x88, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// ADC A, C
	{code: 0x89, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// ADC A, D
	{code: 0x8A, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// ADC A, E
	{code: 0x8B, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// ADC A, H
	{code: 0x8C, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// ADC A, L
	{code: 0x8D, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// ADC A, (HL)
	{code: 0x8E, ticks: 2, mnc: ADC, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// ADC A, A
	{code: 0x8F, ticks: 1, mnc: ADC, oprs: Operands{OprRegA, OprRegA}, length: 1},

	// SUB A, B
	{code: 0x90, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// SUB A, C
	{code: 0x91, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// SUB A, D
	{code: 0x92, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// SUB A, E
	{code: 0x93, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// SUB A, H
	{code: 0x94, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// SUB A, L
	{code: 0x95, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// SUB A, (HL)
	{code: 0x96, ticks: 2, mnc: SUB, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// SUB A, A
	{code: 0x97, ticks: 1, mnc: SUB, oprs: Operands{OprRegA, OprRegA}, length: 1},
	// SBC A, B
	{code: 0x98, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// SBC A, C
	{code: 0x99, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// SBC A, D
	{code: 0x9A, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// SBC A, E
	{code: 0x9B, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// SBC A, H
	{code: 0x9C, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// SBC A, L
	{code: 0x9D, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// SBC A, (HL)
	{code: 0x9E, ticks: 2, mnc: SBC, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// SBC A, A
	{code: 0x9F, ticks: 1, mnc: SBC, oprs: Operands{OprRegA, OprRegA}, length: 1},

	// AND A, B
	{code: 0xA0, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// AND A, C
	{code: 0xA1, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// AND A, D
	{code: 0xA2, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// AND A, E
	{code: 0xA3, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// AND A, H
	{code: 0xA4, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// AND A, L
	{code: 0xA5, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// AND A, (HL)
	{code: 0xA6, ticks: 2, mnc: AND, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// AND A, A
	{code: 0xA7, ticks: 1, mnc: AND, oprs: Operands{OprRegA, OprRegA}, length: 1},
	// XOR A, B
	{code: 0xA8, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// XOR A, C
	{code: 0xA9, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// XOR A, D
	{code: 0xAA, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// XOR A, E
	{code: 0xAB, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// XOR A, H
	{code: 0xAC, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// XOR A, L
	{code: 0xAD, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// XOR A, (HL)
	{code: 0xAE, ticks: 2, mnc: XOR, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// XOR A, A
	{code: 0xAF, ticks: 1, mnc: XOR, oprs: Operands{OprRegA, OprRegA}, length: 1},

	// OR A, B
	{code: 0xB0, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// OR A, C
	{code: 0xB1, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// OR A, D
	{code: 0xB2, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// OR A, E
	{code: 0xB3, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// OR A, H
	{code: 0xB4, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// OR A, L
	{code: 0xB5, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// OR A, (HL)
	{code: 0xB6, ticks: 2, mnc: OR, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// OR A, A
	{code: 0xB7, ticks: 1, mnc: OR, oprs: Operands{OprRegA, OprRegA}, length: 1},
	// CP A, B
	{code: 0xB8, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegB}, length: 1},
	// CP A, C
	{code: 0xB9, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegC}, length: 1},
	// CP A, D
	{code: 0xBA, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegD}, length: 1},
	// CP A, E
	{code: 0xBB, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegE}, length: 1},
	// CP A, H
	{code: 0xBC, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegH}, length: 1},
	// CP A, L
	{code: 0xBD, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegL}, length: 1},
	// CP A, (HL)
	{code: 0xBE, ticks: 2, mnc: CP, oprs: Operands{OprRegA, OprIRegHL}, length: 1},
	// CP A, A
	{code: 0xBF, ticks: 1, mnc: CP, oprs: Operands{OprRegA, OprRegA}, length: 1},

	// RET NZ
	{code: 0xC0, ticks: 2, mnc: RET, oprs: Operands{OprFlagNZ}, length: 1},
	// POP BC
	{code: 0xC1, ticks: 3, mnc: POP, oprs: Operands{OprRegBC}, length: 1},
	// JP NZ, $FFFF
	{code: 0xC2, ticks: 3, mnc: JP, oprs: Operands{OprFlagNZ, OprU16}, length: 3},
	// JP $FFFF
	{code: 0xC3, ticks: 3, mnc: JP, oprs: Operands{OprU16}, length: 3},
	// CALL NZ, $FFFF
	{code: 0xC4, ticks: 3, mnc: CALL, oprs: Operands{OprFlagNZ, OprU16}, length: 3},
	// PUSH BC
	{code: 0xC5, ticks: 4, mnc: PUSH, oprs: Operands{OprRegBC}, length: 1},
	// ADD A, $FF
	{code: 0xC6, ticks: 2, mnc: ADD, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $00
	{code: 0xC7, ticks: 4, mnc: RST, oprs: Operands{OprVec00}, length: 1},
	// RET Z
	{code: 0xC8, ticks: 2, mnc: RET, oprs: Operands{OprFlagZ}, length: 1},
	// RET
	{code: 0xC9, ticks: 4, mnc: RET, oprs: Operands{}, length: 1},
	// JP Z, $FFFF
	{code: 0xCA, ticks: 3, mnc: JP, oprs: Operands{OprFlagZ, OprU16}, length: 3},
	// PREFIX CB
	{code: 0xCB, ticks: 1, mnc: PrefixCB, oprs: Operands{}, length: 1},
	// CALL Z, $FFFF
	{code: 0xCC, ticks: 3, mnc: CALL, oprs: Operands{OprFlagZ, OprU16}, length: 3},
	// CALL $FFFF
	{code: 0xCD, ticks: 6, mnc: CALL, oprs: Operands{OprU16}, length: 3},
	// ADC A, $FF
	{code: 0xCE, ticks: 2, mnc: ADC, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $08
	{code: 0xCF, ticks: 4, mnc: RST, oprs: Operands{OprVec08}, length: 1},

	// RET NC
	{code: 0xD0, ticks: 2, mnc: RET, oprs: Operands{OprFlagNC}, length: 1},
	// POP DE
	{code: 0xD1, ticks: 3, mnc: POP, oprs: Operands{OprRegDE}, length: 1},
	// JP NC, $FFFF
	{code: 0xD2, ticks: 3, mnc: JP, oprs: Operands{OprFlagNC, OprU16}, length: 3},
	// ILLEGAL OP
	{code: 0xD3, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// CALL NC, $FFFF
	{code: 0xD4, ticks: 3, mnc: CALL, oprs: Operands{OprFlagNC, OprU16}, length: 3},
	// PUSH DE
	{code: 0xD5, ticks: 4, mnc: PUSH, oprs: Operands{OprRegDE}, length: 1},
	// SUB A, $FF
	{code: 0xD6, ticks: 2, mnc: SUB, oprs: Operands{OprRegA, OprU8}, length: 1},
	// RST $10
	{code: 0xD7, ticks: 4, mnc: RST, oprs: Operands{OprVec10}, length: 1},
	// RET C
	{code: 0xD8, ticks: 2, mnc: RET, oprs: Operands{OprRegC}, length: 1},
	// RETI
	{code: 0xD9, ticks: 4, mnc: RETI, oprs: Operands{}, length: 1},
	// JP C, $FFFF
	{code: 0xDA, ticks: 3, mnc: JP, oprs: Operands{OprFlagC, OprU16}, length: 3},
	// ILLEGAL OP
	{code: 0xDB, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// CALL C, $FFFF
	{code: 0xDC, ticks: 3, mnc: CALL, oprs: Operands{OprRegC, OprU16}, length: 3},
	// ILLEGAL OP
	{code: 0xDD, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// SBC A, $FF
	{code: 0xDE, ticks: 2, mnc: SBC, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $18
	{code: 0xDF, ticks: 1, mnc: RST, oprs: Operands{OprVec18}, length: 1},

	// LD (FF00 + $FF), A
	{code: 0xE0, ticks: 2, mnc: LD, oprs: Operands{OprFFU8, OprRegA}, length: 2},
	// POP HL
	{code: 0xE1, ticks: 3, mnc: POP, oprs: Operands{OprRegHL}, length: 1},
	// LD (FF00 + C), A
	{code: 0xE2, ticks: 1, mnc: LD, oprs: Operands{OprIRegC, OprRegA}, length: 1},
	// ILLEGAL OP
	{code: 0xE3, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xE4, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// PUSH HL
	{code: 0xE5, ticks: 4, mnc: PUSH, oprs: Operands{OprRegHL}, length: 1},
	// AND A, $FF
	{code: 0xE6, ticks: 2, mnc: AND, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $20
	{code: 0xE7, ticks: 4, mnc: RST, oprs: Operands{OprVec20}, length: 1},
	// ADD SP, $FF
	{code: 0xE8, ticks: 0, mnc: ADD, oprs: Operands{OprRegSP, OprU8}, length: 2},
	// JP HL
	{code: 0xE9, ticks: 4, mnc: JP, oprs: Operands{OprRegHL}, length: 1},
	// LD ($FFFF), A
	{code: 0xEA, ticks: 3, mnc: LD, oprs: Operands{OprInd, OprRegA}, length: 3},
	// ILLEGAL OP
	{code: 0xEB, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xEC, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xED, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// XOR A, $FF
	{code: 0xEE, ticks: 2, mnc: XOR, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $28
	{code: 0xEF, ticks: 4, mnc: RST, oprs: Operands{OprVec28}, length: 1},

	// LD A, (FF00 + $FF)
	{code: 0xF0, ticks: 2, mnc: LD, oprs: Operands{OprRegA, OprFFU8}, length: 2},
	// POP AF
	{code: 0xF1, ticks: 3, mnc: POP, oprs: Operands{OprRegAF}, length: 1},
	// LD A, (FF00 + C)
	{code: 0xF2, ticks: 1, mnc: LD, oprs: Operands{OprRegA, OprIRegC}, length: 1},
	// DI
	{code: 0xF3, ticks: 1, mnc: DI, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xF4, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// PUSH AF
	{code: 0xF5, ticks: 4, mnc: PUSH, oprs: Operands{OprRegAF}, length: 1},
	// OR A, $FF
	{code: 0xF6, ticks: 0, mnc: OR, oprs: Operands{OprRegAF, OprU8}, length: 2},
	// RST $30
	{code: 0xF7, ticks: 0, mnc: RST, oprs: Operands{OprVec30}, length: 1},
	// LD HL, SP + $FF
	{code: 0xF8, ticks: 3, mnc: LD, oprs: Operands{OprRegHL, OprSPI8}, length: 2},
	// LD SP, HL
	{code: 0xF9, ticks: 4, mnc: LD, oprs: Operands{OprRegSP, OprRegHL}, length: 1},
	// LD A, ($FFFF)
	{code: 0xFA, ticks: 3, mnc: LD, oprs: Operands{OprRegA, OprInd}, length: 3},
	// EI
	{code: 0xFB, ticks: 1, mnc: EI, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xFC, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// ILLEGAL OP
	{code: 0xFD, ticks: 0, mnc: IllegalOp, oprs: Operands{}, length: 1},
	// CP A, $FF
	{code: 0xFE, ticks: 2, mnc: CP, oprs: Operands{OprRegA, OprU8}, length: 2},
	// RST $38
	{code: 0xFF, ticks: 4, mnc: RST, oprs: Operands{OprVec38}, length: 1},
}

var cpOpCodes = [OPCodeSize]OpCode{
	// RLC B
	{code: 0x00, ticks: 2, mnc: RLC, oprs: Operands{OprRegB}, length: 2},
	// RLC C
	{code: 0x01, ticks: 2, mnc: RLC, oprs: Operands{OprRegC}, length: 2},
	// RLC D
	{code: 0x02, ticks: 2, mnc: RLC, oprs: Operands{OprRegD}, length: 2},
	// RLC E
	{code: 0x03, ticks: 2, mnc: RLC, oprs: Operands{OprRegE}, length: 2},
	// RLC H
	{code: 0x04, ticks: 2, mnc: RLC, oprs: Operands{OprRegH}, length: 2},
	// RLC L
	{code: 0x05, ticks: 2, mnc: RLC, oprs: Operands{OprRegL}, length: 2},
	// RLC (HL)
	{code: 0x06, ticks: 4, mnc: RLC, oprs: Operands{OprIRegHL}, length: 2},
	// RLC A
	{code: 0x07, ticks: 2, mnc: RLC, oprs: Operands{OprRegA}, length: 2},
	// RRC B
	{code: 0x08, ticks: 2, mnc: RRC, oprs: Operands{OprRegB}, length: 2},
	// RRC C
	{code: 0x09, ticks: 2, mnc: RRC, oprs: Operands{OprRegC}, length: 2},
	// RRC D
	{code: 0x0A, ticks: 2, mnc: RRC, oprs: Operands{OprRegD}, length: 2},
	// RRC E
	{code: 0x0B, ticks: 2, mnc: RRC, oprs: Operands{OprRegE}, length: 2},
	// RRC H
	{code: 0x0C, ticks: 2, mnc: RRC, oprs: Operands{OprRegH}, length: 2},
	// RRC L
	{code: 0x0D, ticks: 2, mnc: RRC, oprs: Operands{OprRegL}, length: 2},
	// RRC (HL)
	{code: 0x0E, ticks: 4, mnc: RRC, oprs: Operands{OprIRegHL}, length: 2},
	// RRC A
	{code: 0x0F, ticks: 2, mnc: RRC, oprs: Operands{OprRegA}, length: 2},

	// RL B
	{code: 0x10, ticks: 2, mnc: RL, oprs: Operands{OprRegB}, length: 2},
	// RL C
	{code: 0x11, ticks: 2, mnc: RL, oprs: Operands{OprRegC}, length: 2},
	// RL D
	{code: 0x12, ticks: 2, mnc: RL, oprs: Operands{OprRegD}, length: 2},
	// RL E
	{code: 0x13, ticks: 2, mnc: RL, oprs: Operands{OprRegE}, length: 2},
	// RL H
	{code: 0x14, ticks: 2, mnc: RL, oprs: Operands{OprRegH}, length: 2},
	// RL L
	{code: 0x15, ticks: 2, mnc: RL, oprs: Operands{OprRegL}, length: 2},
	// RL (HL)
	{code: 0x16, ticks: 4, mnc: RL, oprs: Operands{OprIRegHL}, length: 2},
	// RL A
	{code: 0x17, ticks: 2, mnc: RL, oprs: Operands{OprRegA}, length: 2},
	// RR B
	{code: 0x18, ticks: 2, mnc: RR, oprs: Operands{OprRegB}, length: 2},
	// RR C
	{code: 0x19, ticks: 2, mnc: RR, oprs: Operands{OprRegC}, length: 2},
	// RR D
	{code: 0x1A, ticks: 2, mnc: RR, oprs: Operands{OprRegD}, length: 2},
	// RR E
	{code: 0x1B, ticks: 2, mnc: RR, oprs: Operands{OprRegE}, length: 2},
	// RR H
	{code: 0x1C, ticks: 2, mnc: RR, oprs: Operands{OprRegH}, length: 2},
	// RR L
	{code: 0x1D, ticks: 2, mnc: RR, oprs: Operands{OprRegL}, length: 2},
	// RR (HL)
	{code: 0x1E, ticks: 4, mnc: RR, oprs: Operands{OprIRegHL}, length: 2},
	// RR A
	{code: 0x1F, ticks: 2, mnc: RR, oprs: Operands{OprRegA}, length: 2},

	// SLA B
	{code: 0x20, ticks: 2, mnc: SLA, oprs: Operands{OprRegB}, length: 2},
	// SLA C
	{code: 0x21, ticks: 2, mnc: SLA, oprs: Operands{OprRegC}, length: 2},
	// SLA D
	{code: 0x22, ticks: 2, mnc: SLA, oprs: Operands{OprRegD}, length: 2},
	// SLA E
	{code: 0x23, ticks: 2, mnc: SLA, oprs: Operands{OprRegE}, length: 2},
	// SLA H
	{code: 0x24, ticks: 2, mnc: SLA, oprs: Operands{OprRegH}, length: 2},
	// SLA L
	{code: 0x25, ticks: 2, mnc: SLA, oprs: Operands{OprRegL}, length: 2},
	// SLA (HL)
	{code: 0x26, ticks: 4, mnc: SLA, oprs: Operands{OprIRegHL}, length: 2},
	// SLA A
	{code: 0x27, ticks: 2, mnc: SLA, oprs: Operands{OprRegA}, length: 2},
	// SRA B
	{code: 0x28, ticks: 2, mnc: SRA, oprs: Operands{OprRegB}, length: 2},
	// SRA C
	{code: 0x29, ticks: 2, mnc: SRA, oprs: Operands{OprRegC}, length: 2},
	// SRA D
	{code: 0x2A, ticks: 2, mnc: SRA, oprs: Operands{OprRegD}, length: 2},
	// SRA E
	{code: 0x2B, ticks: 2, mnc: SRA, oprs: Operands{OprRegE}, length: 2},
	// SRA H
	{code: 0x2C, ticks: 2, mnc: SRA, oprs: Operands{OprRegH}, length: 2},
	// SRA L
	{code: 0x2D, ticks: 2, mnc: SRA, oprs: Operands{OprRegL}, length: 2},
	// SRA (HL)
	{code: 0x2E, ticks: 4, mnc: SRA, oprs: Operands{OprIRegHL}, length: 2},
	// SRA A
	{code: 0x2F, ticks: 2, mnc: SRA, oprs: Operands{OprRegA}, length: 2},

	// SWAP B
	{code: 0x30, ticks: 2, mnc: SWAP, oprs: Operands{OprRegB}, length: 2},
	// SWAP C
	{code: 0x31, ticks: 2, mnc: SWAP, oprs: Operands{OprRegC}, length: 2},
	// SWAP D
	{code: 0x32, ticks: 2, mnc: SWAP, oprs: Operands{OprRegD}, length: 2},
	// SWAP E
	{code: 0x33, ticks: 2, mnc: SWAP, oprs: Operands{OprRegE}, length: 2},
	// SWAP H
	{code: 0x34, ticks: 2, mnc: SWAP, oprs: Operands{OprRegH}, length: 2},
	// SWAP L
	{code: 0x35, ticks: 2, mnc: SWAP, oprs: Operands{OprRegL}, length: 2},
	// SWAP (HL)
	{code: 0x36, ticks: 4, mnc: SWAP, oprs: Operands{OprIRegHL}, length: 2},
	// SWAP A
	{code: 0x37, ticks: 2, mnc: SWAP, oprs: Operands{OprRegA}, length: 2},
	// SRL B
	{code: 0x38, ticks: 2, mnc: SRL, oprs: Operands{OprRegB}, length: 2},
	// SRL C
	{code: 0x39, ticks: 2, mnc: SRL, oprs: Operands{OprRegC}, length: 2},
	// SRL D
	{code: 0x3A, ticks: 2, mnc: SRL, oprs: Operands{OprRegD}, length: 2},
	// SRL E
	{code: 0x3B, ticks: 2, mnc: SRL, oprs: Operands{OprRegE}, length: 2},
	// SRL H
	{code: 0x3C, ticks: 2, mnc: SRL, oprs: Operands{OprRegH}, length: 2},
	// SRL L
	{code: 0x3D, ticks: 2, mnc: SRL, oprs: Operands{OprRegL}, length: 2},
	// SRL (HL)
	{code: 0x3E, ticks: 4, mnc: SRL, oprs: Operands{OprIRegHL}, length: 2},
	// SRL A
	{code: 0x3F, ticks: 2, mnc: SRL, oprs: Operands{OprRegA}, length: 2},

	// BIT 0, B
	{code: 0x40, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegB}, length: 2},
	// BIT 0, C
	{code: 0x41, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegC}, length: 2},
	// BIT 0, D
	{code: 0x42, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegD}, length: 2},
	// BIT 0, E
	{code: 0x43, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegE}, length: 2},
	// BIT 0, H
	{code: 0x44, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegH}, length: 2},
	// BIT 0, L
	{code: 0x45, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegL}, length: 2},
	// BIT 0, (HL)
	{code: 0x46, ticks: 3, mnc: BIT, oprs: Operands{OprBit0, OprIRegHL}, length: 2},
	// BIT 0, A
	{code: 0x47, ticks: 2, mnc: BIT, oprs: Operands{OprBit0, OprRegA}, length: 2},
	// BIT 1, B
	{code: 0x48, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegB}, length: 2},
	// BIT 1, C
	{code: 0x49, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegC}, length: 2},
	// BIT 1, D
	{code: 0x4A, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegD}, length: 2},
	// BIT 1, E
	{code: 0x4B, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegE}, length: 2},
	// BIT 1, H
	{code: 0x4C, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegH}, length: 2},
	// BIT 1, L
	{code: 0x4D, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegL}, length: 2},
	// BIT 1, (HL)
	{code: 0x4E, ticks: 3, mnc: BIT, oprs: Operands{OprBit1, OprIRegHL}, length: 2},
	// BIT 1, A
	{code: 0x4F, ticks: 2, mnc: BIT, oprs: Operands{OprBit1, OprRegA}, length: 2},

	// BIT 2, B
	{code: 0x50, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegB}, length: 2},
	// BIT 2, C
	{code: 0x51, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegC}, length: 2},
	// BIT 2, D
	{code: 0x52, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegD}, length: 2},
	// BIT 2, E
	{code: 0x53, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegE}, length: 2},
	// BIT 2, H
	{code: 0x54, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegH}, length: 2},
	// BIT 2, L
	{code: 0x55, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegL}, length: 2},
	// BIT 2, (HL)
	{code: 0x56, ticks: 3, mnc: BIT, oprs: Operands{OprBit2, OprIRegHL}, length: 2},
	// BIT 2, A
	{code: 0x57, ticks: 2, mnc: BIT, oprs: Operands{OprBit2, OprRegA}, length: 2},
	// BIT 3, B
	{code: 0x58, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegB}, length: 2},
	// BIT 3, C
	{code: 0x59, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegC}, length: 2},
	// BIT 3, D
	{code: 0x5A, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegD}, length: 2},
	// BIT 3, E
	{code: 0x5B, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegE}, length: 2},
	// BIT 3, H
	{code: 0x5C, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegH}, length: 2},
	// BIT 3, L
	{code: 0x5D, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegL}, length: 2},
	// BIT 3, (HL)
	{code: 0x5E, ticks: 3, mnc: BIT, oprs: Operands{OprBit3, OprIRegHL}, length: 2},
	// BIT 3, A
	{code: 0x5F, ticks: 2, mnc: BIT, oprs: Operands{OprBit3, OprRegA}, length: 2},

	// BIT 4, B
	{code: 0x60, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegB}, length: 2},
	// BIT 4, C
	{code: 0x61, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegC}, length: 2},
	// BIT 4, D
	{code: 0x62, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegD}, length: 2},
	// BIT 4, E
	{code: 0x63, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegE}, length: 2},
	// BIT 4, H
	{code: 0x64, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegH}, length: 2},
	// BIT 4, L
	{code: 0x65, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegL}, length: 2},
	// BIT 4, (HL)
	{code: 0x66, ticks: 3, mnc: BIT, oprs: Operands{OprBit4, OprIRegHL}, length: 2},
	// BIT 4, A
	{code: 0x67, ticks: 2, mnc: BIT, oprs: Operands{OprBit4, OprRegA}, length: 2},
	// BIT 5, B
	{code: 0x68, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegB}, length: 2},
	// BIT 5, C
	{code: 0x69, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegC}, length: 2},
	// BIT 5, D
	{code: 0x6A, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegD}, length: 2},
	// BIT 5, E
	{code: 0x6B, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegE}, length: 2},
	// BIT 5, H
	{code: 0x6C, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegH}, length: 2},
	// BIT 5, L
	{code: 0x6D, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegL}, length: 2},
	// BIT 5, (HL)
	{code: 0x6E, ticks: 3, mnc: BIT, oprs: Operands{OprBit5, OprIRegHL}, length: 2},
	// BIT 5, A
	{code: 0x6F, ticks: 2, mnc: BIT, oprs: Operands{OprBit5, OprRegA}, length: 2},

	// BIT 6, B
	{code: 0x70, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegB}, length: 2},
	// BIT 6, C
	{code: 0x71, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegC}, length: 2},
	// BIT 6, D
	{code: 0x72, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegD}, length: 2},
	// BIT 6, E
	{code: 0x73, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegE}, length: 2},
	// BIT 6, H
	{code: 0x74, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegH}, length: 2},
	// BIT 6, L
	{code: 0x75, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegL}, length: 2},
	// BIT 6, (HL)
	{code: 0x76, ticks: 3, mnc: BIT, oprs: Operands{OprBit6, OprIRegHL}, length: 2},
	// BIT 6, A
	{code: 0x77, ticks: 2, mnc: BIT, oprs: Operands{OprBit6, OprRegA}, length: 2},
	// BIT 7, B
	{code: 0x78, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegB}, length: 2},
	// BIT 7, C
	{code: 0x79, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegC}, length: 2},
	// BIT 7, D
	{code: 0x7A, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegD}, length: 2},
	// BIT 7, E
	{code: 0x7B, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegE}, length: 2},
	// BIT 7, H
	{code: 0x7C, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegH}, length: 2},
	// BIT 7, L
	{code: 0x7D, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegL}, length: 2},
	// BIT 7, (HL)
	{code: 0x7E, ticks: 3, mnc: BIT, oprs: Operands{OprBit7, OprIRegHL}, length: 2},
	// BIT 7, A
	{code: 0x7F, ticks: 2, mnc: BIT, oprs: Operands{OprBit7, OprRegA}, length: 2},

	// RES 0, B
	{code: 0x80, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegB}, length: 2},
	// RES 0, C
	{code: 0x81, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegC}, length: 2},
	// RES 0, D
	{code: 0x82, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegD}, length: 2},
	// RES 0, E
	{code: 0x83, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegE}, length: 2},
	// RES 0, H
	{code: 0x84, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegH}, length: 2},
	// RES 0, L
	{code: 0x85, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegL}, length: 2},
	// RES 0, (HL)
	{code: 0x86, ticks: 4, mnc: RES, oprs: Operands{OprBit0, OprIRegHL}, length: 2},
	// RES 0, A
	{code: 0x87, ticks: 2, mnc: RES, oprs: Operands{OprBit0, OprRegA}, length: 2},
	// RES 1, B
	{code: 0x88, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegB}, length: 2},
	// RES 1, C
	{code: 0x89, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegC}, length: 2},
	// RES 1, D
	{code: 0x8A, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegD}, length: 2},
	// RES 1, E
	{code: 0x8B, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegE}, length: 2},
	// RES 1, H
	{code: 0x8C, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegH}, length: 2},
	// RES 1, L
	{code: 0x8D, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegL}, length: 2},
	// RES 1, (HL)
	{code: 0x8E, ticks: 4, mnc: RES, oprs: Operands{OprBit1, OprIRegHL}, length: 2},
	// RES 1, A
	{code: 0x8F, ticks: 2, mnc: RES, oprs: Operands{OprBit1, OprRegA}, length: 2},

	// RES 2, B
	{code: 0x90, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegB}, length: 2},
	// RES 2, C
	{code: 0x91, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegC}, length: 2},
	// RES 2, D
	{code: 0x92, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegD}, length: 2},
	// RES 2, E
	{code: 0x93, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegE}, length: 2},
	// RES 2, H
	{code: 0x94, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegH}, length: 2},
	// RES 2, L
	{code: 0x95, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegL}, length: 2},
	// RES 2, (HL)
	{code: 0x96, ticks: 4, mnc: RES, oprs: Operands{OprBit2, OprIRegHL}, length: 2},
	// RES 2, A
	{code: 0x97, ticks: 2, mnc: RES, oprs: Operands{OprBit2, OprRegA}, length: 2},
	// RES 3, B
	{code: 0x98, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegB}, length: 2},
	// RES 3, C
	{code: 0x99, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegC}, length: 2},
	// RES 3, D
	{code: 0x9A, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegD}, length: 2},
	// RES 3, E
	{code: 0x9B, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegE}, length: 2},
	// RES 3, H
	{code: 0x9C, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegH}, length: 2},
	// RES 3, L
	{code: 0x9D, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegL}, length: 2},
	// RES 3, (HL)
	{code: 0x9E, ticks: 4, mnc: RES, oprs: Operands{OprBit3, OprIRegHL}, length: 2},
	// RES 3, A
	{code: 0x9F, ticks: 2, mnc: RES, oprs: Operands{OprBit3, OprRegA}, length: 2},

	// RES 4, B
	{code: 0xA0, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegB}, length: 2},
	// RES 4, C
	{code: 0xA1, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegC}, length: 2},
	// RES 4, D
	{code: 0xA2, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegD}, length: 2},
	// RES 4, E
	{code: 0xA3, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegE}, length: 2},
	// RES 4, H
	{code: 0xA4, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegH}, length: 2},
	// RES 4, L
	{code: 0xA5, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegL}, length: 2},
	// RES 4, (HL)
	{code: 0xA6, ticks: 4, mnc: RES, oprs: Operands{OprBit4, OprIRegHL}, length: 2},
	// RES 4, A
	{code: 0xA7, ticks: 2, mnc: RES, oprs: Operands{OprBit4, OprRegA}, length: 2},
	// RES 5, B
	{code: 0xA8, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegB}, length: 2},
	// RES 5, C
	{code: 0xA9, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegC}, length: 2},
	// RES 5, D
	{code: 0xAA, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegD}, length: 2},
	// RES 5, E
	{code: 0xAB, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegE}, length: 2},
	// RES 5, H
	{code: 0xAC, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegH}, length: 2},
	// RES 5, L
	{code: 0xAD, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegL}, length: 2},
	// RES 5, (HL)
	{code: 0xAE, ticks: 4, mnc: RES, oprs: Operands{OprBit5, OprIRegHL}, length: 2},
	// RES 5, A
	{code: 0xAF, ticks: 2, mnc: RES, oprs: Operands{OprBit5, OprRegA}, length: 2},

	// RES 6, B
	{code: 0xB0, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegB}, length: 2},
	// RES 6, C
	{code: 0xB1, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegC}, length: 2},
	// RES 6, D
	{code: 0xB2, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegD}, length: 2},
	// RES 6, E
	{code: 0xB3, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegE}, length: 2},
	// RES 6, H
	{code: 0xB4, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegH}, length: 2},
	// RES 6, L
	{code: 0xB5, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegL}, length: 2},
	// RES 6, (HL)
	{code: 0xB6, ticks: 4, mnc: RES, oprs: Operands{OprBit6, OprIRegHL}, length: 2},
	// RES 6, A
	{code: 0xB7, ticks: 2, mnc: RES, oprs: Operands{OprBit6, OprRegA}, length: 2},
	// RES 7, B
	{code: 0xB8, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegB}, length: 2},
	// RES 7, C
	{code: 0xB9, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegC}, length: 2},
	// RES 7, D
	{code: 0xBA, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegD}, length: 2},
	// RES 7, E
	{code: 0xBB, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegE}, length: 2},
	// RES 7, H
	{code: 0xBC, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegH}, length: 2},
	// RES 7, L
	{code: 0xBD, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegL}, length: 2},
	// RES 7, (HL)
	{code: 0xBE, ticks: 4, mnc: RES, oprs: Operands{OprBit7, OprIRegHL}, length: 2},
	// RES 7, A
	{code: 0xBF, ticks: 2, mnc: RES, oprs: Operands{OprBit7, OprRegA}, length: 2},

	// SET 0, B
	{code: 0xC0, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegB}, length: 2},
	// SET 0, C
	{code: 0xC1, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegC}, length: 2},
	// SET 0, D
	{code: 0xC2, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegD}, length: 2},
	// SET 0, E
	{code: 0xC3, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegE}, length: 2},
	// SET 0, H
	{code: 0xC4, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegH}, length: 2},
	// SET 0, L
	{code: 0xC5, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegL}, length: 2},
	// SET 0, (HL)
	{code: 0xC6, ticks: 4, mnc: SET, oprs: Operands{OprBit0, OprIRegHL}, length: 2},
	// SET 0, A
	{code: 0xC7, ticks: 2, mnc: SET, oprs: Operands{OprBit0, OprRegA}, length: 2},
	// SET 1, B
	{code: 0xC8, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegB}, length: 2},
	// SET 1, C
	{code: 0xC9, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegC}, length: 2},
	// SET 1, D
	{code: 0xCA, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegD}, length: 2},
	// SET 1, E
	{code: 0xCB, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegE}, length: 2},
	// SET 1, H
	{code: 0xCC, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegH}, length: 2},
	// SET 1, L
	{code: 0xCD, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegL}, length: 2},
	// SET 1, (HL)
	{code: 0xCE, ticks: 4, mnc: SET, oprs: Operands{OprBit1, OprIRegHL}, length: 2},
	// SET 1, A
	{code: 0xCF, ticks: 2, mnc: SET, oprs: Operands{OprBit1, OprRegA}, length: 2},

	// SET 2, B
	{code: 0xD0, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegB}, length: 2},
	// SET 2, C
	{code: 0xD1, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegC}, length: 2},
	// SET 2, D
	{code: 0xD2, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegD}, length: 2},
	// SET 2, E
	{code: 0xD3, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegE}, length: 2},
	// SET 2, H
	{code: 0xD4, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegH}, length: 2},
	// SET 2, L
	{code: 0xD5, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegL}, length: 2},
	// SET 2, (HL)
	{code: 0xD6, ticks: 4, mnc: SET, oprs: Operands{OprBit2, OprIRegHL}, length: 2},
	// SET 2, A
	{code: 0xD7, ticks: 2, mnc: SET, oprs: Operands{OprBit2, OprRegA}, length: 2},
	// SET 3, B
	{code: 0xD8, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegB}, length: 2},
	// SET 3, C
	{code: 0xD9, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegC}, length: 2},
	// SET 3, D
	{code: 0xDA, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegD}, length: 2},
	// SET 3, E
	{code: 0xDB, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegE}, length: 2},
	// SET 3, H
	{code: 0xDC, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegH}, length: 2},
	// SET 3, L
	{code: 0xDD, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegL}, length: 2},
	// SET 3, (HL)
	{code: 0xDE, ticks: 4, mnc: SET, oprs: Operands{OprBit3, OprIRegHL}, length: 2},
	// SET 3, A
	{code: 0xDF, ticks: 2, mnc: SET, oprs: Operands{OprBit3, OprRegA}, length: 2},

	// SET 4, B
	{code: 0xE0, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegB}, length: 2},
	// SET 4, C
	{code: 0xE1, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegC}, length: 2},
	// SET 4, D
	{code: 0xE2, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegD}, length: 2},
	// SET 4, E
	{code: 0xE3, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegE}, length: 2},
	// SET 4, H
	{code: 0xE4, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegH}, length: 2},
	// SET 4, L
	{code: 0xE5, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegL}, length: 2},
	// SET 4, (HL)
	{code: 0xE6, ticks: 4, mnc: SET, oprs: Operands{OprBit4, OprIRegHL}, length: 2},
	// SET 4, A
	{code: 0xE7, ticks: 2, mnc: SET, oprs: Operands{OprBit4, OprRegA}, length: 2},
	// SET 5, B
	{code: 0xE8, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegB}, length: 2},
	// SET 5, C
	{code: 0xE9, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegC}, length: 2},
	// SET 5, D
	{code: 0xEA, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegD}, length: 2},
	// SET 5, E
	{code: 0xEB, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegE}, length: 2},
	// SET 5, H
	{code: 0xEC, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegH}, length: 2},
	// SET 5, L
	{code: 0xED, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegL}, length: 2},
	// SET 5, (HL)
	{code: 0xEE, ticks: 4, mnc: SET, oprs: Operands{OprBit5, OprIRegHL}, length: 2},
	// SET 5, A
	{code: 0xEF, ticks: 2, mnc: SET, oprs: Operands{OprBit5, OprRegA}, length: 2},

	// SET 6, B
	{code: 0xF0, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegB}, length: 2},
	// SET 6, C
	{code: 0xF1, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegC}, length: 2},
	// SET 6, D
	{code: 0xF2, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegD}, length: 2},
	// SET 6, E
	{code: 0xF3, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegE}, length: 2},
	// SET 6, H
	{code: 0xF4, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegH}, length: 2},
	// SET 6, L
	{code: 0xF5, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegL}, length: 2},
	// SET 6, (HL)
	{code: 0xF6, ticks: 4, mnc: SET, oprs: Operands{OprBit6, OprIRegHL}, length: 2},
	// SET 6, A
	{code: 0xF7, ticks: 2, mnc: SET, oprs: Operands{OprBit6, OprRegA}, length: 2},
	// SET 7, B
	{code: 0xF8, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegB}, length: 2},
	// SET 7, C
	{code: 0xF9, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegC}, length: 2},
	// SET 7, D
	{code: 0xFA, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegD}, length: 2},
	// SET 7, E
	{code: 0xFB, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegE}, length: 2},
	// SET 7, H
	{code: 0xFC, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegH}, length: 2},
	// SET 7, L
	{code: 0xFD, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegL}, length: 2},
	// SET 7, (HL)
	{code: 0xFE, ticks: 4, mnc: SET, oprs: Operands{OprBit7, OprIRegHL}, length: 2},
	// SET 7, A
	{code: 0xFF, ticks: 2, mnc: SET, oprs: Operands{OprBit7, OprRegA}, length: 2},
}
