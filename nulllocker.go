package raft

type nullLocker struct {
}

func (locker *nullLocker) Lock() {

}
func (locker *nullLocker) Unlock() {

}
