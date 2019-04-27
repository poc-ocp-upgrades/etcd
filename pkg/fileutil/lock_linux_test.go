package fileutil

import "testing"

func TestLockAndUnlockSyscallFlock(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldTryLock, oldLock := linuxTryLockFile, linuxLockFile
	defer func() {
		linuxTryLockFile, linuxLockFile = oldTryLock, oldLock
	}()
	linuxTryLockFile, linuxLockFile = flockTryLockFile, flockLockFile
	TestLockAndUnlock(t)
}
