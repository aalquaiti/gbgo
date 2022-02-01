package cpu

import "testing"

func TestDisassembler_mnc(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"NOP", fields{op: opCodes[0x00], arg1: 0x10, arg2: 0x10}, "NOP"},
		{"STOP", fields{op: opCodes[0x10], arg1: 0x10, arg2: 0x10}, "STOP"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mnc(); got != tt.want {
				t.Errorf("mnc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncConst8(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"JR", fields{op: opCodes[0x18], arg1: 0x10, arg2: 0x10}, "JR 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncConst8(); got != tt.want {
				t.Errorf("mncConst8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncConst16(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"JP", fields{op: opCodes[0xC3], arg1: 0x10, arg2: 0x20}, "JP 2010"},
		{"CALL", fields{op: opCodes[0xCD], arg1: 0x20, arg2: 0x10}, "CALL 1020"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncConst16(); got != tt.want {
				t.Errorf("mncConst16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncReg(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"DEC B", fields{op: opCodes[0x05], arg1: 0x10, arg2: 0x10}, "DEC B"},
		{"INC HL", fields{op: opCodes[0x23], arg1: 0x10, arg2: 0x10}, "INC HL"},
		{"PUSH DE", fields{op: opCodes[0xD5], arg1: 0x10, arg2: 0x10}, "PUSH DE"},
		{"SLA H", fields{op: cpOpCodes[0x24], arg1: 0x10, arg2: 0x10}, "SLA H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncReg(); got != tt.want {
				t.Errorf("mncReg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncFlag(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"RET NZ", fields{op: opCodes[0xC0], arg1: 0x10, arg2: 0x10}, "RET NZ"},
		{"RET Z", fields{op: opCodes[0xC8], arg1: 0x10, arg2: 0x10}, "RET Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncFlag(); got != tt.want {
				t.Errorf("mncFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncVec(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"RST 18", fields{op: opCodes[0xDF], arg1: 0x10, arg2: 0x10}, "RST 18"},
		{"RST 20", fields{op: opCodes[0xE7], arg1: 0x10, arg2: 0x10}, "RST 20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncVec(); got != tt.want {
				t.Errorf("mncVec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncConst8Reg(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"LD (u8), A", fields{op: opCodes[0xE0], arg1: 0x10, arg2: 0x10}, "LD (FF00+10), A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncConst8Reg(); got != tt.want {
				t.Errorf("mncConst8Reg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncConst16Reg(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"LD (u16), A", fields{op: opCodes[0xEA], arg1: 0x10, arg2: 0x20}, "LD (2010), A"},
		{"LD (u16), SP", fields{op: opCodes[0x08], arg1: 0x10, arg2: 0x20}, "LD (2010), SP"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncConst16Reg(); got != tt.want {
				t.Errorf("mncConst16Reg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncRegConst8(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"LD D, u8", fields{op: opCodes[0x16], arg1: 0x10, arg2: 0x20}, "LD D, 10"},
		{"LD L, u8", fields{op: opCodes[0x2E], arg1: 0x10, arg2: 0x20}, "LD L, 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncRegConst8(); got != tt.want {
				t.Errorf("mncRegConst8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncRegConst16(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"LD BC, u16", fields{op: opCodes[0x01], arg1: 0x10, arg2: 0x20}, "LD BC, 2010"},
		{"LD SP, u18", fields{op: opCodes[0x31], arg1: 0x20, arg2: 0x10}, "LD SP, 1020"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncRegConst16(); got != tt.want {
				t.Errorf("mncRegConst16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncRegReg(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"ADD R16 R16", fields{op: opCodes[0x09], arg1: 0x10, arg2: 0x20}, "ADD HL, BC"},
		{"LD R8, IR16", fields{op: opCodes[0x0A], arg1: 0x20, arg2: 0x10}, "LD A, (BC)"},
		{"LD R8, (HL+)", fields{op: opCodes[0x2A], arg1: 0x20, arg2: 0x10}, "LD A, (HL+)"},
		{"LD R8, (HL-)", fields{op: opCodes[0x3A], arg1: 0x20, arg2: 0x10}, "LD A, (HL-)"},
		{"LD R8, (HL)", fields{op: opCodes[0x4E], arg1: 0x20, arg2: 0x10}, "LD C, (HL)"},
		{"LD R8, R8", fields{op: opCodes[0x5C], arg1: 0x20, arg2: 0x10}, "LD E, H"},
		{"LD R, R", fields{op: opCodes[0x6D], arg1: 0x20, arg2: 0x10}, "LD L, L"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncRegReg(); got != tt.want {
				t.Errorf("mncRegReg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncFlagConst8(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"JR NZ, I8", fields{op: opCodes[0x20], arg1: 0x10, arg2: 0x20}, "JR NZ, 10"},
		{"JR Z, I8", fields{op: opCodes[0x28], arg1: 0x10, arg2: 0x20}, "JR Z, 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncFlagConst8(); got != tt.want {
				t.Errorf("mncFlagConst8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncFlagConst16(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"JP NZ, u16", fields{op: opCodes[0xC2], arg1: 0x10, arg2: 0x20}, "JP NZ, 2010"},
		{"JP C, u16", fields{op: opCodes[0xDA], arg1: 0x10, arg2: 0x20}, "JP C, 2010"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncFlagConst16(); got != tt.want {
				t.Errorf("mncFlagConst16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisassembler_mncBitReg(t *testing.T) {
	type fields struct {
		op      OpCode
		arg1    byte
		arg2    byte
		funcs   dsmFunc
		cpFuncs dsmFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Bit #, Reg", fields{op: cpOpCodes[0x5C], arg1: 0x10, arg2: 0x20}, "BIT 3, H"},
		{"Bit #, (HL)", fields{op: cpOpCodes[0x6E], arg1: 0x10, arg2: 0x20}, "BIT 5, (HL)"},
		{"RES #, A", fields{op: cpOpCodes[0xA7], arg1: 0x10, arg2: 0x20}, "RES 4, A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Disassembler{
				op:      tt.fields.op,
				arg1:    tt.fields.arg1,
				arg2:    tt.fields.arg2,
				funcs:   tt.fields.funcs,
				cpFuncs: tt.fields.cpFuncs,
			}
			if got := d.mncBitReg(); got != tt.want {
				t.Errorf("mncBitReg() = %v, want %v", got, tt.want)
			}
		})
	}
}
