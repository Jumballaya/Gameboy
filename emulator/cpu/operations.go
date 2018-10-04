package cpu

import (
	"fmt"

	"github.com/jumballaya/gameboy/emulator"
)

// Operation function
type OpFn func(*CPURegisters, emulator.Memory, []int) // Takes Registers, MMU and the arguments list

// Operation
type Op struct {
	code    int
	cycles  int
	argsLen int
	label   string
	fn      OpFn
}

// String returns a formatted string representation of the Op
func (o *Op) String() string {
	return fmt.Sprintf("%d\t%s", o.code, o.label)
}

// Oplist is a map of functions (operations) for the CPU
type Oplist struct {
	fns       map[int]Op
	registers *CPURegisters
}

// SetRegisters sets the registers
func (ol *Oplist) SetRegisters(r *CPURegisters) {
	ol.registers = r
}

// String prints a string representation of the CPU
func (ol *Oplist) String() string {
	out := ""
	for _, op := range ol.fns {
		out += "\n" + op.String()
	}
	return out
}
