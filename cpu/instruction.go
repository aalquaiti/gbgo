package cpu

import (
	"fmt"
	"github.com/aalquaiti/gbgo/util"
)

// Instruction defines a CPU instruction
type Instruction struct {
	ticks   uint8         // how many m-ticks required
	execute func() string // the instruction to perform
}

// Initialise instructions
func initInstructions() {
	// region Instruction Set

	// NOP
	inst[0x00] = Instruction{1, nop}
	// LD BC, $FFFF
	inst[0x01] = Instruction{3, ldbc16}
	// LD (BC), A
	inst[0x02] = Instruction{2, ldbca}
	// INC BC
	inst[0x03] = Instruction{2, incbc}
	// INC B
	inst[0x04] = Instruction{1, incb}
	// DEC B
	inst[0x05] = Instruction{1, decb}
	// LD B, $FF
	inst[0x06] = Instruction{2, ldb}
	// RLCA
	inst[0x07] = Instruction{1, rlca}
	// LD ($FFFF), SP
	inst[0x08] = Instruction{5, ldmemsp}
	// ADD HL, BC
	inst[0x09] = Instruction{2, addhlbc}
	// LD A, (BC)
	inst[0x0A] = Instruction{2, ldabc}
	// DEC BC
	inst[0x0B] = Instruction{2, decbc}
	// INC C
	inst[0x0C] = Instruction{1, incc}
	// DEC C
	inst[0x0D] = Instruction{1, decc}
	// LD C, $FF
	inst[0x0E] = Instruction{2, ldc}
	// RRCA
	inst[0x0F] = Instruction{1, rrca}

	// STOP
	inst[0x10] = Instruction{1, stop}
	// LD DE, $FFFF
	inst[0x11] = Instruction{3, ldde16}
	// LD (DE), A
	inst[0x12] = Instruction{2, lddea}
	// INC DE
	inst[0x13] = Instruction{2, incde}
	// INC D
	inst[0x14] = Instruction{1, incd}
	// DEC D
	inst[0x15] = Instruction{1, decd}
	// LD D, $FF
	inst[0x16] = Instruction{2, ldd}
	// RLA
	inst[0x17] = Instruction{1, rla}
	// JR, $FF
	inst[0x18] = Instruction{1, jr}
	// ADD HL, DE
	inst[0x19] = Instruction{2, addhlde}
	// LD A, (DE)
	inst[0x1A] = Instruction{2, ldade}
	// DEC DE
	inst[0x1B] = Instruction{2, decde}
	// INC E
	inst[0x1C] = Instruction{1, ince}
	// DEC E
	inst[0x1D] = Instruction{1, dece}
	// LD E, $FF
	inst[0x1E] = Instruction{2, lde}
	// RRA
	inst[0x1F] = Instruction{1, rra}

	// JR NZ, $FF
	inst[0x20] = Instruction{2, jrnz}
	// LD HL, $FFFF
	inst[0x21] = Instruction{3, ldhl16}
	// LD (HLI), A
	inst[0x22] = Instruction{2, ldhli}
	// INC HL
	inst[0x23] = Instruction{2, inchl}
	// INC H
	inst[0x24] = Instruction{1, inch}
	// DEC H
	inst[0x25] = Instruction{1, dech}
	// LD H, $FF
	inst[0x26] = Instruction{2, ldh}
	// DAA
	inst[0x27] = Instruction{1, daa}
	// JR Z, $FF
	inst[0x28] = Instruction{2, jrz}
	// ADD HL, HL
	inst[0x29] = Instruction{2, addhlhl}
	// LD A, (HLI)
	inst[0x2A] = Instruction{2, ldahli}
	// DEC HL
	inst[0x2B] = Instruction{2, dechl}
	// INC L
	inst[0x2C] = Instruction{1, incl}
	// DEC L
	inst[0x2D] = Instruction{1, decl}
	// LD L, $FF
	inst[0x2E] = Instruction{2, ldl}
	// CPL
	inst[0x2F] = Instruction{1, cpl}

	// JR NC, $FF
	inst[0x30] = Instruction{2, jrnc}
	// LD SP, $FFFF
	inst[0x31] = Instruction{3, ldsp}
	// LD (HLD), A
	inst[0x32] = Instruction{2, ldhlda}
	// INC SP
	inst[0x33] = Instruction{2, incsp}
	// INC (HL)
	inst[0x34] = Instruction{3, inchlind}
	// DEC (HL)
	inst[0x35] = Instruction{3, dechlind}
	// LD (HL), $FF
	inst[0x36] = Instruction{3, ldhl8}
	// SCF
	inst[0x37] = Instruction{1, scf}
	// JR C, $FF
	inst[0x38] = Instruction{2, jrc}
	// ADD HL, SP
	inst[0x39] = Instruction{2, addhlsp}
	// LD A, (HLD)
	inst[0x3A] = Instruction{2, ldahld}
	// DEC SP
	inst[0x3B] = Instruction{2, decsp}
	// INC A
	inst[0x3C] = Instruction{1, inca}
	// DEC A
	inst[0x3D] = Instruction{1, deca}
	// LD A, $FF
	inst[0x3E] = Instruction{2, lda}
	// CCF
	inst[0x3F] = Instruction{1, ccf}

	// LD B, B
	inst[0x40] = Instruction{1, ldbb}
	// LD B, C
	inst[0x41] = Instruction{1, ldbc}
	// LD B, D
	inst[0x42] = Instruction{1, ldbd}
	// LD B, E
	inst[0x43] = Instruction{1, ldbe}
	// LD B, H
	inst[0x44] = Instruction{1, ldbh}
	// LD B, L
	inst[0x45] = Instruction{1, ldbl}
	// LD B, (HL)
	inst[0x46] = Instruction{2, ldbhl}
	// LD B, A
	inst[0x47] = Instruction{1, ldba}
	// LD C, B
	inst[0x48] = Instruction{1, ldcb}
	// LD C, C
	inst[0x49] = Instruction{1, ldcc}
	// LD C, D
	inst[0x4A] = Instruction{1, ldcd}
	// LD C, E
	inst[0x4B] = Instruction{1, ldce}
	// LD C, H
	inst[0x4C] = Instruction{1, ldch}
	// LD C, L
	inst[0x4D] = Instruction{1, ldcl}
	// LD C, (HL)
	inst[0x4E] = Instruction{2, ldchl}
	// LD C, A
	inst[0x4F] = Instruction{1, ldca}

	// LD D, B
	inst[0x50] = Instruction{1, lddb}
	// LD D, C
	inst[0x51] = Instruction{1, lddc}
	// LD D, D
	inst[0x52] = Instruction{1, lddd}
	// LD D, E
	inst[0x53] = Instruction{1, ldde}
	// LD D, H
	inst[0x54] = Instruction{1, lddh}
	// LD D, L
	inst[0x55] = Instruction{1, lddl}
	// LD D, (HL)
	inst[0x56] = Instruction{2, lddhl}
	// LD D, A
	inst[0x57] = Instruction{1, ldda}
	// LD E, B
	inst[0x58] = Instruction{1, ldeb}
	// LD E, C
	inst[0x59] = Instruction{1, ldec}
	// LD E, D
	inst[0x5A] = Instruction{1, lded}
	// LD E, E
	inst[0x5B] = Instruction{1, ldee}
	// LD E, H
	inst[0x5C] = Instruction{1, ldeh}
	// LD E, L
	inst[0x5D] = Instruction{1, ldel}
	// LD E, (HL)
	inst[0x5E] = Instruction{2, ldehl}
	// LD E, A
	inst[0x5F] = Instruction{1, ldea}

	// LD H, B
	inst[0x60] = Instruction{1, ldhb}
	// LD H, C
	inst[0x61] = Instruction{1, ldhc}
	// LD H, D
	inst[0x62] = Instruction{1, ldhd}
	// LD H, E
	inst[0x63] = Instruction{1, ldhe}
	// LD H, H
	inst[0x64] = Instruction{1, ldhh}
	// LD H, L
	inst[0x65] = Instruction{1, ldhl}
	// LD H, (HL)
	inst[0x66] = Instruction{2, ldhhl}
	// LD H, A
	inst[0x67] = Instruction{1, ldha}
	// LD L, B
	inst[0x68] = Instruction{1, ldlb}
	// LD L, C
	inst[0x69] = Instruction{1, ldlc}
	// LD L, D
	inst[0x6A] = Instruction{1, ldld}
	// LD L, E
	inst[0x6B] = Instruction{1, ldle}
	// LD L, H
	inst[0x6C] = Instruction{1, ldlh}
	// LD L, L
	inst[0x6D] = Instruction{1, ldll}
	// LD L, (HL)
	inst[0x6E] = Instruction{2, ldlhl}
	// LD L, A
	inst[0x6F] = Instruction{1, ldla}

	// LD (HL), B
	inst[0x70] = Instruction{1, ldhlb}
	// LD (HL), C
	inst[0x71] = Instruction{1, ldhlc}
	// LD (HL), D
	inst[0x72] = Instruction{1, ldhld}
	// LD (HL), E
	inst[0x73] = Instruction{1, ldhle}
	// LD (HL), H
	inst[0x74] = Instruction{1, ldhlh}
	// LD (HL), L
	inst[0x75] = Instruction{1, ldhll}
	// HALT
	inst[0x76] = Instruction{1, halt}
	// LD (HL), A
	inst[0x77] = Instruction{2, ldhla}
	// LD A, B
	inst[0x78] = Instruction{1, ldab}
	// LD A, C
	inst[0x79] = Instruction{1, ldac}
	// LD A, D
	inst[0x7A] = Instruction{1, ldad}
	// LD A, E
	inst[0x7B] = Instruction{1, ldae}
	// LD A, H
	inst[0x7C] = Instruction{1, ldah}
	// LD A, L
	inst[0x7D] = Instruction{1, ldal}
	// LD A, (HL)
	inst[0x7E] = Instruction{2, ldahl}
	// LD A, A
	inst[0x7F] = Instruction{1, ldaa}

	// ADD A, B
	inst[0x80] = Instruction{1, addab}
	// ADD A, C
	inst[0x81] = Instruction{1, addac}
	// ADD A, D
	inst[0x82] = Instruction{1, addad}
	// ADD A, E
	inst[0x83] = Instruction{1, addae}
	// ADD A, H
	inst[0x84] = Instruction{1, addah}
	// ADD A, L
	inst[0x85] = Instruction{1, addal}
	// ADD A, (HL)
	inst[0x86] = Instruction{2, addahl}
	// ADD A, A
	inst[0x87] = Instruction{1, addaa}
	// ADC A, B
	inst[0x88] = Instruction{1, adcab}
	// ADC A, C
	inst[0x89] = Instruction{1, adcac}
	// ADC A, D
	inst[0x8A] = Instruction{1, adcad}
	// ADC A, E
	inst[0x8B] = Instruction{1, adcae}
	// ADC A, H
	inst[0x8C] = Instruction{1, adcah}
	// ADC A, L
	inst[0x8D] = Instruction{1, adcal}
	// ADC A, (HL)
	inst[0x8E] = Instruction{2, adcahl}
	// ADC A, A
	inst[0x8F] = Instruction{1, adcaa}

	// SUB A, B
	inst[0x90] = Instruction{1, subab}
	// SUB A, C
	inst[0x91] = Instruction{1, subac}
	// SUB A, D
	inst[0x92] = Instruction{1, subad}
	// SUB A, E
	inst[0x93] = Instruction{1, subae}
	// SUB A, H
	inst[0x94] = Instruction{1, subah}
	// SUB A, L
	inst[0x95] = Instruction{1, subal}
	// SUB A, (HL)
	inst[0x96] = Instruction{2, subahl}
	// SUB A, A
	inst[0x97] = Instruction{1, subaa}
	// SBC A, B
	inst[0x98] = Instruction{1, sbcab}
	// SBC A, C
	inst[0x99] = Instruction{1, sbcac}
	// SBC A, D
	inst[0x9A] = Instruction{1, sbcad}
	// SBC A, E
	inst[0x9B] = Instruction{1, sbcae}
	// SBC A, H
	inst[0x9C] = Instruction{1, sbcah}
	// SBC A, L
	inst[0x9D] = Instruction{1, sbcal}
	// SBC A, (HL)
	inst[0x9E] = Instruction{2, sbcahl}
	// SBC A, A
	inst[0x9F] = Instruction{1, sbcaa}

	// AND A, B
	inst[0xA0] = Instruction{1, andab}
	// AND A, C
	inst[0xA1] = Instruction{1, andac}
	// AND A, D
	inst[0xA2] = Instruction{1, andad}
	// AND A, E
	inst[0xA3] = Instruction{1, andae}
	// AND A, H
	inst[0xA4] = Instruction{1, andah}
	// AND A, L
	inst[0xA5] = Instruction{1, andal}
	// AND A, (HL)
	inst[0xA6] = Instruction{2, andahl}
	// AND A, A
	inst[0xA7] = Instruction{1, andaa}
	// XOR A, B
	inst[0xA8] = Instruction{1, xorab}
	// XOR A, C
	inst[0xA9] = Instruction{1, xorac}
	// XOR A, D
	inst[0xAA] = Instruction{1, xorad}
	// XOR A, E
	inst[0xAB] = Instruction{1, xorae}
	// XOR A, H
	inst[0xAC] = Instruction{1, xorah}
	// XOR A, L
	inst[0xAD] = Instruction{1, xoral}
	// XOR A, (HL)
	inst[0xAE] = Instruction{2, xorahl}
	// XOR A, A
	inst[0xAF] = Instruction{1, xoraa}

	// OR A, B
	inst[0xB0] = Instruction{1, orab}
	// OR A, C
	inst[0xB1] = Instruction{1, orac}
	// OR A, D
	inst[0xB2] = Instruction{1, orad}
	// OR A, E
	inst[0xB3] = Instruction{1, orae}
	// OR A, H
	inst[0xB4] = Instruction{1, orah}
	// OR A, L
	inst[0xB5] = Instruction{1, oral}
	// OR A, (HL)
	inst[0xB6] = Instruction{2, orahl}
	// OR A, A
	inst[0xB7] = Instruction{1, oraa}
	// CP A, B
	inst[0xB8] = Instruction{1, cpab}
	// CP A, C
	inst[0xB9] = Instruction{1, cpac}
	// CP A, D
	inst[0xBA] = Instruction{1, cpad}
	// CP A, E
	inst[0xBB] = Instruction{1, cpae}
	// CP A, H
	inst[0xBC] = Instruction{1, cpah}
	// CP A, L
	inst[0xBD] = Instruction{1, cpal}
	// CP A, (HL)
	inst[0xBE] = Instruction{2, cpahl}
	// CP A, A
	inst[0xBF] = Instruction{1, cpaa}

	// RET NZ
	inst[0xC0] = Instruction{2, retnz}
	// POP BC
	inst[0xC1] = Instruction{3, popbc}
	// JP NZ, $FFFF
	inst[0xC2] = Instruction{3, jpnz}
	// JP $FFFF
	inst[0xC3] = Instruction{3, jp}
	// CALL NZ, $FFFF
	inst[0xC4] = Instruction{3, callnz}
	// PUSH BC
	inst[0xC5] = Instruction{4, pushbc}
	// ADD A, $FF
	inst[0xC6] = Instruction{2, adda8}
	// RST $00
	inst[0xC7] = Instruction{4, rst00}
	// RET Z
	inst[0xC8] = Instruction{2, retz}
	// RET
	inst[0xC9] = Instruction{4, ret}
	// JP Z, $FFFF
	inst[0xCA] = Instruction{3, jpz}
	// PREFIX CB
	inst[0xCB] = Instruction{1, prefixcb}
	// CALL Z, $FFFF
	inst[0xCC] = Instruction{3, callz}
	// CALL $FFFF
	inst[0xCD] = Instruction{6, call}
	// ADC A, $FF
	inst[0xCE] = Instruction{2, adca8}
	// RST $08
	inst[0xCF] = Instruction{4, rst08}

	// RET NC
	inst[0xD0] = Instruction{2, retnc}
	// POP DE
	inst[0xD1] = Instruction{3, popde}
	// JP NC, $FFFF
	inst[0xD2] = Instruction{1, jpnc}
	//
	inst[0xD3] = Instruction{0, illegalop}
	// CALL NC, $FFFF
	inst[0xD4] = Instruction{3, callnc}
	// PUSH DE
	inst[0xD5] = Instruction{4, pushde}
	// SUB A, $FF
	inst[0xD6] = Instruction{2, suba8}
	// RST $10
	inst[0xD7] = Instruction{4, rst10}
	// RET C
	inst[0xD8] = Instruction{2, retc}
	// RETI
	inst[0xD9] = Instruction{4, reti}
	// JP C, $FFFF
	inst[0xDA] = Instruction{3, jpc}
	//
	inst[0xDB] = Instruction{0, illegalop}
	// CALL C, $FFFF
	inst[0xDC] = Instruction{3, callc}
	//
	inst[0xDD] = Instruction{0, illegalop}
	// SBC A, $FF
	inst[0xDE] = Instruction{2, sbca8}
	// RST $18
	inst[0xDF] = Instruction{4, rst18}

	// LD (FF00 + $FF), A
	inst[0xE0] = Instruction{2, ldff8a}
	// POP HL
	inst[0xE1] = Instruction{3, pophl}
	// LD (FF00 + C), A
	inst[0xE2] = Instruction{1, ldffca}
	//
	inst[0xE3] = Instruction{0, illegalop}
	//
	inst[0xE4] = Instruction{3, illegalop}
	// PUSH HL
	inst[0xE5] = Instruction{4, pushhl}
	// AND A, $FF
	inst[0xE6] = Instruction{2, anda8}
	// RST $20
	inst[0xE7] = Instruction{4, rst20}
	// ADD SP, $FF
	inst[0xE8] = Instruction{2, addsp}
	// JP HL
	inst[0xE9] = Instruction{4, jphl}
	// LD ($FFFF), A
	inst[0xEA] = Instruction{3, ld16a}
	//
	inst[0xEB] = Instruction{0, illegalop}
	//
	inst[0xEC] = Instruction{3, illegalop}
	//
	inst[0xED] = Instruction{0, illegalop}
	// XOR A, $FF
	inst[0xEE] = Instruction{2, xora8}
	// RST $28
	inst[0xEF] = Instruction{4, rst28}

	// LD A, (FF00 + $FF)
	inst[0xF0] = Instruction{2, ldaff8}
	// POP AF
	inst[0xF1] = Instruction{3, popaf}
	// LD A, (FF00 + C)
	inst[0xF2] = Instruction{1, ldaffc}
	// DI
	inst[0xF3] = Instruction{0, di}
	//
	inst[0xF4] = Instruction{3, illegalop}
	// PUSH AF
	inst[0xF5] = Instruction{4, pushaf}
	// OR A, $FF
	inst[0xF6] = Instruction{2, ora8}
	// RST $30
	inst[0xF7] = Instruction{4, rst30}
	// LD HP, SP + $FF
	inst[0xF8] = Instruction{2, ldhlsp8}
	// LD SP, HL
	inst[0xF9] = Instruction{4, ldsphl}
	// LD A, ($FFFF)
	inst[0xFA] = Instruction{3, lda16}
	// EI
	inst[0xFB] = Instruction{0, ei}
	//
	inst[0xFC] = Instruction{3, illegalop}
	//
	inst[0xFD] = Instruction{0, illegalop}
	// CP A, $FF
	inst[0xFE] = Instruction{2, cpa8}
	// RST $38
	inst[0xFF] = Instruction{4, rst38}

	// endregion Instruction Set

	// region CB Prefixed Instructions

	// RLC B
	cbInst[0x00] = Instruction{2, rlcb}
	// RLC C
	cbInst[0x01] = Instruction{2, rlcc}
	// RLC D
	cbInst[0x02] = Instruction{2, rlcd}
	// RLC E
	cbInst[0x03] = Instruction{2, rlce}
	// RLC H
	cbInst[0x04] = Instruction{2, rlch}
	// RLC L
	cbInst[0x05] = Instruction{2, rlcl}
	// RLC (HL)
	cbInst[0x06] = Instruction{2, rlchl}
	// RLC A
	cbInst[0x07] = Instruction{2, cbrlca}
	// RRC B
	cbInst[0x08] = Instruction{2, rrcb}
	// RRC C
	cbInst[0x09] = Instruction{2, rrcc}
	// RRC D
	cbInst[0x0A] = Instruction{2, rrcd}
	// RRC E
	cbInst[0x0B] = Instruction{2, rrce}
	// RRC H
	cbInst[0x0C] = Instruction{2, rrch}
	// RRC L
	cbInst[0x0D] = Instruction{2, rrcl}
	// RRC (HL)
	cbInst[0x0E] = Instruction{2, rrchl}
	// RRC A
	cbInst[0x0F] = Instruction{2, cbrrca}

	// RL B
	cbInst[0x10] = Instruction{2, rlb}
	// RL C
	cbInst[0x11] = Instruction{2, rlc}
	// RL D
	cbInst[0x12] = Instruction{2, rld}
	// RL E
	cbInst[0x13] = Instruction{2, rle}
	// RL H
	cbInst[0x14] = Instruction{2, rlh}
	// RL L
	cbInst[0x15] = Instruction{2, rll}
	// RL (HL)
	cbInst[0x16] = Instruction{2, rlhl}
	// RL A
	cbInst[0x17] = Instruction{2, cbrla}
	// RR B
	cbInst[0x18] = Instruction{2, rrb}
	// RR C
	cbInst[0x19] = Instruction{2, rrc}
	// RR D
	cbInst[0x1A] = Instruction{2, rrd}
	// RR E
	cbInst[0x1B] = Instruction{2, rre}
	// RR H
	cbInst[0x1C] = Instruction{2, rrh}
	// RR L
	cbInst[0x1D] = Instruction{2, rrl}
	// RR (HL)
	cbInst[0x1E] = Instruction{2, rrhl}
	// RR A
	cbInst[0x1F] = Instruction{2, cbrra}

	// SLA B
	cbInst[0x20] = Instruction{2, slab}
	// SLA C
	cbInst[0x21] = Instruction{2, slac}
	// SLA D
	cbInst[0x22] = Instruction{2, slad}
	// SLA E
	cbInst[0x23] = Instruction{2, slae}
	// SLA H
	cbInst[0x24] = Instruction{2, slah}
	// SLA L
	cbInst[0x25] = Instruction{2, slal}
	// SLA (HL)
	cbInst[0x26] = Instruction{2, slahl}
	// SLA A
	cbInst[0x27] = Instruction{2, slaa}
	// SRA B
	cbInst[0x28] = Instruction{2, srab}
	// SRA C
	cbInst[0x29] = Instruction{2, srac}
	// SRA D
	cbInst[0x2A] = Instruction{2, srad}
	// SRA E
	cbInst[0x2B] = Instruction{2, srae}
	// SRA H
	cbInst[0x2C] = Instruction{2, srah}
	// SRA L
	cbInst[0x2D] = Instruction{2, sral}
	// SRA (HL)
	cbInst[0x2E] = Instruction{2, srahl}
	// SRA A
	cbInst[0x2F] = Instruction{2, sraa}

	// SWAP B
	cbInst[0x30] = Instruction{2, swapb}
	// SWAP C
	cbInst[0x31] = Instruction{2, swapc}
	// SWAP D
	cbInst[0x32] = Instruction{2, swapd}
	// SWAP E
	cbInst[0x33] = Instruction{2, swape}
	// SWAP H
	cbInst[0x34] = Instruction{2, swaph}
	// SWAP L
	cbInst[0x35] = Instruction{2, swapl}
	// SWAP (HL)
	cbInst[0x36] = Instruction{2, swaphl}
	// SWAP A
	cbInst[0x37] = Instruction{2, swapa}
	// SRL B
	cbInst[0x28] = Instruction{2, srlb}
	// SRL C
	cbInst[0x29] = Instruction{2, srlc}
	// SRL D
	cbInst[0x2A] = Instruction{2, srld}
	// SRL E
	cbInst[0x2B] = Instruction{2, srle}
	// SRL H
	cbInst[0x2C] = Instruction{2, srlh}
	// SRL L
	cbInst[0x2D] = Instruction{2, srll}
	// SRL (HL)
	cbInst[0x2E] = Instruction{2, srlhl}
	// SRL A
	cbInst[0x2F] = Instruction{2, srla}

	// SWAP B
	cbInst[0x30] = Instruction{2, swapb}
	// SWAP C
	cbInst[0x31] = Instruction{2, swapc}
	// SWAP D
	cbInst[0x32] = Instruction{2, swapd}
	// SWAP E
	cbInst[0x33] = Instruction{2, swape}
	// SWAP H
	cbInst[0x34] = Instruction{2, swaph}
	// SWAP L
	cbInst[0x35] = Instruction{2, swapl}
	// SWAP (HL)
	cbInst[0x36] = Instruction{2, swaphl}
	// SWAP A
	cbInst[0x37] = Instruction{2, swapa}
	// SRL B
	cbInst[0x38] = Instruction{2, srlb}
	// SRL C
	cbInst[0x39] = Instruction{2, srlc}
	// SRL D
	cbInst[0x3A] = Instruction{2, srld}
	// SRL E
	cbInst[0x3B] = Instruction{2, srle}
	// SRL H
	cbInst[0x3C] = Instruction{2, srlh}
	// SRL L
	cbInst[0x3D] = Instruction{2, srll}
	// SRL (HL)
	cbInst[0x3E] = Instruction{2, srlhl}
	// SRL A
	cbInst[0x3F] = Instruction{2, srla}

	// BIT 0, B
	cbInst[0x40] = Instruction{2, bit0b}
	// BIT 0, C
	cbInst[0x41] = Instruction{2, bit0c}
	// BIT 0, D
	cbInst[0x42] = Instruction{2, bit0d}
	// BIT 0, E
	cbInst[0x43] = Instruction{2, bit0e}
	// BIT 0, H
	cbInst[0x44] = Instruction{2, bit0h}
	// BIT 0, L
	cbInst[0x45] = Instruction{2, bit0l}
	// BIT 0, (HL)
	cbInst[0x46] = Instruction{2, bit0hl}
	// BIT 0, A
	cbInst[0x47] = Instruction{2, bit0a}
	// BIT 1, B
	cbInst[0x48] = Instruction{2, bit1b}
	// BIT 1, C
	cbInst[0x49] = Instruction{2, bit1c}
	// BIT 1, D
	cbInst[0x4A] = Instruction{2, bit1d}
	// BIT 1, E
	cbInst[0x4B] = Instruction{2, bit1e}
	// BIT 1, H
	cbInst[0x4C] = Instruction{2, bit1h}
	// BIT 1, L
	cbInst[0x4D] = Instruction{2, bit1l}
	// BIT 1, (HL)
	cbInst[0x4E] = Instruction{2, bit1hl}
	// BIT 1, A
	cbInst[0x4F] = Instruction{2, bit1a}

	// BIT 2, B
	cbInst[0x50] = Instruction{2, bit2b}
	// BIT 2, C
	cbInst[0x51] = Instruction{2, bit2c}
	// BIT 2, D
	cbInst[0x52] = Instruction{2, bit2d}
	// BIT 2, E
	cbInst[0x53] = Instruction{2, bit2e}
	// BIT 2, H
	cbInst[0x54] = Instruction{2, bit2h}
	// BIT 2, L
	cbInst[0x55] = Instruction{2, bit2l}
	// BIT 2, (HL)
	cbInst[0x56] = Instruction{2, bit2hl}
	// BIT 2, A
	cbInst[0x57] = Instruction{2, bit2a}
	// BIT 3, B
	cbInst[0x58] = Instruction{2, bit3b}
	// BIT 3, C
	cbInst[0x59] = Instruction{2, bit3c}
	// BIT 3, D
	cbInst[0x5A] = Instruction{2, bit3d}
	// BIT 3, E
	cbInst[0x5B] = Instruction{2, bit3e}
	// BIT 3, H
	cbInst[0x5C] = Instruction{2, bit3h}
	// BIT 3, L
	cbInst[0x5D] = Instruction{2, bit3l}
	// BIT 3, (HL)
	cbInst[0x5E] = Instruction{2, bit3hl}
	// BIT 3, A
	cbInst[0x5F] = Instruction{2, bit3a}

	// BIT 4, B
	cbInst[0x60] = Instruction{2, bit4b}
	// BIT 4, C
	cbInst[0x61] = Instruction{2, bit4c}
	// BIT 4, D
	cbInst[0x62] = Instruction{2, bit4d}
	// BIT 4, E
	cbInst[0x63] = Instruction{2, bit4e}
	// BIT 4, H
	cbInst[0x64] = Instruction{2, bit4h}
	// BIT 4, L
	cbInst[0x65] = Instruction{2, bit4l}
	// BIT 4, (HL)
	cbInst[0x66] = Instruction{2, bit4hl}
	// BIT 4, A
	cbInst[0x67] = Instruction{2, bit4a}
	// BIT 5, B
	cbInst[0x68] = Instruction{2, bit5b}
	// BIT 5, C
	cbInst[0x69] = Instruction{2, bit5c}
	// BIT 5, D
	cbInst[0x6A] = Instruction{2, bit5d}
	// BIT 5, E
	cbInst[0x6B] = Instruction{2, bit5e}
	// BIT 5, H
	cbInst[0x6C] = Instruction{2, bit5h}
	// BIT 5, L
	cbInst[0x6D] = Instruction{2, bit5l}
	// BIT 5, (HL)
	cbInst[0x6E] = Instruction{2, bit5hl}
	// BIT 5, A
	cbInst[0x6F] = Instruction{2, bit5a}

	// BIT 6, B
	cbInst[0x70] = Instruction{2, bit6b}
	// BIT 6, C
	cbInst[0x71] = Instruction{2, bit6c}
	// BIT 6, D
	cbInst[0x72] = Instruction{2, bit6d}
	// BIT 6, E
	cbInst[0x73] = Instruction{2, bit6e}
	// BIT 6, H
	cbInst[0x74] = Instruction{2, bit6h}
	// BIT 6, L
	cbInst[0x75] = Instruction{2, bit6l}
	// BIT 6, (HL)
	cbInst[0x76] = Instruction{2, bit6hl}
	// BIT 6, A
	cbInst[0x77] = Instruction{2, bit6a}
	// BIT 7, B
	cbInst[0x78] = Instruction{2, bit7b}
	// BIT 7, C
	cbInst[0x79] = Instruction{2, bit7c}
	// BIT 7, D
	cbInst[0x7A] = Instruction{2, bit7d}
	// BIT 7, E
	cbInst[0x7B] = Instruction{2, bit7e}
	// BIT 7, H
	cbInst[0x7C] = Instruction{2, bit7h}
	// BIT 7, L
	cbInst[0x7D] = Instruction{2, bit7l}
	// BIT 7, (HL)
	cbInst[0x7E] = Instruction{2, bit7hl}
	// BIT 7, A
	cbInst[0x7F] = Instruction{2, bit7a}

	// RES 0, B
	cbInst[0x80] = Instruction{2, res0b}
	// RES 0, C
	cbInst[0x81] = Instruction{2, res0c}
	// RES 0, D
	cbInst[0x82] = Instruction{2, res0d}
	// RES 0, E
	cbInst[0x83] = Instruction{2, res0e}
	// RES 0, H
	cbInst[0x84] = Instruction{2, res0h}
	// RES 0, L
	cbInst[0x85] = Instruction{2, res0l}
	// RES 0, (HL)
	cbInst[0x86] = Instruction{2, res0hl}
	// RES 0, A
	cbInst[0x87] = Instruction{2, res0a}
	// RES 1, B
	cbInst[0x88] = Instruction{2, res1b}
	// RES 1, C
	cbInst[0x89] = Instruction{2, res1c}
	// RES 1, D
	cbInst[0x8A] = Instruction{2, res1d}
	// RES 1, E
	cbInst[0x8B] = Instruction{2, res1e}
	// RES 1, H
	cbInst[0x8C] = Instruction{2, res1h}
	// RES 1, L
	cbInst[0x8D] = Instruction{2, res1l}
	// RES 1, (HL)
	cbInst[0x8E] = Instruction{2, res1hl}
	// RES 1, A
	cbInst[0x8F] = Instruction{2, res1a}

	// RES 2, B
	cbInst[0x90] = Instruction{2, res2b}
	// RES 2, C
	cbInst[0x91] = Instruction{2, res2c}
	// RES 2, D
	cbInst[0x92] = Instruction{2, res2d}
	// RES 2, E
	cbInst[0x93] = Instruction{2, res2e}
	// RES 2, H
	cbInst[0x94] = Instruction{2, res2h}
	// RES 2, L
	cbInst[0x95] = Instruction{2, res2l}
	// RES 2, (HL)
	cbInst[0x96] = Instruction{2, res2hl}
	// RES 2, A
	cbInst[0x97] = Instruction{2, res2a}
	// RES 3, B
	cbInst[0x98] = Instruction{2, res3b}
	// RES 3, C
	cbInst[0x99] = Instruction{2, res3c}
	// RES 3, D
	cbInst[0x9A] = Instruction{2, res3d}
	// RES 3, E
	cbInst[0x9B] = Instruction{2, res3e}
	// RES 3, H
	cbInst[0x9C] = Instruction{2, res3h}
	// RES 3, L
	cbInst[0x9D] = Instruction{2, res3l}
	// RES 3, (HL)
	cbInst[0x9E] = Instruction{2, res3hl}
	// RES 3, A
	cbInst[0x9F] = Instruction{2, res3a}

	// RES 4, B
	cbInst[0xA0] = Instruction{2, res4b}
	// RES 4, C
	cbInst[0xA1] = Instruction{2, res4c}
	// RES 4, D
	cbInst[0xA2] = Instruction{2, res4d}
	// RES 4, E
	cbInst[0xA3] = Instruction{2, res4e}
	// RES 4, H
	cbInst[0xA4] = Instruction{2, res4h}
	// RES 4, L
	cbInst[0xA5] = Instruction{2, res4l}
	// RES 4, (HL)
	cbInst[0xA6] = Instruction{2, res4hl}
	// RES 4, A
	cbInst[0xA7] = Instruction{2, res4a}
	// RES 5, B
	cbInst[0xA8] = Instruction{2, res5b}
	// RES 5, C
	cbInst[0xA9] = Instruction{2, res5c}
	// RES 5, D
	cbInst[0xAA] = Instruction{2, res5d}
	// RES 5, E
	cbInst[0xAB] = Instruction{2, res5e}
	// RES 5, H
	cbInst[0xAC] = Instruction{2, res5h}
	// RES 5, L
	cbInst[0xAD] = Instruction{2, res5l}
	// RES 5, (HL)
	cbInst[0xAE] = Instruction{2, res5hl}
	// RES 5, A
	cbInst[0xAF] = Instruction{2, res5a}

	// RES 6, B
	cbInst[0xB0] = Instruction{2, res6b}
	// RES 6, C
	cbInst[0xB1] = Instruction{2, res6c}
	// RES 6, D
	cbInst[0xB2] = Instruction{2, res6d}
	// RES 6, E
	cbInst[0xB3] = Instruction{2, res6e}
	// RES 6, H
	cbInst[0xB4] = Instruction{2, res6h}
	// RES 6, L
	cbInst[0xB5] = Instruction{2, res6l}
	// RES 6, (HL)
	cbInst[0xB6] = Instruction{2, res6hl}
	// RES 6, A
	cbInst[0xB7] = Instruction{2, res6a}
	// RES 7, B
	cbInst[0xB8] = Instruction{2, res7b}
	// RES 7, C
	cbInst[0xB9] = Instruction{2, res7c}
	// RES 7, D
	cbInst[0xBA] = Instruction{2, res7d}
	// RES 7, E
	cbInst[0xBB] = Instruction{2, res7e}
	// RES 7, H
	cbInst[0xBC] = Instruction{2, res7h}
	// RES 7, L
	cbInst[0xBD] = Instruction{2, res7l}
	// RES 7, (HL)
	cbInst[0xBE] = Instruction{2, res7hl}
	// RES 7, A
	cbInst[0xBF] = Instruction{2, res7a}

	// SET 0, B
	cbInst[0xC0] = Instruction{2, set0b}
	// SET 0, C
	cbInst[0xC1] = Instruction{2, set0c}
	// SET 0, D
	cbInst[0xC2] = Instruction{2, set0d}
	// SET 0, E
	cbInst[0xC3] = Instruction{2, set0e}
	// SET 0, H
	cbInst[0xC4] = Instruction{2, set0h}
	// SET 0, L
	cbInst[0xC5] = Instruction{2, set0l}
	// SET 0, (HL)
	cbInst[0xC6] = Instruction{2, set0hl}
	// SET 0, A
	cbInst[0xC7] = Instruction{2, set0a}
	// SET 1, B
	cbInst[0xC8] = Instruction{2, set1b}
	// SET 1, C
	cbInst[0xC9] = Instruction{2, set1c}
	// SET 1, D
	cbInst[0xCA] = Instruction{2, set1d}
	// SET 1, E
	cbInst[0xCB] = Instruction{2, set1e}
	// SET 1, H
	cbInst[0xCC] = Instruction{2, set1h}
	// SET 1, L
	cbInst[0xCD] = Instruction{2, set1l}
	// SET 1, (HL)
	cbInst[0xCE] = Instruction{2, set1hl}
	// SET 1, A
	cbInst[0xCF] = Instruction{2, set1a}

	// SET 2, B
	cbInst[0xD0] = Instruction{2, set2b}
	// SET 2, C
	cbInst[0xD1] = Instruction{2, set2c}
	// SET 2, D
	cbInst[0xD2] = Instruction{2, set2d}
	// SET 2, E
	cbInst[0xD3] = Instruction{2, set2e}
	// SET 2, H
	cbInst[0xD4] = Instruction{2, set2h}
	// SET 2, L
	cbInst[0xD5] = Instruction{2, set2l}
	// SET 2, (HL)
	cbInst[0xD6] = Instruction{2, set2hl}
	// SET 2, A
	cbInst[0xD7] = Instruction{2, set2a}
	// SET 3, B
	cbInst[0xD8] = Instruction{2, set3b}
	// SET 3, C
	cbInst[0xD9] = Instruction{2, set3c}
	// SET 3, D
	cbInst[0xDA] = Instruction{2, set3d}
	// SET 3, E
	cbInst[0xDB] = Instruction{2, set3e}
	// SET 3, H
	cbInst[0xDC] = Instruction{2, set3h}
	// SET 3, L
	cbInst[0xDD] = Instruction{2, set3l}
	// SET 3, (HL)
	cbInst[0xDE] = Instruction{2, set3hl}
	// SET 3, A
	cbInst[0xDF] = Instruction{2, set3a}

	// SET 4, B
	cbInst[0xE0] = Instruction{2, set4b}
	// SET 4, C
	cbInst[0xE1] = Instruction{2, set4c}
	// SET 4, D
	cbInst[0xE2] = Instruction{2, set4d}
	// SET 4, E
	cbInst[0xE3] = Instruction{2, set4e}
	// SET 4, H
	cbInst[0xE4] = Instruction{2, set4h}
	// SET 4, L
	cbInst[0xE5] = Instruction{2, set4l}
	// SET 4, (HL)
	cbInst[0xE6] = Instruction{2, set4hl}
	// SET 4, A
	cbInst[0xE7] = Instruction{2, set4a}
	// SET 5, B
	cbInst[0xE8] = Instruction{2, set5b}
	// SET 5, C
	cbInst[0xE9] = Instruction{2, set5c}
	// SET 5, D
	cbInst[0xEA] = Instruction{2, set5d}
	// SET 5, E
	cbInst[0xEB] = Instruction{2, set5e}
	// SET 5, H
	cbInst[0xEC] = Instruction{2, set5h}
	// SET 5, L
	cbInst[0xED] = Instruction{2, set5l}
	// SET 5, (HL)
	cbInst[0xEE] = Instruction{2, set5hl}
	// SET 5, A
	cbInst[0xEF] = Instruction{2, set5a}

	// SET 6, B
	cbInst[0xB0] = Instruction{2, set6b}
	// SET 6, C
	cbInst[0xB1] = Instruction{2, set6c}
	// SET 6, D
	cbInst[0xB2] = Instruction{2, set6d}
	// SET 6, E
	cbInst[0xB3] = Instruction{2, set6e}
	// SET 6, H
	cbInst[0xB4] = Instruction{2, set6h}
	// SET 6, L
	cbInst[0xB5] = Instruction{2, set6l}
	// SET 6, (HL)
	cbInst[0xB6] = Instruction{2, set6hl}
	// SET 6, A
	cbInst[0xB7] = Instruction{2, set6a}
	// SET 7, B
	cbInst[0xB8] = Instruction{2, set7b}
	// SET 7, C
	cbInst[0xB9] = Instruction{2, set7c}
	// SET 7, D
	cbInst[0xBA] = Instruction{2, set7d}
	// SET 7, E
	cbInst[0xBB] = Instruction{2, set7e}
	// SET 7, H
	cbInst[0xBC] = Instruction{2, set7h}
	// SET 7, L
	cbInst[0xBD] = Instruction{2, set7l}
	// SET 7, (HL)
	cbInst[0xBE] = Instruction{2, set7hl}
	// SET 7, A
	cbInst[0xBF] = Instruction{2, set7a}

	// endregion CB Prefixed Instructions
}

//region Instruction Functions

func nop() string {
	return "NOP"
}

func ldbc16() string {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	Reg.BC.Set(value)

	fmt.Println()

	return fmt.Sprintf("LD BC, $%.4X", value)
}

func ldbca() string {
	pos := Reg.BC.Get()
	bus.Write(pos, Reg.A.Get())

	return "LD (BC), A"
}

func incbc() string {
	value := Reg.BC.Get()
	value++
	Reg.BC.Set(value)

	return "INC BC"
}

func incb() string {
	incReg(Reg.B.Val())

	return "INC B"
}

func decb() string {
	decReg(Reg.B.Val())

	return "DEC B"
}

func ldb() string {
	Reg.PC.Inc()
	*Reg.B.Val() = bus.Read(Reg.PC.Get())

	return fmt.Sprintf("LD B, $%X", Reg.B)
}

// Rotate Register A left
// Bit 7 shifts to bit 0
// Bit 7 affect the carry Flag
// C <- [7~0] <- [7]
func rlca() string {
	var bit7 bool = Reg.A.Get()&0x80 == 0x80
	// If bit 7 is 1
	*Reg.A.Val() <<= 1
	if bit7 {
		*Reg.A.Val() |= 1
	}
	flags.SetFlagZ(false)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit7)

	return "RLCA"
}

func ldmemsp() string {
	pos := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	bus.Write16(pos, Reg.SP.Get())

	return fmt.Sprintf("LD ($%X), SP", pos)
}

func addhlbc() string {
	addhlReg(Reg.B.Get(), Reg.C.Get())

	return "ADD HL, BC"
}

func ldabc() string {
	pos := Reg.BC.Get()
	*Reg.A.Val() = bus.Read(pos)

	return "LD A, (BC)"
}

func decbc() string {
	value := Reg.BC.Get() - 1
	Reg.BC.Set(value)

	return "DEC BC"
}

func incc() string {
	incReg(Reg.C.Val())

	return "INC C"
}

func decc() string {
	decReg(Reg.C.Val())

	return "DEC C"
}

func ldc() string {
	*Reg.C.Val() = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	return fmt.Sprintf("LD C, $%X", Reg.C.Get())
}

// Rotate Register A right
// Bit 0 shifts to Carry
// [0] -> [7~0] -> C
func rrca() string {
	var bit0 bool = Reg.A.Get()&0x1 == 0x1
	*Reg.A.Val() >>= 1
	if bit0 {
		*Reg.A.Val() |= 0x80
	}
	flags.SetFlagZ(false)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit0)

	return "RRCA"
}

// Enters CPU low power mode.
// In GBC, switches between normal and double CPU speed
func stop() string {
	// TODO implement cpu speed switch

	// After the stop functions comes an operand that is ignored by the cpu
	Reg.PC.Inc()

	return "STOP"
}

func ldde16() string {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	Reg.DE.Set(value)

	return fmt.Sprintf("LD DE, %X", value)
}

func lddea() string {
	pos := Reg.DE.Get()
	bus.Write(pos, Reg.A.Get())

	return "LD (DE), A"
}

func incde() string {
	value := Reg.DE.Get()
	value++
	Reg.DE.Set(value)

	return "INC DE"
}

func incd() string {
	incReg(Reg.D.Val())

	return "INC D"
}

func decd() string {
	decReg(Reg.D.Val())

	return "DEC D"
}

func ldd() string {
	Reg.PC.Inc()
	*Reg.D.Val() = bus.Read(Reg.PC.Get())

	return fmt.Sprintf("LD D, $%X", Reg.D)
}

// Rotate Register A left through Carry
// Previous Carry shifts to bit 0
// Bit 7 shift to Carry
// C <- [7~0] <- C
func rla() string {
	var bit7 bool = Reg.A.Get()&0x80 == 0x80
	*Reg.A.Val() <<= 1
	// If carry flag is 1
	if flags.GetFlagC() {
		*Reg.A.Val() |= 1
	}
	flags.SetFlagZ(false)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit7)

	return "RLA"
}

func jr() string {
	value := jrCond(true, 0)

	return fmt.Sprintf("JR $%X", value)
}

func addhlde() string {
	addhlReg(Reg.D.Get(), Reg.E.Get())

	return "ADD HL, DE"
}

func ldade() string {
	pos := Reg.DE.Get()
	*Reg.A.Val() = bus.Read(pos)

	return "LD A, (DE)"
}

func decde() string {
	value := Reg.DE.Get() - 1
	Reg.DE.Set(value)

	return "DEC DE"
}

func ince() string {
	incReg(Reg.E.Val())

	return "INC E"
}

func dece() string {
	decReg(Reg.E.Val())

	return "DEC E"
}

func lde() string {
	Reg.PC.Inc()
	*Reg.E.Val() = bus.Read(Reg.PC.Get())

	return fmt.Sprintf("LD E, $%X", Reg.E.Get())
}

// Rotate Register A right through Carry
// Previous Carry value shifts to bit 7
// Bit 0 shifts to Carry
// C -> [7~0] -> C
func rra() string {
	var bit0 bool = Reg.A.Get()&0x1 == 0x1
	*Reg.A.Val() >>= 1
	// If carry flag is 1
	if flags.GetFlagC() {
		*Reg.A.Val() |= 0x80
	}
	flags.SetFlagZ(false)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit0)

	return "RRA"
}

func jrnz() string {
	value := jrCond(!flags.GetFlagZ(), 1)

	return fmt.Sprintf("JR NZ, $%X", value)
}

func ldhl16() string {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	Reg.HL.Set(value)

	return fmt.Sprintf("LD HL, %X", value)
}

func ldhli() string {
	pos := Reg.HL.Get()
	bus.Write(pos, Reg.A.Get())
	Reg.HL.Set(pos + 1)

	return "LD (HLI), A"
}

func inchl() string {
	value := Reg.HL.Get()
	value++
	Reg.HL.Set(value)

	return "INC HL"
}

func inch() string {
	incReg(Reg.H.Val())

	return "INC H"
}

func dech() string {
	decReg(Reg.H.Val())

	return "DEC H"
}

func ldh() string {
	Reg.PC.Inc()
	*Reg.H.Val() = bus.Read(Reg.PC.Get())

	return fmt.Sprintf("LD H, $%X", Reg.H.Get())
}

func daa() string {
	// Decimal Adjust the Accumulator to be BCD correct.
	// The process is as follows:
	// 1. Check four Least Significant Bits (LSB)
	// 2. LSB > 9 ||H Flag is Set to One -> Add $06 (or Subtract if N Flag is Set to One)
	// 3. Check four Most Significant Bits (MSB)
	// 4. MSB > 9 || C Flag is Set to One -> Add $60

	// TODO test implementation is correct
	// Use following links as guide
	// http://z80-heaven.wikidot.com/instructions-set:daa
	// http://www.z80.info/z80syntx.htm#DAA
	// https://ehaskins.com/2018-01-30%20Z80%20DAA/
	lsb := Reg.A.Get() & 0x0F
	msb := Reg.A.Get() >> 4

	// TODO Optimise code for better performance
	if lsb > 9 || flags.GetFlagH() {
		if !flags.GetFlagN() {
			*Reg.A.Val() += 0x06
		} else {
			*Reg.A.Val() -= 0x06
		}
	}

	if msb > 9 || flags.GetFlagC() {
		if !flags.GetFlagN() {
			*Reg.A.Val() += 0x60
		} else {

		}

	}

	flags.SetFlagZ(Reg.A.Get() == 0)
	// Carry is set to one When BCD value is over $99, according to definition of DAA
	flags.SetFlagC(Reg.A.Get() > 0x99)

	return "DAA"
}

func jrz() string {
	value := jrCond(flags.GetFlagZ(), 1)

	return fmt.Sprintf("JR Z, $%X", value)
}

func addhlhl() string {
	addhlReg(Reg.H.Get(), Reg.L.Get())

	return "ADD HL, HL"
}

func ldahli() string {
	pos := Reg.HL.Get()
	*Reg.A.Val() = bus.Read(pos)
	Reg.HL.Set(pos + 1)

	return "LD A, (HLI)"
}

func dechl() string {
	value := Reg.HL.Get() - 1
	Reg.HL.Set(value)

	return "DEC HL"
}

func incl() string {
	incReg(Reg.L.Val())

	return "INC L"
}

func decl() string {
	decReg(Reg.L.Val())

	return "DEC L"
}

func ldl() string {
	Reg.PC.Inc()
	*Reg.L.Val() = bus.Read(Reg.PC.Get())

	return fmt.Sprintf("LD L, $%X", Reg.L.Get())
}

func cpl() string {
	*Reg.A.Val() = ^*Reg.A.Val()

	return "CPL"
}

func jrnc() string {
	value := jrCond(!flags.GetFlagC(), 1)

	return fmt.Sprintf("JR NC, $%X", value)
}

func ldsp() string {
	Reg.SP.Set(bus.Read16(Reg.PC.Get()))
	Reg.PC.Inc().Inc()

	return fmt.Sprintf("LD SP, $%X", Reg.SP.Get())
}

func ldhlda() string {
	pos := Reg.HL.Get()
	bus.Write(pos, Reg.A.Get())
	Reg.HL.Set(pos - 1)

	return "LD (HLD), A"
}

func incsp() string {
	Reg.SP.Inc()

	return "INC SP"
}

func inchlind() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	flags.AffectFlagZH(value, value+1)
	flags.SetFlagN(false)
	bus.Write(pos, value+1)

	return "INC (HL)"
}

func dechlind() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	flags.AffectFlagZH(value, value+1)
	flags.SetFlagN(true)
	bus.Write(pos, value+1)

	return "DEC H"
}

func ldhl8() string {
	Reg.PC.Inc()
	value := bus.Read(Reg.PC.Get())
	ldhlind(value)

	return fmt.Sprintf("LD (HL), $%X", Reg.H.Get())
}

// Set Carry Flag
// Flags N and H are set to Zero
func scf() string {
	flags.SetFlagC(true)
	flags.SetFlagN(false)
	flags.SetFlagH(false)

	return "SCF"
}

func jrc() string {
	value := jrCond(flags.GetFlagC(), 1)

	return fmt.Sprintf("JR C, $%X", value)
}

func addhlsp() string {
	addhlReg16(Reg.SP.Get())

	return "ADD HL, SP"
}

func ldahld() string {
	pos := Reg.HL.Get()
	*Reg.A.Val() = bus.Read(pos)
	Reg.HL.Set(pos - 1)

	return "LD A, (HLD)"
}

func decsp() string {
	Reg.SP.Dec()

	return "DEC SP"
}

func inca() string {
	incReg(Reg.A.Val())

	return "INC A"
}

func deca() string {
	decReg(Reg.A.Val())

	return "DEC A"
}

func lda() string {
	*Reg.A.Val() = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	return fmt.Sprintf("LD A, $%.2X", Reg.A.Get())
}

// Complement Carry Flag
func ccf() string {
	flags.SetFlagC(!flags.GetFlagC())

	return "CCF"
}

func ldbb() string {

	return "LD B, B"
}

func ldbc() string {
	*Reg.B.Val() = Reg.C.Get()

	return "LD B, C"
}

func ldbd() string {
	*Reg.B.Val() = Reg.D.Get()

	return "LD B, D"
}

func ldbe() string {
	*Reg.B.Val() = Reg.E.Get()

	return "LD B, E"
}

func ldbh() string {
	*Reg.B.Val() = Reg.H.Get()

	return "LD B, H"
}

func ldbl() string {
	*Reg.B.Val() = Reg.L.Get()

	return "LD B, L"
}

func ldbhl() string {
	*Reg.B.Val() = bus.Read(Reg.HL.Get())

	return "LD B, (HL)"
}

func ldba() string {
	*Reg.B.Val() = Reg.A.Get()

	return "LD B, A"
}

func ldcb() string {
	*Reg.C.Val() = Reg.B.Get()

	return "LD C, B"
}

func ldcc() string {

	return "LD C, C"
}

func ldcd() string {
	*Reg.C.Val() = Reg.D.Get()

	return "LD C, D"
}

func ldce() string {
	*Reg.C.Val() = Reg.E.Get()

	return "LD C, E"
}

func ldch() string {
	*Reg.C.Val() = Reg.H.Get()

	return "LD C, H"
}

func ldcl() string {
	*Reg.C.Val() = Reg.L.Get()

	return "LD C, L"
}

func ldchl() string {
	*Reg.C.Val() = bus.Read(Reg.HL.Get())

	return "LD C, (HL)"
}

func ldca() string {
	*Reg.C.Val() = Reg.A.Get()

	return "LD C, A"
}

func lddb() string {
	*Reg.D.Val() = Reg.B.Get()

	return "LD D, B"
}

func lddc() string {
	*Reg.D.Val() = Reg.C.Get()

	return "LD D, C"
}

func lddd() string {

	return "LD D, D"
}

func ldde() string {
	*Reg.D.Val() = Reg.E.Get()

	return "LD D, E"
}

func lddh() string {
	*Reg.D.Val() = Reg.H.Get()

	return "LD D, H"
}

func lddl() string {
	*Reg.D.Val() = Reg.L.Get()

	return "LD D, L"
}

func lddhl() string {
	*Reg.D.Val() = bus.Read(Reg.HL.Get())

	return "LD D, (HL)"
}

func ldda() string {
	*Reg.D.Val() = Reg.A.Get()

	return "LD D, A"
}

func ldeb() string {
	*Reg.E.Val() = Reg.B.Get()

	return "LD E, B"
}

func ldec() string {
	*Reg.E.Val() = Reg.C.Get()

	return "LD E, C"
}

func lded() string {
	*Reg.E.Val() = Reg.D.Get()

	return "LD E, D"
}

func ldee() string {

	return "LD E, E"
}

func ldeh() string {
	*Reg.E.Val() = Reg.H.Get()

	return "LD E, H"
}

func ldel() string {
	*Reg.E.Val() = Reg.L.Get()

	return "LD E, L"
}

func ldehl() string {
	*Reg.E.Val() = bus.Read(Reg.HL.Get())

	return "LD E, (HL)"
}

func ldea() string {
	*Reg.E.Val() = Reg.A.Get()

	return "LD E, A"
}

func ldhb() string {
	*Reg.H.Val() = Reg.B.Get()

	return "LD H, B"
}

func ldhc() string {
	*Reg.H.Val() = Reg.C.Get()

	return "LD H, C"
}

func ldhd() string {
	*Reg.H.Val() = Reg.D.Get()

	return "LD H, D"
}

func ldhe() string {
	*Reg.H.Val() = Reg.E.Get()

	return "LD H, E"
}

func ldhh() string {

	return "LD H, H"
}

func ldhl() string {
	*Reg.H.Val() = Reg.L.Get()

	return "LD H, L"
}

func ldhhl() string {
	*Reg.H.Val() = bus.Read(Reg.HL.Get())

	return "LD H, (HL)"
}

func ldha() string {
	*Reg.H.Val() = Reg.A.Get()

	return "LD H, A"
}

func ldlb() string {
	*Reg.L.Val() = Reg.B.Get()

	return "LD L, B"
}

func ldlc() string {
	*Reg.L.Val() = Reg.C.Get()

	return "LD L, C"
}

func ldld() string {
	*Reg.L.Val() = Reg.D.Get()

	return "LD L, D"
}

func ldle() string {
	*Reg.L.Val() = Reg.E.Get()

	return "LD L, E"
}

func ldlh() string {
	*Reg.L.Val() = Reg.H.Get()

	return "LD L, H"
}

func ldll() string {

	return "LD L, L"
}

func ldlhl() string {
	*Reg.L.Val() = bus.Read(Reg.HL.Get())

	return "LD L, (HL)"
}

func ldla() string {
	*Reg.H.Val() = Reg.A.Get()

	return "LD L, A"
}

func ldhlb() string {
	ldhlind(Reg.B.Get())

	return "LD (HL), B"
}

func ldhlc() string {
	ldhlind(Reg.C.Get())

	return "LD (HL), C"
}

func ldhld() string {
	ldhlind(Reg.D.Get())

	return "LD (HL), D"
}

func ldhle() string {
	ldhlind(Reg.E.Get())

	return "LD (HL), E"
}

func ldhlh() string {
	ldhlind(Reg.H.Get())

	return "LD (HL), H"
}

func ldhll() string {
	ldhlind(Reg.L.Get())

	return "LD (HL), L"
}

// Halt pauses CPU execution until an interruption take place
func halt() string {
	/*
		Halt stops the CPU execution, and resumes when an interrupt is pending. An interruption is considered
		pending when an interrupt is enabled and its flag is set to one, that is IE && IF !=0 for a certain interrupt.
		The following assumptions take place:
		With IME = 1:
		1. With pending interrupt, cpu will not halt.
		2. The expected behaviour would be the CPU jumping to next instruction
		2. Interrupt handling takes place

		With IME = 0:
		1. If no interrupt pending, halt will execute and cpu will pause until an interrupt becomes pending.
		Interrupt will not be handled as expected with the master interrupt not enabled
		2. If an interrupt is pending,halt immediately exits. Halt bug might take place as explained below

		HALT Bug:
		Take place as IME = 0 with an interrupt is pending. Two of the following scenarios can take place
		1. With no IE instruction before HALT, the byte after halt instruction is read twice
		2. With IE instruction before HALT (with IME delay affect taking place), the interrupt handler takes place.
		The handler, however, returns to halt after serviced, causing the cpu to pause again.
	*/

	// TODO Give the instructions above an index, then refer those index where simulated in cpu and instruction
	// file

	isHalt = true
	return "HALT"
}

func ldhla() string {
	ldhlind(Reg.A.Get())

	return "LD (HL), A"
}

func ldab() string {
	*Reg.A.Val() = Reg.B.Get()

	return "LD A, B"
}

func ldac() string {
	*Reg.A.Val() = Reg.C.Get()

	return "LD A, C"
}

func ldad() string {
	*Reg.A.Val() = Reg.D.Get()

	return "LD A, D"
}

func ldae() string {
	*Reg.A.Val() = Reg.E.Get()

	return "LD A, E"
}

func ldah() string {
	*Reg.A.Val() = Reg.H.Get()

	return "LD A, H"
}

func ldal() string {
	*Reg.A.Val() = Reg.L.Get()

	return "LD A, L"
}

func ldahl() string {
	*Reg.A.Val() = bus.Read(Reg.HL.Get())

	return "LD A, (HL)"
}

func ldaa() string {

	return "LD A, A"
}

func addab() string {
	adda(Reg.B.Get())

	return "ADD A, B"
}

func addac() string {
	adda(Reg.C.Get())

	return "ADD A, C"
}

func addad() string {
	adda(Reg.D.Get())

	return "ADD A, D"
}

func addae() string {
	adda(Reg.E.Get())

	return "ADD A,E"
}

func addah() string {
	adda(Reg.H.Get())

	return "ADD A, H"
}

func addal() string {
	adda(Reg.L.Get())

	return "ADD A, L"
}

func addahl() string {
	adda(bus.Read(Reg.HL.Get()))

	return "ADD A, (HL)"
}

func addaa() string {
	adda(Reg.A.Get())

	return "ADD A, A"
}

func adcab() string {
	adda(Reg.B.Get())

	return "ADC A, B"
}

func adcac() string {
	adca(Reg.C.Get())

	return "ADC A, C"
}

func adcad() string {
	adca(Reg.D.Get())

	return "ADC A, D"
}

func adcae() string {
	adca(Reg.E.Get())

	return "ADC A,E"
}

func adcah() string {
	adca(Reg.H.Get())

	return "ADC A, H"
}

func adcal() string {
	adca(Reg.L.Get())

	return "ADC A, L"
}

func adcahl() string {
	adca(bus.Read(Reg.HL.Get()))

	return "ADC A, (HL)"
}

func adcaa() string {
	adca(Reg.A.Get())

	return "ADC A, A"
}

func subab() string {
	suba(Reg.B.Get())

	return "SUB A, B"
}

func subac() string {
	suba(Reg.C.Get())

	return "SUB A, C"
}

func subad() string {
	suba(Reg.D.Get())

	return "SUB A, D"
}

func subae() string {
	suba(Reg.E.Get())

	return "SUB A, E"
}

func subah() string {
	suba(Reg.H.Get())

	return "SUB A, H"
}

func subal() string {
	suba(Reg.L.Get())

	return "SUB A, L"
}

func subahl() string {
	suba(bus.Read(Reg.HL.Get()))

	return "SUB A, (HL)"
}

func subaa() string {
	suba(Reg.A.Get())

	return "SUB A, A"
}

func sbcab() string {
	sbca(Reg.B.Get())

	return "SBC A, B"
}

func sbcac() string {
	sbca(Reg.C.Get())

	return "SBC A, C"
}

func sbcad() string {
	sbca(Reg.D.Get())

	return "SBC A, D"
}

func sbcae() string {
	sbca(Reg.E.Get())

	return "SBC A, E"
}

func sbcah() string {
	sbca(Reg.H.Get())

	return "SBC A, H"
}

func sbcal() string {
	sbca(Reg.L.Get())

	return "SBC A, L"
}

func sbcahl() string {
	sbca(bus.Read(Reg.HL.Get()))

	return "SBC A, (HL)"
}

func sbcaa() string {
	sbca(Reg.A.Get())

	return "SBC A, A"
}

func andab() string {
	anda(Reg.B.Get())

	return "AND A, B"
}

func andac() string {
	anda(Reg.C.Get())

	return "AND A, C"
}

func andad() string {
	anda(Reg.D.Get())

	return "AND A, D"
}

func andae() string {
	anda(Reg.E.Get())

	return "AND A, E"
}

func andah() string {
	anda(Reg.H.Get())

	return "AND A, H"
}

func andal() string {
	anda(Reg.L.Get())

	return "AND A, L"
}

func andahl() string {
	anda(bus.Read(Reg.HL.Get()))

	return "AND A, (HL)"
}

func andaa() string {
	anda(Reg.A.Get())

	return "AND A, B"
}

func xorab() string {
	xora(Reg.B.Get())

	return "XOR A, B"
}

func xorac() string {
	xora(Reg.C.Get())

	return "XOR A, C"
}

func xorad() string {
	xora(Reg.D.Get())

	return "XOR A, D"
}

func xorae() string {
	xora(Reg.E.Get())

	return "XOR A, E"
}

func xorah() string {
	xora(Reg.H.Get())

	return "XOR A, H"
}

func xoral() string {
	xora(Reg.L.Get())

	return "XOR A, L"
}

func xorahl() string {
	xora(bus.Read(Reg.HL.Get()))

	return "XOR A, (HL)"
}

func xoraa() string {
	xora(Reg.A.Get())

	return "XOR A, A"
}

func orab() string {
	ora(Reg.B.Get())

	return "OR A, B"
}

func orac() string {
	ora(Reg.C.Get())

	return "OR A, C"
}

func orad() string {
	ora(Reg.D.Get())

	return "OR A, D"
}

func orae() string {
	ora(Reg.E.Get())

	return "OR A, E"
}

func orah() string {
	ora(Reg.H.Get())

	return "OR A, H"
}

func oral() string {
	ora(Reg.L.Get())

	return "OR A, L"
}

func orahl() string {
	ora(bus.Read(Reg.HL.Get()))

	return "OR A, (HL)"
}

func oraa() string {
	ora(Reg.A.Get())

	return "OR A, A"
}

func cpab() string {
	cpa(Reg.B.Get())

	return "CP A, B"
}

func cpac() string {
	cpa(Reg.C.Get())

	return "CP A, C"
}

func cpad() string {
	cpa(Reg.C.Get())

	return "CP A, D"
}

func cpae() string {
	cpa(Reg.E.Get())

	return "CP A, E"
}

func cpah() string {
	cpa(Reg.H.Get())

	return "CP A, H"
}

func cpal() string {
	cpa(Reg.L.Get())

	return "CP A, L"
}

func cpahl() string {
	cpa(bus.Read(Reg.HL.Get()))

	return "CP A, (HL)"
}

func cpaa() string {
	cpa(Reg.A.Get())

	return "CP A, A"
}

func retnz() string {
	retCond(!flags.GetFlagZ(), 3)

	return "RET NZ"
}

func popbc() string {
	pop16(Reg.B.Val(), Reg.C.Val())

	return "POP BC"
}

func jpnz() string {
	value := jpCond(!flags.GetFlagZ(), 1)

	return fmt.Sprintf("JP, NZ, $%X", value)
}

func jp() string {
	value := jpCond(true, 0)

	return fmt.Sprintf("JP $%.4X", value)
}

func callnz() string {
	value := callCond(!flags.GetFlagZ(), 3)

	return fmt.Sprintf("CALL NZ, $%X", value)
}

func pushbc() string {
	push16(Reg.B.Get(), Reg.C.Get())

	return "PUSH BC"
}

func adda8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	adda(value)

	return fmt.Sprintf("ADD A, $%X", value)
}

func rst00() string {
	callmem(00)

	return "RST $00"
}

func retz() string {
	retCond(flags.GetFlagZ(), 3)

	return "RET Z"
}

func ret() string {
	retCond(true, 0)

	return "RET"
}

func jpz() string {
	value := jpCond(flags.GetFlagZ(), 1)

	return fmt.Sprintf("JP Z, $%X", value)
}

func prefixcb() string {
	// Fetch instruction
	curOP = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	// Decode
	instruction := cbInst[curOP]

	// Execute Operation
	ticks += instruction.ticks
	return instruction.execute()
}

func callz() string {
	value := callCond(flags.GetFlagZ(), 3)

	return fmt.Sprintf("CALL Z, $%X", value)
}

func call() string {
	value := callCond(true, 0)

	return fmt.Sprintf("CALL $%X", value)
}

func adca8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	adca(value)

	return fmt.Sprintf("ADC A, $%X", value)
}

func rst08() string {
	callmem(0x08)

	return "RST $08"
}

func retnc() string {
	retCond(!flags.GetFlagC(), 3)

	return "RET NC"
}

func popde() string {
	pop16(Reg.D.Val(), Reg.E.Val())

	return "POP DE"
}

func jpnc() string {
	value := jpCond(!flags.GetFlagC(), 1)

	return fmt.Sprintf("JP NC, $%X", value)
}

func callnc() string {
	value := callCond(!flags.GetFlagC(), 3)

	return fmt.Sprintf("CALL NC, $%X", value)
}

func pushde() string {
	push16(Reg.D.Get(), Reg.E.Get())

	return "PUSH DE"
}

func suba8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	suba(value)

	return fmt.Sprintf("SUB A, $%X", value)
}

func rst10() string {
	callmem(0x10)

	return "RST $10"
}

func retc() string {
	retCond(flags.GetFlagC(), 3)

	return "RET C"
}

func reti() string {
	// Equivalent to executing ei() followed by ret()
	Reg.IME = true
	retCond(true, 0)

	return "RETI"
}

func jpc() string {
	value := jpCond(flags.GetFlagC(), 1)

	return fmt.Sprintf("JP C, $%X", value)
}

func callc() string {
	value := callCond(flags.GetFlagC(), 3)

	return fmt.Sprintf("CALL C, $%X", value)
}

func sbca8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	sbca(value)

	return fmt.Sprintf("SBC A, $%X", value)
}

func rst18() string {
	callmem(0x18)

	return "RST $18"
}

func ldff8a() string {
	var value uint8 = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	var address uint16 = 0xFF00 + uint16(value)
	ldmem(address, Reg.A.Get())

	return fmt.Sprintf("LD (FF00 + $%.2X), A", value)
}

func pophl() string {
	pop16(Reg.H.Val(), Reg.L.Val())

	return "POP HL"
}

func ldffca() string {
	var pos uint16 = 0xFF00
	pos += uint16(Reg.C.Get())
	ldmem(pos, Reg.A.Get())

	return "LD (FF00 + C), A"
}

func pushhl() string {
	push16(Reg.H.Get(), Reg.L.Get())

	return "PUSH HL"
}

func anda8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	anda(value)

	return fmt.Sprintf("AND A, $%X", value)
}

func rst20() string {
	callmem(0x20)

	return "RST $20"
}

func addsp() string {
	var value int8 = int8(bus.Read(Reg.PC.Get()))
	Reg.PC.Inc()
	Reg.SP.Set(Reg.SP.Get() + uint16(value))

	return fmt.Sprintf("ADD SP, $%X", value)
}

func jphl() string {
	Reg.PC.Set(Reg.HL.Get())

	return "JP HL"
}

func ld16a() string {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	ldmem(value, Reg.A.Get())

	return fmt.Sprintf("LD ($%X), A", value)
}

func xora8() string {
	value := bus.Read(Reg.PC.Get())
	xora(value)

	return fmt.Sprintf("XOR A, $%X", value)
}

func rst28() string {
	callmem(0x28)

	return "RST $28"
}

func ldaff8() string {
	var pos uint16 = 0xFF00
	var value uint8 = bus.Read(Reg.PC.Get())
	pos += uint16(value)
	Reg.PC.Inc()
	*Reg.A.Val() = bus.Read(pos)

	return fmt.Sprintf("LD A, (FF00 + $%X)", value)
}

func popaf() string {
	pop16(Reg.A.Val(), flags.Val())

	return "POP AF"
}

func ldaffc() string {
	var pos uint16 = 0xFF00
	pos += uint16(Reg.C.Get())
	*Reg.A.Val() = bus.Read(pos)

	return "LD A, (FF00 + C)"
}

func di() string {
	Reg.IME = false
	// Cancels any delayed IME is set by IME
	performIME = false

	return "DI"
}

func pushaf() string {
	push16(Reg.A.Get(), flags.Get())

	return "PUSH AF"
}

func ora8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	ora(value)

	return fmt.Sprintf("OR A, $%X", value)
}

func rst30() string {
	callmem(0x30)

	return "RST $30"
}

func ldhlsp8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	newValue := Reg.SP.Get() + uint16(value)
	flags.AffectFlagHC16(Reg.HL.Get(), newValue)
	Reg.HL.Set(newValue)

	return fmt.Sprintf("LD HL, SP + $%X", value)
}

func ldsphl() string {
	Reg.SP.Set(Reg.HL.Get())

	return "LD SP, HL"
}

func lda16() string {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()
	*Reg.A.Val() = bus.Read(value)

	return fmt.Sprintf("LD A, ($%X)", value)
}

func ei() string {
	performIME = true

	return "EI"
}

func cpa8() string {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()
	cpa(value)

	return fmt.Sprintf("CP A, $%X", value)
}

func rst38() string {
	callmem(0x38)

	return "RST $38"
}

func illegalop() string {
	//TODO make it optional to crash

	return "ILLEGAL OP Used"
}

//endregion Instruction Functions

//region CP Prefixed Instruction Functions

func rlcb() string {
	rlcReg(Reg.B.Val())

	return "RLC B"
}

func rlcc() string {
	rlcReg(Reg.C.Val())

	return "RLC C"
}

func rlcd() string {
	rlcReg(Reg.D.Val())

	return "RLC D"
}

func rlce() string {
	rlcReg(Reg.E.Val())

	return "RLC E"
}

func rlch() string {
	rlcReg(Reg.H.Val())

	return "RLC H"
}

func rlcl() string {
	rlcReg(Reg.L.Val())

	return "RLC L"
}

func rlchl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = rlcVal(value)
	bus.Write(pos, value)

	return "RLC (HL)"
}

func cbrlca() string {
	rlcReg(Reg.A.Val())

	return "RLC A"
}

func rrcb() string {
	rrcReg(Reg.B.Val())

	return "RRC B"
}

func rrcc() string {
	rrcReg(Reg.C.Val())

	return "RRC C"
}

func rrcd() string {
	rrcReg(Reg.D.Val())

	return "RRC D"
}

func rrce() string {
	rrcReg(Reg.E.Val())

	return "RRC E"
}

func rrch() string {
	rrcReg(Reg.H.Val())

	return "RRC H"
}

func rrcl() string {
	rrcReg(Reg.L.Val())

	return "RRC L"
}

func rrchl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = rrcVal(value)
	bus.Write(pos, value)

	return "RRC L"
}

func cbrrca() string {
	rrcReg(Reg.A.Val())

	return "RRC A"
}

func rlb() string {
	rlReg(Reg.B.Val())

	return "RL B"
}

func rlc() string {
	rlReg(Reg.C.Val())

	return "RL C"
}

func rld() string {
	rlReg(Reg.D.Val())

	return "RL D"
}

func rle() string {
	rlReg(Reg.E.Val())

	return "RL E"
}

func rlh() string {
	rlReg(Reg.H.Val())

	return "RL H"
}

func rll() string {
	rlReg(Reg.L.Val())

	return "RL L"
}

func rlhl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = rl(value)
	bus.Write(pos, value)

	return "RL (HL)"
}

func cbrla() string {
	rlReg(Reg.A.Val())

	return "RL A"
}

func rrb() string {
	rrReg(Reg.B.Val())

	return "RR B"
}

func rrc() string {
	rrReg(Reg.C.Val())

	return "RR C"
}

func rrd() string {
	rrReg(Reg.D.Val())

	return "RR D"
}

func rre() string {
	rrReg(Reg.E.Val())

	return "RR E"
}

func rrh() string {
	rrReg(Reg.H.Val())

	return "RR H"
}

func rrl() string {
	rrReg(Reg.L.Val())

	return "RR L"
}

func rrhl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = rr(value)
	bus.Write(pos, value)

	return "RR L"
}

func cbrra() string {
	rrReg(Reg.A.Val())

	return "RR A"
}

func slab() string {
	slaReg(Reg.B.Val())

	return "SLA B"
}

func slac() string {
	slaReg(Reg.C.Val())

	return "SLA C"
}

func slad() string {
	slaReg(Reg.D.Val())

	return "SLA D"
}

func slae() string {
	slaReg(Reg.E.Val())

	return "SLA E"
}

func slah() string {
	slaReg(Reg.H.Val())

	return "SLA H"
}

func slal() string {
	slaReg(Reg.L.Val())

	return "SLA L"
}

func slahl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = sla(value)
	bus.Write(pos, value)

	return "SLA (HL)"
}

func slaa() string {
	slaReg(Reg.A.Val())

	return "SLA A"
}

func srab() string {
	sraReg(Reg.B.Val())

	return "SRA B"
}

func srac() string {
	sraReg(Reg.C.Val())

	return "SRA C"
}

func srad() string {
	sraReg(Reg.D.Val())

	return "SRA D"
}

func srae() string {
	sraReg(Reg.E.Val())

	return "SRA E"
}

func srah() string {
	sraReg(Reg.H.Val())

	return "SRA H"
}

func sral() string {
	sraReg(Reg.L.Val())

	return "SRA L"
}

func srahl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = sra(value)
	bus.Write(pos, value)

	return "SRA (HL)"
}

func sraa() string {
	sraReg(Reg.A.Val())

	return "SRA A"
}

func swapb() string {
	swapReg(Reg.B.Val())

	return "SWAP B"
}

func swapc() string {
	swapReg(Reg.C.Val())

	return "SWAP C"
}

func swapd() string {
	swapReg(Reg.D.Val())

	return "SWAP D"
}

func swape() string {
	swapReg(Reg.E.Val())

	return "SWAP E"
}

func swaph() string {
	swapReg(Reg.H.Val())

	return "SWAP H"
}

func swapl() string {
	swapReg(Reg.L.Val())

	return "SWAP L"
}

func swaphl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = swap(value)
	bus.Write(pos, value)

	return "SWAP (HL)"
}

func swapa() string {
	swapReg(Reg.A.Val())

	return "SWAP A"
}

func srlb() string {
	srlReg(Reg.B.Val())

	return "SRL B"
}

func srlc() string {
	srlReg(Reg.C.Val())

	return "SRL C"
}

func srld() string {
	srlReg(Reg.D.Val())

	return "SRL D"
}

func srle() string {
	srlReg(Reg.E.Val())

	return "SRL E"
}

func srlh() string {
	srlReg(Reg.H.Val())

	return "SRL H"
}

func srll() string {
	srlReg(Reg.L.Val())

	return "SRL L"
}

func srlhl() string {
	pos := Reg.HL.Get()
	value := bus.Read(pos)
	value = srl(value)
	bus.Write(pos, value)

	return "SRL (HL)"
}

func srla() string {
	srlReg(Reg.A.Val())

	return "SRL A"
}

func bit0b() string {
	return bitNumReg(0, Reg.B.Val(), "B")
}

func bit0c() string {
	return bitNumReg(0, Reg.C.Val(), "C")
}

func bit0d() string {
	return bitNumReg(0, Reg.D.Val(), "D")
}

func bit0e() string {
	return bitNumReg(0, Reg.E.Val(), "E")
}

func bit0h() string {
	return bitNumReg(0, Reg.H.Val(), "H")
}

func bit0l() string {
	return bitNumReg(0, Reg.L.Val(), "L")
}

func bit0hl() string {
	return bitNumHL(0)
}

func bit0a() string {
	return bitNumReg(0, Reg.A.Val(), "A")
}

func bit1b() string {
	return bitNumReg(1, Reg.B.Val(), "B")
}

func bit1c() string {
	return bitNumReg(1, Reg.C.Val(), "C")
}

func bit1d() string {
	return bitNumReg(1, Reg.D.Val(), "D")
}

func bit1e() string {
	return bitNumReg(1, Reg.E.Val(), "E")
}

func bit1h() string {
	return bitNumReg(1, Reg.H.Val(), "H")
}

func bit1l() string {
	return bitNumReg(1, Reg.L.Val(), "L")
}

func bit1hl() string {
	return bitNumHL(1)
}

func bit1a() string {
	return bitNumReg(1, Reg.A.Val(), "A")
}

func bit2b() string {
	return bitNumReg(2, Reg.B.Val(), "B")
}

func bit2c() string {
	return bitNumReg(2, Reg.C.Val(), "C")
}

func bit2d() string {
	return bitNumReg(2, Reg.D.Val(), "D")
}

func bit2e() string {
	return bitNumReg(2, Reg.E.Val(), "E")
}

func bit2h() string {
	return bitNumReg(2, Reg.H.Val(), "H")
}

func bit2l() string {
	return bitNumReg(2, Reg.L.Val(), "L")
}

func bit2hl() string {
	return bitNumHL(2)
}

func bit2a() string {
	return bitNumReg(2, Reg.A.Val(), "A")
}

func bit3b() string {
	return bitNumReg(3, Reg.B.Val(), "B")
}

func bit3c() string {
	return bitNumReg(3, Reg.C.Val(), "C")
}

func bit3d() string {
	return bitNumReg(3, Reg.D.Val(), "D")
}

func bit3e() string {
	return bitNumReg(3, Reg.E.Val(), "E")
}

func bit3h() string {
	return bitNumReg(3, Reg.H.Val(), "H")
}

func bit3l() string {
	return bitNumReg(3, Reg.L.Val(), "L")
}

func bit3hl() string {
	return bitNumHL(3)
}

func bit3a() string {
	return bitNumReg(3, Reg.A.Val(), "A")
}

func bit4b() string {
	return bitNumReg(4, Reg.B.Val(), "B")
}

func bit4c() string {
	return bitNumReg(4, Reg.C.Val(), "C")
}

func bit4d() string {
	return bitNumReg(4, Reg.D.Val(), "D")
}

func bit4e() string {
	return bitNumReg(4, Reg.E.Val(), "E")
}

func bit4h() string {
	return bitNumReg(4, Reg.H.Val(), "H")
}

func bit4l() string {
	return bitNumReg(4, Reg.L.Val(), "L")
}

func bit4hl() string {
	return bitNumHL(4)
}

func bit4a() string {
	return bitNumReg(4, Reg.A.Val(), "A")
}

func bit5b() string {
	return bitNumReg(5, Reg.B.Val(), "B")
}

func bit5c() string {
	return bitNumReg(5, Reg.C.Val(), "C")
}

func bit5d() string {
	return bitNumReg(5, Reg.D.Val(), "D")
}

func bit5e() string {
	return bitNumReg(5, Reg.E.Val(), "E")
}

func bit5h() string {
	return bitNumReg(5, Reg.H.Val(), "H")
}

func bit5l() string {
	return bitNumReg(5, Reg.L.Val(), "L")
}

func bit5hl() string {
	return bitNumHL(5)
}

func bit5a() string {
	return bitNumReg(5, Reg.A.Val(), "A")
}

func bit6b() string {
	return bitNumReg(6, Reg.B.Val(), "B")
}

func bit6c() string {
	return bitNumReg(6, Reg.C.Val(), "C")
}

func bit6d() string {
	return bitNumReg(6, Reg.D.Val(), "D")
}

func bit6e() string {
	return bitNumReg(6, Reg.E.Val(), "E")
}

func bit6h() string {
	return bitNumReg(6, Reg.H.Val(), "H")
}

func bit6l() string {
	return bitNumReg(6, Reg.L.Val(), "L")
}

func bit6hl() string {
	return bitNumHL(6)
}

func bit6a() string {
	return bitNumReg(6, Reg.A.Val(), "A")
}

func bit7b() string {
	return bitNumReg(7, Reg.B.Val(), "B")
}

func bit7c() string {
	return bitNumReg(7, Reg.C.Val(), "C")
}

func bit7d() string {
	return bitNumReg(7, Reg.D.Val(), "D")
}

func bit7e() string {
	return bitNumReg(7, Reg.E.Val(), "E")
}

func bit7h() string {
	return bitNumReg(7, Reg.H.Val(), "H")
}

func bit7l() string {
	return bitNumReg(7, Reg.L.Val(), "L")
}

func bit7hl() string {
	return bitNumHL(7)
}

func bit7a() string {
	return bitNumReg(7, Reg.A.Val(), "A")
}

func res0b() string {
	return resNumReg(0, Reg.B.Val(), "B")
}

func res0c() string {
	return resNumReg(0, Reg.C.Val(), "C")
}

func res0d() string {
	return resNumReg(0, Reg.D.Val(), "D")
}

func res0e() string {
	return resNumReg(0, Reg.E.Val(), "E")
}

func res0h() string {
	return resNumReg(0, Reg.H.Val(), "H")
}

func res0l() string {
	return resNumReg(0, Reg.L.Val(), "L")
}

func res0hl() string {
	return resNumHL(0)
}

func res0a() string {
	return resNumReg(0, Reg.A.Val(), "A")
}

func res1b() string {
	return resNumReg(1, Reg.B.Val(), "B")
}

func res1c() string {
	return resNumReg(1, Reg.C.Val(), "C")
}

func res1d() string {
	return resNumReg(1, Reg.D.Val(), "D")
}

func res1e() string {
	return resNumReg(1, Reg.E.Val(), "E")
}

func res1h() string {
	return resNumReg(1, Reg.H.Val(), "H")
}

func res1l() string {
	return resNumReg(1, Reg.L.Val(), "L")
}

func res1hl() string {
	return resNumHL(1)
}

func res1a() string {
	return resNumReg(1, Reg.A.Val(), "A")
}

func res2b() string {
	return resNumReg(2, Reg.B.Val(), "B")
}

func res2c() string {
	return resNumReg(2, Reg.C.Val(), "C")
}

func res2d() string {
	return resNumReg(2, Reg.D.Val(), "D")
}

func res2e() string {
	return resNumReg(2, Reg.E.Val(), "E")
}

func res2h() string {
	return resNumReg(2, Reg.H.Val(), "H")
}

func res2l() string {
	return resNumReg(2, Reg.L.Val(), "L")
}

func res2hl() string {
	return resNumHL(2)
}

func res2a() string {
	return resNumReg(2, Reg.A.Val(), "A")
}

func res3b() string {
	return resNumReg(3, Reg.B.Val(), "B")
}

func res3c() string {
	return resNumReg(3, Reg.C.Val(), "C")
}

func res3d() string {
	return resNumReg(3, Reg.D.Val(), "D")
}

func res3e() string {
	return resNumReg(3, Reg.E.Val(), "E")
}

func res3h() string {
	return resNumReg(3, Reg.H.Val(), "H")
}

func res3l() string {
	return resNumReg(3, Reg.L.Val(), "L")
}

func res3hl() string {
	return resNumHL(3)
}

func res3a() string {
	return resNumReg(3, Reg.A.Val(), "A")
}

func res4b() string {
	return resNumReg(4, Reg.B.Val(), "B")
}

func res4c() string {
	return resNumReg(4, Reg.C.Val(), "C")
}

func res4d() string {
	return resNumReg(4, Reg.D.Val(), "D")
}

func res4e() string {
	return resNumReg(4, Reg.E.Val(), "E")
}

func res4h() string {
	return resNumReg(4, Reg.H.Val(), "H")
}

func res4l() string {
	return resNumReg(4, Reg.L.Val(), "L")
}

func res4hl() string {
	return resNumHL(4)
}

func res4a() string {
	return resNumReg(4, Reg.A.Val(), "A")
}

func res5b() string {
	return resNumReg(5, Reg.B.Val(), "B")
}

func res5c() string {
	return resNumReg(5, Reg.C.Val(), "C")
}

func res5d() string {
	return resNumReg(5, Reg.D.Val(), "D")
}

func res5e() string {
	return resNumReg(5, Reg.E.Val(), "E")
}

func res5h() string {
	return resNumReg(5, Reg.H.Val(), "H")
}

func res5l() string {
	return resNumReg(5, Reg.L.Val(), "L")
}

func res5hl() string {
	return resNumHL(5)
}

func res5a() string {
	return resNumReg(5, Reg.A.Val(), "A")
}

func res6b() string {
	return resNumReg(6, Reg.B.Val(), "B")
}

func res6c() string {
	return resNumReg(6, Reg.C.Val(), "C")
}

func res6d() string {
	return resNumReg(6, Reg.D.Val(), "D")
}

func res6e() string {
	return resNumReg(6, Reg.E.Val(), "E")
}

func res6h() string {
	return resNumReg(6, Reg.H.Val(), "H")
}

func res6l() string {
	return resNumReg(6, Reg.L.Val(), "L")
}

func res6hl() string {
	return resNumHL(6)
}

func res6a() string {
	return resNumReg(6, Reg.A.Val(), "A")
}

func res7b() string {
	return resNumReg(7, Reg.B.Val(), "B")
}

func res7c() string {
	return resNumReg(7, Reg.C.Val(), "C")
}

func res7d() string {
	return resNumReg(7, Reg.D.Val(), "D")
}

func res7e() string {
	return resNumReg(7, Reg.E.Val(), "E")
}

func res7h() string {
	return resNumReg(7, Reg.H.Val(), "H")
}

func res7l() string {
	return resNumReg(7, Reg.L.Val(), "L")
}

func res7hl() string {
	return resNumHL(7)
}

func res7a() string {
	return resNumReg(7, Reg.A.Val(), "A")
}

func set0b() string {
	return setNumReg(0, Reg.B.Val(), "B")
}

func set0c() string {
	return setNumReg(0, Reg.C.Val(), "C")
}

func set0d() string {
	return setNumReg(0, Reg.D.Val(), "D")
}

func set0e() string {
	return setNumReg(0, Reg.E.Val(), "E")
}

func set0h() string {
	return setNumReg(0, Reg.H.Val(), "H")
}

func set0l() string {
	return setNumReg(0, Reg.L.Val(), "L")
}

func set0hl() string {
	return setNumHL(0)
}

func set0a() string {
	return setNumReg(0, Reg.A.Val(), "A")
}

func set1b() string {
	return setNumReg(1, Reg.B.Val(), "B")
}

func set1c() string {
	return setNumReg(1, Reg.C.Val(), "C")
}

func set1d() string {
	return setNumReg(1, Reg.D.Val(), "D")
}

func set1e() string {
	return setNumReg(1, Reg.E.Val(), "E")
}

func set1h() string {
	return setNumReg(1, Reg.H.Val(), "H")
}

func set1l() string {
	return setNumReg(1, Reg.L.Val(), "L")
}

func set1hl() string {
	return setNumHL(1)
}

func set1a() string {
	return setNumReg(1, Reg.A.Val(), "A")
}

func set2b() string {
	return setNumReg(2, Reg.B.Val(), "B")
}

func set2c() string {
	return setNumReg(2, Reg.C.Val(), "C")
}

func set2d() string {
	return setNumReg(2, Reg.D.Val(), "D")
}

func set2e() string {
	return setNumReg(2, Reg.E.Val(), "E")
}

func set2h() string {
	return setNumReg(2, Reg.H.Val(), "H")
}

func set2l() string {
	return setNumReg(2, Reg.L.Val(), "L")
}

func set2hl() string {
	return setNumHL(2)
}

func set2a() string {
	return setNumReg(2, Reg.A.Val(), "A")
}

func set3b() string {
	return setNumReg(3, Reg.B.Val(), "B")
}

func set3c() string {
	return setNumReg(3, Reg.C.Val(), "C")
}

func set3d() string {
	return setNumReg(3, Reg.D.Val(), "D")
}

func set3e() string {
	return setNumReg(3, Reg.E.Val(), "E")
}

func set3h() string {
	return setNumReg(3, Reg.H.Val(), "H")
}

func set3l() string {
	return setNumReg(3, Reg.L.Val(), "L")
}

func set3hl() string {
	return setNumHL(3)
}

func set3a() string {
	return setNumReg(3, Reg.A.Val(), "A")
}

func set4b() string {
	return setNumReg(4, Reg.B.Val(), "B")
}

func set4c() string {
	return setNumReg(4, Reg.C.Val(), "C")
}

func set4d() string {
	return setNumReg(4, Reg.D.Val(), "D")
}

func set4e() string {
	return setNumReg(4, Reg.E.Val(), "E")
}

func set4h() string {
	return setNumReg(4, Reg.H.Val(), "H")
}

func set4l() string {
	return setNumReg(4, Reg.L.Val(), "L")
}

func set4hl() string {
	return setNumHL(4)
}

func set4a() string {
	return setNumReg(4, Reg.A.Val(), "A")
}

func set5b() string {
	return setNumReg(5, Reg.B.Val(), "B")
}

func set5c() string {
	return setNumReg(5, Reg.C.Val(), "C")
}

func set5d() string {
	return setNumReg(5, Reg.D.Val(), "D")
}

func set5e() string {
	return setNumReg(5, Reg.E.Val(), "E")
}

func set5h() string {
	return setNumReg(5, Reg.H.Val(), "H")
}

func set5l() string {
	return setNumReg(5, Reg.L.Val(), "L")
}

func set5hl() string {
	return setNumHL(5)
}

func set5a() string {
	return setNumReg(5, Reg.A.Val(), "A")
}

func set6b() string {
	return setNumReg(6, Reg.B.Val(), "B")
}

func set6c() string {
	return setNumReg(6, Reg.C.Val(), "C")
}

func set6d() string {
	return setNumReg(6, Reg.D.Val(), "D")
}

func set6e() string {
	return setNumReg(6, Reg.E.Val(), "E")
}

func set6h() string {
	return setNumReg(6, Reg.H.Val(), "H")
}

func set6l() string {
	return setNumReg(6, Reg.L.Val(), "L")
}

func set6hl() string {
	return setNumHL(6)
}

func set6a() string {
	return setNumReg(6, Reg.A.Val(), "A")
}

func set7b() string {
	return setNumReg(7, Reg.B.Val(), "B")
}

func set7c() string {
	return setNumReg(7, Reg.C.Val(), "C")
}

func set7d() string {
	return setNumReg(7, Reg.D.Val(), "D")
}

func set7e() string {
	return setNumReg(7, Reg.E.Val(), "E")
}

func set7h() string {
	return setNumReg(7, Reg.H.Val(), "H")
}

func set7l() string {
	return setNumReg(7, Reg.L.Val(), "L")
}

func set7hl() string {
	return setNumHL(7)
}

func set7a() string {
	return setNumReg(7, Reg.A.Val(), "A")
}

//endregion CP Prefixed Instruction Functions

//region Helper functions

// Increment a register by one.
// Affects Flags Z and H. Sets Flag N to 0
func incReg(r8 *uint8) {
	flags.AffectFlagZH(*r8, *r8+1)
	flags.SetFlagN(false)
	*r8++
}

// Decrement a register by one.
// Affects Flags Z and H. Sets Flag N to 0
func decReg(r8 *uint8) {
	flags.AffectFlagZH(*r8, *r8-1)
	flags.SetFlagN(true)
	*r8--
}

// Add value to register HL
// Value comes in most significant byte (high) and least
// significant byte
func addhlReg(high, low uint8) {
	addhlReg16(util.To16(high, low))
}

// Add a 16-bit value to register HL
// Affects Flag H and C. Set Flag N to Zero
func addhlReg16(value uint16) {
	curHL := Reg.HL.Get()
	nextVal := curHL + value
	Reg.HL.Set(nextVal)
	flags.SetFlagN(false)
	flags.AffectFlagHC16(curHL, nextVal)
}

// Relate Jump according to condition. Additional ticks will be added
// if condition met
// Returns byte read after the jump instruction
func jrCond(condition bool, addTicks uint8) uint8 {
	value := bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	if condition {
		// Value is converted to signed 8bit first for relative
		// positioning
		Reg.PC.Set(Reg.PC.Get() + uint16(int8(value)))
		ticks += addTicks
	}

	return value
}

// Jumps to position according to condition. Additional ticks will be added
// if condition met
// Returns 16-bit read after the jump instruction
func jpCond(condition bool, addTicks uint8) uint16 {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()

	if condition {
		Reg.PC.Set(value)
		ticks += addTicks
	}

	return value
}

// Load value to indirect address pointed by register HL
func ldhlind(value uint8) {
	bus.Write(Reg.HL.Get(), value)
}

// Adds value to the Accumulator
// Affects Flags Z, H and C
// Set Flag N to Zero
func adda(value uint8) {
	curVal := Reg.A.Get()
	*Reg.A.Val() += value
	flags.AffectFlagZHC(curVal, Reg.A.Get())
	flags.SetFlagN(false)
}

// Adds value plus carry to the Accumulator
// Affects Flags Z, H and C
// Set Flag N to Zero
func adca(value uint8) {
	if flags.GetFlagC() {
		value++
	}
	adda(value)
}

// Subtracts value from the Accumulator
// Affects Flags Z, H and C
// Set Flag N to One
func suba(value uint8) {
	curVal := Reg.A.Get()
	*Reg.A.Val() -= value
	flags.AffectFlagZHC(curVal, Reg.A.Get())
	flags.SetFlagN(false)
}

// Subtracts (value plus carry) from the Accumulator
// Affects Flags Z, H and C
// Set Flag N to One
func sbca(value uint8) {
	if flags.GetFlagC() {
		value++
	}
	suba(value)
}

// Bitwise AND between Accumulator and given value
// Affects Flag Z
// Set Flags N and C to Zero
// Set Flag H to One
func anda(value uint8) {
	*Reg.A.Val() &= value
	flags.SetFlagZ(Reg.A.Get() == 0)
	flags.SetFlagN(false)
	flags.SetFlagC(false)
	flags.SetFlagH(true)
}

// Bitwise XOR between Accumulator and given value
// Affects Flag Z
// Set Flags N, H and C to Zero
func xora(value uint8) {
	*Reg.A.Val() ^= value
	flags.SetFlagZ(Reg.A.Get() == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(false)
}

// Bitwise OR between Accumulator and given value
// Affects Flag Z
// Set Flags N, H and C to Zero
func ora(value uint8) {
	*Reg.A.Val() |= value
	flags.SetFlagZ(Reg.A.Get() == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(false)
}

// Subtracts value from Accumulator without storing result
// Affects Flag Z, H and C
// Set Flag N to One
func cpa(value uint8) {
	result := Reg.A.Get() - value
	flags.AffectFlagZHC(Reg.A.Get(), result)
}

// Return from subroutine if condition met.
func retCond(condition bool, addTicks uint8) {
	if !condition {
		return
	}

	ticks += addTicks
	Reg.SP.Inc()
	low := bus.Read(Reg.SP.Get())
	Reg.SP.Inc()
	high := bus.Read(Reg.SP.Get())
	Reg.PC.Set(util.To16(high, low))
	//fmt.Printf("Return to %s ", Reg.PC)
}

// Load value from memory at location SP and SP + 1
// low <- [SP]
// high <- [SP+1]
// SP is increment by two afterward
func pop16(high, low *uint8) {
	Reg.SP.Inc()
	*high, *low = util.From16(bus.Read16(Reg.SP.Get()))
	Reg.SP.Inc()
}

// Store value to the Stack
// [SP] <- high
// [SP - 1] <- low
// SP is decrement by two afterward
func push16(high, low uint8) {
	bus.Write(Reg.SP.Get(), high)
	bus.Write(Reg.SP.Get()-1, low)
	Reg.SP.Dec().Dec()
}

// Calls a subroutine according to condition. Additional ticks will be added
// if condition met
// Returns 16-bit read after the call instruction
func callCond(condition bool, addTicks uint8) uint16 {
	value := bus.Read16(Reg.PC.Get())
	Reg.PC.Inc().Inc()

	if condition {
		callmem(value)
		ticks += addTicks
	}

	return value
}

// Calls a subroutine
// Current PC value is pushed to stack and PC is set to value
func callmem(value uint16) {
	push16(util.From16(Reg.PC.Get()))
	Reg.PC.Set(value)
}

// Load value to memory location
func ldmem(pos uint16, value uint8) {
	bus.Write(pos, value)
}

// Rotate Left Circular a Register
// Bit 7 shifts to bit 0
// Bit 7 affect the carry Flag
// C <- [7~0] <- [7]
func rlcReg(r8 *uint8) {
	*r8 = rlcVal(*r8)
}

// Rotate Left Circular an 8-bit value
// Bit 7 shifts to bit 0
// Bit 7 affect the carry Flag
// C <- [7~0] <- [7]
func rlcVal(value uint8) uint8 {
	var bit7 bool = value&0x80 == 0x80
	// If bit 7 is 1
	value <<= 1
	if bit7 {
		value |= 1
	}
	flags.SetFlagZ(value == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit7)

	return value
}

// Rotate Right Circular a Register
// Bit 0 shifts to Carry
// [0] -> [7~0] -> C
func rrcReg(r8 *uint8) {
	*r8 = rrcVal(*r8)
}

// Rotate Right Circular an 8-bit value
// Bit 0 shifts to Carry
// [0] -> [7~0] -> C
func rrcVal(value uint8) uint8 {
	var bit0 bool = value&0x1 == 0x1
	value >>= 1
	if bit0 {
		value |= 0x80
	}
	flags.SetFlagZ(false)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit0)

	return value
}

// Rotate a Register left through Carry
// Previous Carry shifts to bit 0
// Bit 7 shift to Carry
// C <- [7~0] <- C
func rlReg(r8 *uint8) {
	*r8 = rl(*r8)
}

// Rotate an 8-bit value left through Carry
// Previous Carry shifts to bit 0
// Bit 7 shift to Carry
// C <- [7~0] <- C
func rl(value uint8) uint8 {
	oldCarry := flags.GetFlagC()
	value = sla(value)
	// If carry flag is 1
	if oldCarry {
		value |= 1
	}

	return value
}

// Rotate a Register right through Carry
// Previous Carry value shifts to bit 7
// Bit 0 shifts to Carry
// C -> [7~0] -> C
func rrReg(r8 *uint8) {
	*r8 = rr(*r8)
}

// Rotate an 8-bit value right through Carry
// Previous Carry value shifts to bit 7
// Bit 0 shifts to Carry
// C -> [7~0] -> C
func rr(value uint8) uint8 {
	oldCarry := flags.GetFlagC()
	value = sra(value)
	// If carry flag is 1
	if oldCarry {
		value |= 0x80
	}

	return value
}

// Shift Left Arithmetic a Register
// Bit 7 shift to Carry
// C <- [7~0]
func slaReg(r8 *uint8) {
	*r8 = sla(*r8)
}

// Shift Left Arithmetic an 8-bit value
// Bit 7 shift to Carry
// C <- [7~0]
func sla(value uint8) uint8 {
	var bit7 bool = value&0x80 == 0x80
	value <<= 1
	flags.SetFlagZ(value == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit7)

	return value
}

// Shift Right Arithmetic a Register
// Bit 0 shifts to Carry
// Bit 7 value doesn't change
//  [7]-> [7~0] -> C
func sraReg(r8 *uint8) {
	*r8 = sra(*r8)
}

// Shift Right Arithmetic an 8-bit value
// Bit 0 shifts to Carry
// Bit 7 value doesn't change
//  [7]-> [7~0] -> C
func sra(value uint8) uint8 {
	var bit0 bool = value&0x1 == 0x1
	value = value&0x80 | (value >> 1)
	flags.SetFlagZ(value == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit0)

	return value
}

// Swap upper four bits with lower four bits for a R	egister
// [7654] <- [3~0] || [7~5] -> [3210]
func swapReg(r8 *uint8) {
	*r8 = swap(*r8)
}

// Swap upper four bits with lower four bits for an 8-bit value
// [7654] <- [3~0] || [7~5] -> [3210]
func swap(value uint8) uint8 {
	return value<<4 | value>>4
}

// Shift Right Logic a Register
// Bit 0 shifts to Carry
// [7~0] -> C
func srlReg(r8 *uint8) {
	*r8 = srl(*r8)
}

// Shift Right Logic an 8-bit value
// Bit 0 shifts to Carry
// [7~0] -> C
func srl(value uint8) uint8 {
	var bit0 bool = value&0x1 == 0x1
	value >>= 1
	flags.SetFlagZ(value == 0)
	flags.SetFlagN(false)
	flags.SetFlagH(false)
	flags.SetFlagC(bit0)

	return value
}

// Checks whether bit of a Register is set
// pos: Bit position
// r8: Register
// name: Register Name
// Return string in format "BIT pos, name"
func bitNumReg(pos uint8, r8 *uint8, name string) string {
	bit(pos, *r8)

	return fmt.Sprintf("BIT %d, %s", pos, name)
}

// Checks whether bit of a value at memory address is set
// pos: Bit position
// Return string in format "BIT pos, (HL)"
func bitNumHL(pos uint8) string {
	addr := Reg.HL.Get()
	value := bus.Read(addr)
	bit(pos, value)

	return fmt.Sprintf("BIT %d, (HL)", pos)
}

// Checks whether bit at given position of an 8-bit value is
// set or not.
// Set Flag Z to One if bit was not set
// Set Flag N to Zero
// Set Flag H to One
func bit(pos uint8, value uint8) {
	var mask uint8 = 0x01 << pos
	var isSet bool = value&mask == mask
	flags.SetFlagZ(!isSet)
	flags.SetFlagN(false)
	flags.SetFlagH(true)
}

// Set a bit to zero at given position of a register
// pos: Bit position
// r8: Register
// name: Register Name
// Return string in format "RES pos, name"
func resNumReg(pos uint8, r8 *uint8, name string) string {
	*r8 = res(pos, *r8)

	return fmt.Sprintf("RES %d, %s", pos, name)
}

// Set a bit to zero at given position of a value at memory address
// pos: Bit position
// Return string in format "RES pos, (HL)"
func resNumHL(pos uint8) string {
	addr := Reg.HL.Get()
	value := bus.Read(addr)
	value = res(pos, value)
	bus.Write(addr, value)

	return fmt.Sprintf("RES %d, (HL)", pos)
}

// Set a bit to zero at given position of an 8-bit value
func res(pos uint8, value uint8) uint8 {
	var mask uint8 = 0x01 << pos
	mask = ^mask

	return value & mask
}

// Set a bit to one at given position of a register
// pos: Bit position
// r8: Register
// name: Register Name
// Return string in format "RES pos, name"
func setNumReg(pos uint8, r8 *uint8, name string) string {
	*r8 = set(pos, *r8)

	return fmt.Sprintf("RES %d, %s", pos, name)
}

// Set a bit to one at given position of a value at memory address
// pos: Bit position
// Return string in format "RES pos, (HL)"
func setNumHL(pos uint8) string {
	addr := Reg.HL.Get()
	value := bus.Read(addr)
	value = set(pos, value)
	bus.Write(addr, value)

	return fmt.Sprintf("RES %d, (HL)", pos)
}

// Set a bit to zero at given position of an 8-bit value
func set(pos uint8, value uint8) uint8 {
	var mask uint8 = 0x01 << pos

	return value | mask
}

//endregion Helper Functions
