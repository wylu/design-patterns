package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
