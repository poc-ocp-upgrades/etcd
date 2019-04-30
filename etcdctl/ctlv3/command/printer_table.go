package command

import (
	"os"
	v3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/snapshot"
	"github.com/olekukonko/tablewriter"
)

type tablePrinter struct{ printer }

func (tp *tablePrinter) MemberList(r v3.MemberListResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hdr, rows := makeMemberListTable(r)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hdr)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
func (tp *tablePrinter) EndpointHealth(r []epHealth) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hdr, rows := makeEndpointHealthTable(r)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hdr)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
func (tp *tablePrinter) EndpointStatus(r []epStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hdr, rows := makeEndpointStatusTable(r)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hdr)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
func (tp *tablePrinter) EndpointHashKV(r []epHashKV) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hdr, rows := makeEndpointHashKVTable(r)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hdr)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
func (tp *tablePrinter) DBStatus(r snapshot.Status) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hdr, rows := makeDBStatusTable(r)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hdr)
	for _, row := range rows {
		table.Append(row)
	}
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.Render()
}
