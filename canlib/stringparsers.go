package canlib

import "strconv"

func ParseFrameFromStringUnix(data string) CanFrame {
	// can0	1234124	[4]	11 11 11 23 00 00 00 00
	var frame CanFrame

	// Extract ArbitrationIf from string
	d, _ := strconv.ParseUint("0x"+data[8:16], 0, 32)
	frame.ArbitrationId = uint32(d)

	// Extract Size from string
	d, _ = strconv.ParseUint(data[25:26], 0, 32)
	frame.Dlc = byte(d)

	// Extract Data from string
	for i:= 0; i < int(frame.Dlc); i++ {
		d, _ = strconv.ParseUint("0x" + data[32+i*3:34+i*3], 0, 8)
		frame.Data = append(frame.Data, byte(d))
	}

	return frame
}

func ParseFrameFromStringCanBoat(data string) CanFrame {
	var frame CanFrame

	//<0x09f80100> [8] 04 7e 54 19 24 33 55 d5

	// Extract ArbitrationIf from string
	d, _ := strconv.ParseUint(data[1:11], 0, 32)
	frame.ArbitrationId = uint32(d)

	// Extract Size from string
	d, _ = strconv.ParseUint(data[14:15], 0, 32)
	frame.Dlc = byte(d)

	// Extract Data from string
	for i:= 0; i < int(frame.Dlc); i++ {
		d, _ = strconv.ParseUint("0x" + data[17+i*3:19+i*3], 0, 8)
		frame.Data = append(frame.Data, byte(d))
	}

	return frame
}
