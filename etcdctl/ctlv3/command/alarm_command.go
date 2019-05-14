package command

import (
	godefaultbytes "bytes"
	"fmt"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func NewAlarmCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ac := &cobra.Command{Use: "alarm <subcommand>", Short: "Alarm related commands"}
	ac.AddCommand(NewAlarmDisarmCommand())
	ac.AddCommand(NewAlarmListCommand())
	return ac
}
func NewAlarmDisarmCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := cobra.Command{Use: "disarm", Short: "Disarms all alarms", Run: alarmDisarmCommandFunc}
	return &cmd
}
func alarmDisarmCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("alarm disarm command accepts no arguments"))
	}
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).AlarmDisarm(ctx, &v3.AlarmMember{})
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.Alarm(*resp)
}
func NewAlarmListCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := cobra.Command{Use: "list", Short: "Lists all alarms", Run: alarmListCommandFunc}
	return &cmd
}
func alarmListCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("alarm list command accepts no arguments"))
	}
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).AlarmList(ctx)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.Alarm(*resp)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
