package main

import (
	"fmt"
)

// Nice example - https://refactoring.guru/design-patterns/strategy/go/example
// Strategy is a behavioral design pattern that turns a set of <behaviors/algos into objects> and makes them interchangeable inside original context object.

// More complex example using behaviours
// https://faun.pub/head-first-design-patterns-using-go-1-welcome-to-design-patterns-the-strategy-pattern-6cbd940e113a
// Use Strategy:
//   when you want to use different variants of an algorithm within an object and be able to switch from one algorithm to another during runtime.
//   when you have a lot of similar classes that only differ in the way they execute some behaviour.
//   when your class has a massive conditional operator that switches between different variants of the same algorithm.

type Simulator interface {
	RunAlgo()
	SetAlgo(algo Algorithm)
}

type BasicSimulator struct {
	algo Algorithm
}

type Algorithm interface {
	Run()
}

func (bs *BasicSimulator) RunAlgo() {
	// delegate Run() call to the algo
	bs.algo.Run()
}

func (bs *BasicSimulator) SetAlgo(algo Algorithm) {
	bs.algo = algo
}

type KNN struct{}

func (*KNN) Run() {
	fmt.Println("Running KNN")
}

type LinearRegression struct{}

func (*LinearRegression) Run() {
	fmt.Println("Running LR")
}

func NewBasicSimulator(algo Algorithm) Simulator {
	return &BasicSimulator{algo: algo}
}

func NewLinearRegression() Algorithm {
	return &LinearRegression{}
}

func NewKNN() Algorithm {
	return &KNN{}
}

func main() {

	// use algo 1
	lr := NewLinearRegression()
	simulator := NewBasicSimulator(lr)
	simulator.RunAlgo()

	// dynamically switch to algo 2
	knn := NewKNN()
	simulator.SetAlgo(knn)
	simulator.RunAlgo()
}
