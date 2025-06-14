package event

type EventType string

const EventLinkClick EventType = "link.click"

type Event struct {
	Data any
	Type EventType
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (eventBus *EventBus) Publish(event Event) {
	eventBus.bus <- event
}

func (eventBus *EventBus) Subscribe() <-chan Event {
	return  eventBus.bus
}
