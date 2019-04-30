package clientv3

import "context"

func readyWait(rpcCtx, clientCtx context.Context, ready <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-ready:
		return nil
	case <-rpcCtx.Done():
		return rpcCtx.Err()
	case <-clientCtx.Done():
		return clientCtx.Err()
	}
}
