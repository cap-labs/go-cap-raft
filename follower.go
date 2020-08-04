package raft

import (
	"context"

	"github.com/cap-labs/go-cap/raft"
)

type followerState struct {
}

func (state *followerState) MarshalJSON() ([]byte, error) {
	return []byte(`"follower"`), nil
}

func (state *followerState) Tick(context.Context, *ConsensusNode) error {
	return nil
}

func (state *followerState) Vote(context.Context, *ConsensusNode, *raft.RequestVote) (*raft.ResponseVote, error) {
	return nil, nil
}

func (state *followerState) AppendEntries(context.Context, *ConsensusNode, *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error) {
	return nil, nil
}

func (state *followerState) InstallSnapshot(context.Context, *ConsensusNode, *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error) {
	return nil, nil
}

func (state *followerState) Exec(ctx context.Context, cn *ConsensusNode, in *raft.RequestExec) (*raft.ResponseExec, error) {
	return &raft.ResponseExec{}, nil
}
func (state *followerState) NewCluster(ctx context.Context, cn *ConsensusNode, in *raft.Cluster) (*raft.ResponseCluster, error) {
	return &raft.ResponseCluster{}, nil
}
