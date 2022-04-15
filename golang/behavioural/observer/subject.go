package main

type subject interface {
	register(observer)
	deregister(observer)
	notifyAll()
}
