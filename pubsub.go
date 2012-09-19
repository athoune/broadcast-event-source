package main

var clientCount = 0
var clients map[int]chan string = make(map[int]chan string)

type client struct {
	id      int
	channel chan string
}

func Suscribe() client {
	clientCount += 1
	ch := make(chan string)
	clients[clientCount] = ch
	cl := client{clientCount, ch}
	return cl
}

func (c *client) Leave() {
	delete(clients, c.id)
	close(c.channel)
}

func Publish(m string) {
	for i := range clients {
		clients[i] <- m
	}
}
