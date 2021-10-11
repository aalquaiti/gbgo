package gbgo

import "testing"

func init() {
	reg = &Register{}
}

func TestRegisterGetF(t *testing.T) {
	// Regisger F is not supposed to be set directly, to ensure bit 0-3
	// are always set to Zero
	reg.F = 0b10111001

	var expected uint8 = 0b10110000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Register F not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterGetFlagZ(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagZ()

	if expected != actual {
		t.Errorf("Flag Z not is functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
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
}

func TestRegisterGetFlagN(t *testing.T) {
	reg.F = 0b11111111

	var expected bool = true
	var actual bool = reg.GetFlagN()

	if expected != actual {
		t.Errorf("Flag N not is functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagN(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagN(false)
	var expected uint8 = 0b10110000
	var actual uint8 = reg.GetF()

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
		t.Errorf("Flag H not is functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagH(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagH(false)
	var expected uint8 = 0b11010000
	var actual uint8 = reg.GetF()

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
		t.Errorf("Flag C not is functioning as expected.\nExpected = %t"+
			"\nActual = %t", expected, actual)
	}
}

func TestREgisterSetFlagC(t *testing.T) {
	reg.SetF(0xFF)
	reg.SetFlagC(false)
	var expected uint8 = 0b11100000
	var actual uint8 = reg.GetF()

	if expected != actual {
		t.Errorf("Flag C not is functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestRegisterSetF(t *testing.T) {
	reg.SetF(0b10111001)

	var expected uint8 = 0b10110000
	var actual uint8 = reg.F

	if expected != actual {
		t.Errorf("Register F not functioning as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}

}

func TestRegisterGetBC(t *testing.T) {
	reg.B = 0xFF
	reg.C = 0xFE
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
