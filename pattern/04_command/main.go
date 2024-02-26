package main

import "ositlar.com/pattern/04_command/tv"

func main() {
	tv1 := tv.NewTv()

	onCommand := tv.NewOnCommand(tv1)
	offCommand := tv.NewOffCommand(tv1)

	onButton := tv.NewButton(onCommand)
	offButton := tv.NewButton(offCommand)

	onButton.Press()
	offButton.Press()
}
