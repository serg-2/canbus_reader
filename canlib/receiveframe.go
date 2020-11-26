package canlib

import (
	"encoding/binary"

	"golang.org/x/sys/unix"
)

func (i Interface) RecvFrame() (CanFrame, error) {

	f := CanFrame{}

	// read SocketCAN frame from device
	frameBytes := make([]byte, 16)
	_, err := unix.Read(i.SocketFd, frameBytes)
	if err != nil {
		return f, err
	}

	// bytes 0-3: arbitration ID
	f.ArbitrationId = uint32(binary.LittleEndian.Uint32(frameBytes[0:4]))
	// remove bit 31: extended ID flag
	f.ArbitrationId = f.ArbitrationId & 0x7FFFFFFF
	// byte 4: data length code
	f.Dlc = frameBytes[4]
	// data
	f.Data = make([]byte, 8)
	copy(f.Data, frameBytes[8:])

	return f, nil
}
