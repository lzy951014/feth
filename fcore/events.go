/*
 * @Author: Liuzongyun 845666459@qq.com
 * @Date: 2024-11-12 15:37:44
 * @LastEditors: Liuzongyun 845666459@qq.com
 * @LastEditTime: 2024-11-12 17:05:39
 * @FilePath: /feth/fcore/events.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/ethereum/go-ethereum/common"
	// "github.com/lzy951014/feth/fcore/types"
	typess "github.com/lzy951014/feth/fcore/types"
)

// NewTxsEvent is posted when a batch of transactions enters the transaction pool.
type NewTxsEvent struct{ Txs []*typess.Transaction }

// ReannoTxsEvent is posted when a batch of local pending transactions exceed a specified duration.
type ReannoTxsEvent struct{ Txs []*typess.Transaction }

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *typess.Block }

// RemovedLogsEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs []*typess.Log }

// NewVoteEvent is posted when a batch of votes enters the vote pool.
type NewVoteEvent struct{ Vote *typess.VoteEnvelope }

// FinalizedHeaderEvent is posted when a finalized header is reached.
type FinalizedHeaderEvent struct{ Header *typess.Header }

type ChainEvent struct {
	Block *typess.Block
	Hash  common.Hash
	Logs  []*typess.Log
}

type ChainSideEvent struct {
	Block *typess.Block
}

type ChainHeadEvent struct{ Block *typess.Block }
