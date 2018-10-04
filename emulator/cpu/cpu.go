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
	ops       map[string]*Oplist
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

	main, extern := makeOpList(cpu.Registers)
	cpu.ops = map[string]*Oplist{
		"main":   main,
		"extern": extern,
	}

	return cpu
}

// Reset sets all the values back to their starting positions
func (cpu *CPU) Reset() {
	cpu.Registers = newRegisters()
	main, extern := makeOpList(cpu.Registers)
	cpu.ops = map[string]*Oplist{
		"main":   main,
		"extern": extern,
	}
	cpu.Clock = 0
	cpu.Halt = 0
	cpu.Stop = 0
}

// Exec executes a single cycle of the CPU
func (cpu *CPU) Exec() int {
	pc := cpu.Registers.GetPC()
	opCode := cpu.mmu.ReadByte(pc)
	pc += 1

	var op Op
	if opCode == 0xcb {
		opCode = cpu.mmu.ReadByte(pc)
		pc += 1
		op = cpu.ops["extern"].fns[opCode]
	} else {
		op = cpu.ops["main"].fns[opCode]
	}

	argsLen := op.argsLen
	var args []int
	for i := 0; i < argsLen; i++ {
		args = append(args, cpu.mmu.ReadByte(pc))
		pc += 1
	}

	if op.label != "NOP" {
		fmt.Println(op.label)
	}

	cpu.Registers.SetPC(pc)
	op.fn(cpu.Registers, cpu.mmu, args)

	return op.cycles
}
