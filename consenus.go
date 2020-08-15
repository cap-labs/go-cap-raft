package raft

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/cap-labs/go-cap/raft"
	"github.com/libs4go/errors"
	"github.com/libs4go/scf4go"
	"github.com/libs4go/slf4go"
	"github.com/libs4go/smf4go"
)

// ScopeOfAPIError .
const errVendor = "go-cap-raft"

// errors
var (
	ErrParams = errors.New("required parameter error", errors.WithVendor(errVendor), errors.WithCode(-1))
)

type consenusImpl struct {
	slf4go.Logger     `json:"-"`
	Network           raft.Network                           `inject:"raft.network" json:"-"`
	ClusterManager    raft.ClusterManager                    `inject:"raft.cluster.manager" json:"-"`
	LogStore          raft.LogStore                          `inject:"raft.logstore" json:"-"`
	Locker            sync.Locker                            `inject:"raft.locker"`
	ID                string                                 `json:"id"`
	ElectionTimeout   time.Duration                          `json:"election-timeout"`
	Leader            string                                 `json:"leader"`
	LeaderUpdateTimpe time.Time                              `json:"last-leader-update-time"`
	RemoteSMs         map[string]raft.RaftStateMachineClient `json:"remote-state-machines"`
	State             consenusState                          `json:"state"`
	Tick              time.Duration                          `json:"tick"`
}

// New create new raft consenus service for smf4go
func New(config scf4go.Config) (smf4go.Service, error) {

	id := config.Get("id").String("")

	if id == "" {
		return nil, errors.Wrap(ErrParams, "expect local consensus node id")
	}

	return &consenusImpl{
		Logger:    slf4go.Get("raft-consensus"),
		ID:        id,
		RemoteSMs: make(map[string]raft.RaftStateMachineClient),
		Tick:      config.Get("tick").Duration(10 * time.Millisecond),
	}, nil
}

func (c *consenusImpl) Start() error {

	cluster, err := c.ClusterManager.Get()

	if err != nil {
		return err
	}

	c.ElectionTimeout = time.Duration(cluster.ElectionTimeout) * time.Millisecond

	// calc randomized election timeout

	c.ElectionTimeout += time.Duration(rand.Int63n(int64(cluster.ElectionTimeout))) * time.Millisecond

	var local *raft.Peer = nil

	// create remote state machine rpc clients
	for _, peer := range cluster.Peers {
		if peer.Id != c.ID {
			client, err := c.Network.Connect(peer)

			if err != nil {
				return errors.Wrap(err, "connnect to peer %s error", peer.Id)
			}

			c.RemoteSMs[peer.Id] = client
		} else {
			local = peer
		}
	}

	if local == nil {
		return errors.Wrap(ErrParams, "cluster config miss local peer '%s'", c.ID)
	}

	c.State = &followerState{}

	c.LeaderUpdateTimpe = time.Now()

	c.Network.Serve(local, c)

	go c.tickLoop()

	return nil
}

func (c *consenusImpl) tickLoop() {
	ticker := time.NewTicker(c.Tick)

	defer ticker.Stop()

	for range ticker.C {
		c.Locker.Lock()
		c.State.Tick(c)
		c.Locker.Unlock()
	}
}

// consensus apis

func (c *consenusImpl) Command(ctx context.Context, in *raft.CommandRequest) (*raft.CommandResponse, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.State.Command(ctx, c, in)
}

func (c *consenusImpl) ChangeCluster(ctx context.Context, in *raft.Cluster) (*raft.ChangeClusterResponse, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.State.ChangeCluster(ctx, c, in)
}

func (c *consenusImpl) Vote(ctx context.Context, in *raft.RequestVote) (*raft.ResponseVote, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.State.Vote(ctx, c, in)
}

func (c *consenusImpl) AppendEntries(ctx context.Context, in *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.State.AppendEntries(ctx, c, in)
}

func (c *consenusImpl) InstallSnapshot(ctx context.Context, in *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()

	return c.State.InstallSnapshot(ctx, c, in)
}
