// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main chain33程序入口
package main

import (
	"github.com/33cn/chain33/common/db"
	"fmt"
	"flag"
)
var (
	dir = flag.String("d", "datadir/mavltree", "the inspect db dir")
	name = flag.String("n", "store", "the name of db")
)


func main() {
	flag.Parse()
	fmt.Println("dir:", *dir, "name:", *name)
	ldb, err := db.NewGoLevelDB(*name, *dir, 128)
	if err != nil {
		panic(fmt.Sprintln("open db fail", err))
	}
	//go func() {
		fmt.Println("start compact")
		err = ldb.CompactRange(nil, nil)
		fmt.Println("end compact", err)
	//}()

	//i := 0
	//fmt.Println("start set key value")
	//for {
	//	err := ldb.Set([]byte(fmt.Sprintln("key-", i)), []byte(fmt.Sprintln("value-", i)))
	//	if err != nil {
	//		//fmt.Println("set ", i, "error", err)
	//	} else {
	//		//fmt.Println("set success ", i)
	//	}
	//	i++
	//	time.Sleep(time.Microsecond * 1000)
	//}
	//fmt.Println("end set key value")
}
