package cpu

import "github.com/jumballaya/gameboy/emulator"

func makeOpList(r *Registers) (*Oplist, *Oplist) {
	// Main OpList
	mainList := &Oplist{
		registers: r,
		fns: map[int]Op{

			// Load Commands

			// LD B, n
			0x06: Op{
				code:    0x06,
				cycles:  8,
				label:   "LD B, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(args[0]) },
			},

			// LD C, n
			0x0e: Op{
				code:    0x0e,
				cycles:  8,
				label:   "LD C, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(args[0]) },
			},

			// LD D, n
			0x16: Op{
				code:    0x16,
				cycles:  8,
				label:   "LD D, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(args[0]) },
			},

			// LD E, n
			0x1e: Op{
				code:    0x1e,
				cycles:  8,
				label:   "LD E, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(args[0]) },
			},

			// LD H, n
			0x26: Op{
				code:    0x26,
				cycles:  8,
				label:   "LD H, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(args[0]) },
			},

			// LD L, n
			0x2e: Op{
				code:    0x2e,
				cycles:  8,
				label:   "LD L, n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(args[0]) },
			},

			// LD B, B
			0x40: Op{
				code:    0x40,
				cycles:  4,
				label:   "LD B, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetB()) },
			},

			// LD B, C
			0x41: Op{
				code:    0x41,
				cycles:  4,
				label:   "LD B, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetC()) },
			},

			// LD B, D
			0x42: Op{
				code:    0x42,
				cycles:  4,
				label:   "LD B, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetD()) },
			},

			// LD B, E
			0x43: Op{
				code:    0x43,
				cycles:  4,
				label:   "LD B, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetE()) },
			},

			// LD B, H
			0x44: Op{
				code:    0x44,
				cycles:  4,
				label:   "LD B, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetH()) },
			},

			// LD B, L
			0x45: Op{
				code:    0x45,
				cycles:  4,
				label:   "LD B, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetL()) },
			},

			// LD B, (HL)
			0x46: Op{
				code:    0x46,
				cycles:  8,
				label:   "LD B, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(m.ReadByte(r.GetHL())) },
			},

			// LD C, B
			0x48: Op{
				code:    0x48,
				cycles:  4,
				label:   "LD C, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetB()) },
			},

			// LD C, C
			0x49: Op{
				code:    0x49,
				cycles:  4,
				label:   "LD C, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetC()) },
			},

			// LD C, D
			0x4a: Op{
				code:    0x4a,
				cycles:  4,
				label:   "LD C, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetD()) },
			},

			// LD C, E
			0x4b: Op{
				code:    0x4b,
				cycles:  4,
				label:   "LD C, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetE()) },
			},

			// LD C, H
			0x4c: Op{
				code:    0x4c,
				cycles:  4,
				label:   "LD C, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetH()) },
			},

			// LD C, L
			0x4d: Op{
				code:    0x4d,
				cycles:  4,
				label:   "LD C, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetL()) },
			},

			// LD C, (HL)
			0x4e: Op{
				code:    0x4e,
				cycles:  8,
				label:   "LD C, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(m.ReadByte(r.GetHL())) },
			},

			// LD D, B
			0x50: Op{
				code:    0x50,
				cycles:  4,
				label:   "LD D, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetB()) },
			},

			// LD D, C
			0x51: Op{
				code:    0x51,
				cycles:  4,
				label:   "LD D, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetC()) },
			},

			// LD D, D
			0x52: Op{
				code:    0x52,
				cycles:  4,
				label:   "LD D, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetD()) },
			},

			// LD D, E
			0x53: Op{
				code:    0x53,
				cycles:  4,
				label:   "LD D, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetE()) },
			},

			// LD D, H
			0x54: Op{
				code:    0x54,
				cycles:  4,
				label:   "LD D, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetH()) },
			},

			// LD D, L
			0x55: Op{
				code:    0x55,
				cycles:  4,
				label:   "LD D, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetL()) },
			},

			// LD D, (HL)
			0x56: Op{
				code:    0x56,
				cycles:  8,
				label:   "LD D, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(m.ReadByte(r.GetHL())) },
			},

			// LD E, B
			0x58: Op{
				code:    0x58,
				cycles:  4,
				label:   "LD E, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetB()) },
			},

			// LD E, C
			0x59: Op{
				code:    0x59,
				cycles:  4,
				label:   "LD E, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetC()) },
			},

			// LD E, D
			0x5a: Op{
				code:    0x5a,
				cycles:  4,
				label:   "LD E, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetD()) },
			},

			// LD E, E
			0x5b: Op{
				code:    0x5b,
				cycles:  4,
				label:   "LD E, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetE()) },
			},

			// LD E, H
			0x5c: Op{
				code:    0x5c,
				cycles:  4,
				label:   "LD E, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetH()) },
			},

			// LD E, L
			0x5d: Op{
				code:    0x5d,
				cycles:  4,
				label:   "LD E, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetL()) },
			},

			// LD E, (HL)
			0x5e: Op{
				code:    0x5e,
				cycles:  8,
				label:   "LD E, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(m.ReadByte(r.GetHL())) },
			},

			// LD H, B
			0x60: Op{
				code:    0x60,
				cycles:  4,
				label:   "LD H, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetB()) },
			},

			// LD H, C
			0x61: Op{
				code:    0x61,
				cycles:  4,
				label:   "LD H, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetC()) },
			},

			// LD H, D
			0x62: Op{
				code:    0x62,
				cycles:  4,
				label:   "LD H, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetD()) },
			},

			// LD H, E
			0x63: Op{
				code:    0x63,
				cycles:  4,
				label:   "LD H, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetE()) },
			},

			// LD H, H
			0x64: Op{
				code:    0x64,
				cycles:  4,
				label:   "LD H, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetH()) },
			},

			// LD H, L
			0x65: Op{
				code:    0x65,
				cycles:  4,
				label:   "LD H, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetL()) },
			},

			// LD H, (HL)
			0x66: Op{
				code:    0x66,
				cycles:  8,
				label:   "LD H, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(m.ReadByte(r.GetHL())) },
			},

			// LD L, B
			0x68: Op{
				code:    0x68,
				cycles:  4,
				label:   "LD L, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetB()) },
			},

			// LD L, C
			0x69: Op{
				code:    0x69,
				cycles:  4,
				label:   "LD L, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetC()) },
			},

			// LD L, D
			0x6a: Op{
				code:    0x6a,
				cycles:  4,
				label:   "LD L, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetD()) },
			},

			// LD L, E
			0x6b: Op{
				code:    0x6b,
				cycles:  4,
				label:   "LD L, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetE()) },
			},

			// LD L, H
			0x6c: Op{
				code:    0x6c,
				cycles:  4,
				label:   "LD L, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetH()) },
			},

			// LD L, L
			0x6d: Op{
				code:    0x6d,
				cycles:  4,
				label:   "LD L, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetL()) },
			},

			// LD L, (HL)
			0x6e: Op{
				code:    0x6e,
				cycles:  8,
				label:   "LD L, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(m.ReadByte(r.GetHL())) },
			},

			// LD (HL), B
			0x70: Op{
				code:    0x70,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), B",
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetB()) },
			},

			// LD (HL), C
			0x71: Op{
				code:    0x71,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), C",
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetC()) },
			},

			// LD (HL), D
			0x72: Op{
				code:    0x72,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), D",
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetD()) },
			},

			// LD (HL), E
			0x73: Op{
				code:    0x73,
				cycles:  8,
				argsLen: 0,
				label:   "(HL), E",
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetE()) },
			},

			// LD (HL), H
			0x74: Op{
				code:    0x74,
				cycles:  8,
				label:   "(HL), B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetH()) },
			},

			// LD (HL), L
			0x75: Op{
				code:    0x75,
				cycles:  8,
				label:   "(HL), L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetL()) },
			},

			// LD (HL), n
			0x36: Op{
				code:    0x36,
				cycles:  12,
				label:   "(HL), n",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), args[0]) },
			},

			// LD A, A
			0x7f: Op{
				code:    0x7f,
				cycles:  4,
				label:   "LD A, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetA()) },
			},

			// LD A, B
			0x78: Op{
				code:    0x78,
				cycles:  4,
				label:   "LD A, B",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetB()) },
			},

			// LD A, C
			0x79: Op{
				code:    0x79,
				cycles:  4,
				label:   "LD A, C",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetC()) },
			},

			// LD A, D
			0x7a: Op{
				code:    0x7a,
				cycles:  4,
				label:   "LD A, D",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetD()) },
			},

			// LD A, E
			0x7b: Op{
				code:    0x7b,
				cycles:  4,
				label:   "LD A, E",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetE()) },
			},

			// LD A, H
			0x7c: Op{
				code:    0x7c,
				cycles:  4,
				label:   "LD A, H",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetH()) },
			},

			// LD A, L
			0x7d: Op{
				code:    0x7d,
				cycles:  4,
				label:   "LD A, L",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(r.GetL()) },
			},

			// LD A, (HL)
			0x7e: Op{
				code:    0x7d,
				cycles:  8,
				label:   "LD A, (HL)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(m.ReadByte(r.GetHL())) },
			},

			// LD A, (BC)
			0x0a: Op{
				code:    0x0a,
				cycles:  8,
				label:   "LD A, (BC)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(m.ReadByte(r.GetBC())) },
			},

			// LD A, (DE)
			0x1a: Op{
				code:    0x1a,
				cycles:  8,
				label:   "LD A, (DE)",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(m.ReadByte(r.GetDE())) },
			},

			// LD A, (nn)
			0xfa: Op{
				code:    0xfa,
				cycles:  16,
				label:   "LD A, (nn)",
				argsLen: 2,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(ToWord(args[0], args[1])) },
			},

			// LD A, #
			0x3e: Op{
				code:    0x3e,
				cycles:  16,
				label:   "LD A, #",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetA(args[0]) },
			},

			// LD B, A
			0x47: Op{
				code:    0x47,
				cycles:  4,
				label:   "LD B, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetB(r.GetA()) },
			},

			// LD C, A
			0x4f: Op{
				code:    0x4f,
				cycles:  4,
				label:   "LD C, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetC(r.GetA()) },
			},

			// LD D, A
			0x57: Op{
				code:    0x57,
				cycles:  4,
				label:   "LD D, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetD(r.GetA()) },
			},

			// LD E, A
			0x5f: Op{
				code:    0x5f,
				cycles:  4,
				label:   "LD E, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetE(r.GetA()) },
			},

			// LD H, A
			0x67: Op{
				code:    0x67,
				cycles:  4,
				label:   "LD H, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetH(r.GetA()) },
			},

			// LD L, A
			0x6f: Op{
				code:    0x6f,
				cycles:  4,
				label:   "LD L, A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetL(r.GetA()) },
			},

			// LD (BC), A
			0x02: Op{
				code:    0x02,
				cycles:  8,
				label:   "LD (BC), A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetBC(), r.GetA()) },
			},

			// LD (DE), A
			0x12: Op{
				code:    0x12,
				cycles:  8,
				label:   "LD (DE), A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetDE(), r.GetA()) },
			},

			// LD (HL), A
			0x77: Op{
				code:    0x77,
				cycles:  8,
				label:   "LD (HL), A",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) { m.WriteByte(r.GetHL(), r.GetA()) },
			},

			// LD (nn), A
			0xea: Op{
				code:    0xea,
				cycles:  16,
				label:   "LD (nn), A",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(ToWord(args[0], args[1]), r.GetA())
				},
			},

			// LD A, (C)
			0xf2: Op{
				code:    0xf2,
				cycles:  8,
				label:   "LD A, (C)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(m.ReadByte(0xff00 + r.GetC()))
				},
			},

			// LD (C), A
			0xe2: Op{
				code:    0xe2,
				cycles:  8,
				label:   "LD (C), A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(0xff00+r.GetC(), r.GetA())
				},
			},

			// LD A, (HLD)
			0x3a: Op{
				code:    0x3a,
				cycles:  8,
				label:   "LD A, (HLD)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(m.ReadByte(r.DecHL()))
				},
			},

			// LD (HLD), A
			0x32: Op{
				code:    0x32,
				cycles:  8,
				label:   "LD (HLD), A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(r.DecHL(), r.GetA())
				},
			},

			// LD A, (HLI)
			0x2a: Op{
				code:    0x2a,
				cycles:  8,
				label:   "LD A, (HLD)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(m.ReadByte(r.IncHL()))
				},
			},

			// LD (HLI), A
			0x22: Op{
				code:    0x22,
				cycles:  8,
				label:   "LD (HLD), A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(r.IncHL(), r.GetA())
				},
			},

			// LDH (n), A
			0xe0: Op{
				code:    0xe0,
				cycles:  12,
				label:   "LDH (n), A",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(0xfff0+args[0], r.GetA())
				},
			},

			// LDH A, (n)
			0xf0: Op{
				code:    0xf0,
				cycles:  12,
				label:   "LDH A, (n)",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(m.ReadByte(0xfff0 + args[0]))
				},
			},

			// LD BC, nn
			0x01: Op{
				code:    0x01,
				cycles:  12,
				label:   "LD BC, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetBC(ToWord(args[0], args[1]))
				},
			},

			// LD DE, nn
			0x11: Op{
				code:    0x01,
				cycles:  12,
				label:   "LD DE, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetDE(ToWord(args[0], args[1]))
				},
			},

			// LD HL, nn
			0x21: Op{
				code:    0x21,
				cycles:  12,
				label:   "LD HL, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(ToWord(args[0], args[1]))
				},
			},

			// LD SP, nn
			0x31: Op{
				code:    0x31,
				cycles:  12,
				label:   "LD SP, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetSP(ToWord(args[0], args[1]))
				},
			},

			// LD SP, HL
			0xf9: Op{
				code:    0xf9,
				cycles:  8,
				label:   "LD SP, HL",
				argsLen: 2,
				fn:      func(r *Registers, m emulator.Memory, args []int) { r.SetSP(r.GetHL()) },
			},

			// LDHL SP, n
			0xf8: Op{
				code:    0xf8,
				cycles:  12,
				label:   "LDHL SP, n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(AddSignedByteToWord(r.flags, r.GetSP(), args[0]))
				},
			},

			// LD (nn), SP
			0x08: Op{
				code:    0x08,
				cycles:  20,
				label:   "LD (nn), SP",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					argWord := ToWord(args[0], args[1])
					m.WriteByte(argWord, LowerByte(r.GetSP()))
					m.WriteByte((argWord+1)&0xffff, UpperByte(r.GetSP()))
				},
			},

			/****** STACK OPERATIONS ********/

			// PUSH AF
			0xf5: Op{
				code:    0xf5,
				cycles:  16,
				label:   "PUSH AF",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Push(r, m, r.GetAF())
				},
			},

			// PUSH BC
			0xc5: Op{
				code:    0xc5,
				cycles:  16,
				label:   "PUSH BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Push(r, m, r.GetBC())
				},
			},

			// PUSH DE
			0xd5: Op{
				code:    0xd5,
				cycles:  16,
				label:   "PUSH DE",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Push(r, m, r.GetDE())
				},
			},

			// PUSH HL
			0xe5: Op{
				code:    0xe5,
				cycles:  16,
				label:   "PUSH HL",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Push(r, m, r.GetHL())
				},
			},

			// POP AF
			0xf1: Op{
				code:    0xf1,
				cycles:  12,
				label:   "POP AF",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetAF(Pop(r, m))
				},
			},

			// POP BC
			0xc1: Op{
				code:    0xc1,
				cycles:  12,
				label:   "POP BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetBC(Pop(r, m))
				},
			},

			// POP DE
			0xd1: Op{
				code:    0xd1,
				cycles:  12,
				label:   "POP DE",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetDE(Pop(r, m))
				},
			},

			// POP HL
			0xe1: Op{
				code:    0xe1,
				cycles:  12,
				label:   "POP HL",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(Pop(r, m))
				},
			},

			// ADD A, A
			0x87: Op{
				code:    0x87,
				cycles:  4,
				label:   "ADD A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// ADD A, B
			0x80: Op{
				code:    0x80,
				cycles:  4,
				label:   "ADD A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// ADD A, C
			0x81: Op{
				code:    0x81,
				cycles:  4,
				label:   "ADD A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// ADD A, D
			0x82: Op{
				code:    0x82,
				cycles:  4,
				label:   "ADD A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// ADD A, E
			0x83: Op{
				code:    0x83,
				cycles:  4,
				label:   "ADD A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// ADD A, H
			0x84: Op{
				code:    0x84,
				cycles:  4,
				label:   "ADD A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// ADD A, L
			0x85: Op{
				code:    0x85,
				cycles:  4,
				label:   "ADD A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// ADD A, (HL)
			0x86: Op{
				code:    0x86,
				cycles:  8,
				label:   "ADD A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// ADD A, #
			0xc6: Op{
				code:    0xc6,
				cycles:  8,
				label:   "ADD A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytes(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// ADC A, A
			0x8f: Op{
				code:    0x8f,
				cycles:  4,
				label:   "ADC A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// ADC A, B
			0x88: Op{
				code:    0x88,
				cycles:  4,
				label:   "ADC A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// ADC A, C
			0x89: Op{
				code:    0x89,
				cycles:  4,
				label:   "ADC A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// ADC A, D
			0x8a: Op{
				code:    0x8a,
				cycles:  4,
				label:   "ADC A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// ADC A, E
			0x8b: Op{
				code:    0x8b,
				cycles:  4,
				label:   "ADC A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// ADC A, H
			0x8c: Op{
				code:    0x8c,
				cycles:  4,
				label:   "ADC A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// ADC A, L
			0x8d: Op{
				code:    0x8d,
				cycles:  4,
				label:   "ADC A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// ADC A, (HL)
			0x8e: Op{
				code:    0x8e,
				cycles:  8,
				label:   "ADC A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// ADC A, #
			0xce: Op{
				code:    0xce,
				cycles:  8,
				label:   "ADC A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(AddBytesAndCarry(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// SUB A, A
			0x97: Op{
				code:    0x97,
				cycles:  4,
				label:   "SUB A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// SUB A, B
			0x90: Op{
				code:    0x90,
				cycles:  4,
				label:   "SUB A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// SUB A, C
			0x91: Op{
				code:    0x91,
				cycles:  4,
				label:   "SUB A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// SUB A, D
			0x92: Op{
				code:    0x92,
				cycles:  4,
				label:   "SUB A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// SUB A, E
			0x93: Op{
				code:    0x93,
				cycles:  4,
				label:   "SUB A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// SUB A, H
			0x94: Op{
				code:    0x94,
				cycles:  4,
				label:   "SUB A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// SUB A, L
			0x95: Op{
				code:    0x95,
				cycles:  4,
				label:   "SUB A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// SUB A, (HL)
			0x96: Op{
				code:    0x96,
				cycles:  8,
				label:   "SUB A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// SUB A, #
			0xd6: Op{
				code:    0xd6,
				cycles:  8,
				label:   "SUB A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytes(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// SBC A, A
			0x9f: Op{
				code:    0x9f,
				cycles:  4,
				label:   "SBC A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// SBC A, B
			0x98: Op{
				code:    0x98,
				cycles:  4,
				label:   "SBC A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// SBC A, C
			0x99: Op{
				code:    0x99,
				cycles:  4,
				label:   "SBC A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// SBC A, D
			0x9a: Op{
				code:    0x9a,
				cycles:  4,
				label:   "SBC A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// SBC A, E
			0x9b: Op{
				code:    0x9b,
				cycles:  4,
				label:   "SBC A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// SBC A, H
			0x9c: Op{
				code:    0x9c,
				cycles:  4,
				label:   "SBC A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// SBC A, L
			0x9d: Op{
				code:    0x9d,
				cycles:  4,
				label:   "SBC A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// SBC A, (HL)
			0x9e: Op{
				code:    0x9e,
				cycles:  8,
				label:   "SBC A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(SubBytesAndCarry(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// AND A, A
			0xa7: Op{
				code:    0xa7,
				cycles:  4,
				label:   "AND A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// AND A, B
			0xa0: Op{
				code:    0xa0,
				cycles:  4,
				label:   "AND A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// AND A, C
			0xa1: Op{
				code:    0xa1,
				cycles:  4,
				label:   "AND A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// AND A, D
			0xa2: Op{
				code:    0xa2,
				cycles:  4,
				label:   "AND A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// AND A, E
			0xa3: Op{
				code:    0xa3,
				cycles:  4,
				label:   "AND A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// AND A, H
			0xa4: Op{
				code:    0xa4,
				cycles:  4,
				label:   "AND A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// AND A, L
			0xa5: Op{
				code:    0xa5,
				cycles:  4,
				label:   "AND A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// AND A, (HL)
			0xa6: Op{
				code:    0xa6,
				cycles:  8,
				label:   "AND A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// AND A, #
			0xe6: Op{
				code:    0xe6,
				cycles:  8,
				label:   "AND A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(And(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// OR A, A
			0xb7: Op{
				code:    0xb7,
				cycles:  4,
				label:   "OR A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// OR A, B
			0xb0: Op{
				code:    0xb0,
				cycles:  4,
				label:   "OR A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// OR A, C
			0xb1: Op{
				code:    0xb1,
				cycles:  4,
				label:   "OR A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// OR A, D
			0xb2: Op{
				code:    0xb2,
				cycles:  4,
				label:   "OR A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// OR A, E
			0xb3: Op{
				code:    0xb3,
				cycles:  4,
				label:   "OR A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// OR A, H
			0xb4: Op{
				code:    0xb4,
				cycles:  4,
				label:   "OR A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// OR A, L
			0xb5: Op{
				code:    0xb5,
				cycles:  4,
				label:   "OR A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// OR A, (HL)
			0xb6: Op{
				code:    0xb6,
				cycles:  8,
				label:   "OR A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// OR A, #
			0xf6: Op{
				code:    0xf6,
				cycles:  8,
				label:   "OR A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Or(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// XOR A, A
			0xaf: Op{
				code:    0xaf,
				cycles:  4,
				label:   "XOR A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetA()))
				},
			},

			// XOR A, B
			0xa8: Op{
				code:    0xa8,
				cycles:  4,
				label:   "XOR A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetB()))
				},
			},

			// XOR A, C
			0xa9: Op{
				code:    0xa9,
				cycles:  4,
				label:   "XOR A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetC()))
				},
			},

			// XOR A, D
			0xaa: Op{
				code:    0xaa,
				cycles:  4,
				label:   "XOR A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetD()))
				},
			},

			// XOR A, E
			0xab: Op{
				code:    0xab,
				cycles:  4,
				label:   "XOR A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetE()))
				},
			},

			// XOR A, H
			0xac: Op{
				code:    0xac,
				cycles:  4,
				label:   "XOR A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetH()))
				},
			},

			// XOR A, L
			0xad: Op{
				code:    0xad,
				cycles:  4,
				label:   "XOR A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), r.GetL()))
				},
			},

			// XOR A, (HL)
			0xae: Op{
				code:    0xae,
				cycles:  8,
				label:   "XOR A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL())))
				},
			},

			// XOR A, #
			0xee: Op{
				code:    0xee,
				cycles:  8,
				label:   "XOR A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(Xor(r.GetFlags(), r.GetA(), args[0]))
				},
			},

			// CP A, A
			0xbf: Op{
				code:    0xbf,
				cycles:  4,
				label:   "CP A, A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetA())
				},
			},

			// CP A, B
			0xb8: Op{
				code:    0xb8,
				cycles:  4,
				label:   "CP A, B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetB())
				},
			},

			// CP A, C
			0xb9: Op{
				code:    0xb9,
				cycles:  4,
				label:   "CP A, C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetC())
				},
			},

			// CP A, D
			0xba: Op{
				code:    0xba,
				cycles:  4,
				label:   "CP A, D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetD())
				},
			},

			// CP A, E
			0xbb: Op{
				code:    0xbb,
				cycles:  4,
				label:   "CP A, E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetE())
				},
			},

			// CP A, H
			0xbc: Op{
				code:    0xbc,
				cycles:  4,
				label:   "CP A, H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetH())
				},
			},

			// CP A, L
			0xbd: Op{
				code:    0xbd,
				cycles:  4,
				label:   "CP A, L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), r.GetL())
				},
			},

			// CP A, (HL)
			0xbe: Op{
				code:    0xbe,
				cycles:  8,
				label:   "CP A, (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), m.ReadByte(r.GetHL()))
				},
			},

			// CP A, #
			0xfe: Op{
				code:    0xfe,
				cycles:  8,
				label:   "CP A, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					SubBytes(r.GetFlags(), r.GetA(), args[0])
				},
			},

			// INC A
			0x3c: Op{
				code:    0x3c,
				cycles:  4,
				label:   "INC A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(IncByte(r.GetFlags(), r.GetA()))
				},
			},

			// INC B
			0x04: Op{
				code:    0x04,
				cycles:  4,
				label:   "INC B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetB(IncByte(r.GetFlags(), r.GetB()))
				},
			},

			// INC C
			0x0c: Op{
				code:    0x0c,
				cycles:  4,
				label:   "INC C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetC(IncByte(r.GetFlags(), r.GetC()))
				},
			},

			// INC D
			0x14: Op{
				code:    0x14,
				cycles:  4,
				label:   "INC D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetD(IncByte(r.GetFlags(), r.GetD()))
				},
			},

			// INC E
			0x1c: Op{
				code:    0x1c,
				cycles:  4,
				label:   "INC E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetE(IncByte(r.GetFlags(), r.GetE()))
				},
			},

			// INC H
			0x24: Op{
				code:    0x24,
				cycles:  4,
				label:   "INC H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetH(IncByte(r.GetFlags(), r.GetH()))
				},
			},

			// INC L
			0x2c: Op{
				code:    0x2c,
				cycles:  4,
				label:   "INC L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetL(IncByte(r.GetFlags(), r.GetL()))
				},
			},

			// INC (HL)
			0x34: Op{
				code:    0x34,
				cycles:  12,
				label:   "INC (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(r.GetHL(), IncByte(r.GetFlags(), m.ReadByte(r.GetHL())))
				},
			},

			// DEC A
			0x3d: Op{
				code:    0x3d,
				cycles:  4,
				label:   "DEC A",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(DecByte(r.GetFlags(), r.GetA()))
				},
			},

			// DEC B
			0x05: Op{
				code:    0x05,
				cycles:  4,
				label:   "DEC B",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetB(DecByte(r.GetFlags(), r.GetB()))
				},
			},

			// DEC C
			0x0d: Op{
				code:    0x0d,
				cycles:  4,
				label:   "DEC C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetC(DecByte(r.GetFlags(), r.GetC()))
				},
			},

			// DEC D
			0x15: Op{
				code:    0x15,
				cycles:  4,
				label:   "DEC D",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetD(DecByte(r.GetFlags(), r.GetD()))
				},
			},

			// DEC E
			0x1d: Op{
				code:    0x1d,
				cycles:  4,
				label:   "DEC E",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetE(DecByte(r.GetFlags(), r.GetE()))
				},
			},

			// DEC H
			0x25: Op{
				code:    0x25,
				cycles:  4,
				label:   "DEC H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetH(DecByte(r.GetFlags(), r.GetH()))
				},
			},

			// DEC L
			0x2d: Op{
				code:    0x2d,
				cycles:  4,
				label:   "DEC L",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetL(DecByte(r.GetFlags(), r.GetL()))
				},
			},

			// DEC (HL)
			0x35: Op{
				code:    0x35,
				cycles:  12,
				label:   "DEC (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					m.WriteByte(r.GetHL(), DecByte(r.GetFlags(), m.ReadByte(r.GetHL())))
				},
			},

			// ADD HL, BC
			0x09: Op{
				code:    0x09,
				cycles:  8,
				label:   "ADD HL, BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(AddWords(r.GetFlags(), r.GetHL(), r.GetBC()))
				},
			},

			// ADD HL, DE
			0x19: Op{
				code:    0x19,
				cycles:  8,
				label:   "ADD HL, DE",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(AddWords(r.GetFlags(), r.GetHL(), r.GetDE()))
				},
			},

			// ADD HL, HL
			0x29: Op{
				code:    0x29,
				cycles:  8,
				label:   "ADD HL, BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(AddWords(r.GetFlags(), r.GetHL(), r.GetHL()))
				},
			},

			// ADD HL, SP
			0x39: Op{
				code:    0x39,
				cycles:  8,
				label:   "ADD HL, SP",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL(AddWords(r.GetFlags(), r.GetHL(), r.GetSP()))
				},
			},

			// ADD SP, #
			0xe8: Op{
				code:    0xe8,
				cycles:  16,
				label:   "ADD SP, #",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetSP(AddSignedByteToWord(r.GetFlags(), r.GetSP(), args[0]))
				},
			},

			// INC BC
			0x03: Op{
				code:    0x03,
				cycles:  8,
				label:   "INC BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetBC((r.GetBC() + 1) & 0xffff)
				},
			},

			// INC DE
			0x13: Op{
				code:    0x13,
				cycles:  8,
				label:   "INC DE",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetDE((r.GetDE() + 1) & 0xffff)
				},
			},

			// INC HL
			0x23: Op{
				code:    0x23,
				cycles:  8,
				label:   "INC HL",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL((r.GetHL() + 1) & 0xffff)
				},
			},

			// INC SP
			0x33: Op{
				code:    0x33,
				cycles:  8,
				label:   "INC SP",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetSP((r.GetSP() + 1) & 0xffff)
				},
			},

			// DEC BC
			0x0b: Op{
				code:    0x0b,
				cycles:  8,
				label:   "DEC BC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetBC((r.GetBC() - 1) & 0xffff)
				},
			},

			// DEC DE
			0x1b: Op{
				code:    0x1b,
				cycles:  8,
				label:   "DEC DE",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetDE((r.GetDE() - 1) & 0xffff)
				},
			},

			// DEC HL
			0x2b: Op{
				code:    0x2b,
				cycles:  8,
				label:   "DEC HL",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetHL((r.GetHL() - 1) & 0xffff)
				},
			},

			// DEC SP
			0x3b: Op{
				code:    0x3b,
				cycles:  8,
				label:   "DEC SP",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetSP((r.GetSP() - 1) & 0xffff)
				},
			},

			// DAA
			0x27: Op{
				code:    0x27,
				cycles:  4,
				label:   "DAA",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					res := r.GetA()
					flags := r.GetFlags()
					if (res&0x0f) > 9 || flags.IsH() {
						res += 0x06
					}
					if (res&0xf0) > 0x90 || flags.IsC() {
						res += 0x60
						flags.SetC(true)
					}
					res &= 0xff
					flags.SetZ(res == 0)
					flags.SetH(false)
					r.SetA(res)
				},
			},

			// CPL
			0x2f: Op{
				code:    0x2f,
				cycles:  4,
				label:   "CPL",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					f := r.GetFlags()

					f.SetN(true)
					f.SetH(true)

					r.SetA(^(r.GetA()) & 0xff)
				},
			},

			// CCF
			0x3f: Op{
				code:    0x3f,
				cycles:  4,
				label:   "CCF",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					f := r.GetFlags()

					f.SetN(false)
					f.SetH(false)

					f.SetC(!f.IsC())
				},
			},

			// SCF
			0x37: Op{
				code:    0x37,
				cycles:  4,
				label:   "SCF",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					f := r.GetFlags()

					f.SetN(false)
					f.SetH(false)
					f.SetC(true)
				},
			},

			// NOP
			0x00: Op{
				code:    0x00,
				cycles:  4,
				label:   "NOP",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},

			// HALT
			0x76: Op{
				code:    0x76,
				cycles:  4,
				label:   "HALT",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},

			// STOP
			0x10: Op{
				code:    0x10,
				cycles:  4,
				label:   "STOP",
				argsLen: 1,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},

			// DI
			0xf3: Op{
				code:    0xf3,
				cycles:  4,
				label:   "DI",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},

			// EI
			0xfb: Op{
				code:    0xfb,
				cycles:  4,
				label:   "EI",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},

			// RLCA
			0x07: Op{
				code:    0x07,
				cycles:  4,
				label:   "RLCA",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(RotateLeft(r.GetFlags(), r.GetA()))
				},
			},

			// RLA
			0x17: Op{
				code:    0x17,
				cycles:  4,
				label:   "RLA",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(RotateLeftThroughCarry(r.GetFlags(), r.GetA()))
				},
			},

			// RRCA
			0x0f: Op{
				code:    0x0f,
				cycles:  4,
				label:   "RRCA",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(RotateRight(r.GetFlags(), r.GetA()))
				},
			},

			// RRA
			0x1f: Op{
				code:    0x1f,
				cycles:  4,
				label:   "RRA",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetA(RotateRightThroughCarry(r.GetFlags(), r.GetA()))
				},
			},

			/************** JUMPS **************/

			// JMP nn
			0xc3: Op{
				code:    0xc3,
				cycles:  12,
				label:   "JMP nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetPC(ToWord(args[0], args[1]))
				},
			},

			// JP NZ, nn
			0xc2: Op{
				code:    0xc2,
				cycles:  12,
				label:   "JP NZ, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsZ() {
						r.SetPC(ToWord(args[0], args[1]))
					}
				},
			},

			// JP Z, nn
			0xca: Op{
				code:    0xca,
				cycles:  12,
				label:   "JP Z, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsZ() {
						r.SetPC(ToWord(args[0], args[1]))
					}
				},
			},

			// JP NC, nn
			0xd2: Op{
				code:    0xd2,
				cycles:  12,
				label:   "JP NC, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsC() {
						r.SetPC(ToWord(args[0], args[1]))
					}
				},
			},

			// JP C, nn
			0xda: Op{
				code:    0xda,
				cycles:  12,
				label:   "JP C, nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsC() {
						r.SetPC(ToWord(args[0], args[1]))
					}
				},
			},

			// JP (HL)
			0xe9: Op{
				code:    0xe9,
				cycles:  4,
				label:   "JP (HL)",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.SetPC(r.GetHL())
				},
			},

			// JR n
			0x18: Op{
				code:    0x18,
				cycles:  8,
				label:   "JR n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					r.AddToPC(args[0])
				},
			},

			// JR NZ, n
			0x20: Op{
				code:    0x20,
				cycles:  8,
				label:   "JR NZ, n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsZ() {
						r.AddToPC(args[0])
					}
				},
			},

			// JR Z, n
			0x28: Op{
				code:    0x28,
				cycles:  8,
				label:   "JR Z, n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsZ() {
						r.AddToPC(args[0])
					}
				},
			},

			// JR NC, n
			0x30: Op{
				code:    0x30,
				cycles:  8,
				label:   "JR NC, n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsC() {
						r.AddToPC(args[0])
					}
				},
			},

			// JR C, n
			0x38: Op{
				code:    0x38,
				cycles:  8,
				label:   "JR C, n",
				argsLen: 1,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsC() {
						r.AddToPC(args[0])
					}
				},
			},

			// CALL nn
			0xcd: Op{
				code:    0xcd,
				cycles:  12,
				label:   "CALL nn",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Call(r, m, ToWord(args[0], args[1]))
				},
			},

			// CALL NZ, n
			0xc4: Op{
				code:    0xc4,
				cycles:  12,
				label:   "CALL NZ, n",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsZ() {
						Call(r, m, ToWord(args[0], args[1]))
					}
				},
			},

			// CALL Z, n
			0xcc: Op{
				code:    0xcc,
				cycles:  12,
				label:   "CALL Z, n",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsZ() {
						Call(r, m, ToWord(args[0], args[1]))
					}
				},
			},

			// CALL NC, n
			0xd4: Op{
				code:    0xd4,
				cycles:  12,
				label:   "CALL NC, n",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsC() {
						Call(r, m, ToWord(args[0], args[1]))
					}
				},
			},

			// CALL C, n
			0xdc: Op{
				code:    0xdc,
				cycles:  12,
				label:   "CALL C, n",
				argsLen: 2,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsC() {
						Call(r, m, ToWord(args[0], args[1]))
					}
				},
			},

			// RST 00H
			0xc7: Op{
				code:    0xc7,
				cycles:  32,
				label:   "RST 00H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x00)
				},
			},

			// RST 08H
			0xcf: Op{
				code:    0xcf,
				cycles:  32,
				label:   "RST 08H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x08)
				},
			},

			// RST 10H
			0xd7: Op{
				code:    0xd7,
				cycles:  32,
				label:   "RST 10H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x10)
				},
			},

			// RST 18H
			0xdf: Op{
				code:    0xdf,
				cycles:  32,
				label:   "RST 18H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x18)
				},
			},

			// RST 20H
			0xe7: Op{
				code:    0xe7,
				cycles:  32,
				label:   "RST 20H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x20)
				},
			},

			// RST 28H
			0xef: Op{
				code:    0xef,
				cycles:  32,
				label:   "RST 28H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x28)
				},
			},

			// RST 30H
			0xf7: Op{
				code:    0xf7,
				cycles:  32,
				label:   "RST 30H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x30)
				},
			},

			// RST 38H
			0xff: Op{
				code:    0xff,
				cycles:  32,
				label:   "RST 38H",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Reset(r, m, 0x38)
				},
			},

			// RET
			0xc9: Op{
				code:    0xc9,
				cycles:  8,
				label:   "RET",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					Ret(r, m)
				},
			},

			// RET NZ
			0xc0: Op{
				code:    0xc0,
				cycles:  8,
				label:   "RET NZ",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsZ() {
						Ret(r, m)
					}
				},
			},

			// RET Z
			0xc8: Op{
				code:    0xc8,
				cycles:  8,
				label:   "RET Z",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsZ() {
						Ret(r, m)
					}
				},
			},

			// RET NC
			0xd0: Op{
				code:    0xd0,
				cycles:  8,
				label:   "RET NC",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if !r.GetFlags().IsC() {
						Ret(r, m)
					}
				},
			},

			// RET C
			0xd8: Op{
				code:    0xd8,
				cycles:  8,
				label:   "RET C",
				argsLen: 0,
				fn: func(r *Registers, m emulator.Memory, args []int) {
					if r.GetFlags().IsC() {
						Ret(r, m)
					}
				},
			},

			// RETI
			0xd9: Op{
				code:    0xd9,
				cycles:  8,
				label:   "RETI",
				argsLen: 0,
				fn:      func(r *Registers, m emulator.Memory, args []int) {},
			},
		},
	}

	extList := &Oplist{
		registers: r,
		fns: map[int]Op{

			// SWAP A
			0x37: Op{
				code: 0x37,
			},

			// SWAP B
			0x30: Op{
				code: 0x30,
			},

			// SWAP C
			0x31: Op{
				code: 0x31,
			},

			// SWAP D
			0x32: Op{
				code: 0x32,
			},

			// SWAP E
			0x33: Op{
				code: 0x33,
			},

			// SWAP H
			0x34: Op{
				code: 0x34,
			},

			// SWAP L
			0x35: Op{
				code: 0x35,
			},

			// SWAP (HL)
			0x36: Op{
				code: 0x36,
			},
		},
	}

	return mainList, extList
}
