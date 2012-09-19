package main
/* Stolen from https://github.com/bitly/go-notify/blob/master/notify.go */
import (
	"sync"
)

var clientCount = 0
var clientsLock sync.RWMutex
var clients map[int]chan string = make(map[int]chan string)

type client struct {
	id      int
	channel chan string
}

func Suscribe() client {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	clientCount += 1
	ch := make(chan string)
	clients[clientCount] = ch
	cl := client{clientCount, ch}
	return cl
}

func (c *client) Leave() {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	delete(clients, c.id)
	close(c.channel)
}

func Publish(m string) {
	clientsLock.RLock()
	defer clientsLock.RUnlock()
	for i := range clients {
		clients[i] <- m
	}
}
