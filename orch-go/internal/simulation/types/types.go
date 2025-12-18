package types

import (
	"context"
	"time"
)

// SimulationTime represents the virtual time in the simulation.
// It is distinct from wall-clock time to allow for speed control.
type SimulationTime time.Time

// TickInfo contains information preserved for each tick of the simulation.
type TickInfo struct {
	CurrentTime SimulationTime
	TickNumber  uint64
	Duration    time.Duration // Duration since last tick in virtual time
}

// SimulationContext provides context for agents during a tick.
type SimulationContext interface {
	context.Context
	Time() SimulationTime
	Tick() uint64
}
