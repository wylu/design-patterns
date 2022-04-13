package main

type onCommond struct {
	device device
}

func (o *onCommond) execute() {
	o.device.on()
}
