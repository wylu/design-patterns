package main

func main() {
	tv := &tv{}
	onCommond := &onCommond{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommond}
	onButton.press()

	offButton := &button{command: offCommand}
	offButton.press()
}
