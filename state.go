package raft

import (
	"context"
	"encoding/json"

	"github.com/cap-labs/go-cap/raft"
)

type consenusState interface {
	json.Marshaler

	Tick(c *consenusImpl) error

	Command(ctx context.Context, c *consenusImpl, in *raft.CommandRequest) (*raft.CommandResponse, error)

	ChangeCluster(ctx context.Context, c *consenusImpl, in *raft.Cluster) (*raft.ChangeClusterResponse, error)

	Vote(ctx context.Context, c *consenusImpl, in *raft.RequestVote) (*raft.ResponseVote, error)

	AppendEntries(ctx context.Context, c *consenusImpl, in *raft.RequestAppendEntries) (*raft.ResponseAppendEntries, error)

	InstallSnapshot(ctx context.Context, c *consenusImpl, in *raft.RequestInstallSnapshot) (*raft.ResponseInstallSnapshot, error)
}
