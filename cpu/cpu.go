package cpu

import (
	"github.com/aalquaiti/gbgo/io"
	"github.com/aalquaiti/gbgo/util"
)

const (
	INST_SIZE = 0x100      // Instruction Size
	DMG_HZ    = 0x100000   // Game Boy Frequency (m-ticks)
	CGB_HZ    = DMG_HZ * 2 // Game Boy Color Frequency (m-ticks)
	DIV_RATE  = 0xFFF      // Dividor Increment Rate
)

// Game Boy CPU Mode
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
	curOP = 0
}

// Step ticks a cpu until a op is executed
func Step() {
	steps++
	// currentPC := Reg.PC.Get()
	// output := Tick()
	Tick()

	for ticks > 0 {
		Tick()
	}

	// fmt.Printf("%04d  PC:%04X \t %s \n", steps, currentPC, output)
}

// timer handles functionality related to divider and timer
func timer() {
	// When count reaches Dividor rate
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
	var timeReached bool = cycles&timeRate == timeRate

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

// irq haanndles Interrupt request
func irq() {
	// Checks is Master Interrupt is enabled,
	// Ignores intterupts if disabled
	if !Reg.IME {
		return
	}

	// Disable further Intterupts. This is a CPU behaviour when an
	// interrupt is to be executed. So further interrupts must be enabled
	// by the program (Usually using RETI instruction when returning from
	// an interrupt vector)
	Reg.IME = false

	if bus.IsVblank() && bus.IrqVblank() {
		bus.SetIrQVblank(false)
		push16(util.From16(Reg.PC.Get()))
		Reg.PC.Set(0x40)
	} else if bus.IsLCDStat() && bus.IrqLCDStat() {
		bus.SetIRQLCDStat(false)
		push16(util.From16(Reg.PC.Get()))
		Reg.PC.Set(0x48)
	} else if bus.IsTimerInt() && bus.IrqTimer() {
		bus.SetIRQTimer(false)
		push16(util.From16(Reg.PC.Get()))
		Reg.PC.Set(0x50)
	} else if bus.IsSerialInt() && bus.IrqSerial() {
		bus.SetIrqSerial(false)
		push16(util.From16(Reg.PC.Get()))
		Reg.PC.Set(0x58)
	} else if bus.IsJoypadInt() && bus.IrqJoypad() {
		bus.SetIrqJoypad(false)
		push16(util.From16(Reg.PC.Get()))
		Reg.PC.Set(0x60)
	}

}

// Tick emulates machine ticks (m-ticks). Each m-tick is equivalent to four
// system tick
// Goes through a fetch-decode-execute cycle
// Returns Representation of performed instruction if completed, or empty
// string if in the middle of execution
func Tick() string {

	// Count is important for handling time. Therefore it should not
	// exceed cpu frequency value
	if cycles == CPU_FREQ() {
		cycles = 0
	}

	// m-ticks needs to be finished before executing
	// the cycle again
	if ticks > 0 {
		ticks--
		cycles++
		return ""
	}

	// TODO check if timer should be handled with each tick, instead with
	// the current simulated bulks of ticks
	timer()
	irq()

	// Fetch instruction
	curOP = bus.Read(Reg.PC.Get())
	Reg.PC.Inc()

	// Decode
	instruction := inst[curOP]

	// Execute Operation
	// TODO de-assemble and print executed operation
	ticks += instruction.ticks
	output := instruction.execute()

	timer()
	ticks--
	cycles++

	return output
}
