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

func (state *followerState) Command(ctx context.Context, c *consenusImpl, in *raft.CommandRequest) (*raft.CommandResponse, error) {
	return nil, nil
}

func (state *followerState) ChangeCluster(ctx context.Context, c *consenusImpl, in *raft.Cluster) (*raft.ChangeClusterResponse, error) {
	return nil, nil
}

func (state *followerState) Vote(ctx context.Context, c *consenusImpl, in *raft.RequestVote) (*raft.ResponseVote, error) {
	return nil, nil
}

func (state *followerState) AppendEntries(ctx context.Context, c *consenusImpl, in *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error) {
	return nil, nil
}

func (state *followerState) InstallSnapshot(ctx context.Context, c *consenusImpl, in *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error) {
	return nil, nil
}

func (state *followerState) Tick(c *consenusImpl) error {
	return nil
}
