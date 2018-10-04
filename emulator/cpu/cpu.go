package cpu

import (
	"fmt"

	"github.com/jumballaya/gameboy/emulator"
)

// CPU is the emulated Z80 processor
type CPU struct {
	Registers *Registers
	Clock     int
	Halt      int
	Stop      int
	ops       *Oplist
	mmu       emulator.Memory
	gpu       emulator.Memory
}

// New creates a new CPU
func New(mmu emulator.Memory, gpu emulator.Memory) *CPU {
	cpu := &CPU{
		Registers: newRegisters(),
		Clock:     0,
		Halt:      0,
		Stop:      0,
		gpu:       gpu,
		mmu:       mmu,
	}
	cpu.ops, _ = makeOpList(cpu.Registers)

	return cpu
}

// Reset sets all the values back to their starting positions
func (cpu *CPU) Reset() {
	cpu.Registers = newRegisters()
	cpu.ops, _ = makeOpList(cpu.Registers)
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

	fmt.Println(opCode)
}
