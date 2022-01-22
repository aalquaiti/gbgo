package cpu

// Initialise instructions
func (c *CPU) initInstructions() {
	// region OpCode SetBit

	c.inst = [OPCodeSize]OpCode{}

	//// region CB Prefixed Instructions
	//
	//// RLC B
	//c.cbInst[0x00] = OpCode{2, rlcb}
	//// RLC C
	//c.cbInst[0x01] = OpCode{2, rlcc}
	//// RLC D
	//c.cbInst[0x02] = OpCode{2, rlcd}
	//// RLC E
	//c.cbInst[0x03] = OpCode{2, rlce}
	//// RLC H
	//c.cbInst[0x04] = OpCode{2, rlch}
	//// RLC L
	//c.cbInst[0x05] = OpCode{2, rlcl}
	//// RLC (HL)
	//c.cbInst[0x06] = OpCode{2, rlchl}
	//// RLC A
	//c.cbInst[0x07] = OpCode{2, cbrlca}
	//// RRC B
	//c.cbInst[0x08] = OpCode{2, rrcb}
	//// RRC C
	//c.cbInst[0x09] = OpCode{2, rrcc}
	//// RRC D
	//c.cbInst[0x0A] = OpCode{2, rrcd}
	//// RRC E
	//c.cbInst[0x0B] = OpCode{2, rrce}
	//// RRC H
	//c.cbInst[0x0C] = OpCode{2, rrch}
	//// RRC L
	//c.cbInst[0x0D] = OpCode{2, rrcl}
	//// RRC (HL)
	//c.cbInst[0x0E] = OpCode{2, rrchl}
	//// RRC A
	//c.cbInst[0x0F] = OpCode{2, cbrrca}
	//
	//// RL B
	//c.cbInst[0x10] = OpCode{2, rlb}
	//// RL C
	//c.cbInst[0x11] = OpCode{2, rlc}
	//// RL D
	//c.cbInst[0x12] = OpCode{2, rld}
	//// RL E
	//c.cbInst[0x13] = OpCode{2, rle}
	//// RL H
	//c.cbInst[0x14] = OpCode{2, rlh}
	//// RL L
	//c.cbInst[0x15] = OpCode{2, rll}
	//// RL (HL)
	//c.cbInst[0x16] = OpCode{2, rlhl}
	//// RL A
	//c.cbInst[0x17] = OpCode{2, cbrla}
	//// RR B
	//c.cbInst[0x18] = OpCode{2, rrb}
	//// RR C
	//c.cbInst[0x19] = OpCode{2, rrc}
	//// RR D
	//c.cbInst[0x1A] = OpCode{2, rrd}
	//// RR E
	//c.cbInst[0x1B] = OpCode{2, rre}
	//// RR H
	//c.cbInst[0x1C] = OpCode{2, rrh}
	//// RR L
	//c.cbInst[0x1D] = OpCode{2, rrl}
	//// RR (HL)
	//c.cbInst[0x1E] = OpCode{2, rrhl}
	//// RR A
	//c.cbInst[0x1F] = OpCode{2, cbrra}
	//
	//// SLA B
	//c.cbInst[0x20] = OpCode{2, slab}
	//// SLA C
	//c.cbInst[0x21] = OpCode{2, slac}
	//// SLA D
	//c.cbInst[0x22] = OpCode{2, slad}
	//// SLA E
	//c.cbInst[0x23] = OpCode{2, slae}
	//// SLA H
	//c.cbInst[0x24] = OpCode{2, slah}
	//// SLA L
	//c.cbInst[0x25] = OpCode{2, slal}
	//// SLA (HL)
	//c.cbInst[0x26] = OpCode{2, slahl}
	//// SLA A
	//c.cbInst[0x27] = OpCode{2, slaa}
	//// SRA B
	//c.cbInst[0x28] = OpCode{2, srab}
	//// SRA C
	//c.cbInst[0x29] = OpCode{2, srac}
	//// SRA D
	//c.cbInst[0x2A] = OpCode{2, srad}
	//// SRA E
	//c.cbInst[0x2B] = OpCode{2, srae}
	//// SRA H
	//c.cbInst[0x2C] = OpCode{2, srah}
	//// SRA L
	//c.cbInst[0x2D] = OpCode{2, sral}
	//// SRA (HL)
	//c.cbInst[0x2E] = OpCode{2, srahl}
	//// SRA A
	//c.cbInst[0x2F] = OpCode{2, sraa}
	//
	//// SWAP B
	//c.cbInst[0x30] = OpCode{2, swapb}
	//// SWAP C
	//c.cbInst[0x31] = OpCode{2, swapc}
	//// SWAP D
	//c.cbInst[0x32] = OpCode{2, swapd}
	//// SWAP E
	//c.cbInst[0x33] = OpCode{2, swape}
	//// SWAP H
	//c.cbInst[0x34] = OpCode{2, swaph}
	//// SWAP L
	//c.cbInst[0x35] = OpCode{2, swapl}
	//// SWAP (HL)
	//c.cbInst[0x36] = OpCode{2, swaphl}
	//// SWAP A
	//c.cbInst[0x37] = OpCode{2, swapa}
	//// SRL B
	//c.cbInst[0x28] = OpCode{2, srlb}
	//// SRL C
	//c.cbInst[0x29] = OpCode{2, srlc}
	//// SRL D
	//c.cbInst[0x2A] = OpCode{2, srld}
	//// SRL E
	//c.cbInst[0x2B] = OpCode{2, srle}
	//// SRL H
	//c.cbInst[0x2C] = OpCode{2, srlh}
	//// SRL L
	//c.cbInst[0x2D] = OpCode{2, srll}
	//// SRL (HL)
	//c.cbInst[0x2E] = OpCode{2, srlhl}
	//// SRL A
	//c.cbInst[0x2F] = OpCode{2, srla}
	//
	//// SWAP B
	//c.cbInst[0x30] = OpCode{2, swapb}
	//// SWAP C
	//c.cbInst[0x31] = OpCode{2, swapc}
	//// SWAP D
	//c.cbInst[0x32] = OpCode{2, swapd}
	//// SWAP E
	//c.cbInst[0x33] = OpCode{2, swape}
	//// SWAP H
	//c.cbInst[0x34] = OpCode{2, swaph}
	//// SWAP L
	//c.cbInst[0x35] = OpCode{2, swapl}
	//// SWAP (HL)
	//c.cbInst[0x36] = OpCode{2, swaphl}
	//// SWAP A
	//c.cbInst[0x37] = OpCode{2, swapa}
	//// SRL B
	//c.cbInst[0x38] = OpCode{2, srlb}
	//// SRL C
	//c.cbInst[0x39] = OpCode{2, srlc}
	//// SRL D
	//c.cbInst[0x3A] = OpCode{2, srld}
	//// SRL E
	//c.cbInst[0x3B] = OpCode{2, srle}
	//// SRL H
	//c.cbInst[0x3C] = OpCode{2, srlh}
	//// SRL L
	//c.cbInst[0x3D] = OpCode{2, srll}
	//// SRL (HL)
	//c.cbInst[0x3E] = OpCode{2, srlhl}
	//// SRL A
	//c.cbInst[0x3F] = OpCode{2, srla}
	//
	//// BIT 0, B
	//c.cbInst[0x40] = OpCode{2, bit0b}
	//// BIT 0, C
	//c.cbInst[0x41] = OpCode{2, bit0c}
	//// BIT 0, D
	//c.cbInst[0x42] = OpCode{2, bit0d}
	//// BIT 0, E
	//c.cbInst[0x43] = OpCode{2, bit0e}
	//// BIT 0, H
	//c.cbInst[0x44] = OpCode{2, bit0h}
	//// BIT 0, L
	//c.cbInst[0x45] = OpCode{2, bit0l}
	//// BIT 0, (HL)
	//c.cbInst[0x46] = OpCode{2, bit0hl}
	//// BIT 0, A
	//c.cbInst[0x47] = OpCode{2, bit0a}
	//// BIT 1, B
	//c.cbInst[0x48] = OpCode{2, bit1b}
	//// BIT 1, C
	//c.cbInst[0x49] = OpCode{2, bit1c}
	//// BIT 1, D
	//c.cbInst[0x4A] = OpCode{2, bit1d}
	//// BIT 1, E
	//c.cbInst[0x4B] = OpCode{2, bit1e}
	//// BIT 1, H
	//c.cbInst[0x4C] = OpCode{2, bit1h}
	//// BIT 1, L
	//c.cbInst[0x4D] = OpCode{2, bit1l}
	//// BIT 1, (HL)
	//c.cbInst[0x4E] = OpCode{2, bit1hl}
	//// BIT 1, A
	//c.cbInst[0x4F] = OpCode{2, bit1a}
	//
	//// BIT 2, B
	//c.cbInst[0x50] = OpCode{2, bit2b}
	//// BIT 2, C
	//c.cbInst[0x51] = OpCode{2, bit2c}
	//// BIT 2, D
	//c.cbInst[0x52] = OpCode{2, bit2d}
	//// BIT 2, E
	//c.cbInst[0x53] = OpCode{2, bit2e}
	//// BIT 2, H
	//c.cbInst[0x54] = OpCode{2, bit2h}
	//// BIT 2, L
	//c.cbInst[0x55] = OpCode{2, bit2l}
	//// BIT 2, (HL)
	//c.cbInst[0x56] = OpCode{2, bit2hl}
	//// BIT 2, A
	//c.cbInst[0x57] = OpCode{2, bit2a}
	//// BIT 3, B
	//c.cbInst[0x58] = OpCode{2, bit3b}
	//// BIT 3, C
	//c.cbInst[0x59] = OpCode{2, bit3c}
	//// BIT 3, D
	//c.cbInst[0x5A] = OpCode{2, bit3d}
	//// BIT 3, E
	//c.cbInst[0x5B] = OpCode{2, bit3e}
	//// BIT 3, H
	//c.cbInst[0x5C] = OpCode{2, bit3h}
	//// BIT 3, L
	//c.cbInst[0x5D] = OpCode{2, bit3l}
	//// BIT 3, (HL)
	//c.cbInst[0x5E] = OpCode{2, bit3hl}
	//// BIT 3, A
	//c.cbInst[0x5F] = OpCode{2, bit3a}
	//
	//// BIT 4, B
	//c.cbInst[0x60] = OpCode{2, bit4b}
	//// BIT 4, C
	//c.cbInst[0x61] = OpCode{2, bit4c}
	//// BIT 4, D
	//c.cbInst[0x62] = OpCode{2, bit4d}
	//// BIT 4, E
	//c.cbInst[0x63] = OpCode{2, bit4e}
	//// BIT 4, H
	//c.cbInst[0x64] = OpCode{2, bit4h}
	//// BIT 4, L
	//c.cbInst[0x65] = OpCode{2, bit4l}
	//// BIT 4, (HL)
	//c.cbInst[0x66] = OpCode{2, bit4hl}
	//// BIT 4, A
	//c.cbInst[0x67] = OpCode{2, bit4a}
	//// BIT 5, B
	//c.cbInst[0x68] = OpCode{2, bit5b}
	//// BIT 5, C
	//c.cbInst[0x69] = OpCode{2, bit5c}
	//// BIT 5, D
	//c.cbInst[0x6A] = OpCode{2, bit5d}
	//// BIT 5, E
	//c.cbInst[0x6B] = OpCode{2, bit5e}
	//// BIT 5, H
	//c.cbInst[0x6C] = OpCode{2, bit5h}
	//// BIT 5, L
	//c.cbInst[0x6D] = OpCode{2, bit5l}
	//// BIT 5, (HL)
	//c.cbInst[0x6E] = OpCode{2, bit5hl}
	//// BIT 5, A
	//c.cbInst[0x6F] = OpCode{2, bit5a}
	//
	//// BIT 6, B
	//c.cbInst[0x70] = OpCode{2, bit6b}
	//// BIT 6, C
	//c.cbInst[0x71] = OpCode{2, bit6c}
	//// BIT 6, D
	//c.cbInst[0x72] = OpCode{2, bit6d}
	//// BIT 6, E
	//c.cbInst[0x73] = OpCode{2, bit6e}
	//// BIT 6, H
	//c.cbInst[0x74] = OpCode{2, bit6h}
	//// BIT 6, L
	//c.cbInst[0x75] = OpCode{2, bit6l}
	//// BIT 6, (HL)
	//c.cbInst[0x76] = OpCode{2, bit6hl}
	//// BIT 6, A
	//c.cbInst[0x77] = OpCode{2, bit6a}
	//// BIT 7, B
	//c.cbInst[0x78] = OpCode{2, bit7b}
	//// BIT 7, C
	//c.cbInst[0x79] = OpCode{2, bit7c}
	//// BIT 7, D
	//c.cbInst[0x7A] = OpCode{2, bit7d}
	//// BIT 7, E
	//c.cbInst[0x7B] = OpCode{2, bit7e}
	//// BIT 7, H
	//c.cbInst[0x7C] = OpCode{2, bit7h}
	//// BIT 7, L
	//c.cbInst[0x7D] = OpCode{2, bit7l}
	//// BIT 7, (HL)
	//c.cbInst[0x7E] = OpCode{2, bit7hl}
	//// BIT 7, A
	//c.cbInst[0x7F] = OpCode{2, bit7a}
	//
	//// RES 0, B
	//c.cbInst[0x80] = OpCode{2, res0b}
	//// RES 0, C
	//c.cbInst[0x81] = OpCode{2, res0c}
	//// RES 0, D
	//c.cbInst[0x82] = OpCode{2, res0d}
	//// RES 0, E
	//c.cbInst[0x83] = OpCode{2, res0e}
	//// RES 0, H
	//c.cbInst[0x84] = OpCode{2, res0h}
	//// RES 0, L
	//c.cbInst[0x85] = OpCode{2, res0l}
	//// RES 0, (HL)
	//c.cbInst[0x86] = OpCode{2, res0hl}
	//// RES 0, A
	//c.cbInst[0x87] = OpCode{2, res0a}
	//// RES 1, B
	//c.cbInst[0x88] = OpCode{2, res1b}
	//// RES 1, C
	//c.cbInst[0x89] = OpCode{2, res1c}
	//// RES 1, D
	//c.cbInst[0x8A] = OpCode{2, res1d}
	//// RES 1, E
	//c.cbInst[0x8B] = OpCode{2, res1e}
	//// RES 1, H
	//c.cbInst[0x8C] = OpCode{2, res1h}
	//// RES 1, L
	//c.cbInst[0x8D] = OpCode{2, res1l}
	//// RES 1, (HL)
	//c.cbInst[0x8E] = OpCode{2, res1hl}
	//// RES 1, A
	//c.cbInst[0x8F] = OpCode{2, res1a}
	//
	//// RES 2, B
	//c.cbInst[0x90] = OpCode{2, res2b}
	//// RES 2, C
	//c.cbInst[0x91] = OpCode{2, res2c}
	//// RES 2, D
	//c.cbInst[0x92] = OpCode{2, res2d}
	//// RES 2, E
	//c.cbInst[0x93] = OpCode{2, res2e}
	//// RES 2, H
	//c.cbInst[0x94] = OpCode{2, res2h}
	//// RES 2, L
	//c.cbInst[0x95] = OpCode{2, res2l}
	//// RES 2, (HL)
	//c.cbInst[0x96] = OpCode{2, res2hl}
	//// RES 2, A
	//c.cbInst[0x97] = OpCode{2, res2a}
	//// RES 3, B
	//c.cbInst[0x98] = OpCode{2, res3b}
	//// RES 3, C
	//c.cbInst[0x99] = OpCode{2, res3c}
	//// RES 3, D
	//c.cbInst[0x9A] = OpCode{2, res3d}
	//// RES 3, E
	//c.cbInst[0x9B] = OpCode{2, res3e}
	//// RES 3, H
	//c.cbInst[0x9C] = OpCode{2, res3h}
	//// RES 3, L
	//c.cbInst[0x9D] = OpCode{2, res3l}
	//// RES 3, (HL)
	//c.cbInst[0x9E] = OpCode{2, res3hl}
	//// RES 3, A
	//c.cbInst[0x9F] = OpCode{2, res3a}
	//
	//// RES 4, B
	//c.cbInst[0xA0] = OpCode{2, res4b}
	//// RES 4, C
	//c.cbInst[0xA1] = OpCode{2, res4c}
	//// RES 4, D
	//c.cbInst[0xA2] = OpCode{2, res4d}
	//// RES 4, E
	//c.cbInst[0xA3] = OpCode{2, res4e}
	//// RES 4, H
	//c.cbInst[0xA4] = OpCode{2, res4h}
	//// RES 4, L
	//c.cbInst[0xA5] = OpCode{2, res4l}
	//// RES 4, (HL)
	//c.cbInst[0xA6] = OpCode{2, res4hl}
	//// RES 4, A
	//c.cbInst[0xA7] = OpCode{2, res4a}
	//// RES 5, B
	//c.cbInst[0xA8] = OpCode{2, res5b}
	//// RES 5, C
	//c.cbInst[0xA9] = OpCode{2, res5c}
	//// RES 5, D
	//c.cbInst[0xAA] = OpCode{2, res5d}
	//// RES 5, E
	//c.cbInst[0xAB] = OpCode{2, res5e}
	//// RES 5, H
	//c.cbInst[0xAC] = OpCode{2, res5h}
	//// RES 5, L
	//c.cbInst[0xAD] = OpCode{2, res5l}
	//// RES 5, (HL)
	//c.cbInst[0xAE] = OpCode{2, res5hl}
	//// RES 5, A
	//c.cbInst[0xAF] = OpCode{2, res5a}
	//
	//// RES 6, B
	//c.cbInst[0xB0] = OpCode{2, res6b}
	//// RES 6, C
	//c.cbInst[0xB1] = OpCode{2, res6c}
	//// RES 6, D
	//c.cbInst[0xB2] = OpCode{2, res6d}
	//// RES 6, E
	//c.cbInst[0xB3] = OpCode{2, res6e}
	//// RES 6, H
	//c.cbInst[0xB4] = OpCode{2, res6h}
	//// RES 6, L
	//c.cbInst[0xB5] = OpCode{2, res6l}
	//// RES 6, (HL)
	//c.cbInst[0xB6] = OpCode{2, res6hl}
	//// RES 6, A
	//c.cbInst[0xB7] = OpCode{2, res6a}
	//// RES 7, B
	//c.cbInst[0xB8] = OpCode{2, res7b}
	//// RES 7, C
	//c.cbInst[0xB9] = OpCode{2, res7c}
	//// RES 7, D
	//c.cbInst[0xBA] = OpCode{2, res7d}
	//// RES 7, E
	//c.cbInst[0xBB] = OpCode{2, res7e}
	//// RES 7, H
	//c.cbInst[0xBC] = OpCode{2, res7h}
	//// RES 7, L
	//c.cbInst[0xBD] = OpCode{2, res7l}
	//// RES 7, (HL)
	//c.cbInst[0xBE] = OpCode{2, res7hl}
	//// RES 7, A
	//c.cbInst[0xBF] = OpCode{2, res7a}
	//
	//// SET 0, B
	//c.cbInst[0xC0] = OpCode{2, set0b}
	//// SET 0, C
	//c.cbInst[0xC1] = OpCode{2, set0c}
	//// SET 0, D
	//c.cbInst[0xC2] = OpCode{2, set0d}
	//// SET 0, E
	//c.cbInst[0xC3] = OpCode{2, set0e}
	//// SET 0, H
	//c.cbInst[0xC4] = OpCode{2, set0h}
	//// SET 0, L
	//c.cbInst[0xC5] = OpCode{2, set0l}
	//// SET 0, (HL)
	//c.cbInst[0xC6] = OpCode{2, set0hl}
	//// SET 0, A
	//c.cbInst[0xC7] = OpCode{2, set0a}
	//// SET 1, B
	//c.cbInst[0xC8] = OpCode{2, set1b}
	//// SET 1, C
	//c.cbInst[0xC9] = OpCode{2, set1c}
	//// SET 1, D
	//c.cbInst[0xCA] = OpCode{2, set1d}
	//// SET 1, E
	//c.cbInst[0xCB] = OpCode{2, set1e}
	//// SET 1, H
	//c.cbInst[0xCC] = OpCode{2, set1h}
	//// SET 1, L
	//c.cbInst[0xCD] = OpCode{2, set1l}
	//// SET 1, (HL)
	//c.cbInst[0xCE] = OpCode{2, set1hl}
	//// SET 1, A
	//c.cbInst[0xCF] = OpCode{2, set1a}
	//
	//// SET 2, B
	//c.cbInst[0xD0] = OpCode{2, set2b}
	//// SET 2, C
	//c.cbInst[0xD1] = OpCode{2, set2c}
	//// SET 2, D
	//c.cbInst[0xD2] = OpCode{2, set2d}
	//// SET 2, E
	//c.cbInst[0xD3] = OpCode{2, set2e}
	//// SET 2, H
	//c.cbInst[0xD4] = OpCode{2, set2h}
	//// SET 2, L
	//c.cbInst[0xD5] = OpCode{2, set2l}
	//// SET 2, (HL)
	//c.cbInst[0xD6] = OpCode{2, set2hl}
	//// SET 2, A
	//c.cbInst[0xD7] = OpCode{2, set2a}
	//// SET 3, B
	//c.cbInst[0xD8] = OpCode{2, set3b}
	//// SET 3, C
	//c.cbInst[0xD9] = OpCode{2, set3c}
	//// SET 3, D
	//c.cbInst[0xDA] = OpCode{2, set3d}
	//// SET 3, E
	//c.cbInst[0xDB] = OpCode{2, set3e}
	//// SET 3, H
	//c.cbInst[0xDC] = OpCode{2, set3h}
	//// SET 3, L
	//c.cbInst[0xDD] = OpCode{2, set3l}
	//// SET 3, (HL)
	//c.cbInst[0xDE] = OpCode{2, set3hl}
	//// SET 3, A
	//c.cbInst[0xDF] = OpCode{2, set3a}
	//
	//// SET 4, B
	//c.cbInst[0xE0] = OpCode{2, set4b}
	//// SET 4, C
	//c.cbInst[0xE1] = OpCode{2, set4c}
	//// SET 4, D
	//c.cbInst[0xE2] = OpCode{2, set4d}
	//// SET 4, E
	//c.cbInst[0xE3] = OpCode{2, set4e}
	//// SET 4, H
	//c.cbInst[0xE4] = OpCode{2, set4h}
	//// SET 4, L
	//c.cbInst[0xE5] = OpCode{2, set4l}
	//// SET 4, (HL)
	//c.cbInst[0xE6] = OpCode{2, set4hl}
	//// SET 4, A
	//c.cbInst[0xE7] = OpCode{2, set4a}
	//// SET 5, B
	//c.cbInst[0xE8] = OpCode{2, set5b}
	//// SET 5, C
	//c.cbInst[0xE9] = OpCode{2, set5c}
	//// SET 5, D
	//c.cbInst[0xEA] = OpCode{2, set5d}
	//// SET 5, E
	//c.cbInst[0xEB] = OpCode{2, set5e}
	//// SET 5, H
	//c.cbInst[0xEC] = OpCode{2, set5h}
	//// SET 5, L
	//c.cbInst[0xED] = OpCode{2, set5l}
	//// SET 5, (HL)
	//c.cbInst[0xEE] = OpCode{2, set5hl}
	//// SET 5, A
	//c.cbInst[0xEF] = OpCode{2, set5a}
	//
	//// SET 6, B
	//c.cbInst[0xB0] = OpCode{2, set6b}
	//// SET 6, C
	//c.cbInst[0xB1] = OpCode{2, set6c}
	//// SET 6, D
	//c.cbInst[0xB2] = OpCode{2, set6d}
	//// SET 6, E
	//c.cbInst[0xB3] = OpCode{2, set6e}
	//// SET 6, H
	//c.cbInst[0xB4] = OpCode{2, set6h}
	//// SET 6, L
	//c.cbInst[0xB5] = OpCode{2, set6l}
	//// SET 6, (HL)
	//c.cbInst[0xB6] = OpCode{2, set6hl}
	//// SET 6, A
	//c.cbInst[0xB7] = OpCode{2, set6a}
	//// SET 7, B
	//c.cbInst[0xB8] = OpCode{2, set7b}
	//// SET 7, C
	//c.cbInst[0xB9] = OpCode{2, set7c}
	//// SET 7, D
	//c.cbInst[0xBA] = OpCode{2, set7d}
	//// SET 7, E
	//c.cbInst[0xBB] = OpCode{2, set7e}
	//// SET 7, H
	//c.cbInst[0xBC] = OpCode{2, set7h}
	//// SET 7, L
	//c.cbInst[0xBD] = OpCode{2, set7l}
	//// SET 7, (HL)
	//c.cbInst[0xBE] = OpCode{2, set7hl}
	//// SET 7, A
	//c.cbInst[0xBF] = OpCode{2, set7a}
	//
	//// endregion CB Prefixed Instructions
}

////region OpCode Functions
//
//func (c *CPU) nop() {
//}
//
//// ldRegImm16 Load Immediate 16-bit value to Reg16
//func (c *CPU) ldRegImm16() {
//	value := c.bus.Read16(c.Reg.PC.Get())
//	c.Reg.PC.Inc()
//	c.Reg.PC.Inc()
//	c.curOP.operands[0].(Reg16).Set(value)
//}
//
//// ldIRegImm16 Load Reg8 to indirect memory in Reg16
//func (c *CPU) ldIRegReg() {
//	pos := c.curOP.operands[0].(Reg16).Get()
//	c.bus.Write(pos, c.curOP.operands[1].(Reg8).Get())
//}
//
//// incReg Increment Register
//func (c *CPU) incReg() {
//	c.curOP.operands[0].Inc()
//}
//
//// decReg Decrement Register
//func (c *CPU) decReg() {
//	c.curOP.operands[0].Dec()
//}
//
//// ldRegImm8 Load Immediate 8-bit value to Reg8
//func (c *CPU) ldRegImm8() {
//	value := c.bus.Read(c.Reg.PC.Get())
//	c.Reg.PC.Inc()
//	c.curOP.operands[0].(Reg8).Set(value)
//}
//
//// rlca Rotate c.Register A left
//// Bit 7 shifts to bit 0
//// Bit 7 affect the carry Flag
//// C <- [7~0] <- [7]
//func (c *CPU) rlca() {
//	var bit7 bool = c.Reg.A.Get()&0x80 == 0x80
//	// If bit 7 is 1
//	*c.Reg.A.Val() <<= 1
//	if bit7 {
//		*c.Reg.A.Val() |= 1
//	}
//	c.flags.SetFlagZ(false)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit7)
//}
//
////func (c *CPU) ldmemsp() string {
////	pos := c.bus.Read16(c.Reg.PC.Get())
////	c.Reg.PC.Inc().Inc()
////	c.bus.Write16(pos, c.Reg.SP.Get())
////
////	return fmt.Sprintf("LD ($%X), SP", pos)
////}
////
////func (c *CPU) addhlbc() string {
////	addhlc.Reg(c.Reg.B.Get(), c.Reg.C.Get())
////
////	return "ADD HL, BC"
////}
////
////func (c *CPU) ldabc() string {
////	pos := c.Reg.BC.Get()
////	*c.Reg.A.Val() = c.bus.Read(pos)
////
////	return "LD A, (BC)"
////}
////
////func (c *CPU) decbc() string {
////	value := c.Reg.BC.Get() - 1
////	c.Reg.BC.Set(value)
////
////	return "DEC BC"
////}
////
////func (c *CPU) incc() string {
////	incc.Reg(c.Reg.C.Val())
////
////	return "INC C"
////}
////
////func (c *CPU) decc() string {
////	decc.Reg(c.Reg.C.Val())
////
////	return "DEC C"
////}
////
////func (c *CPU) ldc() string {
////	*c.Reg.C.Val() = c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////
////	return fmt.Sprintf("LD C, $%X", c.Reg.C.Get())
////}
////
////// Rotate c.Register A right
////// Bit 0 shifts to Carry
////// [0] -> [7~0] -> C
////func (c *CPU) rrca() string {
////	var bit0 bool = c.Reg.A.Get()&0x1 == 0x1
////	*c.Reg.A.Val() >>= 1
////	if bit0 {
////		*c.Reg.A.Val() |= 0x80
////	}
////	flags.SetFlagZ(false)
////	flags.SetFlagN(false)
////	flags.SetFlagH(false)
////	flags.SetFlagC(bit0)
////
////	return "RRCA"
////}
////
////// Enters CPU low power mode.
////// In GBC, switches between normal and double CPU speed
////func (c *CPU) stop() string {
////	// TODO implement cpu speed switch
////
////	// After the stop func (c *CPU)tions comes an operand that is ignored by the cpu
////	c.Reg.PC.Inc()
////
////	return "STOP"
////}
////
////func (c *CPU) ldde16() string {
////	value := c.bus.Read16(c.Reg.PC.Get())
////	c.Reg.PC.Inc().Inc()
////	c.Reg.DE.Set(value)
////
////	return fmt.Sprintf("LD DE, %X", value)
////}
////
////func (c *CPU) lddea() string {
////	pos := c.Reg.DE.Get()
////	c.bus.Write(pos, c.Reg.A.Get())
////
////	return "LD (DE), A"
////}
////
////func (c *CPU) incde() string {
////	value := c.Reg.DE.Get()
////	value++
////	c.Reg.DE.Set(value)
////
////	return "INC DE"
////}
////
////func (c *CPU) incd() string {
////	incc.Reg(c.Reg.D.Val())
////
////	return "INC D"
////}
////
////func (c *CPU) decd() string {
////	decc.Reg(c.Reg.D.Val())
////
////	return "DEC D"
////}
////
////func (c *CPU) ldd() string {
////	c.Reg.PC.Inc()
////	*c.Reg.D.Val() = c.bus.Read(c.Reg.PC.Get())
////
////	return fmt.Sprintf("LD D, $%X", c.Reg.D)
////}
////
////// Rotate c.Register A left through Carry
////// Previous Carry shifts to bit 0
////// Bit 7 shift to Carry
////// C <- [7~0] <- C
////func (c *CPU) rla() string {
////	var bit7 bool = c.Reg.A.Get()&0x80 == 0x80
////	*c.Reg.A.Val() <<= 1
////	// If carry flag is 1
////	if flags.GetFlagC() {
////		*c.Reg.A.Val() |= 1
////	}
////	flags.SetFlagZ(false)
////	flags.SetFlagN(false)
////	flags.SetFlagH(false)
////	flags.SetFlagC(bit7)
////
////	return "RLA"
////}
////
////func (c *CPU) jr() string {
////	value := jrCond(true, 0)
////
////	return fmt.Sprintf("JR $%X", value)
////}
////
////func (c *CPU) addhlde() string {
////	addhlc.Reg(c.Reg.D.Get(), c.Reg.E.Get())
////
////	return "ADD HL, DE"
////}
////
////func (c *CPU) ldade() string {
////	pos := c.Reg.DE.Get()
////	*c.Reg.A.Val() = c.bus.Read(pos)
////
////	return "LD A, (DE)"
////}
////
////func (c *CPU) decde() string {
////	value := c.Reg.DE.Get() - 1
////	c.Reg.DE.Set(value)
////
////	return "DEC DE"
////}
////
////func (c *CPU) ince() string {
////	incc.Reg(c.Reg.E.Val())
////
////	return "INC E"
////}
////
////func (c *CPU) dece() string {
////	decc.Reg(c.Reg.E.Val())
////
////	return "DEC E"
////}
////
////func (c *CPU) lde() string {
////	c.Reg.PC.Inc()
////	*c.Reg.E.Val() = c.bus.Read(c.Reg.PC.Get())
////
////	return fmt.Sprintf("LD E, $%X", c.Reg.E.Get())
////}
////
////// Rotate c.Register A right through Carry
////// Previous Carry value shifts to bit 7
////// Bit 0 shifts to Carry
////// C -> [7~0] -> C
////func (c *CPU) rra() string {
////	var bit0 bool = c.Reg.A.Get()&0x1 == 0x1
////	*c.Reg.A.Val() >>= 1
////	// If carry flag is 1
////	if flags.GetFlagC() {
////		*c.Reg.A.Val() |= 0x80
////	}
////	flags.SetFlagZ(false)
////	flags.SetFlagN(false)
////	flags.SetFlagH(false)
////	flags.SetFlagC(bit0)
////
////	return "RRA"
////}
////
////func (c *CPU) jrnz() string {
////	value := jrCond(!flags.GetFlagZ(), 1)
////
////	return fmt.Sprintf("JR NZ, $%X", value)
////}
////
////func (c *CPU) ldhl16() string {
////	value := c.bus.Read16(c.Reg.PC.Get())
////	c.Reg.PC.Inc().Inc()
////	c.Reg.HL.Set(value)
////
////	return fmt.Sprintf("LD HL, %X", value)
////}
////
////func (c *CPU) ldhli() string {
////	pos := c.Reg.HL.Get()
////	c.bus.Write(pos, c.Reg.A.Get())
////	c.Reg.HL.Set(pos + 1)
////
////	return "LD (HLI), A"
////}
////
////func (c *CPU) inchl() string {
////	value := c.Reg.HL.Get()
////	value++
////	c.Reg.HL.Set(value)
////
////	return "INC HL"
////}
////
////func (c *CPU) inch() string {
////	incc.Reg(c.Reg.H.Val())
////
////	return "INC H"
////}
////
////func (c *CPU) dech() string {
////	decc.Reg(c.Reg.H.Val())
////
////	return "DEC H"
////}
////
////func (c *CPU) ldh() string {
////	c.Reg.PC.Inc()
////	*c.Reg.H.Val() = c.bus.Read(c.Reg.PC.Get())
////
////	return fmt.Sprintf("LD H, $%X", c.Reg.H.Get())
////}
////
////func (c *CPU) daa() string {
////	// Decimal Adjust the Accumulator to be BCD correct.
////	// The process is as follows:
////	// 1. Check four Least Significant Bits (LSB)
////	// 2. LSB > 9 ||H Flag is SetBit to One -> Add $06 (or Subtract if N Flag is SetBit to One)
////	// 3. Check four Most Significant Bits (MSB)
////	// 4. MSB > 9 || C Flag is SetBit to One -> Add $60
////
////	// TODO test implementation is correct
////	// Use following links as guide
////	// http://z80-heaven.wikidot.com/instructions-set:daa
////	// http://www.z80.info/z80syntx.htm#DAA
////	// https://ehaskins.com/2018-01-30%20Z80%20DAA/
////	lsb := c.Reg.A.Get() & 0x0F
////	msb := c.Reg.A.Get() >> 4
////
////	// TODO Optimise code for better performance
////	if lsb > 9 || flags.GetFlagH() {
////		if !flags.GetFlagN() {
////			*c.Reg.A.Val() += 0x06
////		} else {
////			*c.Reg.A.Val() -= 0x06
////		}
////	}
////
////	if msb > 9 || flags.GetFlagC() {
////		if !flags.GetFlagN() {
////			*c.Reg.A.Val() += 0x60
////		} else {
////
////		}
////
////	}
////
////	flags.SetFlagZ(c.Reg.A.Get() == 0)
////	// Carry is set to one When BCD value is over $99, according to definition of DAA
////	flags.SetFlagC(c.Reg.A.Get() > 0x99)
////
////	return "DAA"
////}
////
////func (c *CPU) jrz() string {
////	value := jrCond(flags.GetFlagZ(), 1)
////
////	return fmt.Sprintf("JR Z, $%X", value)
////}
////
////func (c *CPU) addhlhl() string {
////	addhlc.Reg(c.Reg.H.Get(), c.Reg.L.Get())
////
////	return "ADD HL, HL"
////}
////
////func (c *CPU) ldahli() string {
////	pos := c.Reg.HL.Get()
////	*c.Reg.A.Val() = c.bus.Read(pos)
////	c.Reg.HL.Set(pos + 1)
////
////	return "LD A, (HLI)"
////}
////
////func (c *CPU) dechl() string {
////	value := c.Reg.HL.Get() - 1
////	c.Reg.HL.Set(value)
////
////	return "DEC HL"
////}
////
////func (c *CPU) incl() string {
////	incc.Reg(c.Reg.L.Val())
////
////	return "INC L"
////}
////
////func (c *CPU) decl() string {
////	decc.Reg(c.Reg.L.Val())
////
////	return "DEC L"
////}
////
////func (c *CPU) ldl() string {
////	c.Reg.PC.Inc()
////	*c.Reg.L.Val() = c.bus.Read(c.Reg.PC.Get())
////
////	return fmt.Sprintf("LD L, $%X", c.Reg.L.Get())
////}
////
////func (c *CPU) cpl() string {
////	*c.Reg.A.Val() = ^*c.Reg.A.Val()
////
////	return "CPL"
////}
////
////func (c *CPU) jrnc() string {
////	value := jrCond(!flags.GetFlagC(), 1)
////
////	return fmt.Sprintf("JR NC, $%X", value)
////}
////
////func (c *CPU) ldsp() string {
////	c.Reg.SP.Set(c.bus.Read16(c.Reg.PC.Get()))
////	c.Reg.PC.Inc().Inc()
////
////	return fmt.Sprintf("LD SP, $%X", c.Reg.SP.Get())
////}
////
////func (c *CPU) ldhlda() string {
////	pos := c.Reg.HL.Get()
////	c.bus.Write(pos, c.Reg.A.Get())
////	c.Reg.HL.Set(pos - 1)
////
////	return "LD (HLD), A"
////}
////
////func (c *CPU) incsp() string {
////	c.Reg.SP.Inc()
////
////	return "INC SP"
////}
////
////func (c *CPU) inchlind() string {
////	pos := c.Reg.HL.Get()
////	value := c.bus.Read(pos)
////	flags.AffectFlagZH(value, value+1)
////	flags.SetFlagN(false)
////	c.bus.Write(pos, value+1)
////
////	return "INC (HL)"
////}
////
////func (c *CPU) dechlind() string {
////	pos := c.Reg.HL.Get()
////	value := c.bus.Read(pos)
////	flags.AffectFlagZH(value, value+1)
////	flags.SetFlagN(true)
////	c.bus.Write(pos, value+1)
////
////	return "DEC H"
////}
////
////func (c *CPU) ldhl8() string {
////	c.Reg.PC.Inc()
////	value := c.bus.Read(c.Reg.PC.Get())
////	ldhlind(value)
////
////	return fmt.Sprintf("LD (HL), $%X", c.Reg.H.Get())
////}
////
////// SetBit Carry Flag
////// Flags N and H are set to Zero
////func (c *CPU) scf() string {
////	flags.SetFlagC(true)
////	flags.SetFlagN(false)
////	flags.SetFlagH(false)
////
////	return "SCF"
////}
////
////func (c *CPU) jrc() string {
////	value := jrCond(flags.GetFlagC(), 1)
////
////	return fmt.Sprintf("JR C, $%X", value)
////}
////
////func (c *CPU) addhlsp() string {
////	addhlc.Reg16(c.Reg.SP.Get())
////
////	return "ADD HL, SP"
////}
////
////func (c *CPU) ldahld() string {
////	pos := c.Reg.HL.Get()
////	*c.Reg.A.Val() = c.bus.Read(pos)
////	c.Reg.HL.Set(pos - 1)
////
////	return "LD A, (HLD)"
////}
////
////func (c *CPU) decsp() string {
////	c.Reg.SP.Dec()
////
////	return "DEC SP"
////}
////
////func (c *CPU) inca() string {
////	incc.Reg(c.Reg.A.Val())
////
////	return "INC A"
////}
////
////func (c *CPU) deca() string {
////	decc.Reg(c.Reg.A.Val())
////
////	return "DEC A"
////}
////
////func (c *CPU) lda() string {
////	*c.Reg.A.Val() = c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////
////	return fmt.Sprintf("LD A, $%.2X", c.Reg.A.Get())
////}
////
////// Complement Carry Flag
////func (c *CPU) ccf() string {
////	c.flags.SetFlagC(!c.flags.GetFlagC())
////
////	return "CCF"
////}
////
////func (c *CPU) ldbb() string {
////
////	return "LD B, B"
////}
////
////func (c *CPU) ldbc() string {
////	*c.Reg.B.Val() = c.Reg.C.Get()
////
////	return "LD B, C"
////}
////
////func (c *CPU) ldbd() string {
////	*c.Reg.B.Val() = c.Reg.D.Get()
////
////	return "LD B, D"
////}
////
////func (c *CPU) ldbe() string {
////	*c.Reg.B.Val() = c.Reg.E.Get()
////
////	return "LD B, E"
////}
////
////func (c *CPU) ldbh() string {
////	*c.Reg.B.Val() = c.Reg.H.Get()
////
////	return "LD B, H"
////}
////
////func (c *CPU) ldbl() string {
////	*c.Reg.B.Val() = c.Reg.L.Get()
////
////	return "LD B, L"
////}
////
////func (c *CPU) ldbhl() string {
////	*c.Reg.B.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD B, (HL)"
////}
////
////func (c *CPU) ldba() string {
////	*c.Reg.B.Val() = c.Reg.A.Get()
////
////	return "LD B, A"
////}
////
////func (c *CPU) ldcb() string {
////	*c.Reg.C.Val() = c.Reg.B.Get()
////
////	return "LD C, B"
////}
////
////func (c *CPU) ldcc() string {
////
////	return "LD C, C"
////}
////
////func (c *CPU) ldcd() string {
////	*c.Reg.C.Val() = c.Reg.D.Get()
////
////	return "LD C, D"
////}
////
////func (c *CPU) ldce() string {
////	*c.Reg.C.Val() = c.Reg.E.Get()
////
////	return "LD C, E"
////}
////
////func (c *CPU) ldch() string {
////	*c.Reg.C.Val() = c.Reg.H.Get()
////
////	return "LD C, H"
////}
////
////func (c *CPU) ldcl() string {
////	*c.Reg.C.Val() = c.Reg.L.Get()
////
////	return "LD C, L"
////}
////
////func (c *CPU) ldchl() string {
////	*c.Reg.C.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD C, (HL)"
////}
////
////func (c *CPU) ldca() string {
////	*c.Reg.C.Val() = c.Reg.A.Get()
////
////	return "LD C, A"
////}
////
////func (c *CPU) lddb() string {
////	*c.Reg.D.Val() = c.Reg.B.Get()
////
////	return "LD D, B"
////}
////
////func (c *CPU) lddc() string {
////	*c.Reg.D.Val() = c.Reg.C.Get()
////
////	return "LD D, C"
////}
////
////func (c *CPU) lddd() string {
////
////	return "LD D, D"
////}
////
////func (c *CPU) ldde() string {
////	*c.Reg.D.Val() = c.Reg.E.Get()
////
////	return "LD D, E"
////}
////
////func (c *CPU) lddh() string {
////	*c.Reg.D.Val() = c.Reg.H.Get()
////
////	return "LD D, H"
////}
////
////func (c *CPU) lddl() string {
////	*c.Reg.D.Val() = c.Reg.L.Get()
////
////	return "LD D, L"
////}
////
////func (c *CPU) lddhl() string {
////	*c.Reg.D.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD D, (HL)"
////}
////
////func (c *CPU) ldda() string {
////	*c.Reg.D.Val() = c.Reg.A.Get()
////
////	return "LD D, A"
////}
////
////func (c *CPU) ldeb() string {
////	*c.Reg.E.Val() = c.Reg.B.Get()
////
////	return "LD E, B"
////}
////
////func (c *CPU) ldec() string {
////	*c.Reg.E.Val() = c.Reg.C.Get()
////
////	return "LD E, C"
////}
////
////func (c *CPU) lded() string {
////	*c.Reg.E.Val() = c.Reg.D.Get()
////
////	return "LD E, D"
////}
////
////func (c *CPU) ldee() string {
////
////	return "LD E, E"
////}
////
////func (c *CPU) ldeh() string {
////	*c.Reg.E.Val() = c.Reg.H.Get()
////
////	return "LD E, H"
////}
////
////func (c *CPU) ldel() string {
////	*c.Reg.E.Val() = c.Reg.L.Get()
////
////	return "LD E, L"
////}
////
////func (c *CPU) ldehl() string {
////	*c.Reg.E.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD E, (HL)"
////}
////
////func (c *CPU) ldea() string {
////	*c.Reg.E.Val() = c.Reg.A.Get()
////
////	return "LD E, A"
////}
////
////func (c *CPU) ldhb() string {
////	*c.Reg.H.Val() = c.Reg.B.Get()
////
////	return "LD H, B"
////}
////
////func (c *CPU) ldhc() string {
////	*c.Reg.H.Val() = c.Reg.C.Get()
////
////	return "LD H, C"
////}
////
////func (c *CPU) ldhd() string {
////	*c.Reg.H.Val() = c.Reg.D.Get()
////
////	return "LD H, D"
////}
////
////func (c *CPU) ldhe() string {
////	*c.Reg.H.Val() = c.Reg.E.Get()
////
////	return "LD H, E"
////}
////
////func (c *CPU) ldhh() string {
////
////	return "LD H, H"
////}
////
////func (c *CPU) ldhl() string {
////	*c.Reg.H.Val() = c.Reg.L.Get()
////
////	return "LD H, L"
////}
////
////func (c *CPU) ldhhl() string {
////	*c.Reg.H.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD H, (HL)"
////}
////
////func (c *CPU) ldha() string {
////	*c.Reg.H.Val() = c.Reg.A.Get()
////
////	return "LD H, A"
////}
////
////func (c *CPU) ldlb() string {
////	*c.Reg.L.Val() = c.Reg.B.Get()
////
////	return "LD L, B"
////}
////
////func (c *CPU) ldlc() string {
////	*c.Reg.L.Val() = c.Reg.C.Get()
////
////	return "LD L, C"
////}
////
////func (c *CPU) ldld() string {
////	*c.Reg.L.Val() = c.Reg.D.Get()
////
////	return "LD L, D"
////}
////
////func (c *CPU) ldle() string {
////	*c.Reg.L.Val() = c.Reg.E.Get()
////
////	return "LD L, E"
////}
////
////func (c *CPU) ldlh() string {
////	*c.Reg.L.Val() = c.Reg.H.Get()
////
////	return "LD L, H"
////}
////
////func (c *CPU) ldll() string {
////
////	return "LD L, L"
////}
////
////func (c *CPU) ldlhl() string {
////	*c.Reg.L.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD L, (HL)"
////}
////
////func (c *CPU) ldla() string {
////	*c.Reg.H.Val() = c.Reg.A.Get()
////
////	return "LD L, A"
////}
////
////func (c *CPU) ldhlb() string {
////	ldhlind(c.Reg.B.Get())
////
////	return "LD (HL), B"
////}
////
////func (c *CPU) ldhlc() string {
////	ldhlind(c.Reg.C.Get())
////
////	return "LD (HL), C"
////}
////
////func (c *CPU) ldhld() string {
////	ldhlind(c.Reg.D.Get())
////
////	return "LD (HL), D"
////}
////
////func (c *CPU) ldhle() string {
////	ldhlind(c.Reg.E.Get())
////
////	return "LD (HL), E"
////}
////
////func (c *CPU) ldhlh() string {
////	ldhlind(c.Reg.H.Get())
////
////	return "LD (HL), H"
////}
////
////func (c *CPU) ldhll() string {
////	ldhlind(c.Reg.L.Get())
////
////	return "LD (HL), L"
////}
////
////// Halt pauses CPU execution until an interruption take place
////func (c *CPU) halt() string {
////	/*
////		Halt stops the CPU execution, and resumes when an interrupt is pending. An interruption is considered
////		pending when an interrupt is enabled and its flag is set to one, that is IE && IF !=0 for a certain interrupt.
////		The following assumptions take place:
////		With IME = 1:
////		1. With pending interrupt, cpu will not halt.
////		2. The expected behaviour would be the CPU jumping to next instruction
////		2. Interrupt handling takes place
////
////		With IME = 0:
////		1. If no interrupt pending, halt will execute and cpu will pause until an interrupt becomes pending.
////		Interrupt will not be handled as expected with the master interrupt not enabled
////		2. If an interrupt is pending,halt immediately exits. Halt bug might take place as explained below
////
////		HALT Bug:
////		Take place as IME = 0 with an interrupt is pending. Two of the following scenarios can take place
////		1. With no IE instruction before HALT, the byte after halt instruction is read twice
////		2. With IE instruction before HALT (with IME delay affect taking place), the interrupt handler takes place.
////		The handler, however, returns to halt after serviced, causing the cpu to pause again.
////	*/
////
////	// TODO Give the instructions above an index, then refer those index where simulated in cpu and instruction
////	// file
////
////	c.isHalt = true
////	return "HALT"
////}
////
////func (c *CPU) ldhla() string {
////	ldhlind(c.Reg.A.Get())
////
////	return "LD (HL), A"
////}
////
////func (c *CPU) ldab() string {
////	*c.Reg.A.Val() = c.Reg.B.Get()
////
////	return "LD A, B"
////}
////
////func (c *CPU) ldac() string {
////	*c.Reg.A.Val() = c.Reg.C.Get()
////
////	return "LD A, C"
////}
////
////func (c *CPU) ldad() string {
////	*c.Reg.A.Val() = c.Reg.D.Get()
////
////	return "LD A, D"
////}
////
////func (c *CPU) ldae() string {
////	*c.Reg.A.Val() = c.Reg.E.Get()
////
////	return "LD A, E"
////}
////
////func (c *CPU) ldah() string {
////	*c.Reg.A.Val() = c.Reg.H.Get()
////
////	return "LD A, H"
////}
////
////func (c *CPU) ldal() string {
////	*c.Reg.A.Val() = c.Reg.L.Get()
////
////	return "LD A, L"
////}
////
////func (c *CPU) ldahl() string {
////	*c.Reg.A.Val() = c.bus.Read(c.Reg.HL.Get())
////
////	return "LD A, (HL)"
////}
////
////func (c *CPU) ldaa() string {
////
////	return "LD A, A"
////}
////
////func (c *CPU) addab() string {
////	adda(c.Reg.B.Get())
////
////	return "ADD A, B"
////}
////
////func (c *CPU) addac() string {
////	adda(c.Reg.C.Get())
////
////	return "ADD A, C"
////}
////
////func (c *CPU) addad() string {
////	adda(c.Reg.D.Get())
////
////	return "ADD A, D"
////}
////
////func (c *CPU) addae() string {
////	adda(c.Reg.E.Get())
////
////	return "ADD A,E"
////}
////
////func (c *CPU) addah() string {
////	adda(c.Reg.H.Get())
////
////	return "ADD A, H"
////}
////
////func (c *CPU) addal() string {
////	adda(c.Reg.L.Get())
////
////	return "ADD A, L"
////}
////
////func (c *CPU) addahl() string {
////	adda(c.bus.Read(c.Reg.HL.Get()))
////
////	return "ADD A, (HL)"
////}
////
////func (c *CPU) addaa() string {
////	adda(c.Reg.A.Get())
////
////	return "ADD A, A"
////}
////
////func (c *CPU) adcab() string {
////	adda(c.Reg.B.Get())
////
////	return "ADC A, B"
////}
////
////func (c *CPU) adcac() string {
////	adca(c.Reg.C.Get())
////
////	return "ADC A, C"
////}
////
////func (c *CPU) adcad() string {
////	adca(c.Reg.D.Get())
////
////	return "ADC A, D"
////}
////
////func (c *CPU) adcae() string {
////	adca(c.Reg.E.Get())
////
////	return "ADC A,E"
////}
////
////func (c *CPU) adcah() string {
////	adca(c.Reg.H.Get())
////
////	return "ADC A, H"
////}
////
////func (c *CPU) adcal() string {
////	adca(c.Reg.L.Get())
////
////	return "ADC A, L"
////}
////
////func (c *CPU) adcahl() string {
////	adca(c.bus.Read(c.Reg.HL.Get()))
////
////	return "ADC A, (HL)"
////}
////
////func (c *CPU) adcaa() string {
////	adca(c.Reg.A.Get())
////
////	return "ADC A, A"
////}
////
////func (c *CPU) subab() string {
////	suba(c.Reg.B.Get())
////
////	return "SUB A, B"
////}
////
////func (c *CPU) subac() string {
////	suba(c.Reg.C.Get())
////
////	return "SUB A, C"
////}
////
////func (c *CPU) subad() string {
////	suba(c.Reg.D.Get())
////
////	return "SUB A, D"
////}
////
////func (c *CPU) subae() string {
////	suba(c.Reg.E.Get())
////
////	return "SUB A, E"
////}
////
////func (c *CPU) subah() string {
////	suba(c.Reg.H.Get())
////
////	return "SUB A, H"
////}
////
////func (c *CPU) subal() string {
////	suba(c.Reg.L.Get())
////
////	return "SUB A, L"
////}
////
////func (c *CPU) subahl() string {
////	suba(c.bus.Read(c.Reg.HL.Get()))
////
////	return "SUB A, (HL)"
////}
////
////func (c *CPU) subaa() string {
////	suba(c.Reg.A.Get())
////
////	return "SUB A, A"
////}
////
////func (c *CPU) sbcab() string {
////	sbca(c.Reg.B.Get())
////
////	return "SBC A, B"
////}
////
////func (c *CPU) sbcac() string {
////	sbca(c.Reg.C.Get())
////
////	return "SBC A, C"
////}
////
////func (c *CPU) sbcad() string {
////	sbca(c.Reg.D.Get())
////
////	return "SBC A, D"
////}
////
////func (c *CPU) sbcae() string {
////	sbca(c.Reg.E.Get())
////
////	return "SBC A, E"
////}
////
////func (c *CPU) sbcah() string {
////	sbca(c.Reg.H.Get())
////
////	return "SBC A, H"
////}
////
////func (c *CPU) sbcal() string {
////	sbca(c.Reg.L.Get())
////
////	return "SBC A, L"
////}
////
////func (c *CPU) sbcahl() string {
////	sbca(c.bus.Read(c.Reg.HL.Get()))
////
////	return "SBC A, (HL)"
////}
////
////func (c *CPU) sbcaa() string {
////	sbca(c.Reg.A.Get())
////
////	return "SBC A, A"
////}
////
////func (c *CPU) andab() string {
////	anda(c.Reg.B.Get())
////
////	return "AND A, B"
////}
////
////func (c *CPU) andac() string {
////	anda(c.Reg.C.Get())
////
////	return "AND A, C"
////}
////
////func (c *CPU) andad() string {
////	anda(c.Reg.D.Get())
////
////	return "AND A, D"
////}
////
////func (c *CPU) andae() string {
////	anda(c.Reg.E.Get())
////
////	return "AND A, E"
////}
////
////func (c *CPU) andah() string {
////	anda(c.Reg.H.Get())
////
////	return "AND A, H"
////}
////
////func (c *CPU) andal() string {
////	anda(c.Reg.L.Get())
////
////	return "AND A, L"
////}
////
////func (c *CPU) andahl() string {
////	anda(c.bus.Read(c.Reg.HL.Get()))
////
////	return "AND A, (HL)"
////}
////
////func (c *CPU) andaa() string {
////	anda(c.Reg.A.Get())
////
////	return "AND A, B"
////}
////
////func (c *CPU) xorab() string {
////	xora(c.Reg.B.Get())
////
////	return "XOR A, B"
////}
////
////func (c *CPU) xorac() string {
////	xora(c.Reg.C.Get())
////
////	return "XOR A, C"
////}
////
////func (c *CPU) xorad() string {
////	xora(c.Reg.D.Get())
////
////	return "XOR A, D"
////}
////
////func (c *CPU) xorae() string {
////	xora(c.Reg.E.Get())
////
////	return "XOR A, E"
////}
////
////func (c *CPU) xorah() string {
////	xora(c.Reg.H.Get())
////
////	return "XOR A, H"
////}
////
////func (c *CPU) xoral() string {
////	xora(c.Reg.L.Get())
////
////	return "XOR A, L"
////}
////
////func (c *CPU) xorahl() string {
////	xora(c.bus.Read(c.Reg.HL.Get()))
////
////	return "XOR A, (HL)"
////}
////
////func (c *CPU) xoraa() string {
////	xora(c.Reg.A.Get())
////
////	return "XOR A, A"
////}
////
////func (c *CPU) orab() string {
////	ora(c.Reg.B.Get())
////
////	return "OR A, B"
////}
////
////func (c *CPU) orac() string {
////	ora(c.Reg.C.Get())
////
////	return "OR A, C"
////}
////
////func (c *CPU) orad() string {
////	ora(c.Reg.D.Get())
////
////	return "OR A, D"
////}
////
////func (c *CPU) orae() string {
////	ora(c.Reg.E.Get())
////
////	return "OR A, E"
////}
////
////func (c *CPU) orah() string {
////	ora(c.Reg.H.Get())
////
////	return "OR A, H"
////}
////
////func (c *CPU) oral() string {
////	ora(c.Reg.L.Get())
////
////	return "OR A, L"
////}
////
////func (c *CPU) orahl() string {
////	ora(c.bus.Read(c.Reg.HL.Get()))
////
////	return "OR A, (HL)"
////}
////
////func (c *CPU) oraa() string {
////	ora(c.Reg.A.Get())
////
////	return "OR A, A"
////}
////
////func (c *CPU) cpab() string {
////	cpa(c.Reg.B.Get())
////
////	return "CP A, B"
////}
////
////func (c *CPU) cpac() string {
////	cpa(c.Reg.C.Get())
////
////	return "CP A, C"
////}
////
////func (c *CPU) cpad() string {
////	cpa(c.Reg.C.Get())
////
////	return "CP A, D"
////}
////
////func (c *CPU) cpae() string {
////	cpa(c.Reg.E.Get())
////
////	return "CP A, E"
////}
////
////func (c *CPU) cpah() string {
////	cpa(c.Reg.H.Get())
////
////	return "CP A, H"
////}
////
////func (c *CPU) cpal() string {
////	cpa(c.Reg.L.Get())
////
////	return "CP A, L"
////}
////
////func (c *CPU) cpahl() string {
////	cpa(c.bus.Read(c.Reg.HL.Get()))
////
////	return "CP A, (HL)"
////}
////
////func (c *CPU) cpaa() string {
////	cpa(c.Reg.A.Get())
////
////	return "CP A, A"
////}
////
////func (c *CPU) retnz() string {
////	retCond(!c.flags.GetFlagZ(), 3)
////
////	return "RET NZ"
////}
////
////func (c *CPU) popbc() string {
////	pop16(c.Reg.B.Val(), c.Reg.C.Val())
////
////	return "POP BC"
////}
////
////func (c *CPU) jpnz() string {
////	value := jpCond(!c.flags.GetFlagZ(), 1)
////
////	return fmt.Sprintf("JP, NZ, $%X", value)
////}
////
////func (c *CPU) jp() string {
////	value := jpCond(true, 0)
////
////	return fmt.Sprintf("JP $%.4X", value)
////}
////
////func (c *CPU) callnz() string {
////	value := callCond(!c.flags.GetFlagZ(), 3)
////
////	return fmt.Sprintf("CALL NZ, $%X", value)
////}
////
////func (c *CPU) pushbc() string {
////	push16(c.Reg.B.Get(), c.Reg.C.Get())
////
////	return "PUSH BC"
////}
////
////func (c *CPU) adda8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	adda(value)
////
////	return fmt.Sprintf("ADD A, $%X", value)
////}
////
////func (c *CPU) rst00() string {
////	callmem(00)
////
////	return "RST $00"
////}
////
////func (c *CPU) retz() string {
////	retCond(c.flags.GetFlagZ(), 3)
////
////	return "RET Z"
////}
////
////func (c *CPU) ret() string {
////	retCond(true, 0)
////
////	return "RET"
////}
////
////func (c *CPU) jpz() string {
////	value := jpCond(c.flags.GetFlagZ(), 1)
////
////	return fmt.Sprintf("JP Z, $%X", value)
////}
////
////func (c *CPU) prefixcb() {
////	// Fetch instruction
////	c.curOP = c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////
////	// Decode
////	instruction := c.cbInst[c.curOP]
////
////	// Execute Operation
////	c.ticks += instruction.ticks
////	instruction.execute()
////}
////
////func (c *CPU) callz() string {
////	value := callCond(c.flags.GetFlagZ(), 3)
////
////	return fmt.Sprintf("CALL Z, $%X", value)
////}
////
////func (c *CPU) call() string {
////	value := callCond(true, 0)
////
////	return fmt.Sprintf("CALL $%X", value)
////}
////
////func (c *CPU) adca8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	adca(value)
////
////	return fmt.Sprintf("ADC A, $%X", value)
////}
////
////func (c *CPU) rst08() string {
////	callmem(0x08)
////
////	return "RST $08"
////}
////
////func (c *CPU) retnc() string {
////	retCond(!c.flags.GetFlagC(), 3)
////
////	return "RET NC"
////}
////
////func (c *CPU) popde() string {
////	pop16(c.Reg.D.Val(), c.Reg.E.Val())
////
////	return "POP DE"
////}
////
////func (c *CPU) jpnc() string {
////	value := jpCond(!c.flags.GetFlagC(), 1)
////
////	return fmt.Sprintf("JP NC, $%X", value)
////}
////
////func (c *CPU) callnc() string {
////	value := callCond(!c.flags.GetFlagC(), 3)
////
////	return fmt.Sprintf("CALL NC, $%X", value)
////}
////
////func (c *CPU) pushde() string {
////	push16(c.Reg.D.Get(), c.Reg.E.Get())
////
////	return "PUSH DE"
////}
////
////func (c *CPU) suba8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	suba(value)
////
////	return fmt.Sprintf("SUB A, $%X", value)
////}
////
////func (c *CPU) rst10() string {
////	callmem(0x10)
////
////	return "RST $10"
////}
////
////func (c *CPU) retc() string {
////	retCond(c.flags.GetFlagC(), 3)
////
////	return "RET C"
////}
////
////func (c *CPU) reti() string {
////	// Equivalent to executing ei() followed by ret()
////	c.Reg.IME = true
////	retCond(true, 0)
////
////	return "RETI"
////}
////
////func (c *CPU) jpc() string {
////	value := jpCond(c.flags.GetFlagC(), 1)
////
////	return fmt.Sprintf("JP C, $%X", value)
////}
////
////func (c *CPU) callc() string {
////	value := callCond(c.flags.GetFlagC(), 3)
////
////	return fmt.Sprintf("CALL C, $%X", value)
////}
////
////func (c *CPU) sbca8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	sbca(value)
////
////	return fmt.Sprintf("SBC A, $%X", value)
////}
////
////func (c *CPU) rst18() string {
////	callmem(0x18)
////
////	return "RST $18"
////}
////
////func (c *CPU) ldff8a() string {
////	var value uint8 = c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	var address uint16 = 0xFF00 + uint16(value)
////	ldmem(address, c.Reg.A.Get())
////
////	return fmt.Sprintf("LD (FF00 + $%.2X), A", value)
////}
////
////func (c *CPU) pophl() string {
////	pop16(c.Reg.H.Val(), c.Reg.L.Val())
////
////	return "POP HL"
////}
////
////func (c *CPU) ldffca() string {
////	var pos uint16 = 0xFF00
////	pos += uint16(c.Reg.C.Get())
////	ldmem(pos, c.Reg.A.Get())
////
////	return "LD (FF00 + C), A"
////}
////
////func (c *CPU) pushhl() string {
////	push16(c.Reg.H.Get(), c.Reg.L.Get())
////
////	return "PUSH HL"
////}
////
////func (c *CPU) anda8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	anda(value)
////
////	return fmt.Sprintf("AND A, $%X", value)
////}
////
////func (c *CPU) rst20() string {
////	callmem(0x20)
////
////	return "RST $20"
////}
////
////func (c *CPU) addsp() string {
////	var value int8 = int8(c.bus.Read(c.Reg.PC.Get()))
////	c.Reg.PC.Inc()
////	c.Reg.SP.Set(c.Reg.SP.Get() + uint16(value))
////
////	return fmt.Sprintf("ADD SP, $%X", value)
////}
////
////func (c *CPU) jphl() string {
////	c.Reg.PC.Set(c.Reg.HL.Get())
////
////	return "JP HL"
////}
////
////func (c *CPU) ld16a() string {
////	value := c.bus.Read16(c.Reg.PC.Get())
////	c.Reg.PC.Inc().Inc()
////	ldmem(value, c.Reg.A.Get())
////
////	return fmt.Sprintf("LD ($%X), A", value)
////}
////
////func (c *CPU) xora8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	xora(value)
////
////	return fmt.Sprintf("XOR A, $%X", value)
////}
////
////func (c *CPU) rst28() string {
////	callmem(0x28)
////
////	return "RST $28"
////}
////
////func (c *CPU) ldaff8() string {
////	var pos uint16 = 0xFF00
////	var value uint8 = c.bus.Read(c.Reg.PC.Get())
////	pos += uint16(value)
////	c.Reg.PC.Inc()
////	*c.Reg.A.Val() = c.bus.Read(pos)
////
////	return fmt.Sprintf("LD A, (FF00 + $%X)", value)
////}
////
////func (c *CPU) popaf() string {
////	pop16(c.Reg.A.Val(), c.flags.Val())
////
////	return "POP AF"
////}
////
////func (c *CPU) ldaffc() string {
////	var pos uint16 = 0xFF00
////	pos += uint16(c.Reg.C.Get())
////	*c.Reg.A.Val() = c.bus.Read(pos)
////
////	return "LD A, (FF00 + C)"
////}
////
////func (c *CPU) di() string {
////	c.Reg.IME = false
////	// Cancels any delayed IME is set by IME
////	c.performIME = false
////
////	return "DI"
////}
////
////func (c *CPU) pushaf() string {
////	push16(c.Reg.A.Get(), c.flags.Get())
////
////	return "PUSH AF"
////}
////
////func (c *CPU) ora8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	ora(value)
////
////	return fmt.Sprintf("OR A, $%X", value)
////}
////
////func (c *CPU) rst30() string {
////	callmem(0x30)
////
////	return "RST $30"
////}
////
////func (c *CPU) ldhlsp8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	newValue := c.Reg.SP.Get() + uint16(value)
////	c.flags.AffectFlagHC16(c.Reg.HL.Get(), newValue)
////	c.Reg.HL.Set(newValue)
////
////	return fmt.Sprintf("LD HL, SP + $%X", value)
////}
////
////func (c *CPU) ldsphl() string {
////	c.Reg.SP.Set(c.Reg.HL.Get())
////
////	return "LD SP, HL"
////}
////
////func (c *CPU) lda16() string {
////	value := c.bus.Read16(c.Reg.PC.Get())
////	c.Reg.PC.Inc().Inc()
////	*c.Reg.A.Val() = c.bus.Read(value)
////
////	return fmt.Sprintf("LD A, ($%X)", value)
////}
////
////func (c *CPU) ei() string {
////	c.performIME = true
////
////	return "EI"
////}
////
////func (c *CPU) cpa8() string {
////	value := c.bus.Read(c.Reg.PC.Get())
////	c.Reg.PC.Inc()
////	cpa(value)
////
////	return fmt.Sprintf("CP A, $%X", value)
////}
////
////func (c *CPU) rst38() string {
////	callmem(0x38)
////
////	return "RST $38"
////}
////
////func (c *CPU) illegalop() string {
////	//TODO make it optional to crash
////
////	return "ILLEGAL OP Used"
////}
////
//////endregion OpCode Functions

////region CP Prefixed OpCode Functions
//
//func rlcb() string {
//	rlcReg(Reg.B.Val())
//
//	return "RLC B"
//}
//
//func rlcc() string {
//	rlcReg(Reg.C.Val())
//
//	return "RLC C"
//}
//
//func rlcd() string {
//	rlcReg(Reg.D.Val())
//
//	return "RLC D"
//}
//
//func rlce() string {
//	rlcReg(Reg.E.Val())
//
//	return "RLC E"
//}
//
//func rlch() string {
//	rlcReg(Reg.H.Val())
//
//	return "RLC H"
//}
//
//func rlcl() string {
//	rlcReg(Reg.L.Val())
//
//	return "RLC L"
//}
//
//func rlchl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = rlcVal(value)
//	bus.Write(pos, value)
//
//	return "RLC (HL)"
//}
//
//func cbrlca() string {
//	rlcReg(Reg.A.Val())
//
//	return "RLC A"
//}
//
//func rrcb() string {
//	rrcReg(Reg.B.Val())
//
//	return "RRC B"
//}
//
//func rrcc() string {
//	rrcReg(Reg.C.Val())
//
//	return "RRC C"
//}
//
//func rrcd() string {
//	rrcReg(Reg.D.Val())
//
//	return "RRC D"
//}
//
//func rrce() string {
//	rrcReg(Reg.E.Val())
//
//	return "RRC E"
//}
//
//func rrch() string {
//	rrcReg(Reg.H.Val())
//
//	return "RRC H"
//}
//
//func rrcl() string {
//	rrcReg(Reg.L.Val())
//
//	return "RRC L"
//}
//
//func rrchl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = rrcVal(value)
//	bus.Write(pos, value)
//
//	return "RRC L"
//}
//
//func cbrrca() string {
//	rrcReg(Reg.A.Val())
//
//	return "RRC A"
//}
//
//func rlb() string {
//	rlReg(Reg.B.Val())
//
//	return "RL B"
//}
//
//func rlc() string {
//	rlReg(Reg.C.Val())
//
//	return "RL C"
//}
//
//func rld() string {
//	rlReg(Reg.D.Val())
//
//	return "RL D"
//}
//
//func rle() string {
//	rlReg(Reg.E.Val())
//
//	return "RL E"
//}
//
//func rlh() string {
//	rlReg(Reg.H.Val())
//
//	return "RL H"
//}
//
//func rll() string {
//	rlReg(Reg.L.Val())
//
//	return "RL L"
//}
//
//func rlhl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = rl(value)
//	bus.Write(pos, value)
//
//	return "RL (HL)"
//}
//
//func cbrla() string {
//	rlReg(Reg.A.Val())
//
//	return "RL A"
//}
//
//func rrb() string {
//	rrReg(Reg.B.Val())
//
//	return "RR B"
//}
//
//func rrc() string {
//	rrReg(Reg.C.Val())
//
//	return "RR C"
//}
//
//func rrd() string {
//	rrReg(Reg.D.Val())
//
//	return "RR D"
//}
//
//func rre() string {
//	rrReg(Reg.E.Val())
//
//	return "RR E"
//}
//
//func rrh() string {
//	rrReg(Reg.H.Val())
//
//	return "RR H"
//}
//
//func rrl() string {
//	rrReg(Reg.L.Val())
//
//	return "RR L"
//}
//
//func rrhl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = rr(value)
//	bus.Write(pos, value)
//
//	return "RR L"
//}
//
//func cbrra() string {
//	rrReg(Reg.A.Val())
//
//	return "RR A"
//}
//
//func slab() string {
//	slaReg(Reg.B.Val())
//
//	return "SLA B"
//}
//
//func slac() string {
//	slaReg(Reg.C.Val())
//
//	return "SLA C"
//}
//
//func slad() string {
//	slaReg(Reg.D.Val())
//
//	return "SLA D"
//}
//
//func slae() string {
//	slaReg(Reg.E.Val())
//
//	return "SLA E"
//}
//
//func slah() string {
//	slaReg(Reg.H.Val())
//
//	return "SLA H"
//}
//
//func slal() string {
//	slaReg(Reg.L.Val())
//
//	return "SLA L"
//}
//
//func slahl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = sla(value)
//	bus.Write(pos, value)
//
//	return "SLA (HL)"
//}
//
//func slaa() string {
//	slaReg(Reg.A.Val())
//
//	return "SLA A"
//}
//
//func srab() string {
//	sraReg(Reg.B.Val())
//
//	return "SRA B"
//}
//
//func srac() string {
//	sraReg(Reg.C.Val())
//
//	return "SRA C"
//}
//
//func srad() string {
//	sraReg(Reg.D.Val())
//
//	return "SRA D"
//}
//
//func srae() string {
//	sraReg(Reg.E.Val())
//
//	return "SRA E"
//}
//
//func srah() string {
//	sraReg(Reg.H.Val())
//
//	return "SRA H"
//}
//
//func sral() string {
//	sraReg(Reg.L.Val())
//
//	return "SRA L"
//}
//
//func srahl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = sra(value)
//	bus.Write(pos, value)
//
//	return "SRA (HL)"
//}
//
//func sraa() string {
//	sraReg(Reg.A.Val())
//
//	return "SRA A"
//}
//
//func swapb() string {
//	swapReg(Reg.B.Val())
//
//	return "SWAP B"
//}
//
//func swapc() string {
//	swapReg(Reg.C.Val())
//
//	return "SWAP C"
//}
//
//func swapd() string {
//	swapReg(Reg.D.Val())
//
//	return "SWAP D"
//}
//
//func swape() string {
//	swapReg(Reg.E.Val())
//
//	return "SWAP E"
//}
//
//func swaph() string {
//	swapReg(Reg.H.Val())
//
//	return "SWAP H"
//}
//
//func swapl() string {
//	swapReg(Reg.L.Val())
//
//	return "SWAP L"
//}
//
//func swaphl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = swap(value)
//	bus.Write(pos, value)
//
//	return "SWAP (HL)"
//}
//
//func swapa() string {
//	swapReg(Reg.A.Val())
//
//	return "SWAP A"
//}
//
//func srlb() string {
//	srlReg(Reg.B.Val())
//
//	return "SRL B"
//}
//
//func srlc() string {
//	srlReg(Reg.C.Val())
//
//	return "SRL C"
//}
//
//func srld() string {
//	srlReg(Reg.D.Val())
//
//	return "SRL D"
//}
//
//func srle() string {
//	srlReg(Reg.E.Val())
//
//	return "SRL E"
//}
//
//func srlh() string {
//	srlReg(Reg.H.Val())
//
//	return "SRL H"
//}
//
//func srll() string {
//	srlReg(Reg.L.Val())
//
//	return "SRL L"
//}
//
//func srlhl() string {
//	pos := Reg.HL.Get()
//	value := bus.Read(pos)
//	value = srl(value)
//	bus.Write(pos, value)
//
//	return "SRL (HL)"
//}
//
//func srla() string {
//	srlReg(Reg.A.Val())
//
//	return "SRL A"
//}
//
//func bit0b() string {
//	return bitNumReg(0, Reg.B.Val(), "B")
//}
//
//func bit0c() string {
//	return bitNumReg(0, Reg.C.Val(), "C")
//}
//
//func bit0d() string {
//	return bitNumReg(0, Reg.D.Val(), "D")
//}
//
//func bit0e() string {
//	return bitNumReg(0, Reg.E.Val(), "E")
//}
//
//func bit0h() string {
//	return bitNumReg(0, Reg.H.Val(), "H")
//}
//
//func bit0l() string {
//	return bitNumReg(0, Reg.L.Val(), "L")
//}
//
//func bit0hl() string {
//	return bitNumHL(0)
//}
//
//func bit0a() string {
//	return bitNumReg(0, Reg.A.Val(), "A")
//}
//
//func bit1b() string {
//	return bitNumReg(1, Reg.B.Val(), "B")
//}
//
//func bit1c() string {
//	return bitNumReg(1, Reg.C.Val(), "C")
//}
//
//func bit1d() string {
//	return bitNumReg(1, Reg.D.Val(), "D")
//}
//
//func bit1e() string {
//	return bitNumReg(1, Reg.E.Val(), "E")
//}
//
//func bit1h() string {
//	return bitNumReg(1, Reg.H.Val(), "H")
//}
//
//func bit1l() string {
//	return bitNumReg(1, Reg.L.Val(), "L")
//}
//
//func bit1hl() string {
//	return bitNumHL(1)
//}
//
//func bit1a() string {
//	return bitNumReg(1, Reg.A.Val(), "A")
//}
//
//func bit2b() string {
//	return bitNumReg(2, Reg.B.Val(), "B")
//}
//
//func bit2c() string {
//	return bitNumReg(2, Reg.C.Val(), "C")
//}
//
//func bit2d() string {
//	return bitNumReg(2, Reg.D.Val(), "D")
//}
//
//func bit2e() string {
//	return bitNumReg(2, Reg.E.Val(), "E")
//}
//
//func bit2h() string {
//	return bitNumReg(2, Reg.H.Val(), "H")
//}
//
//func bit2l() string {
//	return bitNumReg(2, Reg.L.Val(), "L")
//}
//
//func bit2hl() string {
//	return bitNumHL(2)
//}
//
//func bit2a() string {
//	return bitNumReg(2, Reg.A.Val(), "A")
//}
//
//func bit3b() string {
//	return bitNumReg(3, Reg.B.Val(), "B")
//}
//
//func bit3c() string {
//	return bitNumReg(3, Reg.C.Val(), "C")
//}
//
//func bit3d() string {
//	return bitNumReg(3, Reg.D.Val(), "D")
//}
//
//func bit3e() string {
//	return bitNumReg(3, Reg.E.Val(), "E")
//}
//
//func bit3h() string {
//	return bitNumReg(3, Reg.H.Val(), "H")
//}
//
//func bit3l() string {
//	return bitNumReg(3, Reg.L.Val(), "L")
//}
//
//func bit3hl() string {
//	return bitNumHL(3)
//}
//
//func bit3a() string {
//	return bitNumReg(3, Reg.A.Val(), "A")
//}
//
//func bit4b() string {
//	return bitNumReg(4, Reg.B.Val(), "B")
//}
//
//func bit4c() string {
//	return bitNumReg(4, Reg.C.Val(), "C")
//}
//
//func bit4d() string {
//	return bitNumReg(4, Reg.D.Val(), "D")
//}
//
//func bit4e() string {
//	return bitNumReg(4, Reg.E.Val(), "E")
//}
//
//func bit4h() string {
//	return bitNumReg(4, Reg.H.Val(), "H")
//}
//
//func bit4l() string {
//	return bitNumReg(4, Reg.L.Val(), "L")
//}
//
//func bit4hl() string {
//	return bitNumHL(4)
//}
//
//func bit4a() string {
//	return bitNumReg(4, Reg.A.Val(), "A")
//}
//
//func bit5b() string {
//	return bitNumReg(5, Reg.B.Val(), "B")
//}
//
//func bit5c() string {
//	return bitNumReg(5, Reg.C.Val(), "C")
//}
//
//func bit5d() string {
//	return bitNumReg(5, Reg.D.Val(), "D")
//}
//
//func bit5e() string {
//	return bitNumReg(5, Reg.E.Val(), "E")
//}
//
//func bit5h() string {
//	return bitNumReg(5, Reg.H.Val(), "H")
//}
//
//func bit5l() string {
//	return bitNumReg(5, Reg.L.Val(), "L")
//}
//
//func bit5hl() string {
//	return bitNumHL(5)
//}
//
//func bit5a() string {
//	return bitNumReg(5, Reg.A.Val(), "A")
//}
//
//func bit6b() string {
//	return bitNumReg(6, Reg.B.Val(), "B")
//}
//
//func bit6c() string {
//	return bitNumReg(6, Reg.C.Val(), "C")
//}
//
//func bit6d() string {
//	return bitNumReg(6, Reg.D.Val(), "D")
//}
//
//func bit6e() string {
//	return bitNumReg(6, Reg.E.Val(), "E")
//}
//
//func bit6h() string {
//	return bitNumReg(6, Reg.H.Val(), "H")
//}
//
//func bit6l() string {
//	return bitNumReg(6, Reg.L.Val(), "L")
//}
//
//func bit6hl() string {
//	return bitNumHL(6)
//}
//
//func bit6a() string {
//	return bitNumReg(6, Reg.A.Val(), "A")
//}
//
//func bit7b() string {
//	return bitNumReg(7, Reg.B.Val(), "B")
//}
//
//func bit7c() string {
//	return bitNumReg(7, Reg.C.Val(), "C")
//}
//
//func bit7d() string {
//	return bitNumReg(7, Reg.D.Val(), "D")
//}
//
//func bit7e() string {
//	return bitNumReg(7, Reg.E.Val(), "E")
//}
//
//func bit7h() string {
//	return bitNumReg(7, Reg.H.Val(), "H")
//}
//
//func bit7l() string {
//	return bitNumReg(7, Reg.L.Val(), "L")
//}
//
//func bit7hl() string {
//	return bitNumHL(7)
//}
//
//func bit7a() string {
//	return bitNumReg(7, Reg.A.Val(), "A")
//}
//
//func res0b() string {
//	return resNumReg(0, Reg.B.Val(), "B")
//}
//
//func res0c() string {
//	return resNumReg(0, Reg.C.Val(), "C")
//}
//
//func res0d() string {
//	return resNumReg(0, Reg.D.Val(), "D")
//}
//
//func res0e() string {
//	return resNumReg(0, Reg.E.Val(), "E")
//}
//
//func res0h() string {
//	return resNumReg(0, Reg.H.Val(), "H")
//}
//
//func res0l() string {
//	return resNumReg(0, Reg.L.Val(), "L")
//}
//
//func res0hl() string {
//	return resNumHL(0)
//}
//
//func res0a() string {
//	return resNumReg(0, Reg.A.Val(), "A")
//}
//
//func res1b() string {
//	return resNumReg(1, Reg.B.Val(), "B")
//}
//
//func res1c() string {
//	return resNumReg(1, Reg.C.Val(), "C")
//}
//
//func res1d() string {
//	return resNumReg(1, Reg.D.Val(), "D")
//}
//
//func res1e() string {
//	return resNumReg(1, Reg.E.Val(), "E")
//}
//
//func res1h() string {
//	return resNumReg(1, Reg.H.Val(), "H")
//}
//
//func res1l() string {
//	return resNumReg(1, Reg.L.Val(), "L")
//}
//
//func res1hl() string {
//	return resNumHL(1)
//}
//
//func res1a() string {
//	return resNumReg(1, Reg.A.Val(), "A")
//}
//
//func res2b() string {
//	return resNumReg(2, Reg.B.Val(), "B")
//}
//
//func res2c() string {
//	return resNumReg(2, Reg.C.Val(), "C")
//}
//
//func res2d() string {
//	return resNumReg(2, Reg.D.Val(), "D")
//}
//
//func res2e() string {
//	return resNumReg(2, Reg.E.Val(), "E")
//}
//
//func res2h() string {
//	return resNumReg(2, Reg.H.Val(), "H")
//}
//
//func res2l() string {
//	return resNumReg(2, Reg.L.Val(), "L")
//}
//
//func res2hl() string {
//	return resNumHL(2)
//}
//
//func res2a() string {
//	return resNumReg(2, Reg.A.Val(), "A")
//}
//
//func res3b() string {
//	return resNumReg(3, Reg.B.Val(), "B")
//}
//
//func res3c() string {
//	return resNumReg(3, Reg.C.Val(), "C")
//}
//
//func res3d() string {
//	return resNumReg(3, Reg.D.Val(), "D")
//}
//
//func res3e() string {
//	return resNumReg(3, Reg.E.Val(), "E")
//}
//
//func res3h() string {
//	return resNumReg(3, Reg.H.Val(), "H")
//}
//
//func res3l() string {
//	return resNumReg(3, Reg.L.Val(), "L")
//}
//
//func res3hl() string {
//	return resNumHL(3)
//}
//
//func res3a() string {
//	return resNumReg(3, Reg.A.Val(), "A")
//}
//
//func res4b() string {
//	return resNumReg(4, Reg.B.Val(), "B")
//}
//
//func res4c() string {
//	return resNumReg(4, Reg.C.Val(), "C")
//}
//
//func res4d() string {
//	return resNumReg(4, Reg.D.Val(), "D")
//}
//
//func res4e() string {
//	return resNumReg(4, Reg.E.Val(), "E")
//}
//
//func res4h() string {
//	return resNumReg(4, Reg.H.Val(), "H")
//}
//
//func res4l() string {
//	return resNumReg(4, Reg.L.Val(), "L")
//}
//
//func res4hl() string {
//	return resNumHL(4)
//}
//
//func res4a() string {
//	return resNumReg(4, Reg.A.Val(), "A")
//}
//
//func res5b() string {
//	return resNumReg(5, Reg.B.Val(), "B")
//}
//
//func res5c() string {
//	return resNumReg(5, Reg.C.Val(), "C")
//}
//
//func res5d() string {
//	return resNumReg(5, Reg.D.Val(), "D")
//}
//
//func res5e() string {
//	return resNumReg(5, Reg.E.Val(), "E")
//}
//
//func res5h() string {
//	return resNumReg(5, Reg.H.Val(), "H")
//}
//
//func res5l() string {
//	return resNumReg(5, Reg.L.Val(), "L")
//}
//
//func res5hl() string {
//	return resNumHL(5)
//}
//
//func res5a() string {
//	return resNumReg(5, Reg.A.Val(), "A")
//}
//
//func res6b() string {
//	return resNumReg(6, Reg.B.Val(), "B")
//}
//
//func res6c() string {
//	return resNumReg(6, Reg.C.Val(), "C")
//}
//
//func res6d() string {
//	return resNumReg(6, Reg.D.Val(), "D")
//}
//
//func res6e() string {
//	return resNumReg(6, Reg.E.Val(), "E")
//}
//
//func res6h() string {
//	return resNumReg(6, Reg.H.Val(), "H")
//}
//
//func res6l() string {
//	return resNumReg(6, Reg.L.Val(), "L")
//}
//
//func res6hl() string {
//	return resNumHL(6)
//}
//
//func res6a() string {
//	return resNumReg(6, Reg.A.Val(), "A")
//}
//
//func res7b() string {
//	return resNumReg(7, Reg.B.Val(), "B")
//}
//
//func res7c() string {
//	return resNumReg(7, Reg.C.Val(), "C")
//}
//
//func res7d() string {
//	return resNumReg(7, Reg.D.Val(), "D")
//}
//
//func res7e() string {
//	return resNumReg(7, Reg.E.Val(), "E")
//}
//
//func res7h() string {
//	return resNumReg(7, Reg.H.Val(), "H")
//}
//
//func res7l() string {
//	return resNumReg(7, Reg.L.Val(), "L")
//}
//
//func res7hl() string {
//	return resNumHL(7)
//}
//
//func res7a() string {
//	return resNumReg(7, Reg.A.Val(), "A")
//}
//
//func set0b() string {
//	return setNumReg(0, Reg.B.Val(), "B")
//}
//
//func set0c() string {
//	return setNumReg(0, Reg.C.Val(), "C")
//}
//
//func set0d() string {
//	return setNumReg(0, Reg.D.Val(), "D")
//}
//
//func set0e() string {
//	return setNumReg(0, Reg.E.Val(), "E")
//}
//
//func set0h() string {
//	return setNumReg(0, Reg.H.Val(), "H")
//}
//
//func set0l() string {
//	return setNumReg(0, Reg.L.Val(), "L")
//}
//
//func set0hl() string {
//	return setNumHL(0)
//}
//
//func set0a() string {
//	return setNumReg(0, Reg.A.Val(), "A")
//}
//
//func set1b() string {
//	return setNumReg(1, Reg.B.Val(), "B")
//}
//
//func set1c() string {
//	return setNumReg(1, Reg.C.Val(), "C")
//}
//
//func set1d() string {
//	return setNumReg(1, Reg.D.Val(), "D")
//}
//
//func set1e() string {
//	return setNumReg(1, Reg.E.Val(), "E")
//}
//
//func set1h() string {
//	return setNumReg(1, Reg.H.Val(), "H")
//}
//
//func set1l() string {
//	return setNumReg(1, Reg.L.Val(), "L")
//}
//
//func set1hl() string {
//	return setNumHL(1)
//}
//
//func set1a() string {
//	return setNumReg(1, Reg.A.Val(), "A")
//}
//
//func set2b() string {
//	return setNumReg(2, Reg.B.Val(), "B")
//}
//
//func set2c() string {
//	return setNumReg(2, Reg.C.Val(), "C")
//}
//
//func set2d() string {
//	return setNumReg(2, Reg.D.Val(), "D")
//}
//
//func set2e() string {
//	return setNumReg(2, Reg.E.Val(), "E")
//}
//
//func set2h() string {
//	return setNumReg(2, Reg.H.Val(), "H")
//}
//
//func set2l() string {
//	return setNumReg(2, Reg.L.Val(), "L")
//}
//
//func set2hl() string {
//	return setNumHL(2)
//}
//
//func set2a() string {
//	return setNumReg(2, Reg.A.Val(), "A")
//}
//
//func set3b() string {
//	return setNumReg(3, Reg.B.Val(), "B")
//}
//
//func set3c() string {
//	return setNumReg(3, Reg.C.Val(), "C")
//}
//
//func set3d() string {
//	return setNumReg(3, Reg.D.Val(), "D")
//}
//
//func set3e() string {
//	return setNumReg(3, Reg.E.Val(), "E")
//}
//
//func set3h() string {
//	return setNumReg(3, Reg.H.Val(), "H")
//}
//
//func set3l() string {
//	return setNumReg(3, Reg.L.Val(), "L")
//}
//
//func set3hl() string {
//	return setNumHL(3)
//}
//
//func set3a() string {
//	return setNumReg(3, Reg.A.Val(), "A")
//}
//
//func set4b() string {
//	return setNumReg(4, Reg.B.Val(), "B")
//}
//
//func set4c() string {
//	return setNumReg(4, Reg.C.Val(), "C")
//}
//
//func set4d() string {
//	return setNumReg(4, Reg.D.Val(), "D")
//}
//
//func set4e() string {
//	return setNumReg(4, Reg.E.Val(), "E")
//}
//
//func set4h() string {
//	return setNumReg(4, Reg.H.Val(), "H")
//}
//
//func set4l() string {
//	return setNumReg(4, Reg.L.Val(), "L")
//}
//
//func set4hl() string {
//	return setNumHL(4)
//}
//
//func set4a() string {
//	return setNumReg(4, Reg.A.Val(), "A")
//}
//
//func set5b() string {
//	return setNumReg(5, Reg.B.Val(), "B")
//}
//
//func set5c() string {
//	return setNumReg(5, Reg.C.Val(), "C")
//}
//
//func set5d() string {
//	return setNumReg(5, Reg.D.Val(), "D")
//}
//
//func set5e() string {
//	return setNumReg(5, Reg.E.Val(), "E")
//}
//
//func set5h() string {
//	return setNumReg(5, Reg.H.Val(), "H")
//}
//
//func set5l() string {
//	return setNumReg(5, Reg.L.Val(), "L")
//}
//
//func set5hl() string {
//	return setNumHL(5)
//}
//
//func set5a() string {
//	return setNumReg(5, Reg.A.Val(), "A")
//}
//
//func set6b() string {
//	return setNumReg(6, Reg.B.Val(), "B")
//}
//
//func set6c() string {
//	return setNumReg(6, Reg.C.Val(), "C")
//}
//
//func set6d() string {
//	return setNumReg(6, Reg.D.Val(), "D")
//}
//
//func set6e() string {
//	return setNumReg(6, Reg.E.Val(), "E")
//}
//
//func set6h() string {
//	return setNumReg(6, Reg.H.Val(), "H")
//}
//
//func set6l() string {
//	return setNumReg(6, Reg.L.Val(), "L")
//}
//
//func set6hl() string {
//	return setNumHL(6)
//}
//
//func set6a() string {
//	return setNumReg(6, Reg.A.Val(), "A")
//}
//
//func set7b() string {
//	return setNumReg(7, Reg.B.Val(), "B")
//}
//
//func set7c() string {
//	return setNumReg(7, Reg.C.Val(), "C")
//}
//
//func set7d() string {
//	return setNumReg(7, Reg.D.Val(), "D")
//}
//
//func set7e() string {
//	return setNumReg(7, Reg.E.Val(), "E")
//}
//
//func set7h() string {
//	return setNumReg(7, Reg.H.Val(), "H")
//}
//
//func set7l() string {
//	return setNumReg(7, Reg.L.Val(), "L")
//}
//
//func set7hl() string {
//	return setNumHL(7)
//}
//
//func set7a() string {
//	return setNumReg(7, Reg.A.Val(), "A")
//}
//
////endregion CP Prefixed OpCode Functions

////region Helper functions
//
//// Increment a register by one.
//// Affects Flags Z and H. Sets Flag N to 0
//func (c *CPU) incReg(reg Reg8) {
//	c.flags.AffectFlagZH(reg.Get(), reg.Get()+1)
//	c.flags.SetFlagN(false)
//	reg.Inc()
//}
//
//// Decrement a register by one.
//// Affects Flags Z and H. Sets Flag N to 0
//func (c *CPU) decReg(reg Reg8) {
//	c.flags.AffectFlagZH(reg.Get(), reg.Get()-1)
//	c.flags.SetFlagN(true)
//	reg.Dec()
//}
//
//// Add value to register HL
//// Value comes in most significant byte (high) and least
//// significant byte
//func (c *CPU) addhlReg(high, low uint8) {
//	c.addhlReg16(bit2.To16(high, low))
//}
//
//// Add a 16-bit value to register HL
//// Affects Flag H and C. SetBit Flag N to Zero
//func (c *CPU) addhlReg16(value uint16) {
//	curHL := c.Reg.HL.Get()
//	nextVal := curHL + value
//	c.Reg.HL.Set(nextVal)
//	c.flags.SetFlagN(false)
//	c.flags.AffectFlagHC16(curHL, nextVal)
//}
//
//// Relate Jump according to condition. Additional ticks will be added
//// if condition met
//// Returns byte read after the jump instruction
//func (c *CPU) jrCond(condition bool, addTicks uint8) uint8 {
//	value := c.bus.Read(c.Reg.PC.Get())
//	c.Reg.PC.Inc()
//
//	if condition {
//		// Value is converted to signed 8bit first for relative
//		// positioning
//		c.Reg.PC.Set(c.Reg.PC.Get() + uint16(int8(value)))
//		c.ticks += addTicks
//	}
//
//	return value
//}
//
//// Jumps to position according to condition. Additional ticks will be added
//// if condition met
//// Returns 16-bit read after the jump instruction
//func (c *CPU) jpCond(condition bool, addTicks uint8) uint16 {
//	value := c.bus.Read16(c.Reg.PC.Get())
//	c.Reg.PC.Inc().Inc()
//
//	if condition {
//		c.Reg.PC.Set(value)
//		c.ticks += addTicks
//	}
//
//	return value
//}
//
//// Load value to indirect address pointed by register HL
//func (c *CPU) ldhlind(value uint8) {
//	c.bus.Write(c.Reg.HL.Get(), value)
//}
//
//// Adds value to the Accumulator
//// Affects Flags Z, H and C
//// SetBit Flag N to Zero
//func (c *CPU) adda(value uint8) {
//	curVal := c.Reg.A.Get()
//	*c.Reg.A.Val() += value
//	c.flags.AffectFlagZHC(curVal, c.Reg.A.Get())
//	c.flags.SetFlagN(false)
//}
//
//// Adds value plus carry to the Accumulator
//// Affects Flags Z, H and C
//// SetBit Flag N to Zero
//func (c *CPU) adca(value uint8) {
//	if c.flags.GetFlagC() {
//		value++
//	}
//	c.adda(value)
//}
//
//// Subtracts value from the Accumulator
//// Affects Flags Z, H and C
//// SetBit Flag N to One
//func (c *CPU) suba(value uint8) {
//	curVal := c.Reg.A.Get()
//	*c.Reg.A.Val() -= value
//	c.flags.AffectFlagZHC(curVal, c.Reg.A.Get())
//	c.flags.SetFlagN(false)
//}
//
//// Subtracts (value plus carry) from the Accumulator
//// Affects Flags Z, H and C
//// SetBit Flag N to One
//func (c *CPU) sbca(value uint8) {
//	if c.flags.GetFlagC() {
//		value++
//	}
//	c.suba(value)
//}
//
//// Bitwise AND between Accumulator and given value
//// Affects Flag Z
//// SetBit Flags N and C to Zero
//// SetBit Flag H to One
//func (c *CPU) anda(value uint8) {
//	*c.Reg.A.Val() &= value
//	c.flags.SetFlagZ(c.Reg.A.Get() == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagC(false)
//	c.flags.SetFlagH(true)
//}
//
//// Bitwise XOR between Accumulator and given value
//// Affects Flag Z
//// SetBit Flags N, H and C to Zero
//func (c *CPU) xora(value uint8) {
//	*c.Reg.A.Val() ^= value
//	c.flags.SetFlagZ(c.Reg.A.Get() == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(false)
//}
//
//// Bitwise OR between Accumulator and given value
//// Affects Flag Z
//// SetBit Flags N, H and C to Zero
//func (c *CPU) ora(value uint8) {
//	*c.Reg.A.Val() |= value
//	c.flags.SetFlagZ(c.Reg.A.Get() == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(false)
//}
//
//// Subtracts value from Accumulator without storing result
//// Affects Flag Z, H and C
//// SetBit Flag N to One
//func (c *CPU) cpa(value uint8) {
//	result := c.Reg.A.Get() - value
//	c.flags.AffectFlagZHC(c.Reg.A.Get(), result)
//}
//
//// Return from subroutine if condition met.
//func (c *CPU) retCond(condition bool, addTicks uint8) {
//	if !condition {
//		return
//	}
//
//	c.ticks += addTicks
//	c.Reg.SP.Inc()
//	low := c.bus.Read(c.Reg.SP.Get())
//	c.Reg.SP.Inc()
//	high := c.bus.Read(c.Reg.SP.Get())
//	c.Reg.PC.Set(bit2.To16(high, low))
//	//fmt.Printf("Return to %s ", c.Reg.PC)
//}
//
//// Load value from memory at location SP and SP + 1
//// low <- [SP]
//// high <- [SP+1]
//// SP is increment by two afterward
//func (c *CPU) pop16(high, low *uint8) {
//	c.Reg.SP.Inc()
//	*high, *low = bit2.From16(c.bus.Read16(c.Reg.SP.Get()))
//	c.Reg.SP.Inc()
//}
//
//// Store value to the Stack
//// [SP] <- high
//// [SP - 1] <- low
//// SP is decrement by two afterward
//func (c *CPU) push16(high, low uint8) {
//	c.bus.Write(c.Reg.SP.Get(), high)
//	c.bus.Write(c.Reg.SP.Get()-1, low)
//	c.Reg.SP.Dec().Dec()
//}
//
//// Calls a subroutine according to condition. Additional ticks will be added
//// if condition met
//// Returns 16-bit read after the call instruction
//func (c *CPU) callCond(condition bool, addTicks uint8) uint16 {
//	value := c.bus.Read16(c.Reg.PC.Get())
//	c.Reg.PC.Inc().Inc()
//
//	if condition {
//		callmem(value)
//		ticks += addTicks
//	}
//
//	return value
//}
//
//// Calls a subroutine
//// Current PC value is pushed to stack and PC is set to value
//func (c *CPU) callmem(value uint16) {
//	push16(bit2.From16(c.Reg.PC.Get()))
//	c.Reg.PC.Set(value)
//}
//
//// Load value to memory location
//func (c *CPU) ldmem(pos uint16, value uint8) {
//	c.bus.Write(pos, value)
//}
//
//// Rotate Left Circular a c.Register
//// Bit 7 shifts to bit 0
//// Bit 7 affect the carry Flag
//// C <- [7~0] <- [7]
//func (c *CPU) rlcc.Reg(r8 *uint8) {
//	*r8 = rlcVal(*r8)
//}
//
//// Rotate Left Circular an 8-bit value
//// Bit 7 shifts to bit 0
//// Bit 7 affect the carry Flag
//// C <- [7~0] <- [7]
//func (c *CPU) rlcVal(value uint8) uint8 {
//	var bit7 bool = value&0x80 == 0x80
//	// If bit 7 is 1
//	value <<= 1
//	if bit7 {
//		value |= 1
//	}
//	c.flags.SetFlagZ(value == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit7)
//
//	return value
//}
//
//// Rotate Right Circular a Register
//// Bit 0 shifts to Carry
//// [0] -> [7~0] -> C
//func (c *CPU) rrcReg(r8 *uint8) {
//	*r8 = rrcVal(*r8)
//}
//
//// Rotate Right Circular an 8-bit value
//// Bit 0 shifts to Carry
//// [0] -> [7~0] -> C
//func (c *CPU) rrcVal(value uint8) uint8 {
//	var bit0 bool = value&0x1 == 0x1
//	value >>= 1
//	if bit0 {
//		value |= 0x80
//	}
//	c.flags.SetFlagZ(false)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit0)
//
//	return value
//}
//
//// Rotate a c.Register left through Carry
//// Previous Carry shifts to bit 0
//// Bit 7 shift to Carry
//// C <- [7~0] <- C
//func (c *CPU) rlc.Reg(r8 *uint8) {
//	*r8 = rl(*r8)
//}
//
//// Rotate an 8-bit value left through Carry
//// Previous Carry shifts to bit 0
//// Bit 7 shift to Carry
//// C <- [7~0] <- C
//func (c *CPU) rl(value uint8) uint8 {
//	oldCarry := c.flags.GetFlagC()
//	value = sla(value)
//	// If carry flag is 1
//	if oldCarry {
//		value |= 1
//	}
//
//	return value
//}
//
//// Rotate a c.Register right through Carry
//// Previous Carry value shifts to bit 7
//// Bit 0 shifts to Carry
//// C -> [7~0] -> C
//func (c *CPU) rrc.Reg(r8 *uint8) {
//	*r8 = rr(*r8)
//}
//
//// Rotate an 8-bit value right through Carry
//// Previous Carry value shifts to bit 7
//// Bit 0 shifts to Carry
//// C -> [7~0] -> C
//func (c *CPU) rr(value uint8) uint8 {
//	oldCarry := c.flags.GetFlagC()
//	value = sra(value)
//	// If carry flag is 1
//	if oldCarry {
//		value |= 0x80
//	}
//
//	return value
//}
//
//// Shift Left Arithmetic a c.Register
//// Bit 7 shift to Carry
//// C <- [7~0]
//func (c *CPU) slac.Reg(r8 *uint8) {
//	*r8 = sla(*r8)
//}
//
//// Shift Left Arithmetic an 8-bit value
//// Bit 7 shift to Carry
//// C <- [7~0]
//func (c *CPU) sla(value uint8) uint8 {
//	var bit7 bool = value&0x80 == 0x80
//	value <<= 1
//	c.flags.SetFlagZ(value == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit7)
//
//	return value
//}
//
//// Shift Right Arithmetic a c.Register
//// Bit 0 shifts to Carry
//// Bit 7 value doesn't change
////  [7]-> [7~0] -> C
//func (c *CPU) srac.Reg(r8 *uint8) {
//	*r8 = sra(*r8)
//}
//
//// Shift Right Arithmetic an 8-bit value
//// Bit 0 shifts to Carry
//// Bit 7 value doesn't change
////  [7]-> [7~0] -> C
//func (c *CPU) sra(value uint8) uint8 {
//	var bit0 bool = value&0x1 == 0x1
//	value = value&0x80 | (value >> 1)
//	c.flags.SetFlagZ(value == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit0)
//
//	return value
//}
//
//// Swap upper four bits with lower four bits for a c.Register
//// [7654] <- [3~0] || [7~5] -> [3210]
//func (c *CPU) swapc.Reg(r8 *uint8) {
//	*r8 = swap(*r8)
//}
//
//// Swap upper four bits with lower four bits for an 8-bitu value
//// [7654] <- [3~0] || [7~5] -> [3210]
//func (c *CPU) swap(value uint8) uint8 {
//	return value<<4 | value>>4
//}
//
//// Shift Right Logic a c.Register
//// Bit 0 shifts to Carry
//// [7~0] -> C
//func (c *CPU) srlc.Reg(r8 *uint8) {
//	*r8 = srl(*r8)
//}
//
//// Shift Right Logic an 8-bit value
//// Bit 0 shifts to Carry
//// [7~0] -> C
//func (c *CPU) srl(value uint8) uint8 {
//	var bit0 bool = value&0x1 == 0x1
//	value >>= 1
//	c.flags.SetFlagZ(value == 0)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(false)
//	c.flags.SetFlagC(bit0)
//
//	return value
//}
//
//// Checks whether a bit of a c.Register is set
//// pos: Bit position
//// r8: c.Register
//// name: c.Register Name
//// Return string in format "BIT pos, name"
//func (c *CPU) bitNumc.Reg(pos uint8, r8 *uint8, name string) string {
//	bit(pos, *r8)
//
//	return fmt.Sprintf("BIT %d, %s", pos, name)
//}
//
//// Checks whether bit of a value at memory address is set
//// pos: Bit position
//// Return string in format "BIT pos, (HL)"
//func (c *CPU) bitNumHL(pos uint8) string {
//	addr := c.Reg.HL.Get()
//	value := c.bus.Read(addr)
//	bit(pos, value)
//
//	return fmt.Sprintf("BIT %d, (HL)", pos)
//}
//
//// Checks whether bit at given position of an 8-bit value is
//// set or not.
//// SetBit Flag Z to One if bit was not set
//// SetBit Flag N to Zero
//// SetBit Flag H to One
//func (c *CPU) bit(pos uint8, value uint8) {
//	var mask uint8 = 0x01 << pos
//	var isSet bool = value&mask == mask
//	c.flags.SetFlagZ(!isSet)
//	c.flags.SetFlagN(false)
//	c.flags.SetFlagH(true)
//}
//
//// SetBit a bit to zero at given position of a register
//// pos: Bit position
//// r8: c.Register
//// name: c.Register Name
//// Return string in format "RES pos, name"
//func (c *CPU) resNumc.Reg(pos uint8, r8 *uint8, name string) string {
//	*r8 = res(pos, *r8)
//
//	return fmt.Sprintf("RES %d, %s", pos, name)
//}
//
//// SetBit a bit to zero at given position of a value at memory address
//// pos: Bit position
//// Return string in format "RES pos, (HL)"
//func (c *CPU) resNumHL(pos uint8) string {
//	addr := c.Reg.HL.Get()
//	value := c.bus.Read(addr)
//	value = res(pos, value)
//	c.bus.Write(addr, value)
//
//	return fmt.Sprintf("RES %d, (HL)", pos)
//}
//
//// SetBit a bit to zero at given position of an 8-bit value
//func (c *CPU) res(pos uint8, value uint8) uint8 {
//	var mask uint8 = 0x01 << pos
//	mask = ^mask
//
//	return value & mask
//}
//
//// SetBit a bit to one at given position of a register
//// pos: Bit position
//// r8: c.Register
//// name: c.Register Name
//// Return string in format "RES pos, name"
//func (c *CPU) setNumc.Reg(pos uint8, r8 *uint8, name string) string {
//	*r8 = set(pos, *r8)
//
//	return fmt.Sprintf("RES %d, %s", pos, name)
//}
//
//// SetBit a bit to one at given position of a value at memory address
//// pos: Bit position
//// Return string in format "RES pos, (HL)"
//func (c *CPU) setNumHL(pos uint8) string {
//	addr := c.Reg.HL.Get()
//	value := c.bus.Read(addr)
//	value = set(pos, value)
//	c.bus.Write(addr, value)
//
//	return fmt.Sprintf("RES %d, (HL)", pos)
//}
//
//// SetBit a bit to zero at given position of an 8-bit value
//func (c *CPU) set(pos uint8, value uint8) uint8 {
//	var mask uint8 = 0x01 << pos
//
//	return value | mask
//}
//
////endregion Helper Functions
