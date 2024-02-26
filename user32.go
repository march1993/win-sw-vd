package main

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

const (
	MOD_ALT     = 0x0001
	MOD_CONTROL = 0x0002
	MOD_SHIFT   = 0x0004
	MOD_WIN     = 0x0008
)

var (
	// Library
	libuser32 *windows.LazyDLL

	// Functions
	findWindowEx     *windows.LazyProc
	getShellWindow   *windows.LazyProc
	getDesktopWindow *windows.LazyProc
)

type HWND = win.HWND

func init() {
	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")

	// Functions
	findWindowEx = libuser32.NewProc("FindWindowExW")
	getShellWindow = libuser32.NewProc("GetShellWindow")
	getDesktopWindow = libuser32.NewProc("GetDesktopWindow")
}

func GetShellWindow() HWND {
	ret, _, _ := syscall.SyscallN(getShellWindow.Addr())

	return HWND(ret)
}

func GetDesktopWindow() HWND {
	ret, _, _ := syscall.SyscallN(getDesktopWindow.Addr())

	return HWND(ret)
}

func FindWindowEx(parent, child HWND, lpClassName, lpWindowName *uint16) HWND {
	ret, _, _ := syscall.SyscallN(findWindowEx.Addr(),
		uintptr(parent),
		uintptr(child),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)))

	return HWND(ret)
}
