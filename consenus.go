package raft

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/cap-labs/go-cap"
	"github.com/cap-labs/go-cap/raft"
	"github.com/libs4go/errors"
	"github.com/libs4go/slf4go"
)

// ScopeOfAPIError .
const errVendor = "go-cap-raft"

// errors
var (
	ErrParams = errors.New("required parameter error", errors.WithVendor(errVendor), errors.WithCode(-1))
)

type consensusState interface {
	json.Marshaler
	Tick(context.Context, *ConsensusNode) error
	Exec(ctx context.Context, cn *ConsensusNode, in *raft.RequestExec) (*raft.ResponseExec, error)
	NewCluster(ctx context.Context, cn *ConsensusNode, in *raft.Cluster) (*raft.ResponseCluster, error)
	Vote(context.Context, *ConsensusNode, *raft.RequestVote) (*raft.ResponseVote, error)
	AppendEntries(context.Context, *ConsensusNode, *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error)
	InstallSnapshot(context.Context, *ConsensusNode, *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error)
}

// ConsensusNode .
type ConsensusNode struct {
	slf4go.Logger          `json:"-"`                             // mixin slf4go logger
	sync.Locker                                                   // consensus locker
	ID                     string                                 `json:"id"`                 // consensus node id
	ElectionTimeout        time.Duration                          `json:"election_timeout"`   // election timeout
	HeartbeatTimeout       time.Duration                          `json:"heart_beat_timeout"` // leader heartbeat timeout  HeartbeatTimeout * 10 < ElectionTimeout
	RemotePeers            map[string]raft.RaftStateMachineClient `json:"remote"`             // remote raft state machine rpc client
	State                  consensusState                         `json:"state"`              // conseusus state object
	Term                   uint64                                 `json:"term"`               // current raft term
	Store                  raft.LogStore                          `json:"-"`                  // consensus logstore service
	Network                raft.Network                           `json:"-"`                  // consensus network service
	ClusterManager         raft.ClusterManager                    `json:"-"`                  // consensus cluster manager
	LastHeartBeatTimestamp time.Time                              `json:"updated"`            // consensus follower node received last heartbeat timestamp
	Leader                 string                                 `json:"leader"`             // leader id
}

// ConsensusOption .
type ConsensusOption func(*ConsensusNode)

// ElectionTimeout .
func ElectionTimeout(duration time.Duration) ConsensusOption {
	return func(node *ConsensusNode) {
		node.ElectionTimeout = duration
	}
}

// HeartBeatTimeout .
func HeartBeatTimeout(duration time.Duration) ConsensusOption {
	return func(node *ConsensusNode) {
		node.HeartbeatTimeout = duration
	}
}

// ClusterManager .
func ClusterManager(clusterManager raft.ClusterManager) ConsensusOption {
	return func(node *ConsensusNode) {
		node.ClusterManager = clusterManager
	}
}

// Network .
func Network(network raft.Network) ConsensusOption {
	return func(node *ConsensusNode) {
		node.Network = network
	}
}

// LogStore .
func LogStore(store raft.LogStore) ConsensusOption {
	return func(node *ConsensusNode) {
		node.Store = store
	}
}

// Locker set consensus node invoke locker
func Locker(locker sync.Locker) ConsensusOption {
	return func(node *ConsensusNode) {
		node.Locker = locker
	}
}

// New .
func New(
	local *cap.Peer,
	options ...ConsensusOption) (*ConsensusNode, error) {

	node := &ConsensusNode{
		Logger:          slf4go.Get("raft-node"),
		Locker:          &nullLocker{},
		ID:              local.Id,
		ElectionTimeout: time.Millisecond * 150,
		RemotePeers:     make(map[string]raft.RaftStateMachineClient),
		State:           &followerState{},
	}

	for _, option := range options {
		option(node)
	}

	if node.HeartbeatTimeout == 0 {
		node.HeartbeatTimeout = node.ElectionTimeout / 10
	}

	if node.Network == nil {
		return nil, errors.Wrap(ErrParams, "expect Network option")
	}

	if node.Store == nil {
		return nil, errors.Wrap(ErrParams, "expect LogStore option")
	}

	if node.ClusterManager == nil {
		return nil, errors.Wrap(ErrParams, "expect ClusterManager option")
	}

	cluster, err := node.ClusterManager.Get()

	if err != nil {
		return nil, err
	}

	for _, peer := range cluster.Peers {
		client, err := node.Network.Connect(peer)

		if err != nil {
			return nil, err
		}

		node.RemotePeers[peer.Id] = client
	}

	lastEntry, err := node.Store.LastEntry()

	if err != nil {
		return nil, err
	}

	if lastEntry != nil {
		node.Term = lastEntry.Term
	}

	node.LastHeartBeatTimestamp = time.Now()

	node.I("start consenus with config: \n{@config}", node)

	node.Network.Serve(local, node)

	return node, nil
}

// Vote .
func (cn *ConsensusNode) Vote(ctx context.Context, request *raft.RequestVote) (*raft.ResponseVote, error) {
	cn.Lock()
	defer cn.Unlock()

	return cn.State.Vote(ctx, cn, request)
}

// AppendEntries .
func (cn *ConsensusNode) AppendEntries(ctx context.Context, request *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error) {
	cn.Lock()
	defer cn.Unlock()
	return cn.State.AppendEntries(ctx, cn, request)
}

// InstallSnapshot .
func (cn *ConsensusNode) InstallSnapshot(ctx context.Context, request *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error) {
	cn.Lock()
	defer cn.Unlock()
	return cn.State.InstallSnapshot(ctx, cn, request)
}

// Tick invoke consensus tick route
func (cn *ConsensusNode) Tick() {
	cn.Lock()
	defer cn.Unlock()

	err := cn.State.Tick(context.Background(), cn)

	if err != nil {
		cn.E("")
	}
}

// Exec .
func (cn *ConsensusNode) Exec(ctx context.Context, request *raft.RequestExec) (*raft.ResponseExec, error) {
	cn.Lock()
	defer cn.Unlock()

	return cn.State.Exec(ctx, cn, request)
}

// NewCluster .
func (cn *ConsensusNode) NewCluster(ctx context.Context, request *raft.Cluster) (*raft.ResponseCluster, error) {
	cn.Lock()
	defer cn.Unlock()

	return cn.State.NewCluster(ctx, cn, request)
}
