package cpu

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

// Affects Flags Z, H and C according to current and new value
func (r *Register) AffectFlagZHC(curVal, newVal uint8) {
	r.SetFlagZ(newVal == 0)
	// when bit 3 overflow
	halfCarry := (curVal&0b1111 == 0b1111) && (newVal&0b1111 == 0)
	r.SetFlagH(halfCarry)
	carry := (curVal&0b11110000 == 0b11110000) && (newVal&0b11110000 == 0)
	r.SetFlagC(carry)
}

func (r *Register) GetBC() uint16 {
	return to16(r.B, r.C)
}

func (r *Register) SetBC(value uint16) {
	ldReg16(&r.B, &r.C, value)
}

func (r *Register) GetDE() uint16 {
	return to16(r.D, r.E)
}

func (r *Register) SetDE(value uint16) {
	ldReg16(&r.D, &r.E, value)
}

func (r *Register) GetHL() uint16 {
	return to16(r.H, r.L)
}

func (r *Register) SetHL(value uint16) {
	ldReg16(&r.H, &r.L, value)
}

// Helper functions

// Combine two 8-bit values to one 16-bit value
func to16(rHigh, rLow uint8) uint16 {
	return uint16(rHigh)<<8 + uint16(rLow)
}

// Load 16-bit value to a 16-bit Register.
// Register represented two 8-bit high and low registers
func ldReg16(rHigh, rLow *uint8, value uint16) {
	*rLow = uint8(value)
	*rHigh = uint8(value >> 8)
}
