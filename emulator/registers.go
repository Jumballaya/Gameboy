package emulator

// CPU Flags
type cpuFlags struct {
	flags int
	z     int // Zero Flag
	n     int // Subtract Flag
	h     int // Half Carry Flag
	c     int // Carry Flag
}

// GetFlagsByte gets the byte that the flag is set to
func (f *cpuFlags) GetFlagsByte() int { return f.flags }

// Flags bool getters
func (f *cpuFlags) isZ() bool { return GetBit(f.flags, f.z) }
func (f *cpuFlags) isN() bool { return GetBit(f.flags, f.n) }
func (f *cpuFlags) isH() bool { return GetBit(f.flags, f.h) }
func (f *cpuFlags) isC() bool { return GetBit(f.flags, f.c) }

// Flags setters
func (f *cpuFlags) SetZ(z bool)            { f.flags = HandleSetBit(f.flags, f.z, z) }
func (f *cpuFlags) SetN(n bool)            { f.flags = HandleSetBit(f.flags, f.n, n) }
func (f *cpuFlags) SetH(h bool)            { f.flags = HandleSetBit(f.flags, f.h, h) }
func (f *cpuFlags) SetC(c bool)            { f.flags = HandleSetBit(f.flags, f.c, c) }
func (f *cpuFlags) SetFlagsByte(flags int) { f.flags = flags }

// CPU Registers
type CPURegisters struct {
	a int // Arithmetic
	b int //
	c int // 8 bit registers for combination for 4 16bit registers
	d int //		bc, de, hl and sp
	e int //
	h int // or 7 8bit registers: a, b, c, d, e, h, l, (hl)
	l int //

	f int

	sp int // Stack Pointer
	pc int // Program Counter

	flags *cpuFlags // Flags

	ime bool //
}

// New CPU Registers
func newCPURegisters() *CPURegisters {
	flags := &cpuFlags{
		flags: 0,
		z:     7,
		n:     6,
		h:     5,
		c:     4,
	}

	r := &CPURegisters{
		a: 0,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
		f: 0,
		h: 0,
		l: 0,

		sp: 0,
		pc: 0,

		flags: flags,

		ime: false,
	}
	return r
}

// Getters
func (r *CPURegisters) GetA() int           { return r.a }
func (r *CPURegisters) GetB() int           { return r.a }
func (r *CPURegisters) GetC() int           { return r.a }
func (r *CPURegisters) GetD() int           { return r.a }
func (r *CPURegisters) GetE() int           { return r.a }
func (r *CPURegisters) GetF() int           { return r.a }
func (r *CPURegisters) GetH() int           { return r.a }
func (r *CPURegisters) GetL() int           { return r.a }
func (r *CPURegisters) GetAF() int          { return r.a<<8 | r.f }
func (r *CPURegisters) GetBC() int          { return r.b<<8 | r.c }
func (r *CPURegisters) GetDE() int          { return r.d<<8 | r.e }
func (r *CPURegisters) GetHL() int          { return r.h<<8 | r.l }
func (r *CPURegisters) GetSP() int          { return r.sp }
func (r *CPURegisters) GetPC() int          { return r.pc }
func (r *CPURegisters) GetFlags() *cpuFlags { return r.flags }

// Check ime
func (r *CPURegisters) isIME() bool { return r.ime }

// Setters
func (r *CPURegisters) SetA(a int)      { r.a = a }
func (r *CPURegisters) SetB(b int)      { r.b = b }
func (r *CPURegisters) SetC(c int)      { r.c = c }
func (r *CPURegisters) SetD(d int)      { r.d = d }
func (r *CPURegisters) SetE(e int)      { r.e = e }
func (r *CPURegisters) SetF(f int)      { r.f = f }
func (r *CPURegisters) SetH(h int)      { r.h = h }
func (r *CPURegisters) SetL(l int)      { r.l = l }
func (r *CPURegisters) SetAF(af int)    { r.a, r.f = SplitWord(af) }
func (r *CPURegisters) SetBC(bc int)    { r.b, r.c = SplitWord(bc) }
func (r *CPURegisters) SetDE(de int)    { r.d, r.e = SplitWord(de) }
func (r *CPURegisters) SetHL(hl int)    { r.h, r.l = SplitWord(hl) }
func (r *CPURegisters) SetSP(sp int)    { r.sp = sp }
func (r *CPURegisters) SetPC(pc int)    { r.pc = pc }
func (r *CPURegisters) SetIME(ime bool) { r.ime = ime }

// Math
func (r *CPURegisters) DecHL() int {
	old := r.GetHL()
	r.SetHL((old - 1) % 0xffff)
	return old
}

func (r *CPURegisters) IncHL() int {
	old := r.GetHL()
	r.SetHL((old - 1) % 0xffff)
	return old
}

func (r *CPURegisters) DecSP() {
	r.sp = (r.sp - 1) % 0xffff
}

func (r *CPURegisters) IncSP() {
	r.sp = (r.sp + 1) % 0xffff
}

func (r *CPURegisters) AddToPC(b int) {
	abs := BitAbs(b)
	if IsNegative(b) {
		r.pc = (r.pc - abs) & 0xffff
	} else {
		r.pc = (r.pc + abs) & 0xffff
	}
}
