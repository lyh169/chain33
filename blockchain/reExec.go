// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"strings"

	"fmt"

	"github.com/33cn/chain33/common/version"
	"github.com/33cn/chain33/types"
	"math/big"
	"github.com/33cn/chain33/common/difficulty"
	"github.com/33cn/chain33/common"
)


// 重新执行block来擦除分叉,主要用于测试链
func (chain *BlockChain) ReExecEraseFork() {
	if !chain.cfg.EnableEraseFork {
		return
	}
	meta, err := chain.blockStore.GetEraseForkMeta()
	if err != nil {
		panic(err)
	}
	curheight := chain.GetBlockHeight()
	if curheight == -1 {
		meta = &types.UpgradeMeta{
			Version: version.GetEraseForkVersion(),
		}
		err = chain.blockStore.SetEraseForkMeta(meta)
		if err != nil {
			panic(err)
		}
	}
	if chain.needEraseFork(meta) {
		start := meta.Height
		//reExecBlock 的过程中，会每个高度都去更新meta
		chain.execEraseFork(start, curheight)
		meta := &types.UpgradeMeta{
			Starting: false,
			Version:  version.GetEraseForkVersion(),
			Height:   0,
		}
		err = chain.blockStore.SetEraseForkMeta(meta)
		if err != nil {
			panic(err)
		}
	}
}

// ReExecBlock 从对应高度本地重新执行区块
func (chain *BlockChain) execEraseFork(startHeight, curHeight int64) {
	var prevStateHash []byte
	if startHeight > 0 {
		blockdetail, err := chain.GetBlock(startHeight - 1)
		if err != nil {
			panic(fmt.Sprintf("get height=%d err, this not allow fail", startHeight-1))
		}
		prevStateHash = blockdetail.Block.StateHash
	}

	for i := startHeight; i <= curHeight; i++ {
		blockdetail, err := chain.GetBlock(i)
		if err != nil {
			panic(fmt.Sprintf("get height=%d err, this not allow fail", i))
		}
		block := blockdetail.Block
		err = execBlockUpgrade(chain.client, prevStateHash, block, false)
		if err != nil {
			panic(fmt.Sprintf("execBlockEx height=%d err=%s, this not allow fail", i, err.Error()))
		}

		err = chain.storeAndEraseForkBlock()
		if err != nil {
			panic(fmt.Sprintf("store and erase fork block height=%d err=%s, this not allow fail", i, err.Error()))
		}
		prevStateHash = block.StateHash
		//更新高度
		err = chain.upgradeMeta(i)
		if err != nil {
			panic(err)
		}
	}
}

func (chain *BlockChain) needEraseFork(meta *types.UpgradeMeta) bool {
	if meta.Starting { //正在
		return true
	}
	v1 := meta.Version
	v2 := version.GetEraseForkVersion()
	v1arr := strings.Split(v1, ".")
	v2arr := strings.Split(v2, ".")
	if len(v1arr) != 3 || len(v2arr) != 3 {
		panic("upgrade erase fork meta version error")
	}
	return v1arr[0] != v2arr[0]
}

func (chain *BlockChain) upgradeEraseForkMeta(height int64) error {
	meta := &types.UpgradeMeta{
		Starting: true,
		Version:  version.GetEraseForkVersion(),
		Height:   height + 1,
	}
	return chain.blockStore.SetEraseForkMeta(meta)
}

func (chain *BlockChain) storeAndEraseForkBlock(newBldetail, oldBldetail *types.BlockDetail, sync bool) error {
	// 写入磁盘 批量将block信息写入磁盘
	newbatch := chain.blockStore.NewBatch(sync)

	var err error

	sequence, err := chain.blockStore.GetSequenceByHash(oldBldetail.Block.Hash())
	if err != nil {
		chainlog.Error("get old sequence by hash:", "height", oldBldetail.Block.Height, "err", err)
		return  err
	}

	// 1、先删除已经存储的block相关的信息
	_, err = chain.blockStore.DelBlock(newbatch, oldBldetail, sequence)
	if err != nil {
		chainlog.Error("storeAndEraseForkBlock DelBlock:", "height", oldBldetail.Block.Height, "err", err)
		return err
	}

    // 2、保存新执行的blcok
	block := newBldetail.Block
	err = chain.blockStore.AddTxs(newbatch, newBldetail)
	if err != nil {
		panic(fmt.Sprintf("execBlockEx connectBlock readd Txs fail height=%d err=%s, this not allow fail", newBldetail.Block.Height, err.Error()))
	}
	//保存block信息到db中
	_, err = chain.blockStore.SaveBlock(newbatch, newBldetail, sequence)
	if err != nil {
		chainlog.Error("connectBlock SaveBlock:", "height", block.Height, "err", err)
		return err
	}

	//保存block的总难度到db中
	parentHash := newBldetail.Block.GetParentHash()
	difficulty := difficulty.CalcWork(block.Difficulty)
	var blocktd *big.Int
	if block.Height == 0 {
		blocktd = difficulty
	} else {
		parenttd, err := chain.blockStore.GetTdByBlockHash(parentHash)
		if err != nil {
			chainlog.Error("connectBlock GetTdByBlockHash", "height", block.Height, "parentHash", common.ToHex(parentHash))
			return err
		}
		blocktd = new(big.Int).Add(difficulty, parenttd)
	}

	err = chain.blockStore.SaveTdByBlockHash(newbatch, newBldetail.Block.Hash(), blocktd)
	if err != nil {
		chainlog.Error("connectBlock SaveTdByBlockHash:", "height", block.Height, "err", err)
		return err
	}
	err = newbatch.Write()
	if err != nil {
		return err
	}
	return nil
}
