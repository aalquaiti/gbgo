package cpu

import (
	"fmt"
	"github.com/aalquaiti/gbgo/util/bitutil"
	log "github.com/sirupsen/logrus"
)

// Reg8 represents an 8-bit Register
type Reg8 interface {
	fmt.Stringer
	Val() *uint8 // Pointer to the variable holding the value
	Name() string
	Get() uint8
	Set(uint8) Reg8
	Inc() Reg8
	Dec() Reg8
}

// Reg16 represents a 16-bit Register
type Reg16 interface {
	fmt.Stringer
	Name() string
	Get() uint16
	Set(uint16) Reg16
	Inc() Reg16
	Dec() Reg16
}

// Register represents all registers in a CPU
type Register struct {
	A, F, B, C, D, E, H, L Reg8
	AF, BC, DE, HL         Reg16
	SP, PC                 Reg16
	// Interrupt Master Enable Flag
	IME bool
}

func NewRegister() Register {
	reg := Register{E: &reg8Impl{}}

	// 8-bit Registers
	reg.A = &reg8Impl{name: "A"}
	reg.F = new(RegF)
	reg.B = &reg8Impl{name: "B"}
	reg.C = &reg8Impl{name: "C"}
	reg.D = &reg8Impl{name: "D"}
	reg.E = &reg8Impl{name: "E"}
	reg.H = &reg8Impl{name: "H"}
	reg.L = &reg8Impl{name: "L"}

	// 16-bit Registers consisting of two 8-bits
	// This allows changes happening on a 16-bit register to be reflected to its correspondence 8-bit register
	reg.AF = &reg16From8Impl{high: reg.A, low: reg.F}
	reg.BC = &reg16From8Impl{high: reg.B, low: reg.C}
	reg.DE = &reg16From8Impl{high: reg.D, low: reg.E}
	reg.HL = &reg16From8Impl{high: reg.H, low: reg.L}

	// 16-bit Registers
	reg.SP = &reg16Impl{name: "SP"}
	reg.PC = &reg16Impl{name: "PC"}

	return reg
}

// reg8Impl implements Reg8 interface that represents an 8-bit register
type reg8Impl struct {
	value uint8
	name  string
}

func (r *reg8Impl) Val() *uint8 {

	return &r.value
}

func (r *reg8Impl) Name() string {

	return r.name
}

func (r *reg8Impl) Get() uint8 {
	return r.value
}

func (r *reg8Impl) Set(value uint8) Reg8 {
	r.value = value

	return r
}

func (r *reg8Impl) Inc() Reg8 {
	r.value++

	return r
}

func (r *reg8Impl) Dec() Reg8 {
	r.value--

	return r
}

func (r reg8Impl) String() string {
	return fmt.Sprintf("%s=%.2X", r.name, r.value)
}

// reg16From8Impl implements Reg16 interface that represents a 16-bit register
type reg16Impl struct {
	value uint16
	name  string
}

func (r *reg16Impl) Name() string {

	return r.name
}

func (r *reg16Impl) Get() uint16 {
	return r.value
}

func (r *reg16Impl) Set(value uint16) Reg16 {
	r.value = value

	return r
}

func (r *reg16Impl) Inc() Reg16 {
	r.value++

	return r
}

func (r *reg16Impl) Dec() Reg16 {
	r.value--

	return r
}

func (r reg16Impl) String() string {
	return fmt.Sprintf("%s=%.4X", r.name, r.value)
}

// reg16From8Impl implements Reg16 interface that represents two 8-bit register
type reg16From8Impl struct {
	high Reg8
	low  Reg8
}

func (r *reg16From8Impl) Name() string {

	return r.high.Name() + r.low.Name()
}

func (r *reg16From8Impl) Get() uint16 {
	return bitutil.To16(r.high.Get(), r.low.Get())
}

func (r *reg16From8Impl) Set(value uint16) Reg16 {
	high, low := bitutil.From16(value)
	r.high.Set(high)
	r.low.Set(low)

	return r
}

func (r *reg16From8Impl) Inc() Reg16 {
	r.Set(r.Get() + 1)

	return r
}

func (r *reg16From8Impl) Dec() Reg16 {
	r.Set(r.Get() - 1)

	return r
}

func (r reg16From8Impl) String() string {
	return fmt.Sprintf("%s=%.04X", r.Name(), r.Get())
}

// RegF represents Flag Register
type RegF uint8

func (r *RegF) Val() *uint8 {
	return (*uint8)(r)
}

func (RegF) Name() string {
	return "F"
}

// Get returns F Flag value, with first four bits always set to Zero
func (r RegF) Get() uint8 {
	return uint8(r) & 0b11110000
}

// Set Flag value, with first four bits always set to Zero
func (r *RegF) Set(value uint8) Reg8 {
	*r = RegF(value & 0b11110000)

	return (Reg8)(r)
}

func (r *RegF) Inc() Reg8 {
	// Should not be used. Implemented to match Reg8 interface
	log.Fatal("Inc() for RegF should not be used")

	return r
}

func (r *RegF) Dec() Reg8 {
	// Should not be used. Implemented to match Reg8 interface
	log.Fatal("Dec() for RegF should not be used")

	return r
}

func (r RegF) String() string {
	return fmt.Sprintf("F=%.2X", string(r))
}

func (r *RegF) GetFlagZ() bool {
	return bitutil.IsSet(uint8(*r), 7)
	//return (r.value & 0b10000000) == 0b10000000
}

func (r *RegF) SetFlagZ(value bool) {
	*r = RegF(bitutil.Set(uint8(*r), 7, value))
}

func (r *RegF) GetFlagN() bool {
	return bitutil.IsSet(uint8(*r), 6)
	//return (r.value & 0b01000000) == 0b01000000
}

func (r *RegF) SetFlagN(value bool) {
	*r = RegF(bitutil.Set(uint8(*r), 6, value))
}

func (r *RegF) GetFlagH() bool {
	return bitutil.IsSet(uint8(*r), 5)
	//return (r.value & 0b00100000) == 0b00100000
}

func (r *RegF) SetFlagH(value bool) {
	*r = RegF(bitutil.Set(uint8(*r), 5, value))
}

func (r *RegF) GetFlagC() bool {
	return bitutil.IsSet(uint8(*r), 4)
}

func (r *RegF) SetFlagC(value bool) {
	*r = RegF(bitutil.Set(uint8(*r), 4, value))
}

// AffectFlagZH changes Flags Z and H according to current and new value
func (r *RegF) AffectFlagZH(curVal, newVal uint8) {
	r.SetFlagZ(newVal == 0)
	// when bit 3 overflow
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
}

// AffectFlagHC changes flags H and C according to current and new value
func (r *RegF) AffectFlagHC(curVal, newVal uint8) {
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
	carry := (curVal&0b11110000 == 0b11110000) && (newVal&0b11110000 == 0)
	r.SetFlagC(carry)
}

// AffectFlagHC16 changes flags H and C according to 16 bit current and new value
// H is affected by overflow in bit 11, C is affected by overflow in bit 15
func (r *RegF) AffectFlagHC16(curVal, newVal uint16) {
	r.AffectFlagHC(uint8(curVal>>8), uint8(newVal>>8))
}

// AffectFlagZHC changes flags Z, H and C according to current and new value
func (r *RegF) AffectFlagZHC(curVal, newVal uint8) {
	r.SetFlagZ(newVal == 0)
	// when bit 3 overflow
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
	carry := (curVal&0b11110000 == 0b11110000) && (newVal&0b11110000 == 0)
	r.SetFlagC(carry)
}
