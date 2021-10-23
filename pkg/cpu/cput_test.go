package cpu

import (
	"testing"
)

func setup() {
	reg = &Register{}
}

func TestRlca(t *testing.T) {
	setup()

	reg.A = 0b11000011
	rlca()

	if !reg.GetFlagC() {
		t.Error(`Arithmetic was not performed as expected. 
		Expected Flag C = 1`)
	}

	var expected uint8 = 0b10000111
	var actual uint8 = reg.A

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}

func TestRrca(t *testing.T) {
	setup()

	reg.A = 0b11000011
	rrca()

	if !reg.GetFlagC() {
		t.Error(`Arithmetic was not performed as expected. 
		\nExpected Flag C = 1`)
	}

	var expected uint8 = 0b11100001
	var actual uint8 = reg.A

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}

func TestRla(t *testing.T) {
	setup()

	reg.A = 0b11000011

	rla()

	if !reg.GetFlagC() {
		t.Error(`Arithmetic was not performed as expected. 
		Expected Flag C = 1`)
	}

	var expected uint8 = 0b10000110
	var actual uint8 = reg.A

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}

func TestJr(t *testing.T) {
	setup()

	reg.PC = 0x0A
	bus.Write(reg.PC+1, 0x04)
	jr()
	var expected uint8 = 0x0F
	var actual uint8 = uint8(reg.PC)

	if expected != actual {
		t.Errorf(`Jump was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}

	reg.PC = 0x0A
	bus.Write(reg.PC+1, 0xFC) //0xFC = -4
	jr()
	expected = 0x07
	actual = uint8(reg.PC)

	if expected != actual {
		t.Errorf(`Jump was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}

}

func TestRra(t *testing.T) {
	setup()

	reg.A = 0b11000011

	rra()

	if !reg.GetFlagC() {
		t.Error(`Arithmetic was not performed as expected. 
		Expected Flag C = 1`)
	}

	var expected uint8 = 0b01100001
	var actual uint8 = reg.A

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}

func TestIncReg(t *testing.T) {
	setup()

	reg.B = 0xFF
	incmem(&reg.B)
	var expected uint8 = 0
	var actual uint8 = reg.B

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}

func TestDecReg(t *testing.T) {
	setup()

	reg.B = 0xFF
	decmem(&reg.B)
	var expected uint8 = 0
	var actual uint8 = reg.B

	if expected != actual {
		t.Errorf(`Arithmetic was not performed as expected.
		Expected = %X
		Actual = %X`, expected, actual)
	}
}
