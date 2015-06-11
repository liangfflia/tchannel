package tchannel

import "time"

// waitFor will retry f till it returns true for a maximum of timeout.
// It returns true if f returned true, false if timeout was hit.
func waitFor(timeout time.Duration, f func() bool) bool {
	timeoutEnd := time.Now().Add(timeout)

	const maxSleep = time.Millisecond * 50
	sleepFor := time.Millisecond
	for {
		if f() {
			return true
		}

		if time.Now().After(timeoutEnd) {
			return false
		}

		time.Sleep(sleepFor)
		if sleepFor < maxSleep {
			sleepFor *= 2
		}
	}
}
