package emulator

// Operation function
type OpFn func(*CPURegisters, *MMU, []int) // Takes Registers, MMU and the arguments list

// Operation
type Op struct {
	code    int
	cycles  int
	argsLen int
	label   string
	fn      OpFn
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
