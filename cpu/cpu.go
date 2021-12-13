package cpu

import (
	"fmt"
	bit2 "github.com/aalquaiti/gbgo/bit"

	"github.com/aalquaiti/gbgo/io"
)

const (
	INST_SIZE = 0x100      // Instruction Size
	DMG_HZ    = 0x100000   // Game Boy Frequency (m-ticks)
	CGB_HZ    = DMG_HZ * 2 // Game Boy Color Frequency (m-ticks)
	DIV_RATE  = 0xFFF      // Divider Increment Rate
)

// Mode Game Boy CPU Mode
type Mode int

const (
	DMG_MODE Mode = iota
	CGB_MODE
)

var (
	mode   Mode
	bus    io.Bus
	ticks  uint8  // m-ticks remaining for an instruction
	cycles uint32 // m-ticks count. Helps with timer functionality
	steps  uint32 // Counts how many instructions executed
	Reg    Register
	flags  *RegF
	curOP  uint8 // Current Op to execute. Used in cpu fetch phase
	inst   [INST_SIZE]Instruction
	cbInst [INST_SIZE]Instruction
)

// Variables that are used by CPU to help perform some of the instructions
var (
	// IsHalt Used by HALT instruction to inform the cpu to halt until an interrupt occurs. This helps with interruption
	// handling that breaks the halt, and to emulate the HALT bug
	isHalt bool

	//  Used by EI instruction to set IME. The EI has a delay one of one cycle, so IME will be set to one after the
	// execution of the next instruction following EI.
	performIME bool
)

// Variables that should be treated as immutable.
// Access should be through functions
var (
	freq uint32
)

func Frequency() uint32 {
	return freq
}

// Init Initialise CPU
func Init(m Mode, b io.Bus) {
	// TODO use different CPU mode
	mode = m
	bus = b
	initInstructions()
	Reset()
}

// Reset resets CPU to pre-start state
func Reset() {
	ticks = 0
	cycles = 0
	Reg = NewRegister()
	flags, _ = Reg.F.(*RegF)
	curOP = 0
	isHalt = false

	// Power-Up Sequence for DMG
	// TODO Add power-up sequence of CGB
	Reg.A.Set(0x01)
	// TODO Reg F value depends on header checksum. Make amends according
	// Refer to https://gbdev.io/pandocs/Power_Up_Sequence.html
	Reg.F.Set(0xB0)
	Reg.B.Set(0x00)
	Reg.C.Set(0x13)
	Reg.D.Set(0x00)
	Reg.E.Set(0xD8)
	Reg.H.Set(0x01)
	Reg.L.Set(0x4D)
	Reg.PC.Set(0x100)
	Reg.SP.Set(0xFFFE)
}

// Step ticks a cpu until an op is executed
func Step() {
	steps++
	currentPC := Reg.PC.Get()
	af := Reg.AF.String()
	bc := Reg.BC.String()
	de := Reg.DE.String()
	hl := Reg.HL.String()
	pc := Reg.PC.String()
	sp := Reg.SP.String()
	output := Tick()
	// Tick()

	for ticks > 0 {
		Tick()
	}

	fmt.Printf("%.04X:\t%-30s %s, %s, %s, %s, %s, %s\n",
		currentPC, output, bc, de, hl, af, sp, pc)
}

// timer handles functionality related to divider and timer
func timer() {
	// When count reaches Divider rate
	if cycles&DIV_RATE == DIV_RATE {
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
	timeRate := Frequency()
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
	timeReached := cycles&timeRate == timeRate

	if timeReached {
		// When Timer is Enabled and Timer Counter overflow,
		// set Timer counter to value stored in TMA and request a
		// Timer interrupt
		timaCount := bus.Read(io.TimaAddr)
		if bus.IsTacTimerEnabled() && timaCount == 0xFF {
			bus.Write(io.TimaAddr, bus.Read(io.TmaAddr))
			bus.SetIRQTimer(true)
		} else {
			bus.Write(io.TimaAddr, timaCount+1)
		}

	}
}

// irq handles Interrupt request
func irq() {

	// Checks is Master Interrupt is enabled,
	// Ignores interrupts if disabled
	if !Reg.IME {

		// Having an interrupt pending breaks the halt loop if one was requested through HALT operation
		// See more details in halt() function
		if bus.InterruptPending() {
			// TODO simulate the halt bug by reading the instruction the follow HALT twice
			isHalt = false
		}

		return
	}

	// Disable further Interrupts. This is a CPU behaviour when an interrupt is to be executed. So further interrupts
	// must be enabled by the program (Usually using RETI instruction when returning from an interrupt vector)
	Reg.IME = false

	// CPU will continue activity it was paused by a HALt instruction
	isHalt = false
	// TODO check if isHalt is true, follows after IE instruction and an interrupt is pending
	// In this case, simulate the halt bug by setting the PC back to the halt instruction

	// The handler should add five cycles as follows:
	// Two m-cyles before pushing PC
	// Two m-cycles after pushing PC
	// One m-cycle after setting handler vector
	if bus.InterruptPending() {
		ticks += 5
		push16(bit2.From16(Reg.PC.Get()))
	}
	if bus.IsVblank() && bus.IrqVblank() {
		bus.SetIrQVblank(false)

		Reg.PC.Set(0x40)
	} else if bus.IsLCDStat() && bus.IrqLCDStat() {
		bus.SetIRQLCDStat(false)
		Reg.PC.Set(0x48)
	} else if bus.IsTimerInt() && bus.IrqTimer() {
		bus.SetIRQTimer(false)
		Reg.PC.Set(0x50)
	} else if bus.IsSerialInt() && bus.IrqSerial() {
		bus.SetIrqSerial(false)
		Reg.PC.Set(0x58)
	} else if bus.IsJoypadInt() && bus.IrqJoypad() {
		bus.SetIrqJoypad(false)
		Reg.PC.Set(0x60)
	}

}

// Tick emulates machine ticks (m-ticks). Each m-tick is equivalent to four
// system tick
// Goes through a fetch-decode-execute cycle
// Returns Representation of performed instruction if completed, or empty
// string if in the middle of execution
func Tick() string {

	// Count is important for handling time. Therefore, it should not
	// exceed cpu frequency value
	if cycles == Frequency() {
		cycles = 0
	}

	// m-ticks needs to be finished before executing
	// the cycle again
	if ticks > 0 {
		advance()
		return ""
	}

	// TODO check if timer should be handled with each tick, instead with
	// the current simulated bulks of ticks
	timer()
	irq()

	// Check if halt was requested (using HALT operation)
	// Halt can be broken if an interrupt occurs
	if isHalt {
		advance()
		return ""
	}

	// Fetch instruction
	curOP = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	// Decode
	instruction := inst[curOP]

	// Execute Operation
	// TODO de-assemble and print executed operation
	ticks += instruction.ticks
	output := instruction.execute()

	// Emulates the IE instruction Delay
	if performIME == true && curOP != 0xFB {
		Reg.IME = true
		performIME = false
	}

	advance()

	return output
}

// advance a cpu tick
func advance() {
	ticks--
	cycles++
}
