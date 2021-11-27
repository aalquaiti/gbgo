package cpu

import "testing"

func TestRegFSet(t *testing.T) {
	Reg.F.Set(0b10111001)

	var expected uint8 = 0b10110000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Register F is not functioning as expected."+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

}

func TestRegFGet(t *testing.T) {
	// Regisger F is not supposed to be set directly, to ensure bit 0-3
	// are always set to Zero
	Reg.F.value = 0b10111001

	var expected uint8 = 0b10110000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Register F is not functioning as expected."+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}
}

func TestRegFGetFlagZ(t *testing.T) {
	Reg.F.value = 0b11111111

	var expected bool = true
	var actual bool = Reg.F.GetFlagZ()

	if expected != actual {
		t.Errorf("Flag Z is not functioning as expected."+
			"\nExpected = %t\nActual = %t", expected, actual)
	}
}

func TestRegFSetFlagZ(t *testing.T) {
	Reg.F.Set(0xFF)
	Reg.F.SetFlagZ(false)
	var expected uint8 = 0b01110000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag Z not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	Reg.F.Set(0x0)
	Reg.F.SetFlagZ(true)
	expected = 0b10000000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag Z not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegFGetFlagN(t *testing.T) {
	Reg.F.value = 0b11111111

	var expected bool = true
	var actual bool = Reg.F.GetFlagN()

	if expected != actual {
		t.Errorf("Flag N is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestRegFSetFlagN(t *testing.T) {
	Reg.F.Set(0xFF)
	Reg.F.SetFlagN(false)
	var expected uint8 = 0b10110000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag N is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	Reg.F.Set(0x0)
	Reg.F.SetFlagN(true)
	expected = 0b01000000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag N not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegFGetFlagH(t *testing.T) {
	Reg.F.value = 0b11111111

	var expected bool = true
	var actual bool = Reg.F.GetFlagN()

	if expected != actual {
		t.Errorf("Flag H is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestRegFSetFlagH(t *testing.T) {
	Reg.F.Set(0xFF)
	Reg.F.SetFlagH(false)
	var expected uint8 = 0b11010000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag H is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	Reg.F.Set(0x0)
	Reg.F.SetFlagH(true)
	expected = 0b00100000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag H not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegFGetFlagC(t *testing.T) {
	Reg.F.value = 0b11111111

	var expected bool = true
	var actual bool = Reg.F.GetFlagN()

	if expected != actual {
		t.Errorf("Flag C is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestRegFSetFlagC(t *testing.T) {
	Reg.F.Set(0xFF)
	Reg.F.SetFlagC(false)
	var expected uint8 = 0b11100000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag C is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	Reg.F.Set(0x0)
	Reg.F.SetFlagC(true)
	expected = 0b00010000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Flag C not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegFAffectZH(t *testing.T) {
	var value uint8 = 0
	// Test Half carry without value becoming zero
	value = 0xF // i.e. 0b00001111
	Reg.F.Set(0x0)
	Reg.F.AffectFlagZH(value, value+1)
	var expected uint8 = 0b00100000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry with value becoming zero
	value = 0xFF // i.e 0b11111111
	Reg.F.Set(0x0)
	Reg.F.AffectFlagZH(value, value+1)
	expected = 0b10100000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}
}

func TestRegFAffectHC(t *testing.T) {
	var value uint8 = 0
	// Test Half carry (Flag H) without Full carry (Flag C)
	value = 0xF // i.e. 0b00001111
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC(value, value+1)
	var expected uint8 = 0b00100000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags H and C not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Full caryy (Flag C) without half carry (Flag H)
	value = 0xF0 // i.e. 0b11110000
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC(value, value+0x10) // i.e value + 0b00010000
	expected = 0b00010000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Fully carry being set to true
	value = 0xFF // i.e 0b11111111
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC(value, value+1)
	expected = 0b00110000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Full carry being set to False
}

func TestRegFAffectHC16(t *testing.T) {
	var value uint16 = 0
	// Test Half carry (Flag H) without Full carry (Flag C)
	value = 0xF00
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC16(value, value+0x100)
	var expected uint8 = 0b00100000
	var actual uint8 = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags H and C not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Full caryy (Flag C) without half carry (Flag H)
	value = 0xF000
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC16(value, value+0x1000) // i.e value + 0b00010000
	expected = 0b00010000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Fully carry being set to true
	value = 0xFF00
	Reg.F.Set(0x0)
	Reg.F.AffectFlagHC16(value, value+0x100)
	expected = 0b00110000
	actual = Reg.F.Get()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Full carry being set to False
}

func TestRegFGetBC(t *testing.T) {
	*Reg.B.Val() = 0xFE
	*Reg.C.Val() = 0xFF
	var expected uint16 = 0xFEFF
	var actual uint16 = Reg.BC.Get()

	if expected != actual {
		t.Errorf("Register values are not matched.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterSetBC(t *testing.T) {
	var expected uint16 = 0xFEFF
	Reg.BC.Set(expected)
	var actual = Reg.BC.Get()

	if expected != actual {
		t.Errorf("Register values are not matched.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}
