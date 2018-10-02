package emulator

// Default Registers
var baseRegisters = map[string]int{
	"a":  0,
	"b":  0,
	"c":  0,
	"d":  0,
	"e":  0,
	"h":  0,
	"l":  0,
	"f":  0,
	"sp": 0,
	"pc": 0,
	"m":  0,
}

// Oplist is a map of functions (operations) for the CPU
type Oplist map[string]func()

// CPU is the emulated Z80 processor
type CPU struct {
	Registers map[string]int
	Clock     int
	Halt      int
	Stop      int
	ops       Oplist
}

// NewCPU creates a new CPU
func newCPU() *CPU {
	return &CPU{
		Registers: baseRegisters,
		Clock:     0,
		Halt:      0,
		Stop:      0,
		ops:       map[string]func(){},
	}
}

// Reset sets all the values back to their starting positions
func (cpu *CPU) Reset() {
	cpu.Registers = baseRegisters
	cpu.Clock = 0
	cpu.Halt = 0
	cpu.Stop = 0
}

// Exec executes a single cycle of the CPU
func (cpu *CPU) Exec() {}
