package main

import "fmt"

//
// invoker
//

// invoker concrete class
type Invoker struct {
	command Command
}

func (i *Invoker) send() {
	// delegates execution to command interface
	i.command.execute()
}

//
// command
//

// command interface
type Command interface {
	execute()
}

// concrete command 1
type AddLiquidityCommand struct {
	receiver Receiver
}

func (c *AddLiquidityCommand) execute() {
	// commands operate on the same environment
	c.receiver.addLiquidity()
}

// concrete command 2
type RemoveLiquidityCommand struct {
	receiver Receiver
}

func (c *RemoveLiquidityCommand) execute() {
	c.receiver.removeLiquidity()
}

//
// receiver
//

// Shared Receiver (or environment)
type Receiver interface {
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
