package emulator

// Emulator structure
type Emulator struct {
	cpu *CPU
	mmu *MMU
	gpu *GPU
}

// New creates a new Emulator
func New() *Emulator {
	cpu := newCPU()
	mmu := newMMU()
	gpu := newGPU()

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
func (e *Emulator) Start() {}

// Pause pauses the emulator
func (e *Emulator) Pause() {}
