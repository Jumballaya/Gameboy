package emulator

func makeOpList(r *CPURegisters) *Oplist {
	// Main OpList
	CPUOpList := &Oplist{
		registers: r,
		fns: map[int]Op{

			// Load Commands

			// LD B, n
			0x06: Op{
				code:    0x06,
				cycles:  8,
				label:   "LD B, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(args[0]) },
			},

			// LD C, n
			0x0e: Op{
				code:    0x0e,
				cycles:  8,
				label:   "LD C, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(args[0]) },
			},

			// LD D, n
			0x16: Op{
				code:    0x16,
				cycles:  8,
				label:   "LD D, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(args[0]) },
			},

			// LD E, n
			0x1e: Op{
				code:    0x1e,
				cycles:  8,
				label:   "LD E, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(args[0]) },
			},

			// LD H, n
			0x26: Op{
				code:    0x26,
				cycles:  8,
				label:   "LD H, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(args[0]) },
			},

			// LD L, n
			0x2e: Op{
				code:    0x2e,
				cycles:  8,
				label:   "LD L, n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(args[0]) },
			},

			// LD B, B
			0x40: Op{
				code:    0x40,
				cycles:  4,
				label:   "LD B, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetB()) },
			},

			// LD B, C
			0x41: Op{
				code:    0x41,
				cycles:  4,
				label:   "LD B, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetC()) },
			},

			// LD B, D
			0x42: Op{
				code:    0x42,
				cycles:  4,
				label:   "LD B, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetD()) },
			},

			// LD B, E
			0x43: Op{
				code:    0x43,
				cycles:  4,
				label:   "LD B, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetE()) },
			},

			// LD B, H
			0x44: Op{
				code:    0x44,
				cycles:  4,
				label:   "LD B, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetH()) },
			},

			// LD B, L
			0x45: Op{
				code:    0x45,
				cycles:  4,
				label:   "LD B, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetL()) },
			},

			// LD B, (HL)
			0x46: Op{
				code:    0x46,
				cycles:  8,
				label:   "LD B, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(m.ReadByte(r.GetHL())) },
			},

			// LD C, B
			0x48: Op{
				code:    0x48,
				cycles:  4,
				label:   "LD C, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetB()) },
			},

			// LD C, C
			0x49: Op{
				code:    0x49,
				cycles:  4,
				label:   "LD C, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetC()) },
			},

			// LD C, D
			0x4a: Op{
				code:    0x4a,
				cycles:  4,
				label:   "LD C, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetD()) },
			},

			// LD C, E
			0x4b: Op{
				code:    0x4b,
				cycles:  4,
				label:   "LD C, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetE()) },
			},

			// LD C, H
			0x4c: Op{
				code:    0x4c,
				cycles:  4,
				label:   "LD C, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetH()) },
			},

			// LD C, L
			0x4d: Op{
				code:    0x4d,
				cycles:  4,
				label:   "LD C, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetL()) },
			},

			// LD C, (HL)
			0x4e: Op{
				code:    0x4e,
				cycles:  8,
				label:   "LD C, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(m.ReadByte(r.GetHL())) },
			},

			// LD D, B
			0x50: Op{
				code:    0x50,
				cycles:  4,
				label:   "LD D, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetB()) },
			},

			// LD D, C
			0x51: Op{
				code:    0x51,
				cycles:  4,
				label:   "LD D, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetC()) },
			},

			// LD D, D
			0x52: Op{
				code:    0x52,
				cycles:  4,
				label:   "LD D, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetD()) },
			},

			// LD D, E
			0x53: Op{
				code:    0x53,
				cycles:  4,
				label:   "LD D, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetE()) },
			},

			// LD D, H
			0x54: Op{
				code:    0x54,
				cycles:  4,
				label:   "LD D, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetH()) },
			},

			// LD D, L
			0x55: Op{
				code:    0x55,
				cycles:  4,
				label:   "LD D, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetL()) },
			},

			// LD D, (HL)
			0x56: Op{
				code:    0x56,
				cycles:  8,
				label:   "LD D, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(m.ReadByte(r.GetHL())) },
			},

			// LD E, B
			0x58: Op{
				code:    0x58,
				cycles:  4,
				label:   "LD E, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetB()) },
			},

			// LD E, C
			0x59: Op{
				code:    0x59,
				cycles:  4,
				label:   "LD E, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetC()) },
			},

			// LD E, D
			0x5a: Op{
				code:    0x5a,
				cycles:  4,
				label:   "LD E, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetD()) },
			},

			// LD E, E
			0x5b: Op{
				code:    0x5b,
				cycles:  4,
				label:   "LD E, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetE()) },
			},

			// LD E, H
			0x5c: Op{
				code:    0x5c,
				cycles:  4,
				label:   "LD E, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetH()) },
			},

			// LD E, L
			0x5d: Op{
				code:    0x5d,
				cycles:  4,
				label:   "LD E, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetL()) },
			},

			// LD E, (HL)
			0x5e: Op{
				code:    0x5e,
				cycles:  8,
				label:   "LD E, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(m.ReadByte(r.GetHL())) },
			},

			// LD H, B
			0x60: Op{
				code:    0x60,
				cycles:  4,
				label:   "LD H, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetB()) },
			},

			// LD H, C
			0x61: Op{
				code:    0x61,
				cycles:  4,
				label:   "LD H, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetC()) },
			},

			// LD H, D
			0x62: Op{
				code:    0x62,
				cycles:  4,
				label:   "LD H, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetD()) },
			},

			// LD H, E
			0x63: Op{
				code:    0x63,
				cycles:  4,
				label:   "LD H, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetE()) },
			},

			// LD H, H
			0x64: Op{
				code:    0x64,
				cycles:  4,
				label:   "LD H, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetH()) },
			},

			// LD H, L
			0x65: Op{
				code:    0x65,
				cycles:  4,
				label:   "LD H, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetL()) },
			},

			// LD H, (HL)
			0x66: Op{
				code:    0x66,
				cycles:  8,
				label:   "LD H, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(m.ReadByte(r.GetHL())) },
			},

			// LD L, B
			0x68: Op{
				code:    0x68,
				cycles:  4,
				label:   "LD L, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetB()) },
			},

			// LD L, C
			0x69: Op{
				code:    0x69,
				cycles:  4,
				label:   "LD L, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetC()) },
			},

			// LD L, D
			0x6a: Op{
				code:    0x6a,
				cycles:  4,
				label:   "LD L, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetD()) },
			},

			// LD L, E
			0x6b: Op{
				code:    0x6b,
				cycles:  4,
				label:   "LD L, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetE()) },
			},

			// LD L, H
			0x6c: Op{
				code:    0x6c,
				cycles:  4,
				label:   "LD L, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetH()) },
			},

			// LD L, L
			0x6d: Op{
				code:    0x6d,
				cycles:  4,
				label:   "LD L, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetL()) },
			},

			// LD L, (HL)
			0x6e: Op{
				code:    0x6e,
				cycles:  8,
				label:   "LD L, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(m.ReadByte(r.GetHL())) },
			},

			// LD (HL), B
			0x70: Op{
				code:    0x70,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), B",
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetB()) },
			},

			// LD (HL), C
			0x71: Op{
				code:    0x71,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), C",
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetC()) },
			},

			// LD (HL), D
			0x72: Op{
				code:    0x72,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), D",
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetD()) },
			},

			// LD (HL), E
			0x73: Op{
				code:    0x73,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), E",
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetE()) },
			},

			// LD (HL), H
			0x74: Op{
				code:    0x74,
				cycles:  8,
				label:   "(HL), B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetH()) },
			},

			// LD (HL), L
			0x75: Op{
				code:    0x75,
				cycles:  8,
				label:   "(HL), L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetL()) },
			},

			// LD (HL), n
			0x76: Op{
				code:    0x76,
				cycles:  12,
				label:   "(HL), n",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), args[0]) },
			},

			// LD A, A
			0x7f: Op{
				code:    0x7f,
				cycles:  4,
				label:   "LD A, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetA()) },
			},

			// LD A, B
			0x78: Op{
				code:    0x78,
				cycles:  4,
				label:   "LD A, B",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetB()) },
			},

			// LD A, C
			0x79: Op{
				code:    0x79,
				cycles:  4,
				label:   "LD A, C",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetC()) },
			},

			// LD A, D
			0x7a: Op{
				code:    0x7a,
				cycles:  4,
				label:   "LD A, D",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetD()) },
			},

			// LD A, E
			0x7b: Op{
				code:    0x7b,
				cycles:  4,
				label:   "LD A, E",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetE()) },
			},

			// LD A, H
			0x7c: Op{
				code:    0x7c,
				cycles:  4,
				label:   "LD A, H",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetH()) },
			},

			// LD A, L
			0x7d: Op{
				code:    0x7d,
				cycles:  4,
				label:   "LD A, L",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetL()) },
			},

			// LD A, (HL)
			0x7e: Op{
				code:    0x7d,
				cycles:  8,
				label:   "LD A, (HL)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(m.ReadByte(r.GetHL())) },
			},

			// LD A, (BC)
			0x0a: Op{
				code:    0x0a,
				cycles:  8,
				label:   "LD A, (BC)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(m.ReadByte(r.GetBC())) },
			},

			// LD A, (DE)
			0x1a: Op{
				code:    0x1a,
				cycles:  8,
				label:   "LD A, (DE)",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(m.ReadByte(r.GetDE())) },
			},

			// LD A, (nn)
			0xfa: Op{
				code:    0xfa,
				cycles:  16,
				label:   "LD A, (nn)",
				argsLen: 2,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(ToWord(args[0], args[1])) },
			},

			// LD A, #
			0x3e: Op{
				code:    0x3e,
				cycles:  16,
				label:   "LD A, #",
				argsLen: 1,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetA(args[0]) },
			},

			// LD B, A
			0x47: Op{
				code:    0x47,
				cycles:  4,
				label:   "LD B, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetA()) },
			},

			// LD C, A
			0x4f: Op{
				code:    0x4f,
				cycles:  4,
				label:   "LD C, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetA()) },
			},

			// LD D, A
			0x57: Op{
				code:    0x57,
				cycles:  4,
				label:   "LD D, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetA()) },
			},

			// LD E, A
			0x5f: Op{
				code:    0x5f,
				cycles:  4,
				label:   "LD E, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetE(r.GetA()) },
			},

			// LD H, A
			0x67: Op{
				code:    0x67,
				cycles:  4,
				label:   "LD H, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetH(r.GetA()) },
			},

			// LD L, A
			0x6f: Op{
				code:    0x6f,
				cycles:  4,
				label:   "LD L, A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetL(r.GetA()) },
			},

			// LD (BC), A
			0x02: Op{
				code:    0x02,
				cycles:  8,
				label:   "LD (BC), A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetBC(), r.GetA()) },
			},

			// LD (DE), A
			0x12: Op{
				code:    0x12,
				cycles:  8,
				label:   "LD (DE), A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetDE(), r.GetA()) },
			},

			// LD (HL), A
			0x77: Op{
				code:    0x77,
				cycles:  8,
				label:   "LD (HL), A",
				argsLen: 0,
				fn:      func(r *CPURegisters, m *MMU, args []int) { m.WriteByte(r.GetHL(), r.GetA()) },
			},

			// LD (nn), A
			0xea: Op{
				code:    0xea,
				cycles:  16,
				label:   "LD (nn), A",
				argsLen: 2,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(ToWord(args[0], args[1]), r.GetA())
				},
			},

			// LD A, (C)
			0xf2: Op{
				code:    0xf2,
				cycles:  8,
				label:   "LD A, (C)",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetA(m.ReadByte(0xff00 + r.GetC()))
				},
			},

			// LD (C), A
			0xe2: Op{
				code:    0xe2,
				cycles:  8,
				label:   "LD (C), A",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(0xff00+r.GetC(), r.GetA())
				},
			},

			// LD A, (HLD)
			0x3a: Op{
				code:    0x3a,
				cycles:  8,
				label:   "LD A, (HLD)",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetA(m.ReadByte(r.DecHL()))
				},
			},

			// LD (HLD), A
			0x32: Op{
				code:    0x32,
				cycles:  8,
				label:   "LD (HLD), A",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(r.DecHL(), r.GetA())
				},
			},

			// LD A, (HLI)
			0x2a: Op{
				code:    0x2a,
				cycles:  8,
				label:   "LD A, (HLD)",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetA(m.ReadByte(r.IncHL()))
				},
			},

			// LD (HLI), A
			0x22: Op{
				code:    0x22,
				cycles:  8,
				label:   "LD (HLD), A",
				argsLen: 0,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(r.IncHL(), r.GetA())
				},
			},

			// LDH (n), A
			0xe0: Op{
				code:    0xe0,
				cycles:  12,
				label:   "LDH (n), A",
				argsLen: 1,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(0xfff0+args[0], r.GetA())
				},
			},

			// LDH A, (n)
			0xf0: Op{
				code:    0xf0,
				cycles:  12,
				label:   "LDH A, (n)",
				argsLen: 1,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetA(m.ReadByte(0xfff0 + args[0]))
				},
			},

			// LD BC, nn
			0x01: Op{
				code:    0x01,
				cycles:  12,
				label:   "LD BC, nn",
				argsLen: 2,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetBC(ToWord(args[0], args[1]))
				},
			},

			// LD DE, nn
			0x11: Op{
				code:    0x01,
				cycles:  12,
				label:   "LD DE, nn",
				argsLen: 2,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetDE(ToWord(args[0], args[1]))
				},
			},

			// LD HL, nn
			0x21: Op{
				code:    0x21,
				cycles:  12,
				label:   "LD HL, nn",
				argsLen: 2,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetHL(ToWord(args[0], args[1]))
				},
			},

			// LD SP, nn
			0x31: Op{
				code:    0x31,
				cycles:  12,
				label:   "LD SP, nn",
				argsLen: 2,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetSP(ToWord(args[0], args[1]))
				},
			},

			// LD SP, HL
			0xf9: Op{
				code:    0xf9,
				cycles:  8,
				label:   "LD SP, HL",
				argsLen: 2,
				fn:      func(r *CPURegisters, m *MMU, args []int) { r.SetSP(r.GetHL()) },
			},

			// LDHL SP, n
			0xf8: Op{
				code:    0xf8,
				cycles:  12,
				label:   "LDHL SP, n",
				argsLen: 1,
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetHL(AddSignedByteToWord(r.flags, r.GetSP(), args[0]))
				},
			},
		},
	}

	return CPUOpList
}
