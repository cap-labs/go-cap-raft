package raft

import (
	"github.com/libs4go/scf4go"
	"github.com/libs4go/smf4go"
)

type raftConsensus struct {
	// Network Network `inject:"raft.network"`
	// Storage Storage `inject:"raft.storage"`
}

// New create raft consensus object compatible smf4go api
func New(config scf4go.Config) (smf4go.Service, error) {
	return &raftConsensus{}, nil
}

func (c *raftConsensus) Start() error {
	return nil
}
