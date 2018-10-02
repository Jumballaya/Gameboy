package emulator

// Oplist is a map of functions (operations) for the CPU
type Oplist map[string]func()

// CPU is the emulated Z80 processor
type CPU struct {
	Registers *CPURegisters
	Clock     int
	Halt      int
	Stop      int
	ops       *Oplist
	mmu       *MMU
	gpu       *GPU
}

// NewCPU creates a new CPU
func newCPU(mmu *MMU, gpu *GPU) *CPU {
	return &CPU{
		Registers: newCPURegisters(),
		Clock:     0,
		Halt:      0,
		Stop:      0,
		ops:       &Oplist{},
		gpu:       gpu,
		mmu:       mmu,
	}
}

// Reset sets all the values back to their starting positions
func (cpu *CPU) Reset() {
	cpu.Registers = newCPURegisters()
	cpu.Clock = 0
	cpu.Halt = 0
	cpu.Stop = 0
}

// Exec executes a single cycle of the CPU
func (cpu *CPU) Exec() {
	pc := cpu.Registers.GetPC()
	cpu.Registers.AddToPC(1)
	cpu.Registers.SetPC(cpu.Registers.GetPC() & 65535)
	//cpu.Clock += cpu.Registers.getM()
	opCode := cpu.mmu.ReadByte(pc)
}
