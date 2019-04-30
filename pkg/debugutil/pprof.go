package debugutil

import (
	"net/http"
	"net/http/pprof"
	"runtime"
)

const HTTPPrefixPProf = "/debug/pprof"

func PProfHandlers() map[string]http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if runtime.SetMutexProfileFraction(-1) == 0 {
		runtime.SetMutexProfileFraction(5)
	}
	m := make(map[string]http.Handler)
	m[HTTPPrefixPProf+"/"] = http.HandlerFunc(pprof.Index)
	m[HTTPPrefixPProf+"/profile"] = http.HandlerFunc(pprof.Profile)
	m[HTTPPrefixPProf+"/symbol"] = http.HandlerFunc(pprof.Symbol)
	m[HTTPPrefixPProf+"/cmdline"] = http.HandlerFunc(pprof.Cmdline)
	m[HTTPPrefixPProf+"/trace "] = http.HandlerFunc(pprof.Trace)
	m[HTTPPrefixPProf+"/heap"] = pprof.Handler("heap")
	m[HTTPPrefixPProf+"/goroutine"] = pprof.Handler("goroutine")
	m[HTTPPrefixPProf+"/threadcreate"] = pprof.Handler("threadcreate")
	m[HTTPPrefixPProf+"/block"] = pprof.Handler("block")
	m[HTTPPrefixPProf+"/mutex"] = pprof.Handler("mutex")
	return m
}
