package main

import (
	"sync"
)

// MyCond — упрощённая реализация sync.Cond
type MyCond struct {
	L Locker // пользовательский мьютекс (должен быть указан при создании)

	waitersMtx sync.Mutex      // защищает waiters
	waiters    []chan struct{} // список каналов ожидающих горутин
}

// Locker — интерфейс, совместимый с sync.Locker
type Locker interface {
	Lock()
	Unlock()
}

// NewMyCond создаёт новый MyCond, привязанный к мьютексу l
func NewMyCond(l Locker) *MyCond {
	return &MyCond{L: l}
}

// Wait блокирует вызывающую горутину до тех пор, пока другая горутина
// не вызовет Signal() или Broadcast() на том же MyCond.
// При входе L должен быть захвачен; Wait разблокирует его и блокирует снова перед возвратом.
func (c *MyCond) Wait() {
	// 1. Создаём канал для этой горутины
	ch := make(chan struct{})

	// 2. Регистрируем канал в списке ожидающих (под защитой waitersMtx)
	c.waitersMtx.Lock()
	c.waiters = append(c.waiters, ch)
	c.waitersMtx.Unlock()

	// 3. Разблокируем пользовательский мьютекс
	c.L.Unlock()

	// 4. Ждём сигнала
	<-ch

	// 5. Перехватываем мьютекс обратно перед выходом
	c.L.Lock()
}

// Signal разбуждает одну из ожидающих горутин (если есть)
func (c *MyCond) Signal() {
	c.waitersMtx.Lock()
	defer c.waitersMtx.Unlock()

	if len(c.waiters) == 0 {
		return
	}

	// Берём первый канал (FIFO — можно и последний, но FIFO честнее)
	ch := c.waiters[0]
	c.waiters = c.waiters[1:]

	// Посылаем сигнал — закрываем канал (безопасно, т.к. только один получатель)
	close(ch)
}

// Broadcast разбуждает все ожидающие горутины
func (c *MyCond) Broadcast() {
	c.waitersMtx.Lock()
	defer c.waitersMtx.Unlock()

	// Закрываем все каналы
	for _, ch := range c.waiters {
		close(ch)
	}
	c.waiters = nil // очищаем список
}