package canlib

import (
	"errors"
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Interface struct {
	IfName   string
	SocketFd int
}

type ifreqIndex struct {
	Name  [16]byte
	Index int
}

func ioctlIfreq(fd int, ifreq *ifreqIndex) (err error) {
	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		unix.SIOCGIFINDEX,
		uintptr(unsafe.Pointer(ifreq)),
	)
	if errno != 0 {
		err = fmt.Errorf("ioctl: %v", errno)
	}
	return
}

func getIfIndex(fd int, ifName string) (int, error) {
	ifNameRaw, err := unix.ByteSliceFromString(ifName)
	if err != nil {
		return 0, err
	}
	if len(ifNameRaw) > 16 {
		return 0, errors.New("maximum ifname length is 16 characters")
	}

	ifReq := ifreqIndex{}
	copy(ifReq.Name[:], ifNameRaw)
	err = ioctlIfreq(fd, &ifReq)
	return ifReq.Index, err
}

func NewRawInterface(ifName string) (Interface, error) {
	canIf := Interface{}
	// 1 - raw interface
	fd, err := unix.Socket(unix.AF_CAN, unix.SOCK_RAW, 1)
	if err != nil {
		return canIf, err
	}

	ifIndex, err := getIfIndex(fd, ifName)
	if err != nil {
		return canIf, err
	}

	addr := &unix.SockaddrCAN{Ifindex: ifIndex}
	if err = unix.Bind(fd, addr); err != nil {
		return canIf, err
	}

	canIf.IfName = ifName
	canIf.SocketFd = fd

	return canIf, nil
}

// Close Interface
func (i Interface) Close() error {
	return unix.Close(i.SocketFd)
}
