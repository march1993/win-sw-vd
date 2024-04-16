package main

import (
	"golang.org/x/sys/windows"
)

type OS int
type LRId struct {
	Id2L uintptr
	Id2R uintptr
}

const (
	Win10 OS = iota + 1
	Win11
)

var IdMap = map[OS]*LRId{
	Win11: &LRId{
		Id2L: 0x00000024,
		Id2R: 0x00000025,
	},
	Win10: &LRId{
		Id2L: 0x00000025,
		Id2R: 0x00000026,
	},
}

var Ids *LRId

func init() {
	maj, min, patch := windows.RtlGetNtVersionNumbers()
	if maj != 10 || min != 0 {
		panic("unsupported Windows version")
	}
	if patch < 21996 {
		// Windows 10
		Ids = IdMap[Win10]
	} else {
		// Windows 11
		Ids = IdMap[Win11]
	}
}
