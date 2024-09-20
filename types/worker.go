package types

type Worker interface {
    getMessages() []Message
}

type Message struct {
    eventID string
    description string
}
