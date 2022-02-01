package cpu

// TODO uncomment
//func setup() {
//	Reg = NewRegister()
//	flags, _ = Reg.F.(*RegF)
//}
//
//func TestRlca(t *testing.T) {
//	setup()
//
//	*Reg.A.Val() = 0b11000011
//	rlca()
//
//	if !flags.GetFlagC() {
//		t.Error(`Arithmetic was not performed as expected.
//		Expected Flag C = 1`)
//	}
//
//	var expected uint8 = 0b10000111
//	var actual uint8 = Reg.A.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestRrca(t *testing.T) {
//	setup()
//
//	*Reg.A.Val() = 0b11000011
//	rrca()
//
//	if !flags.GetFlagC() {
//		t.Error(`Arithmetic was not performed as expected.
//		\nExpected Flag C = 1`)
//	}
//
//	var expected uint8 = 0b11100001
//	var actual uint8 = Reg.A.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestRla(t *testing.T) {
//	setup()
//
//	*Reg.A.Val() = 0b11000011
//
//	rla()
//
//	if !flags.GetFlagC() {
//		t.Error(`Arithmetic was not performed as expected.
//		Expected Flag C = 1`)
//	}
//
//	var expected uint8 = 0b10000110
//	var actual uint8 = Reg.A.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestJr(t *testing.T) {
//	setup()
//
//	Reg.PC.Set(0xC000)
//	bus.Write(Reg.PC.Get(), 0x04)
//	jr()
//	var expected uint16 = 0xC005
//	var actual uint16 = Reg.PC.Get()
//
//	if expected != actual {
//		t.Errorf(`Jump was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//
//	Reg.PC.Set(0xC00A)
//	bus.Write(Reg.PC.Get(), 0xFC) //0xFC = -4
//	jr()
//	expected = 0xC007
//	actual = Reg.PC.Get()
//
//	if expected != actual {
//		t.Errorf(`Jump was not performed as expected.
//		Expected = $%X
//		Actual = $%X`, expected, actual)
//	}
//
//}
//
//func TestRra(t *testing.T) {
//	setup()
//
//	*Reg.A.Val() = 0b11000011
//
//	rra()
//
//	if !flags.GetFlagC() {
//		t.Error(`Arithmetic was not performed as expected.
//		Expected Flag C = 1`)
//	}
//
//	var expected uint8 = 0b01100001
//	var actual uint8 = Reg.A.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestIncReg(t *testing.T) {
//	setup()
//
//	*Reg.B.Val() = 0xFF
//	incReg(Reg.B.Val())
//	var expected uint8 = 0
//	var actual uint8 = Reg.B.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestDecReg(t *testing.T) {
//	setup()
//
//	*Reg.B.Val() = 0xFF
//	decReg(Reg.B.Val())
//	var expected uint8 = 0xFE
//	var actual uint8 = Reg.B.Get()
//
//	if expected != actual {
//		t.Errorf(`Arithmetic was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
//
//func TestSwap(t *testing.T) {
//	setup()
//
//	var value uint8 = 0b10100101
//	var expected uint8 = 0b01011010
//	var actual uint8 = swap(value)
//
//	if expected != actual {
//		t.Errorf(`Bit Operation was not performed as expected.
//		Expected = %X
//		Actual = %X`, expected, actual)
//	}
//}
