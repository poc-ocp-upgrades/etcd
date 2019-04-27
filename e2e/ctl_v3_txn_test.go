package e2e

import "testing"

func TestCtlV3TxnInteractiveSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, txnTestSuccess, withInteractive())
}
func TestCtlV3TxnInteractiveSuccessNoTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, txnTestSuccess, withInteractive(), withCfg(configNoTLS))
}
func TestCtlV3TxnInteractiveSuccessClientTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, txnTestSuccess, withInteractive(), withCfg(configClientTLS))
}
func TestCtlV3TxnInteractiveSuccessPeerTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, txnTestSuccess, withInteractive(), withCfg(configPeerTLS))
}
func TestCtlV3TxnInteractiveFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, txnTestFail, withInteractive())
}
func txnTestSuccess(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "key1", "value1", ""); err != nil {
		cx.t.Fatalf("txnTestSuccess ctlV3Put error (%v)", err)
	}
	if err := ctlV3Put(cx, "key2", "value2", ""); err != nil {
		cx.t.Fatalf("txnTestSuccess ctlV3Put error (%v)", err)
	}
	rqs := []txnRequests{{compare: []string{`value("key1") != "value2"`, `value("key2") != "value1"`}, ifSucess: []string{"get key1", "get key2"}, results: []string{"SUCCESS", "key1", "value1", "key2", "value2"}}, {compare: []string{`version("key1") = "1"`, `version("key2") = "1"`}, ifSucess: []string{"get key1", "get key2", `put "key \"with\" space" "value \x23"`}, ifFail: []string{`put key1 "fail"`, `put key2 "fail"`}, results: []string{"SUCCESS", "key1", "value1", "key2", "value2"}}, {compare: []string{`version("key \"with\" space") = "1"`}, ifSucess: []string{`get "key \"with\" space"`}, results: []string{"SUCCESS", `key "with" space`, "value \x23"}}}
	for _, rq := range rqs {
		if err := ctlV3Txn(cx, rq); err != nil {
			cx.t.Fatal(err)
		}
	}
}
func txnTestFail(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "key1", "value1", ""); err != nil {
		cx.t.Fatalf("txnTestSuccess ctlV3Put error (%v)", err)
	}
	rqs := []txnRequests{{compare: []string{`version("key") < "0"`}, ifSucess: []string{`put key "success"`}, ifFail: []string{`put key "fail"`}, results: []string{"FAILURE", "OK"}}, {compare: []string{`value("key1") != "value1"`}, ifSucess: []string{`put key1 "success"`}, ifFail: []string{`put key1 "fail"`}, results: []string{"FAILURE", "OK"}}}
	for _, rq := range rqs {
		if err := ctlV3Txn(cx, rq); err != nil {
			cx.t.Fatal(err)
		}
	}
}

type txnRequests struct {
	compare		[]string
	ifSucess	[]string
	ifFail		[]string
	results		[]string
}

func ctlV3Txn(cx ctlCtx, rqs txnRequests) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "txn")
	if cx.interactive {
		cmdArgs = append(cmdArgs, "--interactive")
	}
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return err
	}
	_, err = proc.Expect("compares:")
	if err != nil {
		return err
	}
	for _, req := range rqs.compare {
		if err = proc.Send(req + "\r"); err != nil {
			return err
		}
	}
	if err = proc.Send("\r"); err != nil {
		return err
	}
	_, err = proc.Expect("success requests (get, put, del):")
	if err != nil {
		return err
	}
	for _, req := range rqs.ifSucess {
		if err = proc.Send(req + "\r"); err != nil {
			return err
		}
	}
	if err = proc.Send("\r"); err != nil {
		return err
	}
	_, err = proc.Expect("failure requests (get, put, del):")
	if err != nil {
		return err
	}
	for _, req := range rqs.ifFail {
		if err = proc.Send(req + "\r"); err != nil {
			return err
		}
	}
	if err = proc.Send("\r"); err != nil {
		return err
	}
	for _, line := range rqs.results {
		_, err = proc.Expect(line)
		if err != nil {
			return err
		}
	}
	return proc.Close()
}
