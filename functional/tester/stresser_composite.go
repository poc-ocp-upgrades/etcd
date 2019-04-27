package tester

import "sync"

type compositeStresser struct{ stressers []Stresser }

func (cs *compositeStresser) Stress() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, s := range cs.stressers {
		if err := s.Stress(); err != nil {
			for j := 0; j < i; j++ {
				cs.stressers[j].Close()
			}
			return err
		}
	}
	return nil
}
func (cs *compositeStresser) Pause() (ems map[string]int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var emu sync.Mutex
	ems = make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(len(cs.stressers))
	for i := range cs.stressers {
		go func(s Stresser) {
			defer wg.Done()
			errs := s.Pause()
			for k, v := range errs {
				emu.Lock()
				ems[k] += v
				emu.Unlock()
			}
		}(cs.stressers[i])
	}
	wg.Wait()
	return ems
}
func (cs *compositeStresser) Close() (ems map[string]int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var emu sync.Mutex
	ems = make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(len(cs.stressers))
	for i := range cs.stressers {
		go func(s Stresser) {
			defer wg.Done()
			errs := s.Close()
			for k, v := range errs {
				emu.Lock()
				ems[k] += v
				emu.Unlock()
			}
		}(cs.stressers[i])
	}
	wg.Wait()
	return ems
}
func (cs *compositeStresser) ModifiedKeys() (modifiedKey int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, stress := range cs.stressers {
		modifiedKey += stress.ModifiedKeys()
	}
	return modifiedKey
}
