package utils

import (
	"os"
	"os/signal"
	"sync"
)

// WaitForCtrlC sets up a blocking waiter that looks for an os.Interrupt
// Useful for killing daemons and "while-looped" programs
func WaitForCtrlC() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
