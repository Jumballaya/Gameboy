package emulator

// SplitBytes takes a 16 bit number and splits it into the upper and lower 8bit sections
func SplitBytes(dByte int) (int, int) {
	upper := dByte >> 8
	lower := dByte & 0xff
	return upper, lower
}

// Combine bytes takes 2 8bit numbers and combines them into a 16bit number
func CombineBytes(upper, lower int) int {
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

// BitAbs is the absolute value implementation for signed byte values
func BitAbs(signed int) int {
	if IsNegative(signed) {
		return 0x100 - signed
	}
	return signed
}

// SetBit sets the bit value at the address or clears the value
func HandleSetBit(val, addr int, toSet bool) int {
	if toSet {
		return SetBit(val, addr)
	}
	return clearBit(val, addr)
}
