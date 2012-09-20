package main

var Gateway chan string = make(chan string)

func main() {
	go startSocket()
	startHttp()
}
