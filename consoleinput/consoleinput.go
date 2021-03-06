package consoleinput

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")

type Handle syscall.Handle

func New() (Handle, error) {
	handle, err := syscall.Open("CONIN$", syscall.O_RDWR, 0)
	if err != nil {
		return Handle(0), err
	}
	return Handle(handle), nil
}

func (handle Handle) Close() error {
	return syscall.Close(syscall.Handle(handle))
}

var getConsoleMode = kernel32.NewProc("GetConsoleMode")

func (handle Handle) GetConsoleMode() uint32 {
	var mode uint32
	getConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	return mode
}

var setConsoleMode = kernel32.NewProc("SetConsoleMode")

func (handle Handle) SetConsoleMode(flag uint32) {
	setConsoleMode.Call(uintptr(handle), uintptr(flag))
}

var flushConsoleInputBuffer = kernel32.NewProc("FlushConsoleInputBuffer")

func (handle Handle) FlushConsoleInputBuffer() error {
	status, _, err := flushConsoleInputBuffer.Call(uintptr(handle))
	if status != 0 {
		return nil
	} else {
		return err
	}
}

var getNumberOfConsoleInputEvents = kernel32.NewProc("GetNumberOfConsoleInputEvents")

func (handle Handle) GetNumberOfEvent() (int, error) {
	var count uint32 = 0
	status, _, err := getNumberOfConsoleInputEvents.Call(uintptr(handle),
		uintptr(unsafe.Pointer(&count)))
	if status != 0 {
		return int(count), nil
	} else {
		return 0, err
	}
}

var waitForSingleObject = kernel32.NewProc("WaitForSingleObject")

func (handle Handle) WaitForSingleObject(msec uintptr) (uintptr, error) {
	status, _, err := waitForSingleObject.Call(uintptr(handle), msec)
	return status, err
}
