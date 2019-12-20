// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main chain33程序入口
package main

import (
	clog "github.com/33cn/chain33/common/log"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
	"time"
	"unsafe"
	"syscall"
	"os"
)

func main() {
	log1 := &types.Log{
		Loglevel:        "dbug",
		LogConsoleLevel: "info",
		LogFile:         "querylogs/query.log",
		MaxFileSize:     400,
		MaxBackups:      100,
		MaxAge:          28,
		LocalTime:       true,
		Compress:        false,
	}
	clog.SetFileLog(log1)
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		panic(err)
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for  {
		lpFreeBytesAvailable := int64(0)
		lpTotalNumberOfBytes := int64(0)
		lpTotalNumberOfFreeBytes := int64(0)
		syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dir))),
			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)
		log.Info("disk calc", "Available(mb):", float64(lpFreeBytesAvailable/1024/1024)/1024.0,
			"Total(mb):", float64(lpTotalNumberOfBytes/1024/1024)/1024.0,
			"Free(mb):", float64(lpTotalNumberOfFreeBytes/1024/1024)/1024.0)
		time.Sleep(time.Second*10)
	}
}
