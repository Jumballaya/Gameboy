package cpu

import "github.com/jumballaya/gameboy/emulator"

// SplitBytes takes a 16 bit number and splits it into the upper and lower 8bit sections
func SplitWord(word int) (int, int) {
	upper := word >> 8
	lower := word & 0xff
	return upper, lower
}

// UpperByte gets the upper byte from the 16bit word
func UpperByte(word int) int {
	return word >> 8
}

// LowerByte gets the lower byte from the 16bit word
func LowerByte(word int) int {
	return word & 0xff
}

// Combine bytes takes 2 8bit numbers and combines them into a 16bit number
func ToWord(upper, lower int) int {
	return (upper << 8) | lower
}

// Check if bit byte is negative
func IsNegative(signed int) bool {
	return (signed & (1 << 7)) != 0
}

// GetBit returns the bit at the location
func GetBit(val, addr int) bool {
	return (val & (1 << uint(addr))) != 0
}

// Internal version of SetBit
func SetBit(val, addr int) int {
	return (val | (1 << uint(addr))) & 0xff
}

// Clear Bit clears the addr
func ClearBit(val, addr int) int {
	return ^(1 << uint(addr)) & val & 0xff
}

// SetBit sets the bit value at the address or clears the value
func HandleSetBit(val, addr int, toSet bool) int {
	if toSet {
		return SetBit(val, addr)
	}
	return ClearBit(val, addr)
}

// BitAbs is the absolute value implementation for signed byte values
func BitAbs(signed int) int {
	if IsNegative(signed) {
		return 0x100 - signed
	}
	return signed
}

// AddSignedByteToWord adds a signed byte to a 16bit (2byte) word
// Answer from: https://stackoverflow.com/questions/5159603/gbz80-how-does-ld-hl-spe-affect-h-and-c-flags/7261149#7261149
func AddSignedByteToWord(f *Flags, word, signed int) int {
	f.SetZ(false)
	f.SetN(false)

	b := BitAbs(signed)

	if IsNegative(signed) {
		f.SetH((word & 0x0f) < (b & 0x0f))
		f.SetC((word & 0xff) < b)
		return (word - b) % 0xffff
	}
	f.SetC((word&0xff)+b > 0xff)
	f.SetH((word&0x0f)+(b&0x0f) > 0x0f)
	return (word + b) % 0xffff
}

// AddBytes adds 2 bytes togethers and sets the correct flags before returning the result
func AddBytes(f *Flags, b1, b2 int) int {
	f.SetZ(((b1 + b2) & 0xff) == 0)
	f.SetN(false)
	f.SetH((b1&0x0f)+(b2&0x0f) > 0x0f)
	f.SetC(b1+b2 > 0xff)
	return (b1 + b2) & 0xff
}

// AddBytesAndCarry is the same as AddBytes with respect to the carry flag (Flags.c)
func AddBytesAndCarry(f *Flags, b1, b2 int) int {
	carry := 0
	if f.IsC() {
		carry = 1
	}

	sum := b1 + b2 + carry

	f.SetZ((sum & 0xff) == 0)
	f.SetN(false)
	f.SetH((b1&0x0f)+(b2&0x0f)+carry > 0x0f)
	f.SetC(sum > 0xff)
	return (sum) & 0xff
}

// AddWords is the same as AddBytes but with 2 16bit (2bytes) words
func AddWords(f *Flags, w1, w2 int) int {
	f.SetN(false)
	f.SetH((w1&0x0fff)+(w2&0x0fff) > 0x0fff)
	f.SetC(w1+w2 > 0xffff)
	return (w1 + w2) & 0xffff
}

// SubBytes takes the difference between two bytes and sets the flags before returning the result
func SubBytes(f *Flags, b1, b2 int) int {
	f.SetZ(((b1 - b2) & 0xff) == 0)
	f.SetN(true)
	f.SetH((0x0f & b2) > (0x0f & b1))
	f.SetC(b2 > b1)
	return (b1 - b2) % 0xff
}

// SubBytesAndCarry is the same as SubBytes with respect to the carry flag (Flags.c)
func SubBytesAndCarry(f *Flags, b1, b2 int) int {
	carry := 0
	if f.IsC() {
		carry = 1
	}

	diff := b1 - b2 - carry

	f.SetZ((diff & 0xff) == 0)
	f.SetN(true)
	f.SetH((0x0f & (b2 + carry)) > (0x0f & b1))
	f.SetC(b2+carry > b1)
	return diff % 0xff
}

// And runs the & operator between two bytes, sets the flags before returning the value
func And(f *Flags, b1, b2 int) int {
	res := b1 & b2
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(true)
	f.SetC(false)
	return res
}

// Or runs the | operator between two bytes, sets the flags before returning the value
func Or(f *Flags, b1, b2 int) int {
	res := b1 | b2
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	f.SetC(false)
	return res
}

// Xor runs the ^ operator between two bytes, sets the flags before returning the value
func Xor(f *Flags, b1, b2 int) int {
	res := b1 ^ b2
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	f.SetC(false)
	return res
}

// IncByte increments the byte and sets the flags before returning the value
func IncByte(f *Flags, b int) int {
	res := (b + 1) & 0xff
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH((0x0f & res) < (0x0f & b))
	return res
}

// DecByte decrements the byte and sets the flags before returning the value
func DecByte(f *Flags, b int) int {
	res := (b - 1) & 0xff
	f.SetZ(res == 0)
	f.SetN(true)
	f.SetH((0x0f & b) == 0)
	return res
}

// Swap swaps the upper and lower bits of the byte, sets flags before returning swapped value
func Swap(f *Flags, b int) int {
	upper := b & 0xf0
	lower := b & 0x0f
	res := (lower << 4) | (upper >> 4)

	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	f.SetC(false)
	return res
}

// RotateRight uses the >> operator, sets the flags before returning the value
func RotateRight(f *Flags, b int) int {
	res := b >> 1
	if (b & 1) == 1 {
		res |= (1 << 7)
		f.SetC(true)
	} else {
		f.SetC(false)
	}

	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// RotateRightThroughCarry uses the >> operater, set the flags with respect to the
// carry flag before returning the value
func RotateRightThroughCarry(f *Flags, b int) int {
	res := b >> 1
	if f.IsC() {
		res |= (1 << 7)
	} else {
		res |= 0
	}

	f.SetC((b & 1) != 0)
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// RotateLeft uses the << operator, sets the flags before returning the value
func RotateLeft(f *Flags, b int) int {
	res := (b << 1) & 0xff
	if b&(1<<7) != 0 {
		res |= 1
		f.SetC(true)
	} else {
		f.SetC(false)
	}

	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// RotateLeftThroughCarry like RotateLeft with respect to the cary flag
func RotateLeftThroughCarry(f *Flags, b int) int {
	res := (b << 1) & 0xff
	if f.IsC() {
		res |= 1
	} else {
		res |= 0
	}

	f.SetC((b & (1 << 7)) != 0)
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// ShiftLeft uses the << operator and sets flags
func ShiftLeft(f *Flags, b int) int {
	res := (b << 1) & 0xff
	f.SetC((b & (1 << 7)) != 0)
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// ShiftRightArithmetic
func ShiftRightArithmetic(f *Flags, b int) int {
	res := (b >> 1) | (b & (1 << 7))

	f.SetC((b & 1) != 0)
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// ShiftRightLogical
func ShiftRightLogical(f *Flags, b int) int {
	res := b >> 1

	f.SetC((b & 1) != 0)
	f.SetZ(res == 0)
	f.SetN(false)
	f.SetH(false)
	return res
}

// Bit sets flags based on the given byte value and bit
func Bit(f *Flags, val, b int) {
	f.SetN(false)
	f.SetH(true)
	if b < 8 {
		f.SetZ((val & (1 << uint(b))) != 0)
	}
}

// Push pushes a word to the top of the stack
func Push(r *Registers, m emulator.Memory, word int) {
	r.DecSP()
	m.WriteByte(r.GetSP(), UpperByte(word))
	r.DecSP()
	m.WriteByte(r.GetSP(), LowerByte(word))
}

// Pop returns the top word from the stack
func Pop(r *Registers, m emulator.Memory) int {
	lower := m.ReadByte(r.GetSP())
	r.IncSP()
	upper := m.ReadByte(r.GetSP())
	return ToWord(upper, lower)
}

// Call calls the given address
func Call(r *Registers, m emulator.Memory, addr int) {
	Push(r, m, (r.GetPC()+3)&0xffff)
	r.SetPC(addr)
}

// Reset resets the value of PC to the given address
func Reset(r *Registers, m emulator.Memory, addr int) {
	Push(r, m, r.GetPC())
	r.SetPC(addr)
}

// Ret sets the PC to the top of the stack
func Ret(r *Registers, m emulator.Memory) {
	r.SetPC(Pop(r, m))
}
