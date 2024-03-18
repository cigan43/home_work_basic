package main

import (
	"fmt"
	"log"
	"sync"
)

// - Создайте группу из нескольких горутин, которые уывеличивают значения счетчика в переменной.
// - Каждый горутина должна выполнить свою задачу и сообщить об этом.
// - Используйте WaitGroup для ожидания завершения всех воркеров
// - Используйте Mutex для синхронизации доступа к общему ресурсу - переменной счетчика выполненных задач.
// - Напишите юнит тесты на реализованные функции;

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func (c *Counter) Value() int {
	return c.num
}

func main() {
	cnt := &Counter{
		num: 0,
	}

	finish := make(chan struct{})

	go Do(cnt, finish)

	select {
	case <-finish:
		log.Printf("Work done with count: %d", cnt.Value())
	}
}

func Do(cnt *Counter, finish chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc()
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}

	wg.Wait()
	finish <- struct{}{}
	close(finish)
}
