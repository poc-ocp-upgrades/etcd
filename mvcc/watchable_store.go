package mvcc

import (
	"sync"
	"time"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

var (
	chanBufLen		= 1024
	maxWatchersPerSync	= 512
)

type watchable interface {
	watch(key, end []byte, startRev int64, id WatchID, ch chan<- WatchResponse, fcs ...FilterFunc) (*watcher, cancelFunc)
	progress(w *watcher)
	rev() int64
}
type watchableStore struct {
	*store
	mu		sync.RWMutex
	victims		[]watcherBatch
	victimc		chan struct{}
	unsynced	watcherGroup
	synced		watcherGroup
	stopc		chan struct{}
	wg		sync.WaitGroup
}
type cancelFunc func()

func New(b backend.Backend, le lease.Lessor, ig ConsistentIndexGetter) ConsistentWatchableKV {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newWatchableStore(b, le, ig)
}
func newWatchableStore(b backend.Backend, le lease.Lessor, ig ConsistentIndexGetter) *watchableStore {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := &watchableStore{store: NewStore(b, le, ig), victimc: make(chan struct{}, 1), unsynced: newWatcherGroup(), synced: newWatcherGroup(), stopc: make(chan struct{})}
	s.store.ReadView = &readView{s}
	s.store.WriteView = &writeView{s}
	if s.le != nil {
		s.le.SetRangeDeleter(func() lease.TxnDelete {
			return s.Write()
		})
	}
	s.wg.Add(2)
	go s.syncWatchersLoop()
	go s.syncVictimsLoop()
	return s
}
func (s *watchableStore) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(s.stopc)
	s.wg.Wait()
	return s.store.Close()
}
func (s *watchableStore) NewWatchStream() WatchStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	watchStreamGauge.Inc()
	return &watchStream{watchable: s, ch: make(chan WatchResponse, chanBufLen), cancels: make(map[WatchID]cancelFunc), watchers: make(map[WatchID]*watcher)}
}
func (s *watchableStore) watch(key, end []byte, startRev int64, id WatchID, ch chan<- WatchResponse, fcs ...FilterFunc) (*watcher, cancelFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wa := &watcher{key: key, end: end, minRev: startRev, id: id, ch: ch, fcs: fcs}
	s.mu.Lock()
	s.revMu.RLock()
	synced := startRev > s.store.currentRev || startRev == 0
	if synced {
		wa.minRev = s.store.currentRev + 1
		if startRev > wa.minRev {
			wa.minRev = startRev
		}
	}
	if synced {
		s.synced.add(wa)
	} else {
		slowWatcherGauge.Inc()
		s.unsynced.add(wa)
	}
	s.revMu.RUnlock()
	s.mu.Unlock()
	watcherGauge.Inc()
	return wa, func() {
		s.cancelWatcher(wa)
	}
}
func (s *watchableStore) cancelWatcher(wa *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		s.mu.Lock()
		if s.unsynced.delete(wa) {
			slowWatcherGauge.Dec()
			break
		} else if s.synced.delete(wa) {
			break
		} else if wa.compacted {
			break
		} else if wa.ch == nil {
			break
		}
		if !wa.victim {
			panic("watcher not victim but not in watch groups")
		}
		var victimBatch watcherBatch
		for _, wb := range s.victims {
			if wb[wa] != nil {
				victimBatch = wb
				break
			}
		}
		if victimBatch != nil {
			slowWatcherGauge.Dec()
			delete(victimBatch, wa)
			break
		}
		s.mu.Unlock()
		time.Sleep(time.Millisecond)
	}
	watcherGauge.Dec()
	wa.ch = nil
	s.mu.Unlock()
}
func (s *watchableStore) Restore(b backend.Backend) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.store.Restore(b)
	if err != nil {
		return err
	}
	for wa := range s.synced.watchers {
		wa.restore = true
		s.unsynced.add(wa)
	}
	s.synced = newWatcherGroup()
	return nil
}
func (s *watchableStore) syncWatchersLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer s.wg.Done()
	for {
		s.mu.RLock()
		st := time.Now()
		lastUnsyncedWatchers := s.unsynced.size()
		s.mu.RUnlock()
		unsyncedWatchers := 0
		if lastUnsyncedWatchers > 0 {
			unsyncedWatchers = s.syncWatchers()
		}
		syncDuration := time.Since(st)
		waitDuration := 100 * time.Millisecond
		if unsyncedWatchers != 0 && lastUnsyncedWatchers > unsyncedWatchers {
			waitDuration = syncDuration
		}
		select {
		case <-time.After(waitDuration):
		case <-s.stopc:
			return
		}
	}
}
func (s *watchableStore) syncVictimsLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer s.wg.Done()
	for {
		for s.moveVictims() != 0 {
		}
		s.mu.RLock()
		isEmpty := len(s.victims) == 0
		s.mu.RUnlock()
		var tickc <-chan time.Time
		if !isEmpty {
			tickc = time.After(10 * time.Millisecond)
		}
		select {
		case <-tickc:
		case <-s.victimc:
		case <-s.stopc:
			return
		}
	}
}
func (s *watchableStore) moveVictims() (moved int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	victims := s.victims
	s.victims = nil
	s.mu.Unlock()
	var newVictim watcherBatch
	for _, wb := range victims {
		for w, eb := range wb {
			rev := w.minRev - 1
			if w.send(WatchResponse{WatchID: w.id, Events: eb.evs, Revision: rev}) {
				pendingEventsGauge.Add(float64(len(eb.evs)))
			} else {
				if newVictim == nil {
					newVictim = make(watcherBatch)
				}
				newVictim[w] = eb
				continue
			}
			moved++
		}
		s.mu.Lock()
		s.store.revMu.RLock()
		curRev := s.store.currentRev
		for w, eb := range wb {
			if newVictim != nil && newVictim[w] != nil {
				continue
			}
			w.victim = false
			if eb.moreRev != 0 {
				w.minRev = eb.moreRev
			}
			if w.minRev <= curRev {
				s.unsynced.add(w)
			} else {
				slowWatcherGauge.Dec()
				s.synced.add(w)
			}
		}
		s.store.revMu.RUnlock()
		s.mu.Unlock()
	}
	if len(newVictim) > 0 {
		s.mu.Lock()
		s.victims = append(s.victims, newVictim)
		s.mu.Unlock()
	}
	return moved
}
func (s *watchableStore) syncWatchers() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.unsynced.size() == 0 {
		return 0
	}
	s.store.revMu.RLock()
	defer s.store.revMu.RUnlock()
	curRev := s.store.currentRev
	compactionRev := s.store.compactMainRev
	wg, minRev := s.unsynced.choose(maxWatchersPerSync, curRev, compactionRev)
	minBytes, maxBytes := newRevBytes(), newRevBytes()
	revToBytes(revision{main: minRev}, minBytes)
	revToBytes(revision{main: curRev + 1}, maxBytes)
	tx := s.store.b.ReadTx()
	tx.Lock()
	revs, vs := tx.UnsafeRange(keyBucketName, minBytes, maxBytes, 0)
	evs := kvsToEvents(wg, revs, vs)
	tx.Unlock()
	var victims watcherBatch
	wb := newWatcherBatch(wg, evs)
	for w := range wg.watchers {
		w.minRev = curRev + 1
		eb, ok := wb[w]
		if !ok {
			s.synced.add(w)
			s.unsynced.delete(w)
			continue
		}
		if eb.moreRev != 0 {
			w.minRev = eb.moreRev
		}
		if w.send(WatchResponse{WatchID: w.id, Events: eb.evs, Revision: curRev}) {
			pendingEventsGauge.Add(float64(len(eb.evs)))
		} else {
			if victims == nil {
				victims = make(watcherBatch)
			}
			w.victim = true
		}
		if w.victim {
			victims[w] = eb
		} else {
			if eb.moreRev != 0 {
				continue
			}
			s.synced.add(w)
		}
		s.unsynced.delete(w)
	}
	s.addVictim(victims)
	vsz := 0
	for _, v := range s.victims {
		vsz += len(v)
	}
	slowWatcherGauge.Set(float64(s.unsynced.size() + vsz))
	return s.unsynced.size()
}
func kvsToEvents(wg *watcherGroup, revs, vals [][]byte) (evs []mvccpb.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, v := range vals {
		var kv mvccpb.KeyValue
		if err := kv.Unmarshal(v); err != nil {
			plog.Panicf("cannot unmarshal event: %v", err)
		}
		if !wg.contains(string(kv.Key)) {
			continue
		}
		ty := mvccpb.PUT
		if isTombstone(revs[i]) {
			ty = mvccpb.DELETE
			kv.ModRevision = bytesToRev(revs[i]).main
		}
		evs = append(evs, mvccpb.Event{Kv: &kv, Type: ty})
	}
	return evs
}
func (s *watchableStore) notify(rev int64, evs []mvccpb.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var victim watcherBatch
	for w, eb := range newWatcherBatch(&s.synced, evs) {
		if eb.revs != 1 {
			plog.Panicf("unexpected multiple revisions in notification")
		}
		if w.send(WatchResponse{WatchID: w.id, Events: eb.evs, Revision: rev}) {
			pendingEventsGauge.Add(float64(len(eb.evs)))
		} else {
			w.minRev = rev + 1
			if victim == nil {
				victim = make(watcherBatch)
			}
			w.victim = true
			victim[w] = eb
			s.synced.delete(w)
			slowWatcherGauge.Inc()
		}
	}
	s.addVictim(victim)
}
func (s *watchableStore) addVictim(victim watcherBatch) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if victim == nil {
		return
	}
	s.victims = append(s.victims, victim)
	select {
	case s.victimc <- struct{}{}:
	default:
	}
}
func (s *watchableStore) rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.store.Rev()
}
func (s *watchableStore) progress(w *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.synced.watchers[w]; ok {
		w.send(WatchResponse{WatchID: w.id, Revision: s.rev()})
	}
}

type watcher struct {
	key		[]byte
	end		[]byte
	victim		bool
	compacted	bool
	restore		bool
	minRev		int64
	id		WatchID
	fcs		[]FilterFunc
	ch		chan<- WatchResponse
}

func (w *watcher) send(wr WatchResponse) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	progressEvent := len(wr.Events) == 0
	if len(w.fcs) != 0 {
		ne := make([]mvccpb.Event, 0, len(wr.Events))
		for i := range wr.Events {
			filtered := false
			for _, filter := range w.fcs {
				if filter(wr.Events[i]) {
					filtered = true
					break
				}
			}
			if !filtered {
				ne = append(ne, wr.Events[i])
			}
		}
		wr.Events = ne
	}
	if !progressEvent && len(wr.Events) == 0 {
		return true
	}
	select {
	case w.ch <- wr:
		return true
	default:
		return false
	}
}
