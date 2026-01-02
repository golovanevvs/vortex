package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// entry — значение в map
type entry struct {
	p unsafe.Pointer // *T, nil, или указывает на expunged
}

var expunged = unsafe.Pointer(new(interface{})) // маркер "удалено"

// Load возвращает значение, если оно не удалено
func (e *entry) load() (value interface{}, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*interface{})(p), true
}

// MySyncMap — упрощённый аналог sync.Map
type MySyncMap struct {
	mu    sync.Mutex
	read  atomic.Pointer[map[interface{}]*entry] // атомарный указатель на read-map
	dirty map[interface{}]*entry                // мутабельная часть
	misses int                                  // счётчик промахов по read
}

func (m *MySyncMap) Load(key interface{}) (value interface{}, ok bool) {
	read := m.read.Load()
	if read != nil {
		if e, ok := (*read)[key]; ok {
			return e.load()
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Может, появилось в dirty?
	if m.dirty != nil {
		if e, ok := m.dirty[key]; ok {
			// Мигрируем в read (опционально)
			m.missLocked()
			return e.load()
		}
	}
	return nil, false
}

func (m *MySyncMap) Store(key, value interface{}) {
	read := m.read.Load()
	if read != nil {
		if e, ok := (*read)[key]; ok && e.tryStore(&value) {
			return
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.dirty != nil {
		if e, ok := m.dirty[key]; ok {
			e.storeLocked(&value)
			return
		}
		// Новый ключ — кладём в dirty
		m.dirty[key] = &entry{p: unsafe.Pointer(&value)}
	} else {
		// Первый раз пишем — создаём dirty из read
		m.dirty = make(map[interface{}]*entry)
		read := m.read.Load()
		if read != nil {
			for k, e := range *read {
				m.dirty[k] = e
			}
		}
		m.dirty[key] = &entry{p: unsafe.Pointer(&value)}
	}
}

// tryStore — попытка атомарно обновить значение
func (e *entry) tryStore(i *interface{}) bool {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

func (e *entry) storeLocked(i *interface{}) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

// missLocked — вызывается при промахе в read
func (m *MySyncMap) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	// Промахов много — мигрируем dirty в read
	read := make(map[interface{}]*entry, len(m.dirty))
	for k, e := range m.dirty {
		read[k] = e
	}
	m.read.Store(&read)
	m.dirty = nil
	m.misses = 0
}

// Delete — помечает как expunged
func (m *MySyncMap) Delete(key interface{}) {
	read := m.read.Load()
	if read != nil {
		if e, ok := (*read)[key]; ok {
			e.delete()
			return
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.dirty != nil {
		if e, ok := m.dirty[key]; ok {
			delete(m.dirty, key)
			e.delete()
		}
	}
}

func (e *entry) delete() {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return
		}
	}
}