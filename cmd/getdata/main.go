// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main chain33程序入口
package main

import (
	"flag"
	"os/exec"
	"strconv"
	clog "github.com/33cn/chain33/common/log"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
	"time"
)

var (
	from = flag.Int64("f", 4302900, "from traverse addr")
	to   = flag.Int64("t", 4304000, "to traverse addr")
	addr = flag.String("a", "1AH9HRd4WBJ824h9PP1jYpvRZ4BSA4oN6Y", "query addr")
	exect = flag.String("e", "ticket", "exec")
	cli   = flag.String("cli", "./bityuan-cli", "exec")
)

func main() {
	flag.Parse()
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

	for i := *from; i <= *to; i++ {
		str := strconv.Itoa(int(i))
		para := []string{
			"account", "balance", "--height", str, "-a", *addr, "-e", *exect,
		}
		rawOut, err := exec.Command(*cli, para[0:]...).CombinedOutput()
		if err != nil {
			panic(err)
		}
		strOut := string(rawOut)
		log.Info("query", "heigth", str, "data", strOut)
		time.Sleep(time.Microsecond*100)
	}
}
