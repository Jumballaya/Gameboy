package emulator

// SplitBytes takes a 16 bit number and splits it into the upper and lower 8bit sections
func SplitWord(dByte int) (int, int) {
	upper := dByte >> 8
	lower := dByte & 0xff
	return upper, lower
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
func clearBit(val, addr int) int {
	return ^(1 << uint(addr)) & val & 0xff
}

// SetBit sets the bit value at the address or clears the value
func HandleSetBit(val, addr int, toSet bool) int {
	if toSet {
		return SetBit(val, addr)
	}
	return clearBit(val, addr)
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
func AddSignedByteToWord(f *cpuFlags, word, signed int) int {
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
