package event

import "sync"

type EventFunc func(Eventer)

var lock = sync.Mutex{}
var subs = []EventFunc{}

func Register(fn EventFunc) {
	lock.Lock()
	subs = append(subs, fn)
	lock.Unlock()
}

func Emmit(data Eventer) {
	wg := sync.WaitGroup{}

	for _, cb := range subs {
		wg.Add(1)
		go func(fn EventFunc, wg *sync.WaitGroup) {
			wg.Done()
			fn(data)
		}(cb, &wg)
	}

	wg.Wait()

}
