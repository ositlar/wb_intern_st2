package tv

import "ositlar.com/pattern/04_command/interfaces"

type offCommand struct {
	devide interfaces.Device
}

func NewOffCommand(device interfaces.Device) *offCommand {
	return &offCommand{
		devide: device,
	}
}

func (c *offCommand) Execute() {
	c.devide.Off()
}
