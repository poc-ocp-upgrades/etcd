package auth

import (
	"github.com/coreos/etcd/auth/authpb"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/adt"
)

func getMergedPerms(tx backend.BatchTx, userName string) *unifiedRangePermissions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := getUser(tx, userName)
	if user == nil {
		plog.Errorf("invalid user name %s", userName)
		return nil
	}
	readPerms := &adt.IntervalTree{}
	writePerms := &adt.IntervalTree{}
	for _, roleName := range user.Roles {
		role := getRole(tx, roleName)
		if role == nil {
			continue
		}
		for _, perm := range role.KeyPermission {
			var ivl adt.Interval
			var rangeEnd []byte
			if len(perm.RangeEnd) != 1 || perm.RangeEnd[0] != 0 {
				rangeEnd = perm.RangeEnd
			}
			if len(perm.RangeEnd) != 0 {
				ivl = adt.NewBytesAffineInterval(perm.Key, rangeEnd)
			} else {
				ivl = adt.NewBytesAffinePoint(perm.Key)
			}
			switch perm.PermType {
			case authpb.READWRITE:
				readPerms.Insert(ivl, struct{}{})
				writePerms.Insert(ivl, struct{}{})
			case authpb.READ:
				readPerms.Insert(ivl, struct{}{})
			case authpb.WRITE:
				writePerms.Insert(ivl, struct{}{})
			}
		}
	}
	return &unifiedRangePermissions{readPerms: readPerms, writePerms: writePerms}
}
func checkKeyInterval(cachedPerms *unifiedRangePermissions, key, rangeEnd []byte, permtyp authpb.Permission_Type) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(rangeEnd) == 1 && rangeEnd[0] == 0 {
		rangeEnd = nil
	}
	ivl := adt.NewBytesAffineInterval(key, rangeEnd)
	switch permtyp {
	case authpb.READ:
		return cachedPerms.readPerms.Contains(ivl)
	case authpb.WRITE:
		return cachedPerms.writePerms.Contains(ivl)
	default:
		plog.Panicf("unknown auth type: %v", permtyp)
	}
	return false
}
func checkKeyPoint(cachedPerms *unifiedRangePermissions, key []byte, permtyp authpb.Permission_Type) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pt := adt.NewBytesAffinePoint(key)
	switch permtyp {
	case authpb.READ:
		return cachedPerms.readPerms.Intersects(pt)
	case authpb.WRITE:
		return cachedPerms.writePerms.Intersects(pt)
	default:
		plog.Panicf("unknown auth type: %v", permtyp)
	}
	return false
}
func (as *authStore) isRangeOpPermitted(tx backend.BatchTx, userName string, key, rangeEnd []byte, permtyp authpb.Permission_Type) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := as.rangePermCache[userName]
	if !ok {
		perms := getMergedPerms(tx, userName)
		if perms == nil {
			plog.Errorf("failed to create a unified permission of user %s", userName)
			return false
		}
		as.rangePermCache[userName] = perms
	}
	if len(rangeEnd) == 0 {
		return checkKeyPoint(as.rangePermCache[userName], key, permtyp)
	}
	return checkKeyInterval(as.rangePermCache[userName], key, rangeEnd, permtyp)
}
func (as *authStore) clearCachedPerm() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	as.rangePermCache = make(map[string]*unifiedRangePermissions)
}
func (as *authStore) invalidateCachedPerm(userName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(as.rangePermCache, userName)
}

type unifiedRangePermissions struct {
	readPerms	*adt.IntervalTree
	writePerms	*adt.IntervalTree
}
