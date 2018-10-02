package emulator

func makeOpList(r *CPURegisters) *Oplist {
	// Main OpList
	CPUOpList := &Oplist{
		registers: r,
		fns: map[int]Op{

			// NOP
			0x00: Op{
				code:    0x00,
				cycles:  1,
				argsLen: 0,
				label:   "NOP",
				fn:      func(r *CPURegisters, m *MMU, args []int) {},
			},

			// LD BC, nn
			0x01: Op{
				code:    0x01,
				cycles:  2,
				argsLen: 0,
				label:   "LD BC, nn",
				fn: func(r *CPURegisters, m *MMU, args []int) {
					r.SetBC(CombineBytes(args[0], args[1]))
				},
			},

			// LD (BC), A
			0x02: Op{
				code:    0x02,
				cycles:  8,
				argsLen: 0,
				label:   "LD (BC),A",
				fn: func(r *CPURegisters, m *MMU, args []int) {
					m.WriteByte(r.GetBC(), r.GetA())
				},
			},
		},
	}

	return CPUOpList
}
