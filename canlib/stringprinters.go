package canlib

import (
	"fmt"
	"time"
)

func PrintFrameUnixStyle(ifname string, frame CanFrame) {
	fmt.Printf(" %s\t%03X\t[%d]\t%s\n", ifname, frame.ArbitrationId, frame.Dlc, dataToString(frame.Data))
}

func PrintFrameCanBoatStyle(frame CanFrame) {
	pri, pgn, src, dst := GetISO11783BitsFromCanId(frame.ArbitrationId)

	ts := time.Now().UTC()
	// Print Beginning
	fmt.Printf("%s.%03d,%d,%d,%d,%d,%d", ts.Format(
		"2006-01-02-15:04:05"), (ts.UnixNano() / 1000000) % 1000, pri, pgn, src, dst, frame.Dlc)

	// Print Data Bytes
	for i := 0; i < int(frame.Dlc); i++ {
		fmt.Printf(",%02x", frame.Data[i])
	}

	// End String
	fmt.Printf("\n")
}

func dataToString(data []byte) string {
	str := ""
	for i := 0; i < len(data); i++ {
		str = fmt.Sprintf("%s%02X ", str, data[i])
	}
	return str
}
