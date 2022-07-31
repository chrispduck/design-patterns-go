package main

import "fmt"


// A pattern that allows objects to notify other objects about a change in state.
// The Observer pattern provides a way to subscribe and unsubscribe to and from these events for any object that implements a subscriber interface.
// For more info see - https://refactoring.guru/design-patterns/observer 

//
// Business Logic concrete instance
//
type app struct {
	publisher Publisher
}

func (app *app) sendMessage(message string) {
	app.publisher.notifySubscribers(message)
}

//
// Publisher
//

// Publisher interface
type Publisher interface {
	registerSubscriber(l Subscriber)
	deregisterSubscriber(l Subscriber)
	notifySubscribers(data interface{})
}

// concrete publisher
type publisher struct {
	subscribers []Subscriber
}

func (p *publisher) registerSubscriber(s Subscriber) {
	p.subscribers = append(p.subscribers, s)
}

func (p *publisher) deregisterSubscriber(s Subscriber) {
	p.subscribers = removeFromslice(p.subscribers, s)
}

func (p *publisher) notifySubscribers(data interface{}) {
	for _, sub := range p.subscribers {
		sub.update(data)
	}
}

//
// Subscriber
//

// Subscriber interface
type Subscriber interface {
	update(data interface{})
	getID() string
}

type baseSubscriber struct {
	id string
}

func (b baseSubscriber) getID() string {
	return b.id
}

// concrete subscribers
type subscriberA struct {
	baseSubscriber
}

func (subscriberA) update(data interface{}) {
	fmt.Printf("subscriber A updated with data %v.\n", data)
}

type subscriberB struct {
	baseSubscriber
}

func (subscriberB) update(data interface{}) {
	fmt.Printf("subscriber B updated with data %v.\n", data)
}

//
// Helpers
//

func NewPublisher() Publisher {
	return &publisher{}
}

func NewSubscriberA(id string) Subscriber {
	return &subscriberA{
		baseSubscriber{
			id: id,
		},
	}
}

func NewSubscriberB(id string) Subscriber {
	return &subscriberB{
		baseSubscriber{
			id: id,
		},
	}
}

func removeFromslice(subscriberList []Subscriber, subscriberToRemove Subscriber) []Subscriber {
	subscriberListLength := len(subscriberList)
	for i, subscriber := range subscriberList {
		if subscriberToRemove.getID() == subscriber.getID() {
			// replace with the last element
			subscriberList[i] = subscriberList[subscriberListLength-1]
			return subscriberList[:subscriberListLength-1]
		}
	}
	// not in the list
	return subscriberList
}

func main() {
	subA := NewSubscriberA("Chris")
	subB := NewSubscriberB("Ben")
	pub := NewPublisher()
	app := app{
		publisher: pub,
	}
	app.publisher.registerSubscriber(subA)
	app.publisher.registerSubscriber(subB)
	app.sendMessage("hello world")
	app.publisher.deregisterSubscriber(subA)
	app.sendMessage("just a message for sub A")

}
