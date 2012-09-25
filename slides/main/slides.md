!SLIDE

# Blunderbluss

A golang example

![English flintlock blunderbluss](English_flintlock_blunderbuss.jpeg)

!SLIDE

# Objective

Developers want to read code, not manual, even with engraving on the cover.

This code is an example of useful async pattern in the real world.

> Http clients as subscribers (HTML5's event-source), telnet as publisher.


!SLIDE
# Go is authoritarian

Go compiles, uses static typing, reindents your code,
and goes on strike if you leave unused libraries.

See its as a level 0 unit testing.

Static typing is good for optimization, and with the magic `:=` syntax, it's painless.

Go use Capitalized words for public variables, and ask you for short variable name.

!SLIDE

Go is permissive, you can split your code in multiple files, there is no scope here.

Go loves your adminsys, application are just one fat binary,
and you can scp it to production, no runtime needed.

Go handles line ending, not like Erlang endings hell, but in its own way

```go
import (
	"fmt"
	"log"
	"net/http"
)
```

!SLIDE
# Go is predictable

But you must be explicit.

!SLIDE
## Close what you open

```go
channel := Subscribe()
defer channel.Leave()
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

!SLIDE
# Oo programation is dead

Nobody loves OOP, even javascript is cheating with it. Any smalltalk programers anywhere?

Use struct to groups attributes

```go
type subscriptions struct {
	count   int
	clients map[int]chan string
	sync.RWMutex //A struct, too
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
func (c *client) Leave() {//python "self" style
	subscriber.Lock()
	defer subscriber.Unlock()
	delete(subscriber.clients, c.id)
	close(c.channel)
}
```

!SLIDE
Heritate some methods

```go
type subscriptions struct {
	count   int
	clients map[int]chan string
	sync.RWMutex
}

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
# No reactor / actor patterns

You can't play with linear shuffled execution on one thread, like in nodejs.

There is no actor pattern (like in Erlang and Scala), just _goroutine_ (light thread).

Thank you for not using shared memory (global variable),
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
# Channel

Channel is a pipe, you pull something one side, and pull it in the other side.

Channel blocks on put (waits for a receiver), but there is also buffered channel,
with some room to fill before blocking.

```go
for {
	m := <-channel.channel
	msg := fmt.Sprintf("data: %s\n\n", m)
	w.Write([]byte(msg))
}
```

!SLIDE
# Memory is on your hands

List has fixed size, just like in Pascal, and you have to copy every elements when it's full.

Map is more cute, just add and remove stuff.

Absence of generic is painful, you have to copy/paste or hacking with go templates.
Or waiting for Java 5.

!SLIDE
There is reference and pointer in Go, but no tricks.

When you launch your application, you ask why there is so few cores
in this computer, and so many useless memory

!SLIDE
# Where go fits well?

Go is young, pioneer still exist. Your brain is large enough to handle most of it.

* Never use go to beat RoR or PHP, there is already too many HTML developer.
* Use go when you need Erlang but have no time.
* Use go when python is too slow and twisted/gevent/zeromq drive you mad.
* Use go when your node code need to run more than 3 months.
* Use go when java is needed but Eclipse is not installed. Be careful, nobody can beat java libraries.
* Don't use go when C++ is needed, because C++ developer only fall in love one time.

Go is low level, loves parralel task, doesn't have so many libraries.
Use it for network services or tasks.
