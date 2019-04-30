package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"sync"
	"go.etcd.io/etcd/etcdserver/api/snap"
)

type kvstore struct {
	proposeC	chan<- string
	mu		sync.RWMutex
	kvStore		map[string]string
	snapshotter	*snap.Snapshotter
}
type kv struct {
	Key	string
	Val	string
}

func newKVStore(snapshotter *snap.Snapshotter, proposeC chan<- string, commitC <-chan *string, errorC <-chan error) *kvstore {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := &kvstore{proposeC: proposeC, kvStore: make(map[string]string), snapshotter: snapshotter}
	s.readCommits(commitC, errorC)
	go s.readCommits(commitC, errorC)
	return s
}
func (s *kvstore) Lookup(key string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.kvStore[key]
	return v, ok
}
func (s *kvstore) Propose(k string, v string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(kv{k, v}); err != nil {
		log.Fatal(err)
	}
	s.proposeC <- buf.String()
}
func (s *kvstore) readCommits(commitC <-chan *string, errorC <-chan error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for data := range commitC {
		if data == nil {
			snapshot, err := s.snapshotter.Load()
			if err == snap.ErrNoSnapshot {
				return
			}
			if err != nil {
				log.Panic(err)
			}
			log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
			if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
				log.Panic(err)
			}
			continue
		}
		var dataKv kv
		dec := gob.NewDecoder(bytes.NewBufferString(*data))
		if err := dec.Decode(&dataKv); err != nil {
			log.Fatalf("raftexample: could not decode message (%v)", err)
		}
		s.mu.Lock()
		s.kvStore[dataKv.Key] = dataKv.Val
		s.mu.Unlock()
	}
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}
func (s *kvstore) getSnapshot() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return json.Marshal(s.kvStore)
}
func (s *kvstore) recoverFromSnapshot(snapshot []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var store map[string]string
	if err := json.Unmarshal(snapshot, &store); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.kvStore = store
	return nil
}
