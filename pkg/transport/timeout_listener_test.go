package transport

import (
	"net"
	"testing"
	"time"
)

func TestNewTimeoutListener(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l, err := NewTimeoutListener("127.0.0.1:0", "http", nil, time.Hour, time.Hour)
	if err != nil {
		t.Fatalf("unexpected NewTimeoutListener error: %v", err)
	}
	defer l.Close()
	tln := l.(*rwTimeoutListener)
	if tln.rdtimeoutd != time.Hour {
		t.Errorf("read timeout = %s, want %s", tln.rdtimeoutd, time.Hour)
	}
	if tln.wtimeoutd != time.Hour {
		t.Errorf("write timeout = %s, want %s", tln.wtimeoutd, time.Hour)
	}
}
func TestWriteReadTimeoutListener(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("unexpected listen error: %v", err)
	}
	wln := rwTimeoutListener{Listener: ln, wtimeoutd: 10 * time.Millisecond, rdtimeoutd: 10 * time.Millisecond}
	stop := make(chan struct{})
	blocker := func() {
		conn, derr := net.Dial("tcp", ln.Addr().String())
		if derr != nil {
			t.Fatalf("unexpected dail error: %v", derr)
		}
		defer conn.Close()
		<-stop
	}
	go blocker()
	conn, err := wln.Accept()
	if err != nil {
		t.Fatalf("unexpected accept error: %v", err)
	}
	defer conn.Close()
	data := make([]byte, 5*1024*1024)
	done := make(chan struct{})
	go func() {
		_, err = conn.Write(data)
		done <- struct{}{}
	}()
	select {
	case <-done:
	case <-time.After(wln.wtimeoutd*10 + time.Second):
		t.Fatal("wait timeout")
	}
	if operr, ok := err.(*net.OpError); !ok || operr.Op != "write" || !operr.Timeout() {
		t.Errorf("err = %v, want write i/o timeout error", err)
	}
	stop <- struct{}{}
	go blocker()
	conn, err = wln.Accept()
	if err != nil {
		t.Fatalf("unexpected accept error: %v", err)
	}
	buf := make([]byte, 10)
	go func() {
		_, err = conn.Read(buf)
		done <- struct{}{}
	}()
	select {
	case <-done:
	case <-time.After(wln.rdtimeoutd * 10):
		t.Fatal("wait timeout")
	}
	if operr, ok := err.(*net.OpError); !ok || operr.Op != "read" || !operr.Timeout() {
		t.Errorf("err = %v, want write i/o timeout error", err)
	}
	stop <- struct{}{}
}
