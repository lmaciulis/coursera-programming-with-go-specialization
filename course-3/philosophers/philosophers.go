package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philCount = 5
	eatCount  = 3
)

type cs struct {
	sync.Mutex
}

type phil struct {
	id              int
	leftCs, rightCs *cs
}

var wg sync.WaitGroup
var css = make([]*cs, philCount)
var phils = make([]*phil, philCount)
var ch = make(chan int, 2)

func init() {
	for i := 0; i < philCount; i++ {
		css[i] = new(cs)
	}

	for i := 0; i < philCount; i++ {
		phils[i] = &phil{
			id:      i + 1,
			leftCs:  css[i],
			rightCs: css[(i+1)%philCount],
		}
	}
}

func main() {
	go host(ch)

	for _, phil := range phils {
		wg.Add(1)
		go phil.eat()
	}

	wg.Wait()
}

func host(ch chan int) {
	count := eatCount * philCount

	wg.Add(1)
	for i := 0; i < count; i++ {
		ch <- 0
		ch <- 0

		<-ch
		<-ch
	}
	wg.Done()
}

func (p phil) eat() {
	for j := 0; j < eatCount; j++ {
		<-ch // wait for permission

		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Millisecond * 100) // in order to mix a little bit order
		fmt.Printf("finishing eating %d\n", p.id)

		ch <- 0 // free the permission

		p.leftCs.Unlock()
		p.rightCs.Unlock()
	}

	wg.Done()
}
