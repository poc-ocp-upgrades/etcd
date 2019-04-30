package rafthttp

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"sync"
	"time"
	stats "go.etcd.io/etcd/etcdserver/api/v2stats"
	"go.etcd.io/etcd/pkg/pbutil"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
	"go.uber.org/zap"
)

const (
	connPerPipeline	= 4
	pipelineBufSize	= 64
)

var errStopped = errors.New("stopped")

type pipeline struct {
	peerID		types.ID
	tr		*Transport
	picker		*urlPicker
	status		*peerStatus
	raft		Raft
	errorc		chan error
	followerStats	*stats.FollowerStats
	msgc		chan raftpb.Message
	wg		sync.WaitGroup
	stopc		chan struct{}
}

func (p *pipeline) start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.stopc = make(chan struct{})
	p.msgc = make(chan raftpb.Message, pipelineBufSize)
	p.wg.Add(connPerPipeline)
	for i := 0; i < connPerPipeline; i++ {
		go p.handle()
	}
	if p.tr != nil && p.tr.Logger != nil {
		p.tr.Logger.Info("started HTTP pipelining with remote peer", zap.String("local-member-id", p.tr.ID.String()), zap.String("remote-peer-id", p.peerID.String()))
	} else {
		plog.Infof("started HTTP pipelining with peer %s", p.peerID)
	}
}
func (p *pipeline) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(p.stopc)
	p.wg.Wait()
	if p.tr != nil && p.tr.Logger != nil {
		p.tr.Logger.Info("stopped HTTP pipelining with remote peer", zap.String("local-member-id", p.tr.ID.String()), zap.String("remote-peer-id", p.peerID.String()))
	} else {
		plog.Infof("stopped HTTP pipelining with peer %s", p.peerID)
	}
}
func (p *pipeline) handle() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer p.wg.Done()
	for {
		select {
		case m := <-p.msgc:
			start := time.Now()
			err := p.post(pbutil.MustMarshal(&m))
			end := time.Now()
			if err != nil {
				p.status.deactivate(failureType{source: pipelineMsg, action: "write"}, err.Error())
				if m.Type == raftpb.MsgApp && p.followerStats != nil {
					p.followerStats.Fail()
				}
				p.raft.ReportUnreachable(m.To)
				if isMsgSnap(m) {
					p.raft.ReportSnapshot(m.To, raft.SnapshotFailure)
				}
				sentFailures.WithLabelValues(types.ID(m.To).String()).Inc()
				continue
			}
			p.status.activate()
			if m.Type == raftpb.MsgApp && p.followerStats != nil {
				p.followerStats.Succ(end.Sub(start))
			}
			if isMsgSnap(m) {
				p.raft.ReportSnapshot(m.To, raft.SnapshotFinish)
			}
			sentBytes.WithLabelValues(types.ID(m.To).String()).Add(float64(m.Size()))
		case <-p.stopc:
			return
		}
	}
}
func (p *pipeline) post(data []byte) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := p.picker.pick()
	req := createPostRequest(u, RaftPrefix, bytes.NewBuffer(data), "application/protobuf", p.tr.URLs, p.tr.ID, p.tr.ClusterID)
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	req = req.WithContext(ctx)
	go func() {
		select {
		case <-done:
		case <-p.stopc:
			waitSchedule()
			cancel()
		}
	}()
	resp, err := p.tr.pipelineRt.RoundTrip(req)
	done <- struct{}{}
	if err != nil {
		p.picker.unreachable(u)
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.picker.unreachable(u)
		return err
	}
	err = checkPostResponse(resp, b, req, p.peerID)
	if err != nil {
		p.picker.unreachable(u)
		if err == errMemberRemoved {
			reportCriticalError(err, p.errorc)
		}
		return err
	}
	return nil
}
func waitSchedule() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	time.Sleep(time.Millisecond)
}
