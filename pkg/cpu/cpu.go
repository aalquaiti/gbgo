package cpu

import (
	"fmt"
	"log"

	"github.com/aalquaiti/gbgo/pkg"
)

const (
	INST_SIZE = 0xFF
)

type Instruction struct {
	// TODO add comments
	ticks   uint
	execute func() string
}

var (
	bus    pkg.Bus
	ticks  uint
	reg    *Register
	curOP  uint8
	nextOP uint8
	inst   [INST_SIZE]Instruction
)

//TODO Add setters and getters for HI LO access

//TODO the lower four bits of Flag Registerare always 0

func init() {

	initInstructions()
}

// Initialise instructions
func initInstructions() {
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
}

// Emulates machine ticks (m-ticks)
func Tick() {
	ticks--

	if ticks > 0 {
		return
	}

	// execute
	// TODO de-assemble and print
	// Read tick numbers after
	inst[curOP].execute()

	// Read next instruction
	curOP = nextOP
	reg.PC++
	nextOP = bus.Read(reg.PC)
}

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
	incmem(&reg.B)

	return "INC B"
}

func decb() string {
	decmem(&reg.B)

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
	reg.SetFlagC(bit7)
	if bit7 {
		reg.A |= 1
	}

	return "RLCA"
}

func ldmemsp() string {
	pos := bus.Read16(reg.PC + 1)
	reg.PC += 2
	bus.Write16(pos, reg.SP)

	return fmt.Sprintf("LD ($%X), SP", pos)
}

func addhlbc() string {
	addhlreg(reg.B, reg.C)

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
	incmem(&reg.C)

	return "INC C"
}

func decc() string {
	decmem(&reg.C)

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
	reg.SetFlagC(bit0)
	if bit0 {
		reg.A |= 0x80
	}

	return "RRCA"
}

func stop() string {
	// TODO implement
	log.Fatal("Not implemented Yet")

	return "NOT implemented"
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
	incmem(&reg.D)

	return "INC D"
}

func decd() string {
	decmem(&reg.D)

	return "DEC D"
}

func ldd() string {
	reg.PC++
	reg.D = bus.Read(reg.PC)

	return fmt.Sprintf("LD D, $%X", reg.D)
}

// Rotate Register A left
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
	reg.SetFlagC(bit7)

	return "RLA"
}

func jr() string {
	value := jrcond(true, 0)

	return fmt.Sprintf("JR $%X", value)
}

func addhlde() string {
	addhlreg(reg.D, reg.E)

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
	incmem(&reg.E)

	return "INC E"
}

func dece() string {
	decmem(&reg.E)

	return "DEC E"
}

func lde() string {
	reg.PC++
	reg.E = bus.Read(reg.PC)

	return fmt.Sprintf("LD E, $%X", reg.E)
}

// Rotate Register A left
// Previous Carry s-===========hifts to bit 7
// Bit 0 shifts to Carry
// C -> [7~0] -> C
func rra() string {
	var bit0 bool = reg.A&0x1 == 0x1
	reg.A >>= 1
	// If carry flag is 1
	if reg.GetFlagC() {
		reg.A |= 0x80
	}
	reg.SetFlagC(bit0)

	return "RRA"
}

func jrnz() string {
	value := jrcond(!reg.GetFlagZ(), 1)

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
	incmem(&reg.H)

	return "INC H"
}

func dech() string {
	decmem(&reg.H)

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
	value := jrcond(reg.GetFlagZ(), 1)

	return fmt.Sprintf("JR Z, $%X", value)
}

func addhlhl() string {
	addhlreg(reg.H, reg.L)

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
	incmem(&reg.L)

	return "INC L"
}

func decl() string {
	decmem(&reg.L)

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
	value := jrcond(!reg.GetFlagC(), 1)

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
	incmem(&bus.Ram[pos])

	return "INC (HL)"
}

func dechlind() string {
	pos := reg.GetHL()
	decmem(&bus.Ram[pos])

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
	value := jrcond(reg.GetFlagC(), 1)

	return fmt.Sprintf("JR C, $%X", value)
}

func addhlsp() string {
	addhlreg16(reg.SP)

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
	incmem(&reg.A)

	return "INC A"
}

func deca() string {
	decmem(&reg.A)

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
	adda(bus.Ram[reg.GetHL()])

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
	adca(bus.Ram[reg.GetHL()])

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
	suba(bus.Ram[reg.GetHL()])

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
	sbca(bus.Ram[reg.GetHL()])

	return "SBC A, (HL)"
}

func sbcaa() string {
	sbca(reg.A)

	return "SBC A, A"
}

// Helper functions

// Increment an 8-bit memory location by one.
// Affects Flags Z and H. Sets Flag N to 0
func incmem(r8 *uint8) {
	reg.AffectFlagZH(*r8, *r8+1)
	reg.SetFlagN(false)
	*r8++
}

// Decrement an 8-bit memory location by one.
// Affects Flags Z and H. Sets Flag N to 0
func decmem(r8 *uint8) {
	reg.AffectFlagZH(*r8, *r8+1)
	reg.SetFlagN(true)
	*r8++
}

// Add value to register HL
// Value comes in most significant byte (high) and least
// significant byte
func addhlreg(high, low uint8) {
	addhlreg16(to16(high, low))
}

// Add a 16-bit value to register HL
// Affects Flag H and C. Set Flag N to Zero
func addhlreg16(value uint16) {
	curHL := reg.GetHL()
	nextVal := curHL + value
	reg.SetHL(nextVal)
	reg.SetFlagN(false)
	reg.AffectFlagHC16(curHL, nextVal)
}

// Jumps according to condition. Additional ticks will be added
// if condition met
// Returns byte read after the jump instruction
func jrcond(condition bool, addTicks uint8) uint8 {
	reg.PC++

	value := bus.Read(reg.PC)
	relPos := int8(value)

	if condition {

		reg.PC += uint16(relPos)
		ticks += uint(addTicks)
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
	sbca(value)
}
