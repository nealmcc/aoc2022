package main

// GameEvent represents a change in the state of the game.
type GameEvent struct {
	Seq         uint64    // a monotonically increasing number, unique within the flow of a game.
	Type        EventType // the type of event this is
	TotalRocks  int       // the number of rocks that have stopped so far.
	TotalHeight int       // the total height of all stopped rocks so far.
	Msg         string    // a short human-friendly description of the event
	Rows        []Row     // a copy of the top section of the board that contains changes in bottom-up order.
	RowsFrom    int       // the row number on the board where the event rows start.
	Error       error     // if an error occurred, this will contain the error.
}

// EventType indicates what type of event has happened.
type EventType int

const (
	GameStartedEvent EventType = iota
	NewRockEvent
	RockMovedEvent
	RockStoppedEvent
	GameStoppedEvent
	ErrorEvent
)

// String implements fmt.Stringer()
func (et EventType) String() string {
	return [...]string{
		"game started",
		"new rock",
		"rock moved",
		"rock stopped",
		"game stopped",
		"error occurred",
	}[et]
}
