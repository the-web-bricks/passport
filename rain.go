package rain

import (
	"math"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

const (
	totalLength          = 64
	timestampLength      = 40
	nodeIDLength         = 12
	localCounterIDLength = 12
)

type NodeConfig struct {
	NodeID       int64
	Counter      int64
	CounterLimit int64
}

// Rain Rain Unique IDs
type Rain struct {
	config *NodeConfig
}

// PreRun Create and configure the Rain
func PreRun() Rain {
	node, _ := strconv.ParseInt(os.Getenv("NODE_ID"), 10, 64)
	limit := int64(math.Pow(2, localCounterIDLength) - 1)

	if node >= int64(math.Pow(2, nodeIDLength)-1) {
		panic("Node ID exceeded the maximum length. Choose something less than %d")
	}

	Config := NodeConfig{NodeID: node, Counter: 1, CounterLimit: limit}
	return Rain{config: &Config}

}

// Next generates a the next unique ID
func (r *Rain) Next() int64 {

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
