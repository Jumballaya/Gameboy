package emulator

// Memory is an interface for dealing with anything that reads/writes bytes to a store
type Memory interface {
	ReadByte(int) int
	WriteByte(addr, val int)
}

// Renderer takes the pixel array and renders it
type Renderer interface {
	Render([]int)
	Pause()
	Stop()
	Start()
	Print() []int
}
