package passport

import (
	"math"
	"time"
)

const (
	totalLength          = 64
	timestampLength      = 40
	nodeIDLength         = 12
	localCounterIDLength = 12
	// SingleNodeMode use when running SingleNodeMode in one machine
	SingleNodeMode = 0
)

// NodeConfig holds information about the current node running the service
type NodeConfig struct {
	NodeID       int64
	Counter      int64
	CounterLimit int64
}

// Passport Passport Unique IDs
type Passport struct {
	config *NodeConfig
}

// PreRun Create and configure the Passport
func PreRun(nodeID int64) Passport {
	limit := int64(math.Pow(2, localCounterIDLength) - 1)

	if nodeID >= int64(math.Pow(2, nodeIDLength)-1) {
		panic("Node ID exceeded the maximum length. Choose something less than %d")
	}

	Config := NodeConfig{NodeID: nodeID, Counter: 1, CounterLimit: limit}
	return Passport{config: &Config}

}

// ID generates a the next unique ID
func (r *Passport) ID() int64 {

	r.config.ValidateCounter()
	nodeID, counter := r.config.NodeID, r.config.Counter
	currentTimeStamp := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)) // adjust the current time

	id := currentTimeStamp << (totalLength - timestampLength)
	id |= nodeID << (totalLength - timestampLength - nodeIDLength)
	id |= counter
	r.config.Counter++
	return id
}

// ValidateCounter make sure local counter doesn't exceed the localCounterIDLength
func (c *NodeConfig) ValidateCounter() {
	if c.Counter >= c.CounterLimit {
		time.Sleep(1 * time.Millisecond)
		c.Counter = 1
	}
}
