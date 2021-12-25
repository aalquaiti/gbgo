package io

import "github.com/aalquaiti/gbgo/gbgoutil"

// IE Represents Interrupt Enable Register
type IE uint8

// IF Represents Interrupt Flag Register
type IF uint8

type TimeReg struct {
	Div  uint8
	Tima uint8
	Tma  uint8
	Tac  uint8
}

// IsVBlank determines if VBlank Interrupt is Enabled
func (i *IE) IsVBlank() bool {
	return gbgoutil.IsBitSet(uint8(*i), 0)
}

func (i *IE) SetVBlank(enable bool) {
	*i = IE(gbgoutil.SetBit(uint8(*i), 0, enable))
}

// IsLCDStat determines if LCD Status Interrupt is Enabled
func (i *IE) IsLCDStat() bool {
	return gbgoutil.IsBitSet(uint8(*i), 1)
}

func (i *IE) SetLCDStat(enable bool) {
	*i = IE(gbgoutil.SetBit(uint8(*i), 1, enable))
}

// IsTimerInt determines if Timer Interrupt is Enabled
func (i *IE) IsTimerInt() bool {
	return gbgoutil.IsBitSet(uint8(*i), 2)
}

func (i *IE) SetTimerInt(enable bool) {
	*i = IE(gbgoutil.SetBit(uint8(*i), 2, enable))
}

// IsSerialInt determines if Serial Interrupt is Enabled
func (i *IE) IsSerialInt() bool {
	return gbgoutil.IsBitSet(uint8(*i), 3)
}

func (i *IE) SetSerialInt(enable bool) {
	*i = IE(gbgoutil.SetBit(uint8(*i), 3, enable))
}

// IsJoypadInt determines if Joypad Interrupt is Enabled
func (i *IE) IsJoypadInt() bool {
	return gbgoutil.IsBitSet(uint8(*i), 4)
}

func (i *IE) SetJoypad(enable bool) {
	*i = IE(gbgoutil.SetBit(uint8(*i), 4, enable))
}

// IrqVBlank determines if VBlank Interrupt is Requested
func (i *IF) IrqVBlank() bool {
	return gbgoutil.IsBitSet(uint8(*i), 0)
}

func (i *IF) SetIrQVblank(enable bool) {
	*i = IF(gbgoutil.SetBit(uint8(*i), 0, enable))
}

// IrqLCDStat determines if LCD Status Interrupt is Requested
func (i *IF) IrqLCDStat() bool {
	return gbgoutil.IsBitSet(uint8(*i), 1)
}

func (i *IF) SetIRQLCDStat(enable bool) {
	*i = IF(gbgoutil.SetBit(uint8(*i), 1, enable))
}

// IrqTimer determines if Timer Interrupt is Requested
func (i *IF) IrqTimer() bool {
	return gbgoutil.IsBitSet(uint8(*i), 2)
}

func (i *IF) SetIRQTimer(enable bool) {
	*i = IF(gbgoutil.SetBit(uint8(*i), 2, enable))
}

// IrqSerial determines if Serial Interrupt is Requested
func (i *IF) IrqSerial() bool {
	return gbgoutil.IsBitSet(uint8(*i), 3)
}

func (i *IF) SetIrqSerial(enable bool) {
	*i = IF(gbgoutil.SetBit(uint8(*i), 3, enable))
}

// IrqJoyPad determines if JoyPad Interrupt is Requested
func (i *IF) IrqJoyPad() bool {
	return gbgoutil.IsBitSet(uint8(*i), 4)
}

func (i *IF) SetIrqJoyPad(enable bool) {
	*i = IF(gbgoutil.SetBit(uint8(*i), 4, enable))
}

// IncDIV Increment Divider Register by one.
// This is used instead of writing to memory as writing to avoid reseting its value, as an expected behaviour
// by the game boy io
func (t *TimeReg) IncDIV() {
	t.Div++
}

// IsTacTimerEnabled determines Timer Control (TAC) bit 2 to determine if Timer is Enabled. When enabled, Timer Counter
// can be incremented. This does not affect Divider Register
func (t *TimeReg) IsTacTimerEnabled() bool {
	return gbgoutil.IsBitSet(t.Tac, 2)
}

// GetTacClockSelect Retrieve Timer Control (TAC) bits 0 and 1 that determine the Clock Selected for Timer Counter
func (t *TimeReg) GetTacClockSelect() uint8 {
	return t.Tac & 0b11
}
