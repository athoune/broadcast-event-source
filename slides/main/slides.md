!SLIDE

# Blunderbluss, a golang example

![English flintlock blunderbluss](English_flintlock_blunderbuss.jpeg)

!SLIDE

# Objective

Developers want to read code, not F** manual.

This code is an example of useful async pattern in the real world.

Http clients as subscribers (HTML5's event-source), telnet as publisher.


!SLIDE
# Object oriented programation is dead

Nobody loves OOP, even javascript is sheating with it. Any smalltalk programers anywhere?

Use struct to groups attributes

```
type subscriptions struct {
	count   int
	clients map[int]chan string
	sync.RWMutex
}
```

!SLIDE
Add a constructor

```
func Subscribe() client {
	subscriber.Lock()
	defer subscriber.Unlock()
	subscriber.count += 1
	ch := make(chan string)
	subscriber.clients[subscriber.count] = ch
	cl := client{subscriber.count, ch}
	return cl
}
```

!SLIDE
Some methods

```
func (c *client) Leave() {
	subscriber.Lock()
	defer subscriber.Unlock()
	delete(subscriber.clients, c.id)
	close(c.channel)
}
```

!SLIDE
Heritate some methods

```
subscriber.Lock() // from sync.RWMutex
```

!SLIDE
_there is no object in go_

# Avoiding callback hell

Go wrap async handler with callback.

```
func doServeStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		Publish(strings.TrimSpace(string(buf)))
		conn.Write([]byte("ok\n"))
	}
}

listener, err := net.Listen("tcp", "localhost:5000")
for {
	conn, err := listener.Accept()
	go doServeStuff(conn)
}
```
