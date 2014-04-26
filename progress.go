// Package progress provides a simple text animation in a goroutine.
package progress

import (
	"fmt"
	"strings"
	"time"
)

// Frames contain the different text animations available.
var Frames = map[string][]string{
	"arrow":     {" ^ ", " > ", " v ", " < "},
	"dots":      {" ... ", " o.. ", " .o. ", " ..o "},
	"ligatures": {" bq ", " dp ", " qb ", " pd "},
	"lines":     {"   ", " - ", " = ", " # ", " = ", " - "},
	"slash":     {" - ", " \\ ", " | ", " / "},
}

// Show starts a text animation in the specified style, returning
// a channel on to tell it to quit.
func Show(style string) chan bool {
	frames := Frames[style]
	backspaces := strings.Repeat("\b", len(frames[0]))
	tick := time.Tick(400 * time.Millisecond)
	quit := make(chan bool)
	go func() {
		for {
			for _, frame := range frames {
				fmt.Print(frame)
				select {
				case <-quit:
					fmt.Print(backspaces)
					quit <- true
					return
				case <-tick:
					fmt.Print(backspaces)
				}
			}
		}
	}()
	return quit
}
