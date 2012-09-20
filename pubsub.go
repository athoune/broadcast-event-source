package main

/* Stolen from https://github.com/bitly/go-notify/blob/master/notify.go */
import (
	"sync"
)

var subscriber = &subscriptions{
	count:   0,
	clients: make(map[int]chan string),
}

type subscriptions struct {
	count   int
	clients map[int]chan string
	sync.RWMutex
}

type client struct {
	id      int
	channel chan string
}

func Subscribe() client {
	subscriber.Lock()
	defer subscriber.Unlock()
	subscriber.count += 1
	ch := make(chan string)
	subscriber.clients[subscriber.count] = ch
	cl := client{subscriber.count, ch}
	return cl
}

func (c *client) Leave() {
	subscriber.Lock()
	defer subscriber.Unlock()
	delete(subscriber.clients, c.id)
	close(c.channel)
}

func Publish(m string) {
	subscriber.RLock()
	defer subscriber.RUnlock()
	for i := range subscriber.clients {
		subscriber.clients[i] <- m
	}
}
