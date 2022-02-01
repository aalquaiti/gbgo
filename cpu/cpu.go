package cpu

import (
	"github.com/aalquaiti/gbgo/io"
)

const (
	DMG_HZ   = 0x100000   // Game Boy Frequency (m-ticks)
	CGB_HZ   = DMG_HZ * 2 // Game Boy Color Frequency (m-ticks)
	DIV_RATE = 0xFFF      // Divider Increment Rate
)

// Mode Game Boy CPU Mode
type Mode int

const (
	DMG_MODE Mode = iota + 1
	CGB_MODE
)

type CPU struct {
	mode   Mode
	bus    io.Bus
	ticks  uint8  // m-ticks remaining for an instruction
	cycles uint32 // m-ticks count. Helps with timer functionality
	steps  uint32 // Counts how many instructions executed
	Reg    Register
	flags  *RegF
	curOP  OpCode // Current Op to execute. Used in cpu fetch phase
	inst   [OPCodeSize]OpCode
	cbInst [OPCodeSize]OpCode

	// Variables that are used by CPU to help perform some of the instructions

	// IsHalt Used by HALT instruction to inform the cpu to halt until an interrupt occurs. This helps with interruption
	// handling that breaks the halt, and to emulate the HALT bug
	isHalt bool

	//  Used by EI instruction to set IME. The EI has a delay one of one cycle, so IME will be set to one after the
	// execution of the next instruction following EI.
	performIME bool

	// Variables that should be treated as immutable.
	// Access should be through functions

	freq uint32
}

// Init Initialise CPU
func NewCPU(mode Mode, bus io.Bus) *CPU {
	// TODO use different CPU mode
	cpu := &CPU{
		mode: mode,
		bus:  bus,
	}
	cpu.initInstructions()
	cpu.Reset()

	return cpu
}

// Reset resets CPU to pre-start state
func (c *CPU) Reset() {
	c.ticks = 0
	c.cycles = 0
	c.Reg = NewRegister()
	c.flags, _ = c.Reg.F.(*RegF)
	c.curOP = c.inst[0x00]
	c.isHalt = false

	// Power-Up Sequence for DMG
	// TODO Add power-up sequence of CGB
	c.Reg.A.Set(0x01)
	// TODO c.Reg F value depends on header checksum. Make amends according
	// Refer to https://gbdev.io/pandocs/Power_Up_Sequence.html
	c.Reg.F.Set(0xB0)
	c.Reg.B.Set(0x00)
	c.Reg.C.Set(0x13)
	c.Reg.D.Set(0x00)
	c.Reg.E.Set(0xD8)
	c.Reg.H.Set(0x01)
	c.Reg.L.Set(0x4D)
	c.Reg.PC.Set(0x100)
	c.Reg.SP.Set(0xFFFE)
}

// Step ticks a cpu until an op is executed
func (c *CPU) Step() {
	c.steps++
	//currentPC := c.Reg.PC.Get()
	//af := c.Reg.AF.String()
	//bc := c.Reg.BC.String()
	//de := c.Reg.DE.String()
	//hl := c.Reg.HL.String()
	//pc := c.Reg.PC.String()
	//sp := c.Reg.SP.String()
	//output := Tick()
	c.Tick()

	for c.ticks > 0 {
		c.Tick()
	}

	//fmt.Printf("%.04X:\t%-30s %s, %s, %s, %s, %s, %s\n", currentPC, output, bc, de, hl, af, sp, pc)
}

// timer handles functionality related to divider and timer
func (c *CPU) timer() {
	// When count reaches Divider rate
	if c.cycles&DIV_RATE == DIV_RATE {
		// TODO emulate CGB double speed effect
		c.bus.Time.IncDIV()
	}

	// When count reaches Timer rate, increment Timer Counter.
	// Timer Rate depends on Timer Control's bits 0 and 1. The following
	// shows frequency in m-ticks:
	// 00: CPU Clock / 0x1000 = 0x400 Hz
	// 01: CPU Clock / 0x10   = 0x10000 Hz
	// 10: CPU Clock / 0x40   = 0x4000 Hz
	// 11: CPU Clock / 0x100  = 0x1000 Hz
	timeRate := c.freq
	switch c.bus.Time.GetTacClockSelect() {
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
	timeReached := c.cycles&timeRate == timeRate

	if timeReached {
		// When Timer is Enabled and Timer Counter overflow,
		// set Timer counter to value stored in TMA and request a
		// Timer interrupt
		timaCount := c.bus.Read(io.AddrTima)
		if c.bus.Time.IsTacTimerEnabled() && timaCount == 0xFF {
			c.bus.Write(io.AddrTima, c.bus.Read(io.AddrTma))
			c.bus.IF.SetIRQTimer(true)
		} else {
			c.bus.Write(io.AddrTima, timaCount+1)
		}

	}
}

// irq handles Interrupt request
func (c *CPU) irq() {

	// Checks is Master Interrupt is enabled,
	// Ignores interrupts if disabled
	if !c.Reg.IME {

		// Having an interrupt pending breaks the halt loop if one was requested through HALT operation
		// See more details in halt() function
		if c.bus.InterruptPending() {
			// TODO simulate the halt bug by reading the instruction the follow HALT twice
			c.isHalt = false
		}

		return
	}

	// Disable further Interrupts. This is a CPU behaviour when an interrupt is to be executed. So further interrupts
	// must be enabled by the program (Usually using RETI instruction when returning from an interrupt vector)
	c.Reg.IME = false

	// CPU will continue activity it was paused by a HALt instruction
	c.isHalt = false
	// TODO check if isHalt is true, follows after IE instruction and an interrupt is pending
	// In this case, simulate the halt bug by setting the PC back to the halt instruction

	// The handler should add five cycles as follows:
	// Two m-cyles before pushing PC
	// Two m-cycles after pushing PC
	// One m-cycle after setting handler vector
	if c.bus.InterruptPending() {
		c.ticks += 5
		// TODO uncomment
		//push16(gbgoutil.From16(c.Reg.PC.Get()))
	}
	if c.bus.IE.IsVBlank() && c.bus.IF.IrqVBlank() {
		c.bus.IF.SetIrQVblank(false)

		c.Reg.PC.Set(0x40)
	} else if c.bus.IE.IsLCDStat() && c.bus.IF.IrqLCDStat() {
		c.bus.IF.SetIRQLCDStat(false)
		c.Reg.PC.Set(0x48)
	} else if c.bus.IE.IsTimerInt() && c.bus.IF.IrqTimer() {
		c.bus.IF.SetIRQTimer(false)
		c.Reg.PC.Set(0x50)
	} else if c.bus.IE.IsSerialInt() && c.bus.IF.IrqSerial() {
		c.bus.IF.SetIrqSerial(false)
		c.Reg.PC.Set(0x58)
	} else if c.bus.IE.IsJoypadInt() && c.bus.IF.IrqJoyPad() {
		c.bus.IF.SetIrqJoyPad(false)
		c.Reg.PC.Set(0x60)
	}

}

// Tick emulates machine ticks (m-ticks). Each m-tick is equivalent to four
// system tick
// Goes through a fetch-decode-execute cycle
// Returns Representation of performed instruction if completed, or empty
// string if in the middle of execution
func (c *CPU) Tick() {

	// Count is important for handling time. Therefore, it should not
	// exceed cpu frequency value
	if c.cycles == c.freq {
		c.cycles = 0
	}

	// m-ticks needs to be finished before executing
	// the cycle again
	if c.ticks > 0 {
		c.advance()
	}

	// TODO check if timer should be handled with each tick, instead with
	// the current simulated bulks of ticks
	c.timer()
	c.irq()

	// Check if halt was requested (using HALT operation)
	// Halt can be broken if an interrupt occurs
	if c.isHalt {
		c.advance()
	}

	// Fetch instruction
	c.curOP = c.inst[c.bus.Read(c.Reg.PC.Get())]
	c.Reg.PC.Inc()

	// Decode

	// Execute Operation
	// TODO de-assemble and print executed operation
	c.ticks += c.curOP.ticks
	// TODO uncomment
	//c.curOP.execute()

	// Emulates the IE instruction Delay
	if c.performIME == true && c.curOP.code != 0xFB {
		c.Reg.IME = true
		c.performIME = false
	}

	c.advance()
}

// advance a cpu tick
func (c *CPU) advance() {
	c.ticks--
	c.cycles++
}
