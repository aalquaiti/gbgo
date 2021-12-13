package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo16(t *testing.T) {
	var low uint8 = 0xFF
	var high uint8 = 0xFE
	var expected uint16 = 0xFEFF
	var actual uint16 = To16(high, low)

	if expected != actual {
		t.Errorf("Function to16 not working as expected.\nExpected = 0x%X"+
			"\nActual = 0x%X", expected, actual)
	}
}

func TestFrom16(t *testing.T) {
	var expectedLow uint8 = 0xFF
	var expectedHigh uint8 = 0xFE
	var expected uint16 = 0xFEFF
	actualHigh, actualLow := From16(expected)

	// if actualLow != expectedLow || actualHigh != expectedHigh {
	// 	t.Errorf("Function From not working as expected.\nExpected = 0x%X"+
	// 		"\nActual = 0x%X", expected, actual)
	// }
	assert.Equal(t, expectedLow, actualLow)
	assert.Equal(t, expectedHigh, actualHigh)
}
