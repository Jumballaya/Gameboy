package emulator

// Memory is an interface for dealing with anything that reads/writes bytes to a store
type Memory interface {
	ReadByte(int) int
	WriteByte(addr, val int)
}
