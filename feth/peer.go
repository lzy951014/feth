/*
 * @Author: Liuzongyun 845666459@qq.com
 * @Date: 2024-11-12 15:08:17
 * @LastEditors: Liuzongyun 845666459@qq.com
 * @LastEditTime: 2024-11-12 15:32:04
 * @FilePath: /feth/feth/peer.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Copyright 2015 The go-ethereum Authors
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

package eth

import (
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/lzy951014/feth/feth/protocols/bsc"
)

// ethPeerInfo represents a short summary of the `eth` sub-protocol metadata known
// about a connected peer.
type ethPeerInfo struct {
	Version uint `json:"version"` // Ethereum protocol version negotiated
}

// ethPeer is a wrapper around eth.Peer to maintain a few extra metadata.
type ethPeer struct {
	*eth.Peer
	snapExt *snapPeer // Satellite `snap` connection
	bscExt  *bscPeer  // Satellite `bsc` connection

}

// info gathers and returns some `eth` protocol metadata known about a peer.
func (p *ethPeer) info() *ethPeerInfo {
	return &ethPeerInfo{
		Version: p.Version(),
	}
}

// snapPeerInfo represents a short summary of the `snap` sub-protocol metadata known
// about a connected peer.
type snapPeerInfo struct {
	Version uint `json:"version"` // Snapshot protocol version negotiated
}

// snapPeer is a wrapper around snap.Peer to maintain a few extra metadata.
type snapPeer struct {
	*snap.Peer
}

// bscPeer is a wrapper around bsc.Peer to maintain a few extra metadata.
type bscPeer struct {
	*bsc.Peer
}

// info gathers and returns some `snap` protocol metadata known about a peer.
func (p *snapPeer) info() *snapPeerInfo {
	return &snapPeerInfo{
		Version: p.Version(),
	}
}

// info gathers and returns some `bsc` protocol metadata known about a peer.
func (p *bscPeer) info() *bscPeerInfo {
	return &bscPeerInfo{
		Version: p.Version(),
	}
}

// bscPeerInfo represents a short summary of the `bsc` sub-protocol metadata known
// about a connected peer.
type bscPeerInfo struct {
	Version uint `json:"version"` // bsc protocol version negotiated
}
