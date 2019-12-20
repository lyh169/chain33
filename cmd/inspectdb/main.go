// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main chain33程序入口
package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/33cn/chain33/common/db"
	clog "github.com/33cn/chain33/common/log"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)


type StorageSize float64

var (
	dir = flag.String("d", "", "the inspect db dir")

	blockLastHeight       = []byte("blockLastHeight")
	bodyPrefix            = []byte("Body:")
	LastSequence          = []byte("LastSequence")
	headerPrefix          = []byte("Header:")
	heightToHeaderPrefix  = []byte("HH:")
	hashPrefix            = []byte("Hash:")
	tdPrefix              = []byte("TD:")
	heightToHashKeyPrefix = []byte("Height:")
	seqToHashKey          = []byte("Seq:")
	seqCBPrefix           = []byte("SCB:")
	seqCBLastNumPrefix    = []byte("SCBL:")
	paraSeqToHashKey      = []byte("ParaSeq:")
	HashToParaSeqPrefix   = []byte("HashToParaSeq:")
	LastParaSequence      = []byte("LastParaSequence")

	// 新增
	HashToSeqPrefix       = []byte("HashToSeq:")
	LocalPrefix           = []byte("LODB")
	TxShortHashPerfix     = []byte("STX:")
	TxHashPerfix          = []byte("TX:")
	TotalFeeKey           = []byte("TotalFeeKey:")
	TxAddrDirHash         = []byte("TxAddrDirHash:")
	TxAddrHash            = []byte("TxAddrHash:")
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

	path := *dir
	if path == "" {
		path = "datadir"
	}
	ldb, err := db.NewGoLevelDB("blockchain", path, 128)
	if err != nil {
		panic(err)
	}
	it := ldb.Iterator(nil, nil, false)
	defer it.Close()

	var (
		// Key-value store statistics
		total             StorageSize
		headerSize        StorageSize
		bodySize          StorageSize
		receiptSize       StorageSize
		HHsize            StorageSize
		hashPrefixSize    StorageSize
		heightToHashSize  StorageSize
		seqToHashSize     StorageSize
		seqCBSize         StorageSize
		paraSeqToHashSize StorageSize
		HashToParaSeqSize StorageSize
		tdSize            StorageSize
		otherSize         StorageSize

		hashToSeqSize     StorageSize
		localSize         StorageSize
		txShortHashSize   StorageSize
		txHashSize        StorageSize
		totalFeeKeySize   StorageSize
		txAddrDirHashSize StorageSize
		txAddrHashSize    StorageSize
	)
	for it.Rewind(); it.Valid(); it.Next() {
		size := StorageSize(len(it.Key()) + len(it.Value()))
		total += size
		switch {
		case bytes.HasPrefix(it.Key(), headerPrefix):
			headerSize += size
		case bytes.HasPrefix(it.Key(), bodyPrefix):
			blockbody := &types.BlockBody{}
			err = types.Decode(it.Value(), blockbody)
			if err == nil {
				for i := 0; i < len(blockbody.Receipts); i++ {
					for j := 0; j < len(blockbody.Receipts[i].Logs); j++ {
						if blockbody.Receipts[i].Logs[j] != nil {
							blockbody.Receipts[i].Logs[j].Log = nil
						}
					}
				}
				value := types.Encode(blockbody)
				if size - StorageSize(len(value)) > 0{
					receiptSize += size - StorageSize(len(value))
				} else {
					log.Error("receiptSize is  < 0")
				}
			}
			bodySize += size
		case bytes.HasPrefix(it.Key(), heightToHeaderPrefix):
			HHsize += size
		case bytes.HasPrefix(it.Key(), hashPrefix):
			hashPrefixSize += size
		case bytes.HasPrefix(it.Key(), heightToHashKeyPrefix):
			heightToHashSize += size
		case bytes.HasPrefix(it.Key(), seqToHashKey):
			seqToHashSize += size
		case bytes.HasPrefix(it.Key(), seqCBPrefix):
			seqCBSize += size
		case bytes.HasPrefix(it.Key(), paraSeqToHashKey):
			paraSeqToHashSize += size
		case bytes.HasPrefix(it.Key(), HashToParaSeqPrefix):
			HashToParaSeqSize += size
		case bytes.HasPrefix(it.Key(), tdPrefix):
			tdSize += size
		case bytes.HasPrefix(it.Key(), HashToSeqPrefix):
			hashToSeqSize += size
		case bytes.HasPrefix(it.Key(), LocalPrefix):
			localSize += size
		case bytes.HasPrefix(it.Key(), TxShortHashPerfix):
			txShortHashSize += size
		case bytes.HasPrefix(it.Key(), TxHashPerfix):
			txHashSize += size
		case bytes.HasPrefix(it.Key(), TotalFeeKey):
			totalFeeKeySize += size
		case bytes.HasPrefix(it.Key(), TxAddrDirHash):
			txAddrDirHashSize += size
		case bytes.HasPrefix(it.Key(), TxAddrHash):
			txAddrHashSize += size
		default:
			otherSize += size
			key := it.Key()
			log.Debug("other ", "key prefix", string(key))
		}
	}

	// Display the database statistic.
	stats := [][]string{
		{"Key-Value store", "Headers", headerSize.String()},
		{"Key-Value store", "Bodies", bodySize.String()},
		{"Key-Value store", "Receipts", receiptSize.String()},
		{"Key-Value store", "HHsize", HHsize.String()},
		{"Key-Value store", "hashPrefixSize", hashPrefixSize.String()},
		{"Key-Value store", "heightToHashSize", heightToHashSize.String()},
		{"Key-Value store", "seqToHashSize", seqToHashSize.String()},
		{"Key-Value store", "seqCBSize", seqCBSize.String()},
		{"Key-Value store", "paraSeqToHashSize", paraSeqToHashSize.String()},
		{"Key-Value store", "HashToParaSeqSize", HashToParaSeqSize.String()},
		{"Key-Value store", "Difficulties", tdSize.String()},

		{"Key-Value store", "hashToSeqSize", hashToSeqSize.String()},
		{"Key-Value store", "localSize", localSize.String()},
		{"Key-Value store", "txShortHashSize", txShortHashSize.String()},
		{"Key-Value store", "txHashSize", txHashSize.String()},
		{"Key-Value store", "totalFeeKeySize", totalFeeKeySize.String()},
		{"Key-Value store", "txAddrDirHashSize", txAddrDirHashSize.String()},
		{"Key-Value store", "txAddrHashSize", txAddrHashSize.String()},

		{"Key-Value store", "otherSize", otherSize.String()},
		{"Key-Value store", "total", total.String()},
	}
	for _, stat := range stats {
		log.Info(stat[0], stat[1], stat[2])
	}
}

// String implements the stringer interface.
func (s StorageSize) String() string {
	if s > 1099511627776 {
		return fmt.Sprintf("%.2f TiB", s/1099511627776)
	} else if s > 1073741824 {
		return fmt.Sprintf("%.2f GiB", s/1073741824)
	} else if s > 1048576 {
		return fmt.Sprintf("%.2f MiB", s/1048576)
	} else if s > 1024 {
		return fmt.Sprintf("%.2f KiB", s/1024)
	} else {
		return fmt.Sprintf("%.2f B", s)
	}
}
