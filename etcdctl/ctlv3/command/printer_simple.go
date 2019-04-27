package command

import (
	"fmt"
	"strings"
	v3 "github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/types"
)

type simplePrinter struct {
	isHex		bool
	valueOnly	bool
}

func (s *simplePrinter) Del(resp v3.DeleteResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println(resp.Deleted)
	for _, kv := range resp.PrevKvs {
		printKV(s.isHex, s.valueOnly, kv)
	}
}
func (s *simplePrinter) Get(resp v3.GetResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, kv := range resp.Kvs {
		printKV(s.isHex, s.valueOnly, kv)
	}
}
func (s *simplePrinter) Put(r v3.PutResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println("OK")
	if r.PrevKv != nil {
		printKV(s.isHex, s.valueOnly, r.PrevKv)
	}
}
func (s *simplePrinter) Txn(resp v3.TxnResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp.Succeeded {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAILURE")
	}
	for _, r := range resp.Responses {
		fmt.Println("")
		switch v := r.Response.(type) {
		case *pb.ResponseOp_ResponseDeleteRange:
			s.Del((v3.DeleteResponse)(*v.ResponseDeleteRange))
		case *pb.ResponseOp_ResponsePut:
			s.Put((v3.PutResponse)(*v.ResponsePut))
		case *pb.ResponseOp_ResponseRange:
			s.Get(((v3.GetResponse)(*v.ResponseRange)))
		default:
			fmt.Printf("unexpected response %+v\n", r)
		}
	}
}
func (s *simplePrinter) Watch(resp v3.WatchResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, e := range resp.Events {
		fmt.Println(e.Type)
		if e.PrevKv != nil {
			printKV(s.isHex, s.valueOnly, e.PrevKv)
		}
		printKV(s.isHex, s.valueOnly, e.Kv)
	}
}
func (s *simplePrinter) Grant(resp v3.LeaseGrantResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("lease %016x granted with TTL(%ds)\n", resp.ID, resp.TTL)
}
func (p *simplePrinter) Revoke(id v3.LeaseID, r v3.LeaseRevokeResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("lease %016x revoked\n", id)
}
func (p *simplePrinter) KeepAlive(resp v3.LeaseKeepAliveResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("lease %016x keepalived with TTL(%d)\n", resp.ID, resp.TTL)
}
func (s *simplePrinter) TimeToLive(resp v3.LeaseTimeToLiveResponse, keys bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp.GrantedTTL == 0 && resp.TTL == -1 {
		fmt.Printf("lease %016x already expired\n", resp.ID)
		return
	}
	txt := fmt.Sprintf("lease %016x granted with TTL(%ds), remaining(%ds)", resp.ID, resp.GrantedTTL, resp.TTL)
	if keys {
		ks := make([]string, len(resp.Keys))
		for i := range resp.Keys {
			ks[i] = string(resp.Keys[i])
		}
		txt += fmt.Sprintf(", attached keys(%v)", ks)
	}
	fmt.Println(txt)
}
func (s *simplePrinter) Leases(resp v3.LeaseLeasesResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("found %d leases\n", len(resp.Leases))
	for _, item := range resp.Leases {
		fmt.Printf("%016x\n", item.ID)
	}
}
func (s *simplePrinter) Alarm(resp v3.AlarmResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, e := range resp.Alarms {
		fmt.Printf("%+v\n", e)
	}
}
func (s *simplePrinter) MemberAdd(r v3.MemberAddResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Member %16x added to cluster %16x\n", r.Member.ID, r.Header.ClusterId)
}
func (s *simplePrinter) MemberRemove(id uint64, r v3.MemberRemoveResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Member %16x removed from cluster %16x\n", id, r.Header.ClusterId)
}
func (s *simplePrinter) MemberUpdate(id uint64, r v3.MemberUpdateResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Member %16x updated in cluster %16x\n", id, r.Header.ClusterId)
}
func (s *simplePrinter) MemberList(resp v3.MemberListResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, rows := makeMemberListTable(resp)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}
func (s *simplePrinter) EndpointStatus(statusList []epStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, rows := makeEndpointStatusTable(statusList)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}
func (s *simplePrinter) EndpointHashKV(hashList []epHashKV) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, rows := makeEndpointHashKVTable(hashList)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}
func (s *simplePrinter) DBStatus(ds dbstatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, rows := makeDBStatusTable(ds)
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}
func (s *simplePrinter) MoveLeader(leader, target uint64, r v3.MoveLeaderResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Leadership transferred from %s to %s\n", types.ID(leader), types.ID(target))
}
func (s *simplePrinter) RoleAdd(role string, r v3.AuthRoleAddResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s created\n", role)
}
func (s *simplePrinter) RoleGet(role string, r v3.AuthRoleGetResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s\n", role)
	fmt.Println("KV Read:")
	printRange := func(perm *v3.Permission) {
		sKey := string(perm.Key)
		sRangeEnd := string(perm.RangeEnd)
		if strings.Compare(sRangeEnd, "\x00") != 0 {
			fmt.Printf("\t[%s, %s)", sKey, sRangeEnd)
		} else {
			fmt.Printf("\t[%s, <open ended>", sKey)
		}
		if strings.Compare(v3.GetPrefixRangeEnd(sKey), sRangeEnd) == 0 {
			fmt.Printf(" (prefix %s)", sKey)
		}
		fmt.Printf("\n")
	}
	for _, perm := range r.Perm {
		if perm.PermType == v3.PermRead || perm.PermType == v3.PermReadWrite {
			if len(perm.RangeEnd) == 0 {
				fmt.Printf("\t%s\n", string(perm.Key))
			} else {
				printRange((*v3.Permission)(perm))
			}
		}
	}
	fmt.Println("KV Write:")
	for _, perm := range r.Perm {
		if perm.PermType == v3.PermWrite || perm.PermType == v3.PermReadWrite {
			if len(perm.RangeEnd) == 0 {
				fmt.Printf("\t%s\n", string(perm.Key))
			} else {
				printRange((*v3.Permission)(perm))
			}
		}
	}
}
func (s *simplePrinter) RoleList(r v3.AuthRoleListResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, role := range r.Roles {
		fmt.Printf("%s\n", role)
	}
}
func (s *simplePrinter) RoleDelete(role string, r v3.AuthRoleDeleteResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s deleted\n", role)
}
func (s *simplePrinter) RoleGrantPermission(role string, r v3.AuthRoleGrantPermissionResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s updated\n", role)
}
func (s *simplePrinter) RoleRevokePermission(role string, key string, end string, r v3.AuthRoleRevokePermissionResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(end) == 0 {
		fmt.Printf("Permission of key %s is revoked from role %s\n", key, role)
		return
	}
	if strings.Compare(end, "\x00") != 0 {
		fmt.Printf("Permission of range [%s, %s) is revoked from role %s\n", key, end, role)
	} else {
		fmt.Printf("Permission of range [%s, <open ended> is revoked from role %s\n", key, role)
	}
}
func (s *simplePrinter) UserAdd(name string, r v3.AuthUserAddResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("User %s created\n", name)
}
func (s *simplePrinter) UserGet(name string, r v3.AuthUserGetResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("User: %s\n", name)
	fmt.Printf("Roles:")
	for _, role := range r.Roles {
		fmt.Printf(" %s", role)
	}
	fmt.Printf("\n")
}
func (s *simplePrinter) UserChangePassword(v3.AuthUserChangePasswordResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println("Password updated")
}
func (s *simplePrinter) UserGrantRole(user string, role string, r v3.AuthUserGrantRoleResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s is granted to user %s\n", role, user)
}
func (s *simplePrinter) UserRevokeRole(user string, role string, r v3.AuthUserRevokeRoleResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Role %s is revoked from user %s\n", role, user)
}
func (s *simplePrinter) UserDelete(user string, r v3.AuthUserDeleteResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("User %s deleted\n", user)
}
func (s *simplePrinter) UserList(r v3.AuthUserListResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, user := range r.Users {
		fmt.Printf("%s\n", user)
	}
}
