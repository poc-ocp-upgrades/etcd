package fileutil

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

const (
	F_OFD_GETLK	= 37
	F_OFD_SETLK	= 37
	F_OFD_SETLKW	= 38
)

var (
	wrlck			= syscall.Flock_t{Type: syscall.F_WRLCK, Whence: int16(io.SeekStart), Start: 0, Len: 0}
	linuxTryLockFile	= flockTryLockFile
	linuxLockFile		= flockLockFile
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	getlk := syscall.Flock_t{Type: syscall.F_RDLCK}
	if err := syscall.FcntlFlock(0, F_OFD_GETLK, &getlk); err == nil {
		linuxTryLockFile = ofdTryLockFile
		linuxLockFile = ofdLockFile
	}
}
func TryLockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return linuxTryLockFile(path, flag, perm)
}
func ofdTryLockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, fmt.Errorf("ofdTryLockFile failed to open %q (%v)", path, err)
	}
	flock := wrlck
	if err = syscall.FcntlFlock(f.Fd(), F_OFD_SETLK, &flock); err != nil {
		f.Close()
		if err == syscall.EWOULDBLOCK {
			err = ErrLocked
		}
		return nil, err
	}
	return &LockedFile{f}, nil
}
func LockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return linuxLockFile(path, flag, perm)
}
func ofdLockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, fmt.Errorf("ofdLockFile failed to open %q (%v)", path, err)
	}
	flock := wrlck
	err = syscall.FcntlFlock(f.Fd(), F_OFD_SETLKW, &flock)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &LockedFile{f}, nil
}
