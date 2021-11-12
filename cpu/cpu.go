package cpu

import (
	"fmt"
	"log"

	"github.com/aalquaiti/gbgo/io"
)

const (
	INST_SIZE = 0x100      // Instruction Size
	DMG_HZ    = 0x100000   // Game Boy Frequency (m-ticks)
	CGB_HZ    = DMG_HZ * 2 // Game Boy Color Frequency (m-ticks)
	DIV_RATE  = 0xFFF      // Dividor Increment Rate
)

type Instruction struct {
	// TODO add comments
	ticks   uint8
	execute func() string
}

// TODO Implement Modes, including CGB double speed mode
type Mode int

var (
	bus    io.Bus
	ticks  uint8  // m-ticks remaining for an instruction
	count  uint32 // m-ticks count. Helps with timer functionality
	reg    *Register
	curOP  uint8 // Current Op to execute. Used in cpu fetch phase
	inst   [INST_SIZE]Instruction
	cbInst [INST_SIZE]Instruction
)

// Variables that should be treated as immutable.
// Access should be through functions
var (
	cpuFreq uint32
)

func CPU_FREQ() uint32 {
	return cpuFreq
}

func init() {

	initInstructions()
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
	inst[0x00] = Instruction{2, rlcb}
	// RLC C
	inst[0x01] = Instruction{2, rlcc}
	// RLC D
	inst[0x02] = Instruction{2, rlcd}
	// RLC E
	inst[0x03] = Instruction{2, rlce}
	// RLC H
	inst[0x04] = Instruction{2, rlch}
	// RLC L
	inst[0x05] = Instruction{2, rlcl}
	// RLC (HL)
	inst[0x06] = Instruction{2, rlchl}
	// RLC A
	inst[0x07] = Instruction{2, cbrlca}
	// RRC B
	inst[0x08] = Instruction{2, rrcb}
	// RRC C
	inst[0x09] = Instruction{2, rrcc}
	// RRC D
	inst[0x0A] = Instruction{2, rrcd}
	// RRC E
	inst[0x0B] = Instruction{2, rrce}
	// RRC H
	inst[0x0C] = Instruction{2, rrch}
	// RRC L
	inst[0x0D] = Instruction{2, rrcl}
	// RRC (HL)
	inst[0x0E] = Instruction{2, rrchl}
	// RRC A
	inst[0x0F] = Instruction{2, cbrrca}

	// RL B
	inst[0x10] = Instruction{2, rlb}
	// RL C
	inst[0x11] = Instruction{2, rlc}
	// RL D
	inst[0x12] = Instruction{2, rld}
	// RL E
	inst[0x13] = Instruction{2, rle}
	// RL H
	inst[0x14] = Instruction{2, rlh}
	// RL L
	inst[0x15] = Instruction{2, rll}
	// RL (HL)
	inst[0x16] = Instruction{2, rlhl}
	// RL A
	inst[0x17] = Instruction{2, cbrla}
	// RR B
	inst[0x18] = Instruction{2, rrb}
	// RR C
	inst[0x19] = Instruction{2, rrc}
	// RR D
	inst[0x1A] = Instruction{2, rrd}
	// RR E
	inst[0x1B] = Instruction{2, rre}
	// RR H
	inst[0x1C] = Instruction{2, rrh}
	// RR L
	inst[0x1D] = Instruction{2, rrl}
	// RR (HL)
	inst[0x1E] = Instruction{2, rrhl}
	// RR A
	inst[0x1F] = Instruction{2, cbrra}

	// SLA B
	inst[0x20] = Instruction{2, slab}
	// SLA C
	inst[0x21] = Instruction{2, slac}
	// SLA D
	inst[0x22] = Instruction{2, slad}
	// SLA E
	inst[0x23] = Instruction{2, slae}
	// SLA H
	inst[0x24] = Instruction{2, slah}
	// SLA L
	inst[0x25] = Instruction{2, slal}
	// SLA (HL)
	inst[0x26] = Instruction{2, slahl}
	// SLA A
	inst[0x27] = Instruction{2, slaa}
	// SRA B
	inst[0x28] = Instruction{2, srab}
	// SRA C
	inst[0x29] = Instruction{2, srac}
	// SRA D
	inst[0x2A] = Instruction{2, srad}
	// SRA E
	inst[0x2B] = Instruction{2, srae}
	// SRA H
	inst[0x2C] = Instruction{2, srah}
	// SRA L
	inst[0x2D] = Instruction{2, sral}
	// SRA (HL)
	inst[0x2E] = Instruction{2, srahl}
	// SRA A
	inst[0x2F] = Instruction{2, sraa}

	// SWAP B
	inst[0x30] = Instruction{2, swapb}
	// SWAP C
	inst[0x31] = Instruction{2, swapc}
	// SWAP D
	inst[0x32] = Instruction{2, swapd}
	// SWAP E
	inst[0x33] = Instruction{2, swape}
	// SWAP H
	inst[0x34] = Instruction{2, swaph}
	// SWAP L
	inst[0x35] = Instruction{2, swapl}
	// SWAP (HL)
	inst[0x36] = Instruction{2, swaphl}
	// SWAP A
	inst[0x37] = Instruction{2, swapa}
	// SRL B
	inst[0x28] = Instruction{2, srlb}
	// SRL C
	inst[0x29] = Instruction{2, srlc}
	// SRL D
	inst[0x2A] = Instruction{2, srld}
	// SRL E
	inst[0x2B] = Instruction{2, srle}
	// SRL H
	inst[0x2C] = Instruction{2, srlh}
	// SRL L
	inst[0x2D] = Instruction{2, srll}
	// SRL (HL)
	inst[0x2E] = Instruction{2, srlhl}
	// SRL A
	inst[0x2F] = Instruction{2, srla}

	// SWAP B
	inst[0x30] = Instruction{2, swapb}
	// SWAP C
	inst[0x31] = Instruction{2, swapc}
	// SWAP D
	inst[0x32] = Instruction{2, swapd}
	// SWAP E
	inst[0x33] = Instruction{2, swape}
	// SWAP H
	inst[0x34] = Instruction{2, swaph}
	// SWAP L
	inst[0x35] = Instruction{2, swapl}
	// SWAP (HL)
	inst[0x36] = Instruction{2, swaphl}
	// SWAP A
	inst[0x37] = Instruction{2, swapa}
	// SRL B
	inst[0x38] = Instruction{2, srlb}
	// SRL C
	inst[0x39] = Instruction{2, srlc}
	// SRL D
	inst[0x3A] = Instruction{2, srld}
	// SRL E
	inst[0x3B] = Instruction{2, srle}
	// SRL H
	inst[0x3C] = Instruction{2, srlh}
	// SRL L
	inst[0x3D] = Instruction{2, srll}
	// SRL (HL)
	inst[0x3E] = Instruction{2, srlhl}
	// SRL A
	inst[0x3F] = Instruction{2, srla}

	// BIT 0, B
	inst[0x40] = Instruction{2, bit0b}
	// BIT 0, C
	inst[0x41] = Instruction{2, bit0c}
	// BIT 0, D
	inst[0x42] = Instruction{2, bit0d}
	// BIT 0, E
	inst[0x43] = Instruction{2, bit0e}
	// BIT 0, H
	inst[0x44] = Instruction{2, bit0h}
	// BIT 0, L
	inst[0x45] = Instruction{2, bit0l}
	// BIT 0, (HL)
	inst[0x46] = Instruction{2, bit0hl}
	// BIT 0, A
	inst[0x47] = Instruction{2, bit0a}
	// BIT 1, B
	inst[0x48] = Instruction{2, bit1b}
	// BIT 1, C
	inst[0x49] = Instruction{2, bit1c}
	// BIT 1, D
	inst[0x4A] = Instruction{2, bit1d}
	// BIT 1, E
	inst[0x4B] = Instruction{2, bit1e}
	// BIT 1, H
	inst[0x4C] = Instruction{2, bit1h}
	// BIT 1, L
	inst[0x4D] = Instruction{2, bit1l}
	// BIT 1, (HL)
	inst[0x4E] = Instruction{2, bit1hl}
	// BIT 1, A
	inst[0x4F] = Instruction{2, bit1a}

	// BIT 2, B
	inst[0x50] = Instruction{2, bit2b}
	// BIT 2, C
	inst[0x51] = Instruction{2, bit2c}
	// BIT 2, D
	inst[0x52] = Instruction{2, bit2d}
	// BIT 2, E
	inst[0x53] = Instruction{2, bit2e}
	// BIT 2, H
	inst[0x54] = Instruction{2, bit2h}
	// BIT 2, L
	inst[0x55] = Instruction{2, bit2l}
	// BIT 2, (HL)
	inst[0x56] = Instruction{2, bit2hl}
	// BIT 2, A
	inst[0x57] = Instruction{2, bit2a}
	// BIT 3, B
	inst[0x58] = Instruction{2, bit3b}
	// BIT 3, C
	inst[0x59] = Instruction{2, bit3c}
	// BIT 3, D
	inst[0x5A] = Instruction{2, bit3d}
	// BIT 3, E
	inst[0x5B] = Instruction{2, bit3e}
	// BIT 3, H
	inst[0x5C] = Instruction{2, bit3h}
	// BIT 3, L
	inst[0x5D] = Instruction{2, bit3l}
	// BIT 3, (HL)
	inst[0x5E] = Instruction{2, bit3hl}
	// BIT 3, A
	inst[0x5F] = Instruction{2, bit3a}

	// BIT 4, B
	inst[0x60] = Instruction{2, bit4b}
	// BIT 4, C
	inst[0x61] = Instruction{2, bit4c}
	// BIT 4, D
	inst[0x62] = Instruction{2, bit4d}
	// BIT 4, E
	inst[0x63] = Instruction{2, bit4e}
	// BIT 4, H
	inst[0x64] = Instruction{2, bit4h}
	// BIT 4, L
	inst[0x65] = Instruction{2, bit4l}
	// BIT 4, (HL)
	inst[0x66] = Instruction{2, bit4hl}
	// BIT 4, A
	inst[0x67] = Instruction{2, bit4a}
	// BIT 5, B
	inst[0x68] = Instruction{2, bit5b}
	// BIT 5, C
	inst[0x69] = Instruction{2, bit5c}
	// BIT 5, D
	inst[0x6A] = Instruction{2, bit5d}
	// BIT 5, E
	inst[0x6B] = Instruction{2, bit5e}
	// BIT 5, H
	inst[0x6C] = Instruction{2, bit5h}
	// BIT 5, L
	inst[0x6D] = Instruction{2, bit5l}
	// BIT 5, (HL)
	inst[0x6E] = Instruction{2, bit5hl}
	// BIT 5, A
	inst[0x6F] = Instruction{2, bit5a}

	// BIT 6, B
	inst[0x70] = Instruction{2, bit6b}
	// BIT 6, C
	inst[0x71] = Instruction{2, bit6c}
	// BIT 6, D
	inst[0x72] = Instruction{2, bit6d}
	// BIT 6, E
	inst[0x73] = Instruction{2, bit6e}
	// BIT 6, H
	inst[0x74] = Instruction{2, bit6h}
	// BIT 6, L
	inst[0x75] = Instruction{2, bit6l}
	// BIT 6, (HL)
	inst[0x76] = Instruction{2, bit6hl}
	// BIT 6, A
	inst[0x77] = Instruction{2, bit6a}
	// BIT 7, B
	inst[0x78] = Instruction{2, bit7b}
	// BIT 7, C
	inst[0x79] = Instruction{2, bit7c}
	// BIT 7, D
	inst[0x7A] = Instruction{2, bit7d}
	// BIT 7, E
	inst[0x7B] = Instruction{2, bit7e}
	// BIT 7, H
	inst[0x7C] = Instruction{2, bit7h}
	// BIT 7, L
	inst[0x7D] = Instruction{2, bit7l}
	// BIT 7, (HL)
	inst[0x7E] = Instruction{2, bit7hl}
	// BIT 7, A
	inst[0x7F] = Instruction{2, bit7a}

	// endregion CB Prefixed Instructions
}

// Handles functionality related to divider and timer
func handleTimer() {
	// When count reaches Dividor rate
	if count&DIV_RATE == DIV_RATE {
		// TODO emulate CGB double speed effect
		bus.IncDIV()
	}

	// When count reaches Timer rate, increment Timer Counter.
	// Timer Rate depends on Timer Control's bits 0 and 1. The following
	// shows frequency in m-ticks:
	// 00: CPU Clock / 0x1000 = 0x400 Hz
	// 01: CPU Clock / 0x10   = 0x10000 Hz
	// 10: CPU Clock / 0x40   = 0x4000 Hz
	// 11: CPU Clock / 0x100  = 0x1000 Hz
	timeRate := CPU_FREQ()
	switch bus.GetTacClockSelect() {
	case 0b00:
		timeRate /= 0x1000
	case 0b01:
		timeRate /= 0x10
	case 0b10:
		timeRate /= 0x40
	case 0b11:
		timeRate /= 0x100
	}

	// Time Rate is decreased by one to AND it with CPU clock count.
	// This will shows if the time should be incremented
	// TODO: Add an example for clarity
	timeRate -= 1
	var timeReached bool = count&timeRate == timeRate

	if timeReached {
		// When Timer is Enabled and Timer Counter overflow,
		// set Timer counter to value stored in TMA and request a
		// Timer interrupt
		timaCount := bus.Read(io.TIMA_ADDR)
		if bus.IsTacTimerEnabled() && timaCount == 0xFF {
			bus.Write(io.TIMA_ADDR, bus.Read(io.TMA_ADDR))
			bus.SetIRQTimer(true)
		} else {
			bus.Write(io.TIMA_ADDR, timaCount+1)
		}

	}
}

// Hanndles Interrupt request
func irq() {
	// Checks is Master Interrupt is enabled,
	// Ignores intterupts if disabled
	if !reg.IME {
		return
	}

	// Disable further Intterupts. This is a CPU behaviour when an
	// interrupt is to be executed. So further interrupts must be enabled
	// by the program (Usually using RETI instruction when returning from
	// an interrupt vector)
	reg.IME = false

	if bus.IsVblank() && bus.IrqVblank() {
		bus.SetIrQVblank(false)
		push16(from16(reg.PC))
		reg.PC = 0x40
	} else if bus.IsLCDStat() && bus.IrqLCDStat() {
		bus.SetIRQLCDStat(false)
		push16(from16(reg.PC))
		reg.PC = 0x48
	} else if bus.IsTimerInt() && bus.IrqTimer() {
		bus.SetIRQTimer(false)
		push16(from16(reg.PC))
		reg.PC = 0x50
	} else if bus.IsSerialInt() && bus.IrqSerial() {
		bus.SetIrqSerial(false)
		push16(from16(reg.PC))
		reg.PC = 0x58
	} else if bus.IsJoypadInt() && bus.IrqJoypad() {
		bus.SetIrqJoypad(false)
		push16(from16(reg.PC))
		reg.PC = 0x60
	}

}

// Emulates machine ticks (m-ticks). Each m-tick is equivalent to four
// system tick
// Goes through a fetch-decode-execute cycle
func Tick() {

	// Count is important for handling time. Therefore it should not
	// exceed cpu frequency value
	if count == CPU_FREQ() {
		count = 0
	}

	// m-ticks needs to be finished before executing
	// the cycle again
	if ticks > 0 {
		ticks--
		count++
		return
	}

	// TODO check if timer should be handled with each tick, instead with
	// the current simulated bulks of ticks
	handleTimer()
	irq()

	// Fetch instruction
	curOP = bus.Read(reg.PC)
	reg.PC++

	// Decode
	instruction := inst[curOP]

	// Execute Operation
	// TODO de-assemble and print executed operation
	ticks += instruction.ticks
	instruction.execute()

	handleTimer()
	ticks--
	count++
}

//region Instruction Functions

func nop() string {
	return "NOP"
}

func ldbc16() string {
	value := bus.Read16(reg.PC + 1)
	reg.PC += 2
	reg.SetBC(value)

	return fmt.Sprintf("LD BC, %X", value)
}

func ldbca() string {
	pos := reg.GetBC()
	bus.Write(pos, reg.A)

	return "LD (BC), A"
}

func incbc() string {
	value := reg.GetBC()
	value++
	reg.SetBC(value)

	return "INC BC"
}

func incb() string {
	incReg(&reg.B)

	return "INC B"
}

func decb() string {
	decReg(&reg.B)

	return "DEC B"
}

func ldb() string {
	reg.PC++
	reg.B = bus.Read(reg.PC)

	return fmt.Sprintf("LD B, $%X", reg.B)
}

// Rotate Register A left
// Bit 7 shifts to bit 0
// Bit 7 affect the carry Flag
// C <- [7~0] <- [7]
func rlca() string {
	var bit7 bool = reg.A&0x80 == 0x80
	// If bit 7 is 1
	reg.A <<= 1
	if bit7 {
		reg.A |= 1
	}
	reg.SetFlagZ(false)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit7)

	return "RLCA"
}

func ldmemsp() string {
	pos := bus.Read16(reg.PC + 1)
	reg.PC += 2
	bus.Write16(pos, reg.SP)

	return fmt.Sprintf("LD ($%X), SP", pos)
}

func addhlbc() string {
	addhlReg(reg.B, reg.C)

	return "ADD HL, BC"
}

func ldabc() string {
	pos := reg.GetBC()
	reg.A = bus.Read(pos)

	return "LD A, (BC)"
}

func decbc() string {
	value := reg.GetBC() - 1
	reg.SetBC(value)

	return "DEC BC"
}

func incc() string {
	incReg(&reg.C)

	return "INC C"
}

func decc() string {
	decReg(&reg.C)

	return "DEC C"
}

func ldc() string {
	reg.PC++
	reg.C = bus.Read(reg.PC)

	return fmt.Sprintf("LD C, $%X", reg.C)
}

// Rotate Register A right
// Bit 0 shifts to Carry
// [0] -> [7~0] -> C
func rrca() string {
	var bit0 bool = reg.A&0x1 == 0x1
	reg.A >>= 1
	if bit0 {
		reg.A |= 0x80
	}
	reg.SetFlagZ(false)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit0)

	return "RRCA"
}

// Enters CPU low power mode.
// In GBC, switches between normal and double CPU speed
func stop() string {
	// TODO implement cpu speed switch

	reg.PC++

	return "STOP"
}

func ldde16() string {
	value := bus.Read16(reg.PC + 1)
	reg.PC += 2
	reg.SetDE(value)

	return fmt.Sprintf("LD DE, %X", value)
}

func lddea() string {
	pos := reg.GetDE()
	bus.Write(pos, reg.A)

	return "LD (DE), A"
}

func incde() string {
	value := reg.GetDE()
	value++
	reg.SetDE(value)

	return "INC DE"
}

func incd() string {
	incReg(&reg.D)

	return "INC D"
}

func decd() string {
	decReg(&reg.D)

	return "DEC D"
}

func ldd() string {
	reg.PC++
	reg.D = bus.Read(reg.PC)

	return fmt.Sprintf("LD D, $%X", reg.D)
}

// Rotate Register A left through Carry
// Previous Carry shifts to bit 0
// Bit 7 shift to Carry
// C <- [7~0] <- C
func rla() string {
	var bit7 bool = reg.A&0x80 == 0x80
	reg.A <<= 1
	// If carry flag is 1
	if reg.GetFlagC() {
		reg.A |= 1
	}
	reg.SetFlagZ(false)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit7)

	return "RLA"
}

func jr() string {
	value := jrCond(true, 0)

	return fmt.Sprintf("JR $%X", value)
}

func addhlde() string {
	addhlReg(reg.D, reg.E)

	return "ADD HL, DE"
}

func ldade() string {
	pos := reg.GetDE()
	reg.A = bus.Read(pos)

	return "LD A, (DE)"
}

func decde() string {
	value := reg.GetDE() - 1
	reg.SetDE(value)

	return "DEC DE"
}

func ince() string {
	incReg(&reg.E)

	return "INC E"
}

func dece() string {
	decReg(&reg.E)

	return "DEC E"
}

func lde() string {
	reg.PC++
	reg.E = bus.Read(reg.PC)

	return fmt.Sprintf("LD E, $%X", reg.E)
}

// Rotate Register A right through Carry
// Previous Carry value shifts to bit 7
// Bit 0 shifts to Carry
// C -> [7~0] -> C
func rra() string {
	var bit0 bool = reg.A&0x1 == 0x1
	reg.A >>= 1
	// If carry flag is 1
	if reg.GetFlagC() {
		reg.A |= 0x80
	}
	reg.SetFlagZ(false)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit0)

	return "RRA"
}

func jrnz() string {
	value := jrCond(!reg.GetFlagZ(), 1)

	return fmt.Sprintf("JR NZ, $%X", value)
}

func ldhl16() string {
	value := bus.Read16(reg.PC + 1)
	reg.PC += 2
	reg.SetHL(value)

	return fmt.Sprintf("LD HL, %X", value)
}

func ldhli() string {
	pos := reg.GetHL()
	bus.Write(pos, reg.A)
	reg.SetHL(pos + 1)

	return "LD (HLI), A"
}

func inchl() string {
	value := reg.GetHL()
	value++
	reg.SetHL(value)

	return "INC HL"
}

func inch() string {
	incReg(&reg.H)

	return "INC H"
}

func dech() string {
	decReg(&reg.H)

	return "DEC H"
}

func ldh() string {
	reg.PC++
	reg.H = bus.Read(reg.PC)

	return fmt.Sprintf("LD H, $%X", reg.H)
}

func daa() string {
	// TODO implement
	log.Fatal("Not implemented Yet")

	return "Not Implemented"
}

func jrz() string {
	value := jrCond(reg.GetFlagZ(), 1)

	return fmt.Sprintf("JR Z, $%X", value)
}

func addhlhl() string {
	addhlReg(reg.H, reg.L)

	return "ADD HL, HL"
}

func ldahli() string {
	pos := reg.GetHL()
	reg.A = bus.Read(pos)
	reg.SetHL(pos + 1)

	return "LD A, (HLI)"
}

func dechl() string {
	value := reg.GetHL() - 1
	reg.SetHL(value)

	return "DEC HL"
}

func incl() string {
	incReg(&reg.L)

	return "INC L"
}

func decl() string {
	decReg(&reg.L)

	return "DEC L"
}

func ldl() string {
	reg.PC++
	reg.L = bus.Read(reg.PC)

	return fmt.Sprintf("LD L, $%X", reg.L)
}

func cpl() string {
	reg.A = ^reg.A

	return "CPL"
}

func jrnc() string {
	value := jrCond(!reg.GetFlagC(), 1)

	return fmt.Sprintf("JR NC, $%X", value)
}

func ldsp() string {
	reg.SP = bus.Read16(reg.PC + 1)
	reg.PC += 2

	return fmt.Sprintf("LD SP, %X", reg.SP)
}

func ldhlda() string {
	pos := reg.GetHL()
	bus.Write(pos, reg.A)
	reg.SetHL(pos - 1)

	return "LD (HLD), A"
}

func incsp() string {
	reg.SP++

	return "INC SP"
}

func inchlind() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	reg.AffectFlagZH(value, value+1)
	reg.SetFlagN(false)
	bus.Write(pos, value+1)

	return "INC (HL)"
}

func dechlind() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	reg.AffectFlagZH(value, value+1)
	reg.SetFlagN(true)
	bus.Write(pos, value+1)

	return "DEC H"
}

func ldhl8() string {
	reg.PC++
	value := bus.Read(reg.PC)
	ldhlind(value)

	return fmt.Sprintf("LD (HL), $%X", reg.H)
}

// Set Carry Flag
// Flags N and H are set to Zero
func scf() string {
	reg.SetFlagC(true)
	reg.SetFlagN(false)
	reg.SetFlagH(false)

	return "SCF"
}

func jrc() string {
	value := jrCond(reg.GetFlagC(), 1)

	return fmt.Sprintf("JR C, $%X", value)
}

func addhlsp() string {
	addhlReg16(reg.SP)

	return "ADD HL, SP"
}

func ldahld() string {
	pos := reg.GetHL()
	reg.A = bus.Read(pos)
	reg.SetHL(pos - 1)

	return "LD A, (HLD)"
}

func decsp() string {
	reg.SP--

	return "DEC SP"
}

func inca() string {
	incReg(&reg.A)

	return "INC A"
}

func deca() string {
	decReg(&reg.A)

	return "DEC A"
}

func lda() string {
	reg.PC++
	reg.A = bus.Read(reg.PC)

	return fmt.Sprintf("LD A, $%X", reg.A)
}

// Complement Carry Flag
func ccf() string {
	reg.SetFlagC(!reg.GetFlagC())

	return "CCF"
}

func ldbb() string {

	return "LD B, B"
}

func ldbc() string {
	reg.B = reg.C

	return "LD B, C"
}

func ldbd() string {
	reg.B = reg.D

	return "LD B, D"
}

func ldbe() string {
	reg.B = reg.E

	return "LD B, E"
}

func ldbh() string {
	reg.B = reg.H

	return "LD B, H"
}

func ldbl() string {
	reg.B = reg.L

	return "LD B, L"
}

func ldbhl() string {
	reg.B = bus.Read(reg.GetHL())

	return "LD B, (HL)"
}

func ldba() string {
	reg.B = reg.A

	return "LD B, A"
}

func ldcb() string {
	reg.C = reg.B

	return "LD C, B"
}

func ldcc() string {

	return "LD C, C"
}

func ldcd() string {
	reg.C = reg.D

	return "LD C, D"
}

func ldce() string {
	reg.C = reg.E

	return "LD C, E"
}

func ldch() string {
	reg.C = reg.H

	return "LD C, H"
}

func ldcl() string {
	reg.C = reg.L

	return "LD C, L"
}

func ldchl() string {
	reg.C = bus.Read(reg.GetHL())

	return "LD C, (HL)"
}

func ldca() string {
	reg.C = reg.A

	return "LD C, A"
}

func lddb() string {
	reg.D = reg.B

	return "LD D, B"
}

func lddc() string {
	reg.D = reg.C

	return "LD D, C"
}

func lddd() string {

	return "LD D, D"
}

func ldde() string {
	reg.D = reg.E

	return "LD D, E"
}

func lddh() string {
	reg.D = reg.H

	return "LD D, H"
}

func lddl() string {
	reg.D = reg.L

	return "LD D, L"
}

func lddhl() string {
	reg.D = bus.Read(reg.GetHL())

	return "LD D, (HL)"
}

func ldda() string {
	reg.D = reg.A

	return "LD D, A"
}

func ldeb() string {
	reg.E = reg.B

	return "LD E, B"
}

func ldec() string {
	reg.E = reg.C

	return "LD E, C"
}

func lded() string {
	reg.E = reg.D

	return "LD E, D"
}

func ldee() string {

	return "LD E, E"
}

func ldeh() string {
	reg.E = reg.H

	return "LD E, H"
}

func ldel() string {
	reg.E = reg.L

	return "LD E, L"
}

func ldehl() string {
	reg.E = bus.Read(reg.GetHL())

	return "LD E, (HL)"
}

func ldea() string {
	reg.E = reg.A

	return "LD E, A"
}

func ldhb() string {
	reg.H = reg.B

	return "LD H, B"
}

func ldhc() string {
	reg.H = reg.C

	return "LD H, C"
}

func ldhd() string {
	reg.H = reg.D

	return "LD H, D"
}

func ldhe() string {
	reg.H = reg.E

	return "LD H, E"
}

func ldhh() string {

	return "LD H, H"
}

func ldhl() string {
	reg.H = reg.L

	return "LD H, L"
}

func ldhhl() string {
	reg.H = bus.Read(reg.GetHL())

	return "LD H, (HL)"
}

func ldha() string {
	reg.H = reg.A

	return "LD H, A"
}

func ldlb() string {
	reg.L = reg.B

	return "LD L, B"
}

func ldlc() string {
	reg.L = reg.C

	return "LD L, C"
}

func ldld() string {
	reg.L = reg.D

	return "LD L, D"
}

func ldle() string {
	reg.L = reg.E

	return "LD L, E"
}

func ldlh() string {
	reg.L = reg.H

	return "LD L, H"
}

func ldll() string {

	return "LD L, L"
}

func ldlhl() string {
	reg.L = bus.Read(reg.GetHL())

	return "LD L, (HL)"
}

func ldla() string {
	reg.H = reg.A

	return "LD L, A"
}

func ldhlb() string {
	ldhlind(reg.B)

	return "LD (HL), B"
}

func ldhlc() string {
	ldhlind(reg.C)

	return "LD (HL), C"
}

func ldhld() string {
	ldhlind(reg.D)

	return "LD (HL), D"
}

func ldhle() string {
	ldhlind(reg.E)

	return "LD (HL), E"
}

func ldhlh() string {
	ldhlind(reg.H)

	return "LD (HL), H"
}

func ldhll() string {
	ldhlind(reg.L)

	return "LD (HL), L"
}

func halt() string {
	//TODO implement
	log.Fatal("Not implemented")

	return "Not Implemented"
}

func ldhla() string {
	ldhlind(reg.A)

	return "LD (HL), A"
}

func ldab() string {
	reg.A = reg.B

	return "LD A, B"
}

func ldac() string {
	reg.A = reg.C

	return "LD A, C"
}

func ldad() string {
	reg.A = reg.D

	return "LD A, D"
}

func ldae() string {
	reg.A = reg.E

	return "LD A, E"
}

func ldah() string {
	reg.A = reg.H

	return "LD A, H"
}

func ldal() string {
	reg.H = reg.L

	return "LD A, L"
}

func ldahl() string {
	reg.A = bus.Read(reg.GetHL())

	return "LD A, (HL)"
}

func ldaa() string {

	return "LD A, A"
}

func addab() string {
	adda(reg.B)

	return "ADD A, B"
}

func addac() string {
	adda(reg.C)

	return "ADD A, C"
}

func addad() string {
	adda(reg.D)

	return "ADD A, D"
}

func addae() string {
	adda(reg.E)

	return "ADD A,E"
}

func addah() string {
	adda(reg.H)

	return "ADD A, H"
}

func addal() string {
	adda(reg.L)

	return "ADD A, L"
}

func addahl() string {
	adda(bus.Read(reg.GetHL()))

	return "ADD A, (HL)"
}

func addaa() string {
	adda(reg.A)

	return "ADD A, A"
}

func adcab() string {
	adda(reg.B)

	return "ADC A, B"
}

func adcac() string {
	adca(reg.C)

	return "ADC A, C"
}

func adcad() string {
	adca(reg.D)

	return "ADC A, D"
}

func adcae() string {
	adca(reg.E)

	return "ADC A,E"
}

func adcah() string {
	adca(reg.H)

	return "ADC A, H"
}

func adcal() string {
	adca(reg.L)

	return "ADC A, L"
}

func adcahl() string {
	adca(bus.Read(reg.GetHL()))

	return "ADC A, (HL)"
}

func adcaa() string {
	adca(reg.A)

	return "ADC A, A"
}

func subab() string {
	suba(reg.B)

	return "SUB A, B"
}

func subac() string {
	suba(reg.C)

	return "SUB A, C"
}

func subad() string {
	suba(reg.D)

	return "SUB A, D"
}

func subae() string {
	suba(reg.E)

	return "SUB A, E"
}

func subah() string {
	suba(reg.H)

	return "SUB A, H"
}

func subal() string {
	suba(reg.L)

	return "SUB A, L"
}

func subahl() string {
	suba(bus.Read(reg.GetHL()))

	return "SUB A, (HL)"
}

func subaa() string {
	suba(reg.A)

	return "SUB A, A"
}

func sbcab() string {
	sbca(reg.B)

	return "SBC A, B"
}

func sbcac() string {
	sbca(reg.C)

	return "SBC A, C"
}

func sbcad() string {
	sbca(reg.D)

	return "SBC A, D"
}

func sbcae() string {
	sbca(reg.E)

	return "SBC A, E"
}

func sbcah() string {
	sbca(reg.H)

	return "SBC A, H"
}

func sbcal() string {
	sbca(reg.L)

	return "SBC A, L"
}

func sbcahl() string {
	sbca(bus.Read(reg.GetHL()))

	return "SBC A, (HL)"
}

func sbcaa() string {
	sbca(reg.A)

	return "SBC A, A"
}

func andab() string {
	anda(reg.B)

	return "AND A, B"
}

func andac() string {
	anda(reg.C)

	return "AND A, C"
}

func andad() string {
	anda(reg.D)

	return "AND A, D"
}

func andae() string {
	anda(reg.E)

	return "AND A, E"
}

func andah() string {
	anda(reg.H)

	return "AND A, H"
}

func andal() string {
	anda(reg.L)

	return "AND A, L"
}

func andahl() string {
	anda(bus.Read(reg.GetHL()))

	return "AND A, (HL)"
}

func andaa() string {
	anda(reg.A)

	return "AND A, B"
}

func xorab() string {
	xora(reg.B)

	return "XOR A, B"
}

func xorac() string {
	xora(reg.C)

	return "XOR A, C"
}

func xorad() string {
	xora(reg.D)

	return "XOR A, D"
}

func xorae() string {
	xora(reg.E)

	return "XOR A, E"
}

func xorah() string {
	xora(reg.H)

	return "XOR A, H"
}

func xoral() string {
	xora(reg.L)

	return "XOR A, L"
}

func xorahl() string {
	xora(bus.Read(reg.GetHL()))

	return "XOR A, (HL)"
}

func xoraa() string {
	xora(reg.A)

	return "XOR A, A"
}

func orab() string {
	ora(reg.B)

	return "OR A, B"
}

func orac() string {
	ora(reg.C)

	return "OR A, C"
}

func orad() string {
	ora(reg.D)

	return "OR A, D"
}

func orae() string {
	ora(reg.E)

	return "OR A, E"
}

func orah() string {
	ora(reg.H)

	return "OR A, H"
}

func oral() string {
	ora(reg.L)

	return "OR A, L"
}

func orahl() string {
	ora(bus.Read(reg.GetHL()))

	return "OR A, (HL)"
}

func oraa() string {
	ora(reg.A)

	return "OR A, A"
}

func cpab() string {
	cpa(reg.B)

	return "CP A, B"
}

func cpac() string {
	cpa(reg.C)

	return "CP A, C"
}

func cpad() string {
	cpa(reg.C)

	return "CP A, D"
}

func cpae() string {
	cpa(reg.E)

	return "CP A, E"
}

func cpah() string {
	cpa(reg.H)

	return "CP A, H"
}

func cpal() string {
	cpa(reg.L)

	return "CP A, L"
}

func cpahl() string {
	cpa(bus.Read(reg.GetHL()))

	return "CP A, (HL)"
}

func cpaa() string {
	cpa(reg.A)

	return "CP A, A"
}

func retnz() string {
	retCond(!reg.GetFlagZ(), 3)

	return "RET NZ"
}

func popbc() string {
	pop16(&reg.B, &reg.C)

	return "POP BC"
}

func jpnz() string {
	value := jpCond(!reg.GetFlagZ(), 1)

	return fmt.Sprintf("JP, NZ, $%X", value)
}

func jp() string {
	value := jpCond(true, 0)

	return fmt.Sprintf("JP $%X", value)
}

func callnz() string {
	value := callCond(!reg.GetFlagZ(), 3)

	return fmt.Sprintf("CALL NZ, $%X", value)
}

func pushbc() string {
	push16(reg.B, reg.C)

	return "PUSH BC"
}

func adda8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	adda(value)

	return fmt.Sprintf("ADD A, $%X", value)
}

func rst00() string {
	callmem(00)

	return "RST $00"
}

func retz() string {
	retCond(reg.GetFlagZ(), 3)

	return "RET Z"
}

func ret() string {
	retCond(true, 0)

	return "RET"
}

func jpz() string {
	value := jpCond(reg.GetFlagZ(), 1)

	return fmt.Sprintf("JP Z, $%X", value)
}

func prefixcb() string {
	// Fetch instruction
	curOP = bus.Read(reg.PC)
	reg.PC++

	// Decode
	instruction := cbInst[curOP]

	// Execute Operation
	ticks += instruction.ticks
	return instruction.execute()
}

func callz() string {
	value := callCond(reg.GetFlagZ(), 3)

	return fmt.Sprintf("CALL Z, $%X", value)
}

func call() string {
	value := callCond(true, 0)

	return fmt.Sprintf("CALL $%X", value)
}

func adca8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	adca(value)

	return fmt.Sprintf("ADC A, $%X", value)
}

func rst08() string {
	callmem(0x08)

	return "RST $08"
}

func retnc() string {
	retCond(!reg.GetFlagC(), 3)

	return "RET NC"
}

func popde() string {
	pop16(&reg.D, &reg.E)

	return "POP DE"
}

func jpnc() string {
	value := jpCond(!reg.GetFlagC(), 1)

	return fmt.Sprintf("JP NC, $%X", value)
}

func callnc() string {
	value := callCond(!reg.GetFlagC(), 3)

	return fmt.Sprintf("CALL NC, $%X", value)
}

func pushde() string {
	push16(reg.D, reg.E)

	return "PUSH DE"
}

func suba8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	suba(value)

	return fmt.Sprintf("SUB A, $%X", value)
}

func rst10() string {
	callmem(0x10)

	return "RST $10"
}

func retc() string {
	retCond(reg.GetFlagC(), 3)

	return "RET C"
}

func reti() string {
	// Equivalent to executing ei() followed by ret()
	reg.IME = true
	retCond(true, 0)

	return "RETI"
}

func jpc() string {
	value := jpCond(reg.GetFlagC(), 1)

	return fmt.Sprintf("JP C, $%X", value)
}

func callc() string {
	value := callCond(reg.GetFlagC(), 3)

	return fmt.Sprintf("CALL C, $%X", value)
}

func sbca8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	sbca(value)

	return fmt.Sprintf("SBC A, $%X", value)
}

func rst18() string {
	callmem(0x18)

	return "RST $18"
}

func ldff8a() string {
	var pos uint16 = 0xFF00
	var value uint8 = bus.Read(reg.PC)
	pos += uint16(value)
	reg.PC++
	ldmem(pos, reg.A)

	return fmt.Sprint("LD (FF00 + $X), A", value)
}

func pophl() string {
	pop16(&reg.H, &reg.L)

	return "POP HL"
}

func ldffca() string {
	var pos uint16 = 0xFF00
	pos += uint16(reg.C)
	ldmem(pos, reg.A)

	return "LD (FF00 + C), A"
}

func pushhl() string {
	push16(reg.H, reg.L)

	return "PUSH HL"
}

func anda8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	anda(value)

	return fmt.Sprintf("AND A, $%X", value)
}

func rst20() string {
	callmem(0x20)

	return "RST $20"
}

func addsp() string {
	var value int8 = int8(bus.Read(reg.PC))
	reg.PC++
	reg.SP += uint16(value)

	return fmt.Sprintf("ADD SP, $%X", value)
}

func jphl() string {
	reg.PC = reg.GetHL()

	return "JP HL"
}

func ld16a() string {
	value := bus.Read16(reg.PC)
	reg.PC += 2
	ldmem(value, reg.A)

	return fmt.Sprintf("LD ($%X), A", value)
}

func xora8() string {
	value := bus.Read(reg.PC)
	xora(value)

	return fmt.Sprintf("XOR A, $%X", value)
}

func rst28() string {
	callmem(0x28)

	return "RST $28"
}

func ldaff8() string {
	var pos uint16 = 0xFF00
	var value uint8 = bus.Read(reg.PC)
	pos += uint16(value)
	reg.PC++
	reg.A = bus.Read(pos)

	return fmt.Sprintf("LD A, (FF00 + $%X)", value)
}

func popaf() string {
	var pos uint16 = 0xFF00
	pos += uint16(reg.C)
	reg.A = bus.Read(pos)

	return "LD A, (FF00 + C)"
}

func ldaffc() string {
	var pos uint16 = 0xFF00
	pos += uint16(reg.C)
	reg.A = bus.Read(pos)

	return "LD A, (FF00 + C)"
}

func di() string {
	reg.IME = false

	return "DI"
}

func pushaf() string {
	push16(reg.A, reg.F)

	return "PUSH AF"
}

func ora8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	ora(value)

	return fmt.Sprintf("OR A, $%X", value)
}

func rst30() string {
	callmem(0x30)

	return "RST $30"
}

func ldhlsp8() string {
	value := bus.Read(reg.PC)
	reg.PC++
	newValue := reg.SP + uint16(value)
	reg.AffectFlagHC16(reg.GetHL(), newValue)
	reg.SetHL(newValue)

	return fmt.Sprintf("LD HL, SP + $%X", value)
}

func ldsphl() string {
	reg.SP = reg.GetHL()

	return "LD SP, HL"
}

func lda16() string {
	value := bus.Read16(reg.PC)
	reg.PC += 2
	reg.A = bus.Read(value)

	return fmt.Sprintf("LD A, ($%X)", value)
}

func ei() string {
	reg.IME = true

	return "EI"
}

func cpa8() string {
	value := bus.Read(reg.PC)
	reg.PC++
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
	rlcReg(&reg.B)

	return "RLC B"
}

func rlcc() string {
	rlcReg(&reg.C)

	return "RLC C"
}

func rlcd() string {
	rlcReg(&reg.D)

	return "RLC D"
}

func rlce() string {
	rlcReg(&reg.E)

	return "RLC E"
}

func rlch() string {
	rlcReg(&reg.H)

	return "RLC H"
}

func rlcl() string {
	rlcReg(&reg.L)

	return "RLC L"
}

func rlchl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = rlcVal(value)
	bus.Write(pos, value)

	return "RLC (HL)"
}

func cbrlca() string {
	rlcReg(&reg.A)

	return "RLC A"
}

func rrcb() string {
	rrcReg(&reg.B)

	return "RRC B"
}

func rrcc() string {
	rrcReg(&reg.C)

	return "RRC C"
}

func rrcd() string {
	rrcReg(&reg.D)

	return "RRC D"
}

func rrce() string {
	rrcReg(&reg.E)

	return "RRC E"
}

func rrch() string {
	rrcReg(&reg.H)

	return "RRC H"
}

func rrcl() string {
	rrcReg(&reg.L)

	return "RRC L"
}

func rrchl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = rrcVal(value)
	bus.Write(pos, value)

	return "RRC L"
}

func cbrrca() string {
	rrcReg(&reg.A)

	return "RRC A"
}

func rlb() string {
	rlReg(&reg.B)

	return "RL B"
}

func rlc() string {
	rlReg(&reg.C)

	return "RL C"
}

func rld() string {
	rlReg(&reg.D)

	return "RL D"
}

func rle() string {
	rlReg(&reg.E)

	return "RL E"
}

func rlh() string {
	rlReg(&reg.H)

	return "RL H"
}

func rll() string {
	rlReg(&reg.L)

	return "RL L"
}

func rlhl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = rl(value)
	bus.Write(pos, value)

	return "RL (HL)"
}

func cbrla() string {
	rlReg(&reg.A)

	return "RL A"
}

func rrb() string {
	rrReg(&reg.B)

	return "RR B"
}

func rrc() string {
	rrReg(&reg.C)

	return "RR C"
}

func rrd() string {
	rrReg(&reg.D)

	return "RR D"
}

func rre() string {
	rrReg(&reg.E)

	return "RR E"
}

func rrh() string {
	rrReg(&reg.H)

	return "RR H"
}

func rrl() string {
	rrReg(&reg.L)

	return "RR L"
}

func rrhl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = rr(value)
	bus.Write(pos, value)

	return "RR L"
}

func cbrra() string {
	rrReg(&reg.A)

	return "RR A"
}

func slab() string {
	slaReg(&reg.B)

	return "SLA B"
}

func slac() string {
	slaReg(&reg.C)

	return "SLA C"
}

func slad() string {
	slaReg(&reg.D)

	return "SLA D"
}

func slae() string {
	slaReg(&reg.E)

	return "SLA E"
}

func slah() string {
	slaReg(&reg.H)

	return "SLA H"
}

func slal() string {
	slaReg(&reg.L)

	return "SLA L"
}

func slahl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = sla(value)
	bus.Write(pos, value)

	return "SLA (HL)"
}

func slaa() string {
	slaReg(&reg.A)

	return "SLA A"
}

func srab() string {
	sraReg(&reg.B)

	return "SRA B"
}

func srac() string {
	sraReg(&reg.C)

	return "SRA C"
}

func srad() string {
	sraReg(&reg.D)

	return "SRA D"
}

func srae() string {
	sraReg(&reg.E)

	return "SRA E"
}

func srah() string {
	sraReg(&reg.H)

	return "SRA H"
}

func sral() string {
	sraReg(&reg.L)

	return "SRA L"
}

func srahl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = sra(value)
	bus.Write(pos, value)

	return "SRA (HL)"
}

func sraa() string {
	sraReg(&reg.A)

	return "SRA A"
}

func swapb() string {
	swapReg(&reg.B)

	return "SWAP B"
}

func swapc() string {
	swapReg(&reg.C)

	return "SWAP C"
}

func swapd() string {
	swapReg(&reg.D)

	return "SWAP D"
}

func swape() string {
	swapReg(&reg.E)

	return "SWAP E"
}

func swaph() string {
	swapReg(&reg.H)

	return "SWAP H"
}

func swapl() string {
	swapReg(&reg.L)

	return "SWAP L"
}

func swaphl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = swap(value)
	bus.Write(pos, value)

	return "SWAP (HL)"
}

func swapa() string {
	swapReg(&reg.A)

	return "SWAP A"
}

func srlb() string {
	srlReg(&reg.B)

	return "SRL B"
}

func srlc() string {
	srlReg(&reg.C)

	return "SRL C"
}

func srld() string {
	srlReg(&reg.D)

	return "SRL D"
}

func srle() string {
	srlReg(&reg.E)

	return "SRL E"
}

func srlh() string {
	srlReg(&reg.H)

	return "SRL H"
}

func srll() string {
	srlReg(&reg.L)

	return "SRL L"
}

func srlhl() string {
	pos := reg.GetHL()
	value := bus.Read(pos)
	value = srl(value)
	bus.Write(pos, value)

	return "SRL (HL)"
}

func srla() string {
	srlReg(&reg.A)

	return "SRL A"
}

func bit0b() string {
	return bitNumReg(0, &reg.B, "B")
}

func bit0c() string {
	return bitNumReg(0, &reg.C, "C")
}

func bit0d() string {
	return bitNumReg(0, &reg.D, "D")
}

func bit0e() string {
	return bitNumReg(0, &reg.E, "E")
}

func bit0h() string {
	return bitNumReg(0, &reg.H, "H")
}

func bit0l() string {
	return bitNumReg(0, &reg.L, "L")
}

func bit0hl() string {
	return bitNumHL(0)
}

func bit0a() string {
	return bitNumReg(0, &reg.A, "A")
}

func bit1b() string {
	return bitNumReg(1, &reg.B, "B")
}

func bit1c() string {
	return bitNumReg(1, &reg.C, "C")
}

func bit1d() string {
	return bitNumReg(1, &reg.D, "D")
}

func bit1e() string {
	return bitNumReg(1, &reg.E, "E")
}

func bit1h() string {
	return bitNumReg(1, &reg.H, "H")
}

func bit1l() string {
	return bitNumReg(1, &reg.L, "L")
}

func bit1hl() string {
	return bitNumHL(1)
}

func bit1a() string {
	return bitNumReg(1, &reg.A, "A")
}

func bit2b() string {
	return bitNumReg(2, &reg.B, "B")
}

func bit2c() string {
	return bitNumReg(2, &reg.C, "C")
}

func bit2d() string {
	return bitNumReg(2, &reg.D, "D")
}

func bit2e() string {
	return bitNumReg(2, &reg.E, "E")
}

func bit2h() string {
	return bitNumReg(2, &reg.H, "H")
}

func bit2l() string {
	return bitNumReg(2, &reg.L, "L")
}

func bit2hl() string {
	return bitNumHL(2)
}

func bit2a() string {
	return bitNumReg(2, &reg.A, "A")
}

func bit3b() string {
	return bitNumReg(3, &reg.B, "B")
}

func bit3c() string {
	return bitNumReg(3, &reg.C, "C")
}

func bit3d() string {
	return bitNumReg(3, &reg.D, "D")
}

func bit3e() string {
	return bitNumReg(3, &reg.E, "E")
}

func bit3h() string {
	return bitNumReg(3, &reg.H, "H")
}

func bit3l() string {
	return bitNumReg(3, &reg.L, "L")
}

func bit3hl() string {
	return bitNumHL(3)
}

func bit3a() string {
	return bitNumReg(3, &reg.A, "A")
}

func bit4b() string {
	return bitNumReg(4, &reg.B, "B")
}

func bit4c() string {
	return bitNumReg(4, &reg.C, "C")
}

func bit4d() string {
	return bitNumReg(4, &reg.D, "D")
}

func bit4e() string {
	return bitNumReg(4, &reg.E, "E")
}

func bit4h() string {
	return bitNumReg(4, &reg.H, "H")
}

func bit4l() string {
	return bitNumReg(4, &reg.L, "L")
}

func bit4hl() string {
	return bitNumHL(4)
}

func bit4a() string {
	return bitNumReg(4, &reg.A, "A")
}

func bit5b() string {
	return bitNumReg(5, &reg.B, "B")
}

func bit5c() string {
	return bitNumReg(5, &reg.C, "C")
}

func bit5d() string {
	return bitNumReg(5, &reg.D, "D")
}

func bit5e() string {
	return bitNumReg(5, &reg.E, "E")
}

func bit5h() string {
	return bitNumReg(5, &reg.H, "H")
}

func bit5l() string {
	return bitNumReg(5, &reg.L, "L")
}

func bit5hl() string {
	return bitNumHL(5)
}

func bit5a() string {
	return bitNumReg(5, &reg.A, "A")
}

func bit6b() string {
	return bitNumReg(6, &reg.B, "B")
}

func bit6c() string {
	return bitNumReg(6, &reg.C, "C")
}

func bit6d() string {
	return bitNumReg(6, &reg.D, "D")
}

func bit6e() string {
	return bitNumReg(6, &reg.E, "E")
}

func bit6h() string {
	return bitNumReg(6, &reg.H, "H")
}

func bit6l() string {
	return bitNumReg(6, &reg.L, "L")
}

func bit6hl() string {
	return bitNumHL(6)
}

func bit6a() string {
	return bitNumReg(6, &reg.A, "A")
}

func bit7b() string {
	return bitNumReg(7, &reg.B, "B")
}

func bit7c() string {
	return bitNumReg(7, &reg.C, "C")
}

func bit7d() string {
	return bitNumReg(7, &reg.D, "D")
}

func bit7e() string {
	return bitNumReg(7, &reg.E, "E")
}

func bit7h() string {
	return bitNumReg(7, &reg.H, "H")
}

func bit7l() string {
	return bitNumReg(7, &reg.L, "L")
}

func bit7hl() string {
	return bitNumHL(7)
}

func bit7a() string {
	return bitNumReg(7, &reg.A, "A")
}

//endregion CP Prefixed Instruction Functions

//region Helper functions

// Increment a register by one.
// Affects Flags Z and H. Sets Flag N to 0
func incReg(r8 *uint8) {
	reg.AffectFlagZH(*r8, *r8+1)
	reg.SetFlagN(false)
	*r8++
}

// Decrement a register by one.
// Affects Flags Z and H. Sets Flag N to 0
func decReg(r8 *uint8) {
	reg.AffectFlagZH(*r8, *r8+1)
	reg.SetFlagN(true)
	*r8++
}

// Add value to register HL
// Value comes in most significant byte (high) and least
// significant byte
func addhlReg(high, low uint8) {
	addhlReg16(to16(high, low))
}

// Add a 16-bit value to register HL
// Affects Flag H and C. Set Flag N to Zero
func addhlReg16(value uint16) {
	curHL := reg.GetHL()
	nextVal := curHL + value
	reg.SetHL(nextVal)
	reg.SetFlagN(false)
	reg.AffectFlagHC16(curHL, nextVal)
}

// Relate Jump according to condition. Additional ticks will be added
// if condition met
// Returns byte read after the jump instruction
func jrCond(condition bool, addTicks uint8) uint8 {
	value := bus.Read(reg.PC + 1)
	reg.PC++

	if condition {
		// Value is converted to signed 8bit first for relative
		// positioning
		reg.PC += uint16(int8(value))
		ticks += addTicks
	}

	return value
}

// Jumps to position according to condition. Additional ticks will be added
// if condition met
// Returns 16-bit read after the jump instruction
func jpCond(condition bool, addTicks uint8) uint16 {
	value := bus.Read16(reg.PC + 1)
	reg.PC += 2

	if condition {
		reg.PC = value
		ticks += addTicks
	}

	return value
}

// Load value to indirect address pointed by register HL
func ldhlind(value uint8) {
	bus.Write(reg.GetHL(), value)
}

// Adds value to the Accumulator
// Affects Flags Z, H and C
// Set Flag N to Zero
func adda(value uint8) {
	curVal := reg.A
	reg.A += value
	reg.AffectFlagZHC(curVal, reg.A)
	reg.SetFlagN(false)
}

// Adds value plus carry to the Accumulator
// Affects Flags Z, H and C
// Set Flag N to Zero
func adca(value uint8) {
	if reg.GetFlagC() {
		value++
	}
	adda(value)
}

// Subtracts value from the Accumulator
// Affects Flags Z, H and C
// Set Flag N to One
func suba(value uint8) {
	curVal := reg.A
	reg.A -= value
	reg.AffectFlagZHC(curVal, reg.A)
	reg.SetFlagN(false)
}

// Subtracts (value plus carry) from the Accumulator
// Affects Flags Z, H and C
// Set Flag N to One
func sbca(value uint8) {
	if reg.GetFlagC() {
		value++
	}
	suba(value)
}

// Bitwise AND between Accumulator and given value
// Affects Flag Z
// Set Flags N and C to Zero
// Set Flag H to One
func anda(value uint8) {
	reg.A &= value
	reg.SetFlagZ(reg.A == 0)
	reg.SetFlagN(false)
	reg.SetFlagC(false)
	reg.SetFlagH(true)
}

// Bitwise XOR between Accumulator and given value
// Affects Flag Z
// Set Flags N, H and C to Zero
func xora(value uint8) {
	reg.A ^= value
	reg.SetFlagZ(reg.A == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(false)
}

// Bitwise OR between Accumulator and given value
// Affects Flag Z
// Set Flags N, H and C to Zero
func ora(value uint8) {
	reg.A |= value
	reg.SetFlagZ(reg.A == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(false)
}

// Subtracts value from Accumulator without storing result
// Affects Flag Z, H and C
// Set Flag N to One
func cpa(value uint8) {
	result := reg.A - value
	reg.AffectFlagZHC(reg.A, result)
}

// Return from subroutine if condition met.
func retCond(condition bool, addTicks uint8) {
	if !condition {
		return
	}

	ticks += addTicks
	reg.PC = bus.Read16(reg.SP)
	reg.SP++
}

// Load value from memory at location SP and SP + 1
// low <- [SP]
// high <- [SP+1]
// SP is increment by two afterward
func pop16(high, low *uint8) {
	*low, *high = from16(bus.Read16(reg.PC))
	reg.SP += 2
}

// Store value to the Stack
// [SP] <- high
// [SP - 1] <- low
// SP is decrement by two afterward
func push16(high, low uint8) {
	bus.Write(reg.SP, high)
	bus.Write(reg.SP-1, low)
	reg.SP -= 2
}

// Calls a subroutine according to condition. Additional ticks will be added
// if condition met
// Returns 16-bit read after the call instruction
func callCond(condition bool, addTicks uint8) uint16 {
	value := bus.Read16(reg.PC + 1)
	reg.PC += 2

	if condition {
		callmem(value)
		ticks += addTicks
	}

	return value
}

// Calls a subroutine
// Current PC value is pushed to stack and PC is set to value
func callmem(value uint16) {
	push16(from16(reg.PC))
	reg.PC = value
}

// Load value to memory location
func ldmem(pos uint16, value uint8) {
	bus.Write(pos, value)
}

//endregion Helper Functions

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
	reg.SetFlagZ(value == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit7)

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
	reg.SetFlagZ(false)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit0)

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
	oldCarry := reg.GetFlagC()
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
	oldCarry := reg.GetFlagC()
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
	reg.SetFlagZ(value == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit7)

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
	reg.SetFlagZ(value == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit0)

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
	reg.SetFlagZ(value == 0)
	reg.SetFlagN(false)
	reg.SetFlagH(false)
	reg.SetFlagC(bit0)

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
// address: Memory address
// Return string in format "BIT pos, (HL)"
func bitNumHL(pos uint8) string {
	addr := reg.GetHL()
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
	reg.SetFlagZ(!isSet)
	reg.SetFlagN(false)
	reg.SetFlagH(true)
}
