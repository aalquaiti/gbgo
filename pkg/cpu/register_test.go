package cpu

import "testing"

func init() {
	reg = &Register{}
}

func TestRegisterSetF(t *testing.T) {
	reg.SetF(0b10111001)

	var expected uint8 = 0b10110000
	var actual uint8 = reg.F

	if expected != actual {
		t.Errorf("Register F is not functioning as expected."+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

}

func TestRegisterGetF(t *testing.T) {
	// Regisger F is not supposed to be set directly, to ensure bit 0-3
	// are always set to Zero
	reg.F = 0b10111001

	var expected uint8 = 0b10110000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Register F is not functioning as expected."+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterGetFlagZ(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagZ()

	if expected != actual {
		t.Errorf("Flag Z is not functioning as expected."+
			"\nExpected = %t\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagZ(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagZ(false)
	var expected uint8 = 0b01110000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Flag Z not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	reg.SetF(0x0)
	reg.SetFlagZ(true)
	expected = 0b10000000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Flag Z not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterGetFlagN(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagN()

	if expected != actual {
		t.Errorf("Flag N is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagN(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagN(false)
	var expected uint8 = 0b10110000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Flag N is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	reg.SetF(0x0)
	reg.SetFlagN(true)
	expected = 0b01000000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Flag N not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterGetFlagH(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagN()

	if expected != actual {
		t.Errorf("Flag H is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagH(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagH(false)
	var expected uint8 = 0b11010000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Flag H is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	reg.SetF(0x0)
	reg.SetFlagH(true)
	expected = 0b00100000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Flag H not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterGetFlagC(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagN()

	if expected != actual {
		t.Errorf("Flag C is not functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagC(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagC(false)
	var expected uint8 = 0b11100000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Flag C is not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

	reg.SetF(0x0)
	reg.SetFlagC(true)
	expected = 0b00010000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Flag C not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterAffectZH(t *testing.T) {
	var value uint8 = 0
	// Test Half carry without value becoming zero
	value = 0xF // i.e. 0b00001111
	reg.SetF(0x0)
	reg.AffectFlagZH(value, value+1)
	var expected uint8 = 0b00100000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry with value becoming zero
	value = 0xFF // i.e 0b11111111
	reg.SetF(0x0)
	reg.AffectFlagZH(value, value+1)
	expected = 0b10100000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterAffectHC(t *testing.T) {
	var value uint8 = 0
	// Test Half carry (Flag H) without Full carry (Flag C)
	value = 0xF // i.e. 0b00001111
	reg.SetF(0x0)
	reg.AffectFlagHC(value, value+1)
	var expected uint8 = 0b00100000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags H and C not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Full caryy (Flag C) without half carry (Flag H)
	value = 0xF0 // i.e. 0b11110000
	reg.SetF(0x0)
	reg.AffectFlagHC(value, value+0x10) // i.e value + 0b00010000
	expected = 0b00010000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Fully carry being set to true
	value = 0xFF // i.e 0b11111111
	reg.SetF(0x0)
	reg.AffectFlagHC(value, value+1)
	expected = 0b00110000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Full carry being set to False
}

func TestRegisterAffectHC16(t *testing.T) {
	var value uint16 = 0
	// Test Half carry (Flag H) without Full carry (Flag C)
	value = 0xF00
	reg.SetF(0x0)
	reg.AffectFlagHC16(value, value+0x100)
	var expected uint8 = 0b00100000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags H and C not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Full caryy (Flag C) without half carry (Flag H)
	value = 0xF000
	reg.SetF(0x0)
	reg.AffectFlagHC16(value, value+0x1000) // i.e value + 0b00010000
	expected = 0b00010000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Fully carry being set to true
	value = 0xFF00
	reg.SetF(0x0)
	reg.AffectFlagHC16(value, value+0x100)
	expected = 0b00110000
	actual = reg.GetF()

	if expected != actual {
		t.Errorf("Affecting Flags Z and H not functioning as expected"+
			"\nExpected = 0x%X\nActual = 0x%X", expected, actual)
	}

	// Test Half carry and Full carry being set to False
}

func TestRegisterGetBC(t *testing.T) {
	reg.B = 0xFE
	reg.C = 0xFF
	var expected uint16 = 0xFEFF
	var actual uint16 = reg.GetBC()

	if expected != actual {
		t.Errorf("Register values are not matched.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterSetBC(t *testing.T) {
	var expected uint16 = 0xFEFF
	reg.SetBC(expected)
	var actual = reg.GetBC()

	if expected != actual {
		t.Errorf("Register values are not matched.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestT016(t *testing.T) {
	var low uint8 = 0xFF
	var high uint8 = 0xFE
	var expected uint16 = 0xFEFF
	var actual uint16 = to16(high, low)

	if expected != actual {
		t.Errorf("Function to16 not working as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestLdReg16(t *testing.T) {
	var expected uint16 = 0xFEFF
	ldReg16(&reg.B, &reg.C, expected)
	var actual = reg.GetBC()

	if expected != actual {
		t.Errorf("Register values are not matched.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}
