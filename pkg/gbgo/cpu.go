package gbgo

import (
	"fmt"
)

const (
	INST_SIZE = 0xFF
)

type Register struct {
	// TODO Add comments
	A, F, B, C, D, E, H, L uint8
	SP, PC                 uint16
}

// Functions related to Flag REgister (F)

// Returns F Flag value, with first four bits always set to Zero
func (r *Register) GetF() uint8 {
	return r.F & 0b11110000
}

// Sets F Flag value, with first four bits aways set to Zero
func (r *Register) SetF(value uint8) {
	r.F = value & 0b11110000
}

func (r *Register) GetFlagZ() bool {
	return (r.F & 0b10000000) == 0b10000000
}

func (r *Register) SetFlagZ(value bool) {
	if value {
		r.F |= 0b10000000
	} else {
		r.F &= 0b01110000
	}
}

func (r *Register) GetFlagN() bool {
	return (r.F & 0b01000000) == 0b01000000
}

func (r *Register) SetFlagN(value bool) {
	if value {
		r.F |= 0b01000000
	} else {
		r.F &= 0b10110000
	}
}

func (r *Register) GetFlagH() bool {
	return (r.F & 0b00100000) == 0b00100000
}

func (r *Register) SetFlagH(value bool) {
	if value {
		r.F |= 0b00100000
	} else {
		r.F &= 0b11010000
	}
}

func (r *Register) GetFlagC() bool {
	return (r.F & 0b00010000) == 0b00010000
}

func (r *Register) SetFlagC(value bool) {
	if value {
		r.F |= 0b00010000
	} else {
		r.F &= 0b11100000
	}
}

// Affects Flags Z and H according to current and new value
func (r *Register) AffectFlagZH(curVal, newVal uint8) {
	r.SetFlagZ(newVal == 0)
	// when bit 3 overflow
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
}

// Affects Flags H and C according to current and new value
func (r *Register) AffectFlagHC(curVal, newVal uint8) {
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
	carry := (curVal&0b11110000 == 0b11110000) && (newVal&0b11110000 == 0)
	r.SetFlagC(carry)
}

// Affects Flags H and C according to 16 bit current and new value
// H is affected by overflow in bit 11, C is affected by overflow in bit 15
func (r *Register) AffectFlagHC16(curVal, newVal uint16) {
	r.AffectFlagHC(uint8(curVal>>8), uint8(newVal>>8))
}

func (r *Register) GetBC() uint16 {
	var result uint16 = uint16(r.C)<<8 + uint16(r.B)

	return result
}

func (r *Register) SetBC(value uint16) {
	r.B = uint8(value)
	r.C = uint8(value >> 8)
}

func (r *Register) GetHL() uint16 {
	var result uint16 = uint16(r.H)<<8 + uint16(r.L)

	return result
}

func (r *Register) SetHL(value uint16) {
	r.L = uint8(value)
	r.H = uint8(value >> 8)
}

type Instruction struct {
	// TODO add comments
	ticks   uint
	execute func() string
}

var (
	bus    Bus
	ticks  uint
	reg    *Register
	curOP  uint8
	nextOP uint8
	inst   [INST_SIZE]Instruction
)

//TODO Add setters and getters for HI LO access

//TODO the lower four bits of Flag Registerare always 0

func init() {

	initInstructions()
}

// Initialise instructions
func initInstructions() {

	// NOP
	inst[0x00] = Instruction{1, nop}
	// LD BC, $FFFF
	inst[0x01] = Instruction{3, ldbc}
	// LD (BC), A
	inst[0x02] = Instruction{2, ldbca}
	// INC BC
	inst[0x03] = Instruction{2, incbc}
	// INC B
	inst[0x04] = Instruction{1, incb}
	// DEC B
	inst[0x05] = Instruction{1, decb}
	// LD B, $FF
	inst[0x06] = Instruction{2, ldb}
	// RLCA
	inst[0x07] = Instruction{1, rlca}
	// LD ($FFFF), SP
	inst[0x08] = Instruction{5, ldmemsp}
	// ADD HL, BC
	inst[0x09] = Instruction{2, addhlbc}
	// LD A, (BC)
	inst[0x0A] = Instruction{2, ldabc}
}

// Emulates machine ticks (m-ticks)
func Tick() {
	ticks--

	if ticks > 0 {
		return
	}

	// execute
	// TODO de-assemble and print
	inst[curOP].execute()

	// Read next instruction
	curOP = nextOP
	reg.PC++
	nextOP = bus.read(reg.PC)
}

func nop() string {
	return "NOP"
}

func ldbc() string {
	value := bus.read16(reg.PC + 1)
	reg.PC += 2
	reg.SetBC(value)

	return fmt.Sprintf("LD BC, %X", value)
}

func ldbca() string {
	pos := reg.GetBC()
	reg.A = bus.read(pos)

	return "LD (BC), A"
}

func incbc() string {
	value := reg.GetBC()
	value++
	reg.SetBC(value)

	return "INC BC"
}

func incb() string {
	reg.AffectFlagZH(reg.B, reg.B+1)
	reg.B++
	reg.SetFlagN(false)

	return "INC B"
}

func decb() string {
	reg.AffectFlagZH(reg.B, reg.B-1)
	reg.B--
	reg.SetFlagN(true)

	return "DEC B"
}

func ldb() string {
	reg.PC++
	reg.B = bus.read(reg.PC)

	return fmt.Sprintf("LD B, $%X", reg.B)
}

func rlca() string {
	// TODO implement
	return ""
}

func ldmemsp() string {
	pos := bus.read16(reg.PC + 1)
	reg.PC += 2
	bus.write16(pos, reg.SP)

	return fmt.Sprintf("LD ($%X), SP", pos)
}

func addhlbc() string {
	curVal := reg.GetHL()
	nextVal := curVal + reg.GetBC()
	reg.SetHL(nextVal)
	reg.SetFlagN(false)
	reg.AffectFlagHC16(curVal, nextVal)

	return "ADD HL, BC"
}

func ldabc() string {
	pos := reg.GetBC()
	reg.A = bus.read(pos)

	return "LD A, (BC)"
}
