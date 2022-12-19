package main

// GameEvent represents a change in the state of the game.
type GameEvent struct {
	Seq         uint64    // a monotonically increasing number, unique within the flow of a game.
	Type        EventType // the type of event this is
	Msg         string    // a short human-friendly description of the event
	ChangedRows []Row     // a copy of the top section of the board, that contains changes, in bottom-up order.
	NumAdded    int       // the amount that the size of the board has increased.
	Error       error     // if an error occurred, this will contain the error.
}

// EventType indicates what type of event has happened.
type EventType int

const (
	GameStartedEvent EventType = iota + 1
	RockPlacedEvent
	RockMovedEvent
	GameStoppedEvent
	ErrorEvent
)

// String implements fmt.Stringer()
func (et EventType) String() string {
	return [...]string{
		"no-op",
		"game started",
		"new rock",
		"rock moved",
		"game stopped",
		"error occurred",
	}[et]
}
