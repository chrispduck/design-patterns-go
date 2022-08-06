package main

import "fmt"

type ComponentName string

const (
	componentNameA ComponentName = "A"
	componentNameB ComponentName = "B"
	componentNameC ComponentName = "C"
)

type IMediator interface {
	notify(componentName ComponentName, event string)
	setComponents(A ComponentA, B ComponentB, C ComponentC)
}
type Mediator struct {
	componentA ComponentA
	componentB ComponentB
	componentC ComponentC
}

func NewMediator() IMediator {
	return &Mediator{}
}

func (m *Mediator) notify(componentName ComponentName, event string) {
	if componentName == componentNameA {
		m.componentA.reactTo(event)
		return
	}
	if componentName == componentNameB {
		m.componentB.reactTo(event)
		return
	}
	if componentName == componentNameC {
		m.componentC.reactTo(event)
	}
}

func (m *Mediator) setComponents(A ComponentA, B ComponentB, C ComponentC) {
	m.componentA = A
	m.componentB = B
	m.componentC = C
}

type ComponentA struct {
	m IMediator
}

func (A *ComponentA) doStuff() {
	A.m.notify(componentNameB, "!! do something for componentA !!")
}
func (ComponentA) reactTo(event string) {
	fmt.Printf("Component A reacting to %s\n", event)
}

type ComponentB struct {
	m IMediator
}

func (B *ComponentB) doStuff() {
	B.m.notify(componentNameC, "!! do something for componentB !!")
}

func (ComponentB) reactTo(event string) {
	fmt.Printf("Component B reacting to %s\n", event)
}

type ComponentC struct {
	m IMediator
}

func (ComponentC) reactTo(event string) {
	fmt.Printf("Component C reacting to %s\n", event)
}

func main() {
	mediator := NewMediator()
	A := ComponentA{m: mediator}
	B := ComponentB{m: mediator}
	C := ComponentC{m: mediator}
	mediator.setComponents(A, B, C)
	A.doStuff()
	B.doStuff()
}
