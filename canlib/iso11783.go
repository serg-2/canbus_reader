package canlib

func GetISO11783BitsFromCanId(id uint32) (uint, uint, uint, uint){
	var dst,pgn uint

	var PF = uint8(id >> 16)
	var PS = uint8(id >> 8)
	var DP = uint8(id >> 24) & 1

	src := uint(uint8(id) >> 0)
	prio := uint(uint8(id >> 26) & 0x7)

	if PF < 240 {
		// PDU1 format, the PS contains the destination address
		dst = uint(PS)
		pgn = uint(DP) << 16 + uint(PF) << 8
	} else {
		dst = 0xff
		pgn = uint(DP) << 16 + uint(PF) << 8 + uint(PS)
	}

	return prio, pgn, src, dst
}