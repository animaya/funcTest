package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// Clear the screen (for simplicity, you can skip this step)
		fmt.Print("\033[H\033[2J")

		// Print the clock face
		fmt.Printf("   12\n  /  \\\n1| %02d |2\n  \\  /\n   6\n", hour)

		// Calculate the positions of the hands
		minAngle := float64(min) * 360 / 60
		hourAngle := (float64(hour%12)*360/12 + float64(min)*360/60/12)
		secAngle := float64(sec) * 360 / 60

		// Print the clock hands
		fmt.Printf("Minute Hand: %s\n", strings.Repeat(" ", int(minAngle)/6)+">")
		fmt.Printf("Hour Hand:   %s\n", strings.Repeat(" ", int(hourAngle)/6)+">")
		fmt.Printf("Second Hand: %s\n", strings.Repeat(" ", int(secAngle)/6)+">")

		time.Sleep(1 * time.Second)
	}
}
