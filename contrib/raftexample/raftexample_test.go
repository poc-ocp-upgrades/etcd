package main

import (
	"fmt"
	"os"
	"testing"
	"github.com/coreos/etcd/raft/raftpb"
)

type cluster struct {
	peers		[]string
	commitC		[]<-chan *string
	errorC		[]<-chan error
	proposeC	[]chan string
	confChangeC	[]chan raftpb.ConfChange
}

func newCluster(n int) *cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	peers := make([]string, n)
	for i := range peers {
		peers[i] = fmt.Sprintf("http://127.0.0.1:%d", 10000+i)
	}
	clus := &cluster{peers: peers, commitC: make([]<-chan *string, len(peers)), errorC: make([]<-chan error, len(peers)), proposeC: make([]chan string, len(peers)), confChangeC: make([]chan raftpb.ConfChange, len(peers))}
	for i := range clus.peers {
		os.RemoveAll(fmt.Sprintf("raftexample-%d", i+1))
		os.RemoveAll(fmt.Sprintf("raftexample-%d-snap", i+1))
		clus.proposeC[i] = make(chan string, 1)
		clus.confChangeC[i] = make(chan raftpb.ConfChange, 1)
		clus.commitC[i], clus.errorC[i], _ = newRaftNode(i+1, clus.peers, false, nil, clus.proposeC[i], clus.confChangeC[i])
	}
	return clus
}
func (clus *cluster) sinkReplay() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range clus.peers {
		for s := range clus.commitC[i] {
			if s == nil {
				break
			}
		}
	}
}
func (clus *cluster) Close() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range clus.peers {
		close(clus.proposeC[i])
		for range clus.commitC[i] {
		}
		if erri := <-clus.errorC[i]; erri != nil {
			err = erri
		}
		os.RemoveAll(fmt.Sprintf("raftexample-%d", i+1))
		os.RemoveAll(fmt.Sprintf("raftexample-%d-snap", i+1))
	}
	return err
}
func (clus *cluster) closeNoErrors(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := clus.Close(); err != nil {
		t.Fatal(err)
	}
}
func TestProposeOnCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := newCluster(3)
	defer clus.closeNoErrors(t)
	clus.sinkReplay()
	donec := make(chan struct{})
	for i := range clus.peers {
		go func(pC chan<- string, cC <-chan *string, eC <-chan error) {
			for n := 0; n < 100; n++ {
				s, ok := <-cC
				if !ok {
					pC = nil
				}
				select {
				case pC <- *s:
					continue
				case err := <-eC:
					t.Fatalf("eC message (%v)", err)
				}
			}
			donec <- struct{}{}
			for range cC {
			}
		}(clus.proposeC[i], clus.commitC[i], clus.errorC[i])
		go func(i int) {
			clus.proposeC[i] <- "foo"
		}(i)
	}
	for range clus.peers {
		<-donec
	}
}
func TestCloseProposerBeforeReplay(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := newCluster(1)
	defer clus.closeNoErrors(t)
}
func TestCloseProposerInflight(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := newCluster(1)
	defer clus.closeNoErrors(t)
	clus.sinkReplay()
	go func() {
		clus.proposeC[0] <- "foo"
		clus.proposeC[0] <- "bar"
	}()
	if c, ok := <-clus.commitC[0]; *c != "foo" || !ok {
		t.Fatalf("Commit failed")
	}
}
