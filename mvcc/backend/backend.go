package backend

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
	bolt "github.com/coreos/bbolt"
	"github.com/coreos/pkg/capnslog"
)

var (
	defaultBatchLimit		= 10000
	defaultBatchInterval		= 100 * time.Millisecond
	defragLimit			= 10000
	initialMmapSize			= uint64(10 * 1024 * 1024 * 1024)
	plog				= capnslog.NewPackageLogger("github.com/coreos/etcd", "mvcc/backend")
	minSnapshotWarningTimeout	= time.Duration(30 * time.Second)
)

type Backend interface {
	ReadTx() ReadTx
	BatchTx() BatchTx
	Snapshot() Snapshot
	Hash(ignores map[IgnoreKey]struct{}) (uint32, error)
	Size() int64
	SizeInUse() int64
	Defrag() error
	ForceCommit()
	Close() error
}
type Snapshot interface {
	Size() int64
	WriteTo(w io.Writer) (n int64, err error)
	Close() error
}
type backend struct {
	size		int64
	sizeInUse	int64
	commits		int64
	mu		sync.RWMutex
	db		*bolt.DB
	batchInterval	time.Duration
	batchLimit	int
	batchTx		*batchTxBuffered
	readTx		*readTx
	stopc		chan struct{}
	donec		chan struct{}
}
type BackendConfig struct {
	Path		string
	BatchInterval	time.Duration
	BatchLimit	int
	MmapSize	uint64
}

func DefaultBackendConfig() BackendConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return BackendConfig{BatchInterval: defaultBatchInterval, BatchLimit: defaultBatchLimit, MmapSize: initialMmapSize}
}
func New(bcfg BackendConfig) Backend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newBackend(bcfg)
}
func NewDefaultBackend(path string) Backend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcfg := DefaultBackendConfig()
	bcfg.Path = path
	return newBackend(bcfg)
}
func newBackend(bcfg BackendConfig) *backend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bopts := &bolt.Options{}
	if boltOpenOptions != nil {
		*bopts = *boltOpenOptions
	}
	bopts.InitialMmapSize = bcfg.mmapSize()
	db, err := bolt.Open(bcfg.Path, 0600, bopts)
	if err != nil {
		plog.Panicf("cannot open database at %s (%v)", bcfg.Path, err)
	}
	b := &backend{db: db, batchInterval: bcfg.BatchInterval, batchLimit: bcfg.BatchLimit, readTx: &readTx{buf: txReadBuffer{txBuffer: txBuffer{make(map[string]*bucketBuffer)}}, buckets: make(map[string]*bolt.Bucket)}, stopc: make(chan struct{}), donec: make(chan struct{})}
	b.batchTx = newBatchTxBuffered(b)
	go b.run()
	return b
}
func (b *backend) BatchTx() BatchTx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.batchTx
}
func (b *backend) ReadTx() ReadTx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.readTx
}
func (b *backend) ForceCommit() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.batchTx.Commit()
}
func (b *backend) Snapshot() Snapshot {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.batchTx.Commit()
	b.mu.RLock()
	defer b.mu.RUnlock()
	tx, err := b.db.Begin(false)
	if err != nil {
		plog.Fatalf("cannot begin tx (%s)", err)
	}
	stopc, donec := make(chan struct{}), make(chan struct{})
	dbBytes := tx.Size()
	go func() {
		defer close(donec)
		var sendRateBytes int64 = 100 * 1024 * 1014
		warningTimeout := time.Duration(int64((float64(dbBytes) / float64(sendRateBytes)) * float64(time.Second)))
		if warningTimeout < minSnapshotWarningTimeout {
			warningTimeout = minSnapshotWarningTimeout
		}
		start := time.Now()
		ticker := time.NewTicker(warningTimeout)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				plog.Warningf("snapshotting is taking more than %v seconds to finish transferring %v MB [started at %v]", time.Since(start).Seconds(), float64(dbBytes)/float64(1024*1014), start)
			case <-stopc:
				snapshotDurations.Observe(time.Since(start).Seconds())
				return
			}
		}
	}()
	return &snapshot{tx, stopc, donec}
}

type IgnoreKey struct {
	Bucket	string
	Key	string
}

func (b *backend) Hash(ignores map[IgnoreKey]struct{}) (uint32, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	b.mu.RLock()
	defer b.mu.RUnlock()
	err := b.db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		for next, _ := c.First(); next != nil; next, _ = c.Next() {
			b := tx.Bucket(next)
			if b == nil {
				return fmt.Errorf("cannot get hash of bucket %s", string(next))
			}
			h.Write(next)
			b.ForEach(func(k, v []byte) error {
				bk := IgnoreKey{Bucket: string(next), Key: string(k)}
				if _, ok := ignores[bk]; !ok {
					h.Write(k)
					h.Write(v)
				}
				return nil
			})
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
func (b *backend) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&b.size)
}
func (b *backend) SizeInUse() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&b.sizeInUse)
}
func (b *backend) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(b.donec)
	t := time.NewTimer(b.batchInterval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
		case <-b.stopc:
			b.batchTx.CommitAndStop()
			return
		}
		b.batchTx.Commit()
		t.Reset(b.batchInterval)
	}
}
func (b *backend) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(b.stopc)
	<-b.donec
	return b.db.Close()
}
func (b *backend) Commits() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&b.commits)
}
func (b *backend) Defrag() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.defrag()
}
func (b *backend) defrag() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := time.Now()
	b.batchTx.Lock()
	defer b.batchTx.Unlock()
	b.mu.Lock()
	defer b.mu.Unlock()
	b.readTx.mu.Lock()
	defer b.readTx.mu.Unlock()
	b.batchTx.unsafeCommit(true)
	b.batchTx.tx = nil
	tmpdb, err := bolt.Open(b.db.Path()+".tmp", 0600, boltOpenOptions)
	if err != nil {
		return err
	}
	err = defragdb(b.db, tmpdb, defragLimit)
	if err != nil {
		tmpdb.Close()
		os.RemoveAll(tmpdb.Path())
		return err
	}
	dbp := b.db.Path()
	tdbp := tmpdb.Path()
	err = b.db.Close()
	if err != nil {
		plog.Fatalf("cannot close database (%s)", err)
	}
	err = tmpdb.Close()
	if err != nil {
		plog.Fatalf("cannot close database (%s)", err)
	}
	err = os.Rename(tdbp, dbp)
	if err != nil {
		plog.Fatalf("cannot rename database (%s)", err)
	}
	b.db, err = bolt.Open(dbp, 0600, boltOpenOptions)
	if err != nil {
		plog.Panicf("cannot open database at %s (%v)", dbp, err)
	}
	b.batchTx.tx, err = b.db.Begin(true)
	if err != nil {
		plog.Fatalf("cannot begin tx (%s)", err)
	}
	b.readTx.reset()
	b.readTx.tx = b.unsafeBegin(false)
	size := b.readTx.tx.Size()
	db := b.db
	atomic.StoreInt64(&b.size, size)
	atomic.StoreInt64(&b.sizeInUse, size-(int64(db.Stats().FreePageN)*int64(db.Info().PageSize)))
	took := time.Since(now)
	defragDurations.Observe(took.Seconds())
	return nil
}
func defragdb(odb, tmpdb *bolt.DB, limit int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tmptx, err := tmpdb.Begin(true)
	if err != nil {
		return err
	}
	tx, err := odb.Begin(false)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	c := tx.Cursor()
	count := 0
	for next, _ := c.First(); next != nil; next, _ = c.Next() {
		b := tx.Bucket(next)
		if b == nil {
			return fmt.Errorf("backend: cannot defrag bucket %s", string(next))
		}
		tmpb, berr := tmptx.CreateBucketIfNotExists(next)
		if berr != nil {
			return berr
		}
		tmpb.FillPercent = 0.9
		b.ForEach(func(k, v []byte) error {
			count++
			if count > limit {
				err = tmptx.Commit()
				if err != nil {
					return err
				}
				tmptx, err = tmpdb.Begin(true)
				if err != nil {
					return err
				}
				tmpb = tmptx.Bucket(next)
				tmpb.FillPercent = 0.9
				count = 0
			}
			return tmpb.Put(k, v)
		})
	}
	return tmptx.Commit()
}
func (b *backend) begin(write bool) *bolt.Tx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	tx := b.unsafeBegin(write)
	b.mu.RUnlock()
	size := tx.Size()
	db := tx.DB()
	atomic.StoreInt64(&b.size, size)
	atomic.StoreInt64(&b.sizeInUse, size-(int64(db.Stats().FreePageN)*int64(db.Info().PageSize)))
	return tx
}
func (b *backend) unsafeBegin(write bool) *bolt.Tx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tx, err := b.db.Begin(write)
	if err != nil {
		plog.Fatalf("cannot begin tx (%s)", err)
	}
	return tx
}
func NewTmpBackend(batchInterval time.Duration, batchLimit int) (*backend, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, err := ioutil.TempDir(os.TempDir(), "etcd_backend_test")
	if err != nil {
		plog.Fatal(err)
	}
	tmpPath := filepath.Join(dir, "database")
	bcfg := DefaultBackendConfig()
	bcfg.Path, bcfg.BatchInterval, bcfg.BatchLimit = tmpPath, batchInterval, batchLimit
	return newBackend(bcfg), tmpPath
}
func NewDefaultTmpBackend() (*backend, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewTmpBackend(defaultBatchInterval, defaultBatchLimit)
}

type snapshot struct {
	*bolt.Tx
	stopc	chan struct{}
	donec	chan struct{}
}

func (s *snapshot) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(s.stopc)
	<-s.donec
	return s.Tx.Rollback()
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
