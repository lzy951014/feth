package trust

import (
	"time"

	"github.com/lzy951014/feth/fp2p/tracker"
)

// requestTracker is a singleton tracker for request times.
var requestTracker = tracker.New(ProtocolName, time.Minute)
