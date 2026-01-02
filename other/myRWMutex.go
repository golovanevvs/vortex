package main

import (
	"sync"
)

type MyRWMutex struct {
	mu             sync.Mutex
	cond           *sync.Cond // условная переменная для сигнализации
	readers        int        // кол-во активных читателей
	writer         bool       // true, если писатель активен
	waitingWriters int        // число писателей, ожидающих входа
}

func NewMyRWMutex() *MyRWMutex {
	m := &MyRWMutex{}
	m.cond = sync.NewCond(&m.mu)
	return m
}

// RLock захватывает блокировку на чтение
func (rw *MyRWMutex) RLock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// Если уже есть писатель или ждут писатели — новые читатели НЕ пускаются!
	// Это предотвращает голодание писателя.
	for rw.writer || rw.waitingWriters > 0 {
		rw.cond.Wait()
	}

	rw.readers++
}

// RUnlock освобождает блокировку чтения
func (rw *MyRWMutex) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.readers--
	if rw.readers < 0 {
		panic("MyRWMutex: RUnlock without RLock")
	}

	// Если читателей не осталось — можно пускать писателя
	if rw.readers == 0 {
		rw.cond.Signal() // или Broadcast, но Signal достаточно
	}
}

// Lock захватывает блокировку на запись
func (rw *MyRWMutex) Lock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// Увеличиваем счётчик ожидающих писателей
	rw.waitingWriters++

	// Ждём, пока не освободится: ни читателей, ни писателя
	for rw.readers > 0 || rw.writer {
		rw.cond.Wait()
	}

	// Теперь мы — писатель
	rw.waitingWriters--
	rw.writer = true
}

// Unlock освобождает блокировку записи
func (rw *MyRWMutex) Unlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	if !rw.writer {
		panic("MyRWMutex: Unlock without Lock")
	}

	rw.writer = false

	// Разблокируем всех: и читателей, и писателей
	// Но по нашей логике — читатели не идут, если waitingWriters > 0,
	// поэтому безопасно использовать Broadcast
	rw.cond.Broadcast()
}
