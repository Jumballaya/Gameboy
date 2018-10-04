package core

import (
	"github.com/jumballaya/gameboy/emulator/cpu"
	"github.com/jumballaya/gameboy/emulator/gpu"
	"github.com/jumballaya/gameboy/emulator/mmu"
)

// Emulator structure
type Emulator struct {
	cpu *cpu.CPU
	mmu *mmu.MMU
	gpu *gpu.GPU
}

// New creates a new Emulator
func New() *Emulator {
	mmu := mmu.New()
	gpu := gpu.New()
	cpu := cpu.New(mmu, gpu)

	e := &Emulator{
		cpu: cpu,
		mmu: mmu,
		gpu: gpu,
	}
	e.Reset()

	return e
}

// Reset resets everything
func (e *Emulator) Reset() {
	e.cpu.Reset()
	e.gpu.Reset()
	e.mmu.Reset()
}

// Load loads the ROM data
func (e *Emulator) Load(rom []byte) error {
	e.mmu.Load(rom)
	return nil
}

// Start starts the emulator and runs the ROM
func (e *Emulator) Start() {
	for {
		e.cpu.Exec()
	}
}

// Pause pauses the emulator
func (e *Emulator) Pause() {}
