package raft

import (
	"testing"

	"github.com/cap-labs/go-cap/raft/mock"
	"github.com/cap-labs/go-cap/raft/testsuite"
	"github.com/libs4go/smf4go/tester"
)

func TestConsensus(t *testing.T) {
	tester.T(t).Run(testMain)
}

func testMain(tester tester.Tester) {
	tester.ConfigPath("./test/test.json")
	// register raft service deps services
	tester.LocalService().Register("raft.locker", mock.NewLocker)
	tester.LocalService().Register("raft.cluster.manager", mock.NewClusterManager)
	tester.LocalService().Register("raft.logstore", mock.NewLogStore)
	tester.LocalService().Register("raft.network", mock.NewNetwork)
	// register raft main service
	tester.LocalService().Register("raft.consensus", New)
	// register raft standard test suite and run
	tester.LocalService().Register("raft.testsuite", testsuite.New)
}
