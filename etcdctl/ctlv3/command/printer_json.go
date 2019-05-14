package command

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonPrinter struct{ printer }

func newJSONPrinter() printer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &jsonPrinter{&printerRPC{newPrinterUnsupported("json"), printJSON}}
}
func (p *jsonPrinter) EndpointStatus(r []epStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	printJSON(r)
}
func (p *jsonPrinter) EndpointHashKV(r []epHashKV) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	printJSON(r)
}
func (p *jsonPrinter) DBStatus(r dbstatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	printJSON(r)
}
func printJSON(v interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Println(string(b))
}
