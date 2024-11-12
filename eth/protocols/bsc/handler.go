package bsc

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

// Handler is a callback to invoke from an outside runner after the boilerplate
// exchanges have passed.
type Handler func(peer *Peer) error

type Backend interface {
	// Chain retrieves the blockchain object to serve data.
	Chain() *core.BlockChain

	// 当对等点加入 `bsc` 协议时，将调用 RunPeer。处理程序
	// 应该做任何对等维护工作、握手和验证。如果全部
	// 传递后，控制权应该交还给 `handler` 来处理
	// 未来的入站消息。
	// RunPeer is invoked when a peer joins on the `bsc` protocol. The handler
	// should do any peer maintenance work, handshakes and validations. If all
	// is passed, control should be given back to the `handler` to process the
	// inbound messages going forward.
	RunPeer(peer *Peer, handler Handler) error

	// PeerInfo retrieves all known `bsc` information about a peer.
	PeerInfo(id enode.ID) interface{}

	//Handle是一个回调函数，当接收到数据包时调用
	//远程对等点。只有协议处理程序未消耗的数据包才会
	//转发到后台。
	// Handle is a callback to be invoked when a data packet is received from
	// the remote peer. Only packets not consumed by the protocol handler will
	// be forwarded to the backend.
	Handle(peer *Peer, packet Packet) error
}

// MakeProtocols constructs the P2P protocol definitions for `bsc`.
func MakeProtocols(backend Backend, dnsdisc enode.Iterator) []p2p.Protocol {
	// Filter the discovery iterator for nodes advertising vote support.
	dnsdisc = enode.Filter(dnsdisc, func(n *enode.Node) bool {
		var vote enrEntry
		return n.Load(&vote) == nil
	})

	protocols := make([]p2p.Protocol, len(ProtocolVersions))
	for i, version := range ProtocolVersions {
		version := version // Closure

		protocols[i] = p2p.Protocol{
			Name:    ProtocolName,
			Version: version,
			Length:  protocolLengths[version],
			Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				peer := NewPeer(version, p, rw)
				defer peer.Close()

				return backend.RunPeer(peer, func(peer *Peer) error {
					return Handle(backend, peer)
				})
			},
			NodeInfo: func() interface{} {
				return nodeInfo(backend.Chain())
			},
			PeerInfo: func(id enode.ID) interface{} {
				return backend.PeerInfo(id)
			},
			Attributes:     []enr.Entry{&enrEntry{}},
			DialCandidates: dnsdisc,
		}
	}
	return protocols
}

// Handle is the callback invoked to manage the life cycle of a `bsc` peer.
// When this function terminates, the peer is disconnected.
func Handle(backend Backend, peer *Peer) error {
	for {
		if err := handleMessage(backend, peer); err != nil {
			peer.Log().Debug("Message handling failed in `bsc`", "err", err)
			return err
		}
	}
}

type msgHandler func(backend Backend, msg Decoder, peer *Peer) error
type Decoder interface {
	Decode(val interface{}) error
}

var bsc1 = map[uint64]msgHandler{
	VotesMsg: handleVotes,
}

// 每当从某个接收到入站消息时，就会调用handleMessage
// `bsc` 协议上的远程对等点。远程连接被断开
// 返回任何错误。
// handleMessage is invoked whenever an inbound message is received from a
// remote peer on the `bsc` protocol. The remote connection is torn down upon
// returning any error.
func handleMessage(backend Backend, peer *Peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := peer.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > maxMessageSize {
		return fmt.Errorf("%w: %v > %v", errMsgTooLarge, msg.Size, maxMessageSize)
	}
	defer msg.Discard()

	var handlers = bsc1
	//跟踪服务请求和运行处理程序所需的时间
	// Track the amount of time it takes to serve the request and run the handler
	if metrics.Enabled {
		h := fmt.Sprintf("%s/%s/%d/%#02x", p2p.HandleHistName, ProtocolName, peer.Version(), msg.Code)
		defer func(start time.Time) {
			sampler := func() metrics.Sample {
				return metrics.ResettingSample(
					metrics.NewExpDecaySample(1028, 0.015),
				)
			}
			metrics.GetOrRegisterHistogramLazy(h, nil, sampler).Update(time.Since(start).Microseconds())
		}(time.Now())
	}
	if handler := handlers[msg.Code]; handler != nil {
		return handler(backend, msg, peer)
	}
	return fmt.Errorf("%w: %v", errInvalidMsgCode, msg.Code)
}

func handleVotes(backend Backend, msg Decoder, peer *Peer) error {
	ann := new(VotesPacket)
	if err := msg.Decode(ann); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	// Schedule all the unknown hashes for retrieval
	peer.markVotes(ann.Votes)
	return backend.Handle(peer, ann)
}

// NodeInfo represents a short summary of the `bsc` sub-protocol metadata
// known about the host peer.
type NodeInfo struct{}

// nodeInfo retrieves some `bsc` protocol metadata about the running host node.
func nodeInfo(_ *core.BlockChain) *NodeInfo {
	return &NodeInfo{}
}
