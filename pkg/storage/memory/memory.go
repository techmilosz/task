package memory

import "sync"

type Memory[T comparable] struct {
	values map[T]struct{}
	mx     sync.Mutex
}

func New[T comparable]() *Memory[T] {
	return &Memory[T]{
		values: map[T]struct{}{},
		mx:     sync.Mutex{},
	}
}

func (m *Memory[T]) Add(val T) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.values[val] = struct{}{}
}

func (m *Memory[T]) Remove(val T) {
	m.mx.Lock()
	defer m.mx.Unlock()
	delete(m.values, val)
}

func (m *Memory[T]) GetAll() []T {
	m.mx.Lock()
	defer m.mx.Unlock()

	res := make([]T, 0, len(m.values))

	for val := range m.values {
		res = append(res, val)
	}

	return res
}
