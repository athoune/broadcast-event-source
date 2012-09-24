!SLIDE

# Blunderbluss, a golang example

![English flintlock blunderbluss](English_flintlock_blunderbuss.jpeg)

!SLIDE

# Objective

Developers want to read code, not manual, even with engraving on the cover.

This code is an example of useful async pattern in the real world.

Http clients as subscribers (HTML5's event-source), telnet as publisher.


!SLIDE
# Object oriented programation is dead

Nobody loves OOP, even javascript is sheating with it. Any smalltalk programers anywhere?

Use struct to groups attributes

```go
type subscriptions struct {
	count   int
	clients map[int]chan string
	sync.RWMutex
}
```

!SLIDE
Add a constructor, a factory indeed

```go
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

```go
func (c *client) Leave() {
	subscriber.Lock()
	defer subscriber.Unlock()
	delete(subscriber.clients, c.id)
	close(c.channel)
}
```

!SLIDE
Heritate some methods

```go
subscriber.Lock() // from sync.RWMutex
```

!SLIDE
> there is no object in go.

!SLIDE
# Avoiding callback hell

Go wrap async handler with callback.

```go
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

!SLIDE
# No reactor, nor actor patterns

You can't play with linear shuffled execution on one thread, like in nodejs.

No actor pattern, like in Erlang and Scala.

Just _goroutine_ (light thread), please, don't use shared memory (global variable),
_sync_ with lock or use message passing with _channel_.

```go
func Publish(m string) {
	subscriber.RLock()
	defer subscriber.RUnlock()
	for i := range subscriber.clients {
		subscriber.clients[i] <- m
	}
}
```

!SLIDE
# Memory is on your hands

List has fixed size, just like in Pascal, and you have to copy every elements when it's full.

Map is more cute, just add and remove stuff.

Absence of generic is painful, you have to copy/paste or hacking with go templates.
Or waiting for Java 5.

There is reference and pointer in Go, but no tricks.

!SLIDE
# Go knows what is good for you

Splitting your code in multiple files is your business, there is no scope here.

`go fmt` indent your code, don't try to argue.

`go build` go on strike when you don't clean your import.

static type and compilation validation are unit test for the lazy.

Everything is typed, but go can guess type with `:=`

Go application are just one fat binary, and you can scp it to production, no runtime needed.

`go` handles line ending, not like Erlang endings hell, but in its own way

```go
import (
	"fmt"
	"log"
	"net/http"
)
```

!SLIDE
# Go is predictable

## Close what you open

```go
channel := Subscribe()
defer func() {
	channel.Leave()
}()
```

!SLIDE
## Collect your trash

You can't throw it anymore.

```go
_, err := conn.Read(buf)
if err != nil {
	fmt.Println("error reading", err.Error())
	return
}
```

!SLIDE
## It's your choice to crash

```go
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
```

`recover` exists, but don't bother with it, multiple return is more clean.

