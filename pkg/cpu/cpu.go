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
	inst[0x01] = Instruction{3, ldbc}
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
	inst[0x11] = Instruction{3, ldde}
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
	inst[0x21] = Instruction{3, ldhl}
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
	inst[0x32] = Instruction{2, ldhld}
	// INC SP
	inst[0x33] = Instruction{2, incsp}
	// INC (HL)
	inst[0x34] = Instruction{3, inchlind}
	// DEC (HL)
	inst[0x35] = Instruction{3, dechlind}
	// LD (HL), $FF
	inst[0x36] = Instruction{3, ldhlind}
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

func ldbc() string {
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

func ldde() string {
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
// Previous Carry shifts to bit 7
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

func ldhl() string {
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

func ldhld() string {
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

func ldhlind() string {
	reg.PC++
	value := bus.Read(reg.PC)
	bus.Write(reg.GetHL(), value)

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
