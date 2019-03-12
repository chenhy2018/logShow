package main

import (
	"fmt"
	// "io"
)

//灯类
type Light struct {
}

func (light Light)on() {
	fmt.Println("light on.")
}

func (light Light)off() {
	fmt.Println("light off")
}

//命令接口
type ICommand interface {
	Execute()
	Display()
}

type Command struct {
	ICommand
}
//开灯命令类
type LightOnCommand struct {
	light *Light
	Command
}

func NewLightOnCommand(clight *Light) *LightOnCommand {
	current := &LightOnCommand{
		light: clight,
	}
	current.Command.ICommand = current.Command
	current.ICommand = current
	return current
}

func (lightcmd LightOnCommand)Execute() {

	// if (lightcmd.light == nil) {
	// 	fmt.Println("Execute, Light is not init.")
	// 	return
	// }
	lightcmd.light.on()
}

func (lightcmd LightOnCommand)Display() {
	fmt.Println("LightOnCommand Display")
}
//遥控器类
type SimpleRemoteControl struct {
	Command *Command
}

func (current *SimpleRemoteControl)setCommand(command *Command) {
	current.Command = command
}

func (current SimpleRemoteControl)buttonWasPressed() {

	if current.Command != nil {
		current.Command.Display()
		current.Command.Execute()
	}
}

func main() {

	//命令接受者
	var light = &Light{}
	//命令
	var lightOnCommand = NewLightOnCommand(light)
	//调用者
	simpleRemoteControl := &SimpleRemoteControl{}
	simpleRemoteControl.setCommand(&lightOnCommand.Command)
	fmt.Println("button was pressed down")
	simpleRemoteControl.buttonWasPressed()
}