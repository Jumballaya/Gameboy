package mmu

// Memory Management Unit struct
type MMU struct {
	bios     []int
	rom      []byte
	cartType int

	mbc       string // Need to figure out a real structure for this (memory bank controller)
	romOffset int
	ramOffset int

	wram [8192]int
	eram [32768]int
	zram [127]int

	inBios int
}

// New creates a new MMU
func New() *MMU {
	mmu := &MMU{
		bios: biosClassic,
	}
	mmu.Reset()
	return mmu
}

// Reset sets all the values back to their starting positions
func (mmu *MMU) Reset() {
	// Set WRAM
	for i := 0; i < 8192; i++ {
		mmu.wram[i] = 0
	}

	// Set ERAM
	for i := 0; i < 32768; i++ {
		mmu.eram[i] = 0
	}

	// Set ZRAM
	for i := 0; i < 127; i++ {
		mmu.zram[i] = 0
	}

	mmu.inBios = 1
	mmu.cartType = 0
	mmu.romOffset = 0x4000
	mmu.ramOffset = 0
}

// Load is basically the same as inserting a cartridge into a gameboy
func (mmu *MMU) Load(rom []byte) {
	mmu.rom = rom
}

// ReadByte reads the memory at the given byte address
func (mmu *MMU) ReadByte(addr int) int { return 0 }

// ReadWord acts like ReadByte except it takes a 16bit address and turns it into 2 8bit ReadByte calls
func (mmu *MMU) ReadWord(addr int) int {
	return mmu.ReadByte(addr) + (mmu.ReadByte(addr+1) << 8)
}

// WriteByte is the opposite of ReadByte and writes the val at the given address
func (mmu *MMU) WriteByte(addr, val int) {}

// WriteWord is similar to ReadWord but with the WriteByte function
func (mmu *MMU) WriteWord(addr, val int) {
	mmu.WriteByte(addr, val&255)
	mmu.WriteByte(addr+1, val>>8)
}
