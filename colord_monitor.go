// monitor.go
package main

import (
	"log"
	"os/exec"
	"regexp"
	"time"

	"github.com/atotto/clipboard"
)

func isValidHexCode(s string) bool {
	match, _ := regexp.MatchString("^#([A-Fa-f0-9]{3}([A-Fa-f0-9]{1})?|[A-Fa-f0-9]{6}([A-Fa-f0-9]{2})?)$", s)
	return match
}

func main() {
	var lastColor string
	for {
		colorStr, err := clipboard.ReadAll()
		if err != nil {
			log.Println("Failed to read from clipboard:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if colorStr != lastColor && isValidHexCode(colorStr) {
			lastColor = colorStr
			// fmt.Printf("Detected color change: %s\n", colorStr)

			// Spawn the display program with the color as an argument
			cmd := exec.Command("colord_display", colorStr)
			if err := cmd.Start(); err != nil {
				log.Printf("Error spawning display program: %v\n", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
