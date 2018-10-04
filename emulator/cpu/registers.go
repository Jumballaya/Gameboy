package cpu

// CPU Flags
type Flags struct {
	flags int
	z     int // Zero Flag
	n     int // Subtract Flag
	h     int // Half Carry Flag
	c     int // Carry Flag
}

// GetFlagsByte gets the byte that the flag is set to
func (f *Flags) GetFlagsByte() int { return f.flags }

// Flags bool getters
func (f *Flags) IsZ() bool { return GetBit(f.flags, f.z) }
func (f *Flags) IsN() bool { return GetBit(f.flags, f.n) }
func (f *Flags) IsH() bool { return GetBit(f.flags, f.h) }
func (f *Flags) IsC() bool { return GetBit(f.flags, f.c) }

// Flags setters
func (f *Flags) SetZ(z bool)            { f.flags = HandleSetBit(f.flags, f.z, z) }
func (f *Flags) SetN(n bool)            { f.flags = HandleSetBit(f.flags, f.n, n) }
func (f *Flags) SetH(h bool)            { f.flags = HandleSetBit(f.flags, f.h, h) }
func (f *Flags) SetC(c bool)            { f.flags = HandleSetBit(f.flags, f.c, c) }
func (f *Flags) SetFlagsByte(flags int) { f.flags = flags }

// CPU Registers
type Registers struct {
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

	flags *Flags // Flags

	ime bool //
}

// New CPU Registers
func newRegisters() *Registers {
	flags := &Flags{
		flags: 0,
		z:     7,
		n:     6,
		h:     5,
		c:     4,
	}

	r := &Registers{
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
func (r *Registers) GetA() int        { return r.a }
func (r *Registers) GetB() int        { return r.a }
func (r *Registers) GetC() int        { return r.a }
func (r *Registers) GetD() int        { return r.a }
func (r *Registers) GetE() int        { return r.a }
func (r *Registers) GetF() int        { return r.a }
func (r *Registers) GetH() int        { return r.a }
func (r *Registers) GetL() int        { return r.a }
func (r *Registers) GetAF() int       { return r.a<<8 | r.f }
func (r *Registers) GetBC() int       { return r.b<<8 | r.c }
func (r *Registers) GetDE() int       { return r.d<<8 | r.e }
func (r *Registers) GetHL() int       { return r.h<<8 | r.l }
func (r *Registers) GetSP() int       { return r.sp }
func (r *Registers) GetPC() int       { return r.pc }
func (r *Registers) GetFlags() *Flags { return r.flags }

// Check ime
func (r *Registers) isIME() bool { return r.ime }

// Setters
func (r *Registers) SetA(a int)      { r.a = a }
func (r *Registers) SetB(b int)      { r.b = b }
func (r *Registers) SetC(c int)      { r.c = c }
func (r *Registers) SetD(d int)      { r.d = d }
func (r *Registers) SetE(e int)      { r.e = e }
func (r *Registers) SetF(f int)      { r.f = f }
func (r *Registers) SetH(h int)      { r.h = h }
func (r *Registers) SetL(l int)      { r.l = l }
func (r *Registers) SetAF(af int)    { r.a, r.f = SplitWord(af) }
func (r *Registers) SetBC(bc int)    { r.b, r.c = SplitWord(bc) }
func (r *Registers) SetDE(de int)    { r.d, r.e = SplitWord(de) }
func (r *Registers) SetHL(hl int)    { r.h, r.l = SplitWord(hl) }
func (r *Registers) SetSP(sp int)    { r.sp = sp }
func (r *Registers) SetPC(pc int)    { r.pc = pc }
func (r *Registers) SetIME(ime bool) { r.ime = ime }

// Math
func (r *Registers) DecHL() int {
	old := r.GetHL()
	r.SetHL((old - 1) % 0xffff)
	return old
}

func (r *Registers) IncHL() int {
	old := r.GetHL()
	r.SetHL((old + 1) % 0xffff)
	return old
}

func (r *Registers) DecSP() {
	r.sp = (r.sp - 1) % 0xffff
}

func (r *Registers) IncSP() {
	r.sp = (r.sp + 1) % 0xffff
}

func (r *Registers) AddToPC(b int) {
	abs := BitAbs(b)
	if IsNegative(b) {
		r.pc = (r.pc - abs) & 0xffff
	} else {
		r.pc = (r.pc + abs) & 0xffff
	}
}
