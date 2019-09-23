package main

import (
	"time"
)

const (
	timestamp      = 42 // In milliseconds precision. The maximum timestamp that can be represented using 42 bits is Wednesday, May 15, 2109 7:35:11.103 AM
	nodeID         = 10 // gives us 1024 node
	localCounterID = 12 // gives us capacity of 4095 hit per timestamp
)

// RUID Rain Unique IDs
type RUID struct {
	value []byte
}

func main() {
	var r *RUID
	println("starting!")
	println("now: ", r.next())
	println("then: ", r.next())
	println("after: ", r.next())
}

func (r *RUID) next() int64 {
	currentTimeStamp := time.Now().UnixNano() / int64(time.Millisecond)

	return currentTimeStamp
}
