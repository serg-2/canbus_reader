package canlib

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func ParseFrame(frameStr string) (CanFrame, error) {
	frame := CanFrame{}
	fields := strings.Split(frameStr, "_")
	arbId, err := strconv.ParseUint(fields[0], 16, 32)
	if err != nil {
		return frame, err
	}
	frame.ArbitrationId = uint32(arbId)

	if len(fields[1])%2 != 0 {
		return frame, fmt.Errorf("invalid frame bytes")
	}
	frame.Dlc = byte(len(fields[1]) / 2)

	frame.Data = make([]byte, frame.Dlc)
	for i := byte(0); i < frame.Dlc; i++ {
		var val, err = strconv.ParseInt(fields[1][i*2:i*2+2], 16, 9)
		if err != nil {
			return frame, err
		}
		frame.Data[i] = byte(val)
	}

	return frame, nil
}

func (i Interface) SendFrame(f CanFrame) error {

	// assemble a SocketCAN frame
	bytesToSend := make([]byte, 16)
	// bytes 0-3: arbitration ID
	if f.ArbitrationId < 0x800 {
		// standard ID
		binary.LittleEndian.PutUint32(bytesToSend[0:4], f.ArbitrationId)
	} else {
		// extended ID
		// set bit 31 (frame format flag (0 = standard 11 bit, 1 = extended 29 bit)
		binary.LittleEndian.PutUint32(bytesToSend[0:4], f.ArbitrationId|1<<31)
	}

	// byte 4: data length code
	bytesToSend[4] = f.Dlc
	// copy message into frameBytes
	copy(bytesToSend[8:], f.Data)

	_, err := unix.Write(i.SocketFd, bytesToSend)
	return err
}
