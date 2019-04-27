package v3rpc

import "github.com/gogo/protobuf/proto"

type codec struct{}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := proto.Marshal(v.(proto.Message))
	sentBytes.Add(float64(len(b)))
	return b, err
}
func (c *codec) Unmarshal(data []byte, v interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	receivedBytes.Add(float64(len(data)))
	return proto.Unmarshal(data, v.(proto.Message))
}
func (c *codec) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "proto"
}
