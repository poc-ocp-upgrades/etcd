package auth

import (
	"go.etcd.io/etcd/auth/authpb"
	"go.etcd.io/etcd/mvcc/backend"
	"go.etcd.io/etcd/pkg/adt"
	"go.uber.org/zap"
)

func getMergedPerms(lg *zap.Logger, tx backend.BatchTx, userName string) *unifiedRangePermissions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := getUser(lg, tx, userName)
	if user == nil {
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
func checkKeyInterval(lg *zap.Logger, cachedPerms *unifiedRangePermissions, key, rangeEnd []byte, permtyp authpb.Permission_Type) bool {
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
		if lg != nil {
			lg.Panic("unknown auth type", zap.String("auth-type", permtyp.String()))
		} else {
			plog.Panicf("unknown auth type: %v", permtyp)
		}
	}
	return false
}
func checkKeyPoint(lg *zap.Logger, cachedPerms *unifiedRangePermissions, key []byte, permtyp authpb.Permission_Type) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pt := adt.NewBytesAffinePoint(key)
	switch permtyp {
	case authpb.READ:
		return cachedPerms.readPerms.Intersects(pt)
	case authpb.WRITE:
		return cachedPerms.writePerms.Intersects(pt)
	default:
		if lg != nil {
			lg.Panic("unknown auth type", zap.String("auth-type", permtyp.String()))
		} else {
			plog.Panicf("unknown auth type: %v", permtyp)
		}
	}
	return false
}
func (as *authStore) isRangeOpPermitted(tx backend.BatchTx, userName string, key, rangeEnd []byte, permtyp authpb.Permission_Type) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := as.rangePermCache[userName]
	if !ok {
		perms := getMergedPerms(as.lg, tx, userName)
		if perms == nil {
			if as.lg != nil {
				as.lg.Warn("failed to create a merged permission", zap.String("user-name", userName))
			} else {
				plog.Errorf("failed to create a unified permission of user %s", userName)
			}
			return false
		}
		as.rangePermCache[userName] = perms
	}
	if len(rangeEnd) == 0 {
		return checkKeyPoint(as.lg, as.rangePermCache[userName], key, permtyp)
	}
	return checkKeyInterval(as.lg, as.rangePermCache[userName], key, rangeEnd, permtyp)
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
