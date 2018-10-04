package mmu

import "github.com/jumballaya/gameboy/emulator"

// Memory Management Unit struct
type MMU struct {
	bios     []int
	rom      []byte
	cartType int

	gpu emulator.Memory

	mbc       string // Need to figure out a real structure for this (memory bank controller)
	romOffset int
	ramOffset int

	wram [8192]int
	eram [32768]int
	zram [127]int

	inBios int
}

// New creates a new MMU
func New(gpu emulator.Memory) *MMU {
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
	mmu.cartType = int(rom[0x0147])
	mmu.rom = rom
}

// ReadByte reads the memory at the given byte address
func (mmu *MMU) ReadByte(addr int) int {
	// Read the correct memory location

	// BIOS
	if addr >= 0x00 && addr <= 0xff {
		return mmu.bios[addr]
	}

	// ROM Bank #0
	if addr < 0x4000 {
		return int(mmu.rom[addr])
	}

	// Switchable ROM Bank
	if addr < 0x8000 {
		return int(mmu.rom[addr])
	}

	// VRAM 8k
	if addr < 0xa000 {
		return mmu.gpu.ReadByte(addr)
	}

	// Switch RAM bank 8k
	if addr < 0xc000 {
		return mmu.zram[addr]
	}

	// Internal RAM 8k
	if addr < 0xe000 {
		return mmu.wram[addr]
	}

	// Echo RAM 8k
	if addr < 0xfe00 {
		// The addresses 0xe000 - 0xf000 appear to access the internal RAM the same as 0xc000 - 0xde00.
		// Writing a byte to 0xe000 will appear at 0xc000 and 0xe000 and vice-versa
		return mmu.eram[addr]
	}

	// 0xfea0

	// 0xff00

	// 0xff4c

	// 0xff80

	// 0xffff

	return int(mmu.rom[addr])
}

// ReadWord acts like ReadByte except it takes a 16bit address and turns it into 2 8bit ReadByte calls
func (mmu *MMU) ReadWord(addr int) int {
	return mmu.ReadByte(addr) + (mmu.ReadByte(addr+1) << 8)
}

// WriteByte is the opposite of ReadByte and writes the val at the given address
func (mmu *MMU) WriteByte(addr, val int) {
	// Write to the correct memory location
	mmu.eram[addr%len(mmu.eram)] = val
}

// WriteWord is similar to ReadWord but with the WriteByte function
func (mmu *MMU) WriteWord(addr, val int) {
	mmu.WriteByte(addr, val&255)
	mmu.WriteByte(addr+1, val>>8)
}
