package engine

import (
	"orch-go/internal/simulation/types"
	"sync"
	"time"
)

// Clock manages the simulation time.
type Clock struct {
	mu              sync.RWMutex
	currentTime     time.Time
	tickInterval    time.Duration
	tickCount       uint64
	speedMultiplier float64
	running         bool
}

// NewClock creates a new simulation clock starting at startTime.
func NewClock(startTime time.Time, tickInterval time.Duration) *Clock {
	return &Clock{
		currentTime:     startTime,
		tickInterval:    tickInterval,
		speedMultiplier: 1.0,
	}
}

// CurrentTime returns the current simulation time.
func (c *Clock) CurrentTime() types.SimulationTime {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return types.SimulationTime(c.currentTime)
}

// Tick advances the clock by one tick interval.
func (c *Clock) Tick() types.TickInfo {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.tickCount++
	// Just advance simply by tick interval for now.
	// Speed multiplier handles how fast Ticks happen in real time, not how much virtual time passes per tick usually,
	// unless we want variable steps. Let's keep fixed time step for ensuring determinism.
	previousTime := c.currentTime
	c.currentTime = c.currentTime.Add(c.tickInterval)

	return types.TickInfo{
		CurrentTime: types.SimulationTime(c.currentTime),
		TickNumber:  c.tickCount,
		Duration:    c.currentTime.Sub(previousTime),
	}
}

func (c *Clock) SetSpeed(multiplier float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.speedMultiplier = multiplier
}
