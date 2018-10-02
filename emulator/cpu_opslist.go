package emulator

func makeOpList(r *CPURegisters) *Oplist {
	// Main OpList
	CPUOpList := &Oplist{
		registers: r,
		fns: map[int]Op{

			// Load Commands

			// LD B, n
			0x06: Op{
				code:   0x06,
				cycles: 8,
				label:  "LD B, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(args[0]) },
			},

			// LD C, n
			0x0e: Op{
				code:   0x0e,
				cycles: 8,
				label:  "LD C, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(args[0]) },
			},

			// LD D, n
			0x16: Op{
				code:   0x16,
				cycles: 8,
				label:  "LD D, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(args[0]) },
			},

			// LD E, n
			0x1e: Op{
				code:   0x1e,
				cycles: 8,
				label:  "LD E, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetE(args[0]) },
			},

			// LD H, n
			0x26: Op{
				code:   0x26,
				cycles: 8,
				label:  "LD H, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetH(args[0]) },
			},

			// LD L, n
			0x2e: Op{
				code:   0x2e,
				cycles: 8,
				label:  "LD L, n",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetL(args[0]) },
			},

			// LD A, A
			0x7f: Op{
				code:   0x7f,
				cycles: 4,
				label:  "LD A, A",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetA()) },
			},

			// LD A, B
			0x78: Op{
				code:   0x78,
				cycles: 4,
				label:  "LD A, B",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetB()) },
			},

			// LD A, C
			0x79: Op{
				code:   0x79,
				cycles: 4,
				label:  "LD A, C",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetC()) },
			},

			// LD A, D
			0x7a: Op{
				code:   0x7a,
				cycles: 4,
				label:  "LD A, D",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetD()) },
			},

			// LD A, E
			0x7b: Op{
				code:   0x7b,
				cycles: 4,
				label:  "LD A, E",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetE()) },
			},

			// LD A, H
			0x7c: Op{
				code:   0x7c,
				cycles: 4,
				label:  "LD A, H",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetH()) },
			},

			// LD A, L
			0x7d: Op{
				code:   0x7d,
				cycles: 4,
				label:  "LD A, L",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(r.GetL()) },
			},

			// LD A, (HL)
			0x7e: Op{
				code:   0x7d,
				cycles: 4,
				label:  "LD A, (HL)",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetA(m.ReadByte(r.GetHL())) },
			},

			// LD B, B
			0x40: Op{
				code:   0x40,
				cycles: 4,
				label:  "LD B, B",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetB()) },
			},

			// LD B, C
			0x41: Op{
				code:   0x41,
				cycles: 4,
				label:  "LD B, C",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetC()) },
			},

			// LD B, D
			0x42: Op{
				code:   0x42,
				cycles: 4,
				label:  "LD B, D",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetD()) },
			},

			// LD B, E
			0x43: Op{
				code:   0x43,
				cycles: 4,
				label:  "LD B, E",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetE()) },
			},

			// LD B, H
			0x44: Op{
				code:   0x44,
				cycles: 4,
				label:  "LD B, H",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetH()) },
			},

			// LD B, L
			0x45: Op{
				code:   0x45,
				cycles: 4,
				label:  "LD B, L",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(r.GetL()) },
			},

			// LD B, (HL)
			0x46: Op{
				code:   0x46,
				cycles: 8,
				label:  "LD B, (HL)",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetB(m.ReadByte(r.GetHL())) },
			},

			// LD C, B
			0x48: Op{
				code:   0x48,
				cycles: 4,
				label:  "LD C, B",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetB()) },
			},

			// LD C, C
			0x49: Op{
				code:   0x49,
				cycles: 4,
				label:  "LD C, C",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetC()) },
			},

			// LD C, D
			0x4a: Op{
				code:   0x4a,
				cycles: 4,
				label:  "LD C, D",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetD()) },
			},

			// LD C, E
			0x4b: Op{
				code:   0x4b,
				cycles: 4,
				label:  "LD C, E",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetE()) },
			},

			// LD C, H
			0x4c: Op{
				code:   0x4c,
				cycles: 4,
				label:  "LD C, H",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetH()) },
			},

			// LD C, L
			0x4d: Op{
				code:   0x4d,
				cycles: 4,
				label:  "LD C, L",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(r.GetL()) },
			},

			// LD C, (HL)
			0x4e: Op{
				code:   0x4e,
				cycles: 8,
				label:  "LD C, (HL)",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetC(m.ReadByte(r.GetHL())) },
			},

			// LD D, B
			0x50: Op{
				code:   0x50,
				cycles: 4,
				label:  "LD D, B",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetB()) },
			},

			// LD D, C
			0x51: Op{
				code:   0x51,
				cycles: 4,
				label:  "LD D, C",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetC()) },
			},

			// LD D, D
			0x52: Op{
				code:   0x52,
				cycles: 4,
				label:  "LD D, D",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetD()) },
			},

			// LD D, E
			0x53: Op{
				code:   0x53,
				cycles: 4,
				label:  "LD D, E",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetE()) },
			},

			// LD D, H
			0x54: Op{
				code:   0x54,
				cycles: 4,
				label:  "LD D, H",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetH()) },
			},

			// LD D, L
			0x55: Op{
				code:   0x55,
				cycles: 4,
				label:  "LD D, L",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(r.GetL()) },
			},

			// LD D, (HL)
			0x56: Op{
				code:   0x56,
				cycles: 8,
				label:  "LD D, (HL)",
				fn:     func(r *CPURegisters, m *MMU, args []int) { r.SetD(m.ReadByte(r.GetHL())) },
			},
		},
	}

	return CPUOpList
}
