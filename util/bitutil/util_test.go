package bitutil

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

func TestSetBit(t *testing.T) {

	// Test that the fourth bit is set to 1
	var value uint8 = 0b11101010
	var expected uint8 = 0b11111010
	var actual uint8 = Set(value, 4, true)

	assert.Equal(t, expected, actual)

	// Test that the third bit is set to 0
	value = 0b11101010
	expected = 0b11100010
	actual = Set(value, 3, false)

	assert.Equal(t, expected, actual)
}

func TestIsSet(t *testing.T) {
	type args struct {
		value    uint8
		position uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"bit set as 1", args{0b11101010, 3}, true},
		{"bit set as 0", args{0b11101010, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsSet(tt.args.value, tt.args.position),
				"IsSet(%v, %v)", tt.args.value, tt.args.position)
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		value    uint8
		position uint8
		set      bool
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"set bit[3] as 1", args{0b11100010, 3, true}, 0b11101010},
		{"set position beyond 7", args{0b11100010, 8, true}, 0b11100010},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Set(tt.args.value, tt.args.position, tt.args.set),
				"Set(%v, %v, %v)", tt.args.value, tt.args.position, tt.args.set)
		})
	}
}
