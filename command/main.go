package main

import "fmt"

//  Use the Command pattern when you want to parametrize objects with operations.
// 	Benefits - decouples invokers from executors - single responsibility, easy to add more commands
// 	Cons - adds complexity as we we require another "command" interface
//  For more info see https://refactoring.guru/design-patterns/command

//
// invoker
//

// invoker concrete class
type Invoker struct {
	command ICommand
}

func (i *Invoker) send() {
	// delegates execution to command interface
	i.command.execute()
}

//
// command
//

// command interface
type ICommand interface {
	execute()
}

// concrete command 1
type AddLiquidityCommand struct {
	receiver IReceiver
}

func (c *AddLiquidityCommand) execute() {
	// commands operate on the same environment
	c.receiver.addLiquidity()
}

// concrete command 2
type RemoveLiquidityCommand struct {
	receiver IReceiver
}

func (c *RemoveLiquidityCommand) execute() {
	c.receiver.removeLiquidity()
}

//
// receiver
//

// Shared Receiver (or environment)
type IReceiver interface {
	addLiquidity()
	removeLiquidity()
}

type receiver struct{}

func (*receiver) addLiquidity() {
	fmt.Println("added liquidity")
}

func (*receiver) removeLiquidity() {
	fmt.Println("removed liquidity")
}

//
// example
//
func main() {
	receiver := &receiver{}
	// Two invoker objects which contain different commands that operate on the same receiver
	invokerAdd := &Invoker{command: &AddLiquidityCommand{receiver: receiver}}
	invokerRemove := &Invoker{command: &RemoveLiquidityCommand{receiver: receiver}}
	invokerAdd.send()
	invokerRemove.send()
}
