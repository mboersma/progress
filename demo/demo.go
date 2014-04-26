package main

import (
	"fmt"
	"time"

	"github.com/mboersma/progress"
)

func main() {
	for style := range progress.Frames {
		fmt.Print(style)
		quit := progress.Show(style)
		time.Sleep(5 * time.Second)
		if quit <- true; <-quit {
			fmt.Print(" ")
		}
	}
}
