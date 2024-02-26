package main

import (
	"fmt"
	"os"
	"slices"
	"syscall"
	"time"

	"github.com/lxn/win"
	"github.com/mitchellh/go-ps"
)

func HotKey(vk, mod uintptr) uintptr {
	return vk<<8 | mod
}

const (
	Ks2L = win.VK_LEFT<<8 | MOD_CONTROL | MOD_WIN
	Ks2R = win.VK_RIGHT<<8 | MOD_CONTROL | MOD_WIN
)

const (
	Id2L = 0x00000024
	Id2R = 0x00000025
)

func main() {
	action := "demo"

	if len(os.Args) == 2 {
		action = os.Args[1]
	}

	explorerPids := []uint32{}
	if list, err := ps.Processes(); err != nil {
		panic(err.Error())
	} else {
		for _, p := range list {
			if p.Executable() == "explorer.exe" {
				explorerPids = append(explorerPids, uint32(p.Pid()))
			}
		}
	}

	fmt.Println("len(explorerPids):", explorerPids)

	targets := []HWND{}

	hdw := win.GetDesktopWindow()
	for c := HWND(0); ; {
		c = FindWindowEx(hdw, c, syscall.StringToUTF16Ptr("WorkerW"), nil)

		if c == 0 {
			break
		}

		pid := uint32(0)
		win.GetWindowThreadProcessId(c, &pid)

		if slices.Contains(explorerPids, pid) {
			targets = append(targets, c)
		}
	}

	fmt.Println("len(targets):", len(targets))

	switch action {
	case "left":
		switchLeft(targets)
	case "right":
		switchRight(targets)
	default:
		fallthrough
	case "demo":
		switchRight(targets)
		time.Sleep(time.Second)
		switchLeft(targets)
	}
}

// to left
func switchLeft(hwnds []HWND) {
	for _, hwnd := range hwnds {
		win.SendMessage(hwnd, win.WM_HOTKEY, Id2L, Ks2L)
	}
}

// to right
func switchRight(hwnds []HWND) {
	for _, hwnd := range hwnds {
		win.SendMessage(hwnd, win.WM_HOTKEY, Id2R, Ks2R)
	}
}
