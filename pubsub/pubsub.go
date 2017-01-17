package pubsub

// SubscribeEvent - event for new subscribers
type SubscribeEvent struct {
	Topic   string
	Channel chan string
}

// UnsubscribeEvent - event for unsubscribe client
type UnsubscribeEvent struct {
	Topic   string
	Channel chan string
}

// BroadcastEvent - broadcast event
type BroadcastEvent struct {
	Topic string
	Msg   string
}

// PubSub - manage publish and subscribers
type PubSub struct {
	Subscribers map[string]map[chan string]bool
	Subscribe   chan SubscribeEvent
	Unsubscribe chan UnsubscribeEvent
	Broadcast   chan BroadcastEvent
}

// NewPubSub - create new PubSub
func NewPubSub() *PubSub {
	return &PubSub{
		make(map[string]map[chan string]bool),
		make(chan SubscribeEvent),
		make(chan UnsubscribeEvent),
		make(chan BroadcastEvent)}
}

// SubscribeClient - subscribe new client
func (pubsub *PubSub) SubscribeClient(event SubscribeEvent) {
	if pubsub.Subscribers[event.Topic] == nil {
		pubsub.Subscribers[event.Topic] = make(map[chan string]bool)
	}
	pubsub.Subscribers[event.Topic][event.Channel] = true
}

// UnsubscribeClient - unsubscribe client
func (pubsub *PubSub) UnsubscribeClient(event UnsubscribeEvent) {
	delete(pubsub.Subscribers[event.Topic], event.Channel)
}

// BroadcastClients - publish message to clients
func (pubsub *PubSub) BroadcastClients(topic, msg string) {
	for subscriber := range pubsub.Subscribers[topic] {
		subscriber <- msg
	}
}

// Start - run pubsub loop
func (pubsub *PubSub) Start() {
	for {
		select {
		case subscriber := <-pubsub.Subscribe:
			pubsub.SubscribeClient(subscriber)
		case unsubscribe := <-pubsub.Unsubscribe:
			pubsub.UnsubscribeClient(unsubscribe)
		case broadcast := <-pubsub.Broadcast:
			go pubsub.BroadcastClients(broadcast.Topic, broadcast.Msg)
		}
	}
}
