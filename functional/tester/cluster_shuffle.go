package tester

import (
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func (clus *Cluster) shuffleCases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(1000)
	n := len(clus.cases)
	cp := coprime(n)
	css := make([]Case, n)
	for i := 0; i < n; i++ {
		css[i] = clus.cases[(cp*i+offset)%n]
	}
	clus.cases = css
	clus.lg.Info("shuffled test failure cases", zap.Int("total", n))
}
func coprime(n int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	coprime := 1
	for i := n / 2; i < n; i++ {
		if gcd(i, n) == 1 {
			coprime = i
			break
		}
	}
	return coprime
}
func gcd(x, y int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
