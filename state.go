package godel

// StateIdentifier defines a simple identifier for a state.
type StateIdentifier string

// States defines the states that are registered.
type States map[StateIdentifier]State

// State is the definition for a state in a machine. A state needs to react to events and that is defined in OnEvent.
type State interface {
	OnEvent(event Event) (StateIdentifier, error) // Define what happens whenever a state encounters an event.
	Identify() StateIdentifier                    // We need a way to identify states.
}
