package rafthttp

import (
	"context"
	"net/http"
	"sync"
	"time"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/logutil"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/pkg/capnslog"
	"github.com/xiang90/probing"
	"golang.org/x/time/rate"
)

var plog = logutil.NewMergeLogger(capnslog.NewPackageLogger("github.com/coreos/etcd", "rafthttp"))

type Raft interface {
	Process(ctx context.Context, m raftpb.Message) error
	IsIDRemoved(id uint64) bool
	ReportUnreachable(id uint64)
	ReportSnapshot(id uint64, status raft.SnapshotStatus)
}
type Transporter interface {
	Start() error
	Handler() http.Handler
	Send(m []raftpb.Message)
	SendSnapshot(m snap.Message)
	AddRemote(id types.ID, urls []string)
	AddPeer(id types.ID, urls []string)
	RemovePeer(id types.ID)
	RemoveAllPeers()
	UpdatePeer(id types.ID, urls []string)
	ActiveSince(id types.ID) time.Time
	ActivePeers() int
	Stop()
}
type Transport struct {
	DialTimeout		time.Duration
	DialRetryFrequency	rate.Limit
	TLSInfo			transport.TLSInfo
	ID			types.ID
	URLs			types.URLs
	ClusterID		types.ID
	Raft			Raft
	Snapshotter		*snap.Snapshotter
	ServerStats		*stats.ServerStats
	LeaderStats		*stats.LeaderStats
	ErrorC			chan error
	streamRt		http.RoundTripper
	pipelineRt		http.RoundTripper
	mu			sync.RWMutex
	remotes			map[types.ID]*remote
	peers			map[types.ID]Peer
	pipelineProber		probing.Prober
	streamProber		probing.Prober
}

func (t *Transport) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	t.streamRt, err = newStreamRoundTripper(t.TLSInfo, t.DialTimeout)
	if err != nil {
		return err
	}
	t.pipelineRt, err = NewRoundTripper(t.TLSInfo, t.DialTimeout)
	if err != nil {
		return err
	}
	t.remotes = make(map[types.ID]*remote)
	t.peers = make(map[types.ID]Peer)
	t.pipelineProber = probing.NewProber(t.pipelineRt)
	t.streamProber = probing.NewProber(t.streamRt)
	if t.DialRetryFrequency == 0 {
		t.DialRetryFrequency = rate.Every(100 * time.Millisecond)
	}
	return nil
}
func (t *Transport) Handler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pipelineHandler := newPipelineHandler(t, t.Raft, t.ClusterID)
	streamHandler := newStreamHandler(t, t, t.Raft, t.ID, t.ClusterID)
	snapHandler := newSnapshotHandler(t, t.Raft, t.Snapshotter, t.ClusterID)
	mux := http.NewServeMux()
	mux.Handle(RaftPrefix, pipelineHandler)
	mux.Handle(RaftStreamPrefix+"/", streamHandler)
	mux.Handle(RaftSnapshotPrefix, snapHandler)
	mux.Handle(ProbingPrefix, probing.NewHandler())
	return mux
}
func (t *Transport) Get(id types.ID) Peer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.peers[id]
}
func (t *Transport) Send(msgs []raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range msgs {
		if m.To == 0 {
			continue
		}
		to := types.ID(m.To)
		t.mu.RLock()
		p, pok := t.peers[to]
		g, rok := t.remotes[to]
		t.mu.RUnlock()
		if pok {
			if m.Type == raftpb.MsgApp {
				t.ServerStats.SendAppendReq(m.Size())
			}
			p.send(m)
			continue
		}
		if rok {
			g.send(m)
			continue
		}
		plog.Debugf("ignored message %s (sent to unknown peer %s)", m.Type, to)
	}
}
func (t *Transport) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, r := range t.remotes {
		r.stop()
	}
	for _, p := range t.peers {
		p.stop()
	}
	t.pipelineProber.RemoveAll()
	t.streamProber.RemoveAll()
	if tr, ok := t.streamRt.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
	if tr, ok := t.pipelineRt.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
	t.peers = nil
	t.remotes = nil
}
func (t *Transport) CutPeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.RLock()
	p, pok := t.peers[id]
	g, gok := t.remotes[id]
	t.mu.RUnlock()
	if pok {
		p.(Pausable).Pause()
	}
	if gok {
		g.Pause()
	}
}
func (t *Transport) MendPeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.RLock()
	p, pok := t.peers[id]
	g, gok := t.remotes[id]
	t.mu.RUnlock()
	if pok {
		p.(Pausable).Resume()
	}
	if gok {
		g.Resume()
	}
}
func (t *Transport) AddRemote(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.remotes == nil {
		return
	}
	if _, ok := t.peers[id]; ok {
		return
	}
	if _, ok := t.remotes[id]; ok {
		return
	}
	urls, err := types.NewURLs(us)
	if err != nil {
		plog.Panicf("newURLs %+v should never fail: %+v", us, err)
	}
	t.remotes[id] = startRemote(t, urls, id)
}
func (t *Transport) AddPeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.peers == nil {
		panic("transport stopped")
	}
	if _, ok := t.peers[id]; ok {
		return
	}
	urls, err := types.NewURLs(us)
	if err != nil {
		plog.Panicf("newURLs %+v should never fail: %+v", us, err)
	}
	fs := t.LeaderStats.Follower(id.String())
	t.peers[id] = startPeer(t, urls, id, fs)
	addPeerToProber(t.pipelineProber, id.String(), us, RoundTripperNameSnapshot, rtts)
	addPeerToProber(t.streamProber, id.String(), us, RoundTripperNameRaftMessage, rtts)
	plog.Infof("added peer %s", id)
}
func (t *Transport) RemovePeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.removePeer(id)
}
func (t *Transport) RemoveAllPeers() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	for id := range t.peers {
		t.removePeer(id)
	}
}
func (t *Transport) removePeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if peer, ok := t.peers[id]; ok {
		peer.stop()
	} else {
		plog.Panicf("unexpected removal of unknown peer '%d'", id)
	}
	delete(t.peers, id)
	delete(t.LeaderStats.Followers, id.String())
	t.pipelineProber.Remove(id.String())
	t.streamProber.Remove(id.String())
	plog.Infof("removed peer %s", id)
}
func (t *Transport) UpdatePeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.peers[id]; !ok {
		return
	}
	urls, err := types.NewURLs(us)
	if err != nil {
		plog.Panicf("newURLs %+v should never fail: %+v", us, err)
	}
	t.peers[id].update(urls)
	t.pipelineProber.Remove(id.String())
	addPeerToProber(t.pipelineProber, id.String(), us, RoundTripperNameSnapshot, rtts)
	t.streamProber.Remove(id.String())
	addPeerToProber(t.streamProber, id.String(), us, RoundTripperNameRaftMessage, rtts)
	plog.Infof("updated peer %s", id)
}
func (t *Transport) ActiveSince(id types.ID) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if p, ok := t.peers[id]; ok {
		return p.activeSince()
	}
	return time.Time{}
}
func (t *Transport) SendSnapshot(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	p := t.peers[types.ID(m.To)]
	if p == nil {
		m.CloseWithError(errMemberNotFound)
		return
	}
	p.sendSnap(m)
}

type Pausable interface {
	Pause()
	Resume()
}

func (t *Transport) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, p := range t.peers {
		p.(Pausable).Pause()
	}
}
func (t *Transport) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, p := range t.peers {
		p.(Pausable).Resume()
	}
}
func (t *Transport) ActivePeers() (cnt int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, p := range t.peers {
		if !p.activeSince().IsZero() {
			cnt++
		}
	}
	return cnt
}

type nopTransporter struct{}

func NewNopTransporter() Transporter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &nopTransporter{}
}
func (s *nopTransporter) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *nopTransporter) Handler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *nopTransporter) Send(m []raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) SendSnapshot(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) AddRemote(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) AddPeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) RemovePeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) RemoveAllPeers() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) UpdatePeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) ActiveSince(id types.ID) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Time{}
}
func (s *nopTransporter) ActivePeers() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (s *nopTransporter) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporter) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type snapTransporter struct {
	nopTransporter
	snapDoneC	chan snap.Message
	snapDir		string
}

func NewSnapTransporter(snapDir string) (Transporter, <-chan snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan snap.Message, 1)
	tr := &snapTransporter{snapDoneC: ch, snapDir: snapDir}
	return tr, ch
}
func (s *snapTransporter) SendSnapshot(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss := snap.New(s.snapDir)
	ss.SaveDBFrom(m.ReadCloser, m.Snapshot.Metadata.Index+1)
	m.CloseWithError(nil)
	s.snapDoneC <- m
}
