package emulator

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := New()
	if cpu == nil {
		t.Fatalf("CPU was not able to be instantiated")
	}
}
