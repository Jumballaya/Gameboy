package emulator

import "testing"

func TestNewCPU(t *testing.T) {
	gpu := newGPU()
	mmu := newMMU()
	cpu := newCPU(mmu, gpu)
	if cpu == nil {
		t.Fatalf("CPU was not able to be instantiated")
	}
}

func TestExecMovesPointerUp(t *testing.T) {
	gpu := newGPU()
	mmu := newMMU()
	cpu := newCPU(mmu, gpu)

	if cpu.Registers.GetPC() != 0 {
		t.Fatalf("CPU was incorrectly initialized, counter at %d, wanted: %d", cpu.Registers.GetPC(), 0)
	}

	cpu.Exec()

	if cpu.Registers.GetPC() != 1 {
		t.Fatalf("CPU exec function does not move pointer up, counter at %d, wanted: %d", cpu.Registers.GetPC(), 1)
	}
}
