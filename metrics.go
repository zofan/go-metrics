package metrics

import "sync"

type Metrics struct {
	metrics map[string]uint64

	mu sync.RWMutex
}

func New() *Metrics {
	return &Metrics{
		metrics: make(map[string]uint64),
	}
}

func (m *Metrics) Inc(name string) {
	m.mu.Lock()
	m.metrics[name] += 1
	m.mu.Unlock()
}

func (m *Metrics) IncN(name string, delta uint64) {
	m.mu.Lock()
	m.metrics[name] += delta
	m.mu.Unlock()
}

func (m *Metrics) Get(name string) uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.metrics[name]
}

func (m *Metrics) GetAll(name string) map[string]uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.metrics
}

func (m *Metrics) Reset() {
	m.mu.Lock()
	m.metrics = make(map[string]uint64)
	m.mu.Unlock()
}
