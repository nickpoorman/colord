// display.go
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

type color struct {
	r, g, b, a float32
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Color argument missing")
	}
	colorStr := os.Args[1]
	// trim the whitespace
	colorStr = strings.TrimSpace(colorStr)

	if !isValidHexCode(colorStr) {
		// Try it with a # symbol prepended
		colorStr = fmt.Sprintf("#%s", colorStr)
		if !isValidHexCode(colorStr) {
			return
		}
	}

	displayTimeStr := os.Args[2]
	if displayTimeStr == "" {
		displayTimeStr = "1"
	}
	displayTime, err := strconv.Atoi(displayTimeStr)
	if err != nil {
		log.Fatalf("Error parsing display time: %v\n", err)
		return
	}

	windowWidthStr := os.Args[3]
	if windowWidthStr == "" {
		windowWidthStr = "50"
	}
	windowWidth, err := strconv.Atoi(windowWidthStr)
	if err != nil {
		log.Fatalf("Error parsing display time: %v\n", err)
		return
	}

	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw: %v", err)
	}
	defer glfw.Terminate()

	// Set the GLFW_DECORATED hint to false to create a window without a title bar
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.False) // Keeps the window from being resizable

	glfw.WindowHint(glfw.Resizable, glfw.False)
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	window, err := glfw.CreateWindow(windowWidth, windowWidth, "Clipboard Color", nil, nil)
	if err != nil {
		panic(err)
	}

	// Position window in the bottom right corner
	window.SetPos(mode.Width-windowWidth, mode.Height-windowWidth)

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize OpenGL: %v", err)
	}

	showWindowChan := make(chan color)
	hideWindowChan := make(chan struct{})

	go func() {
		if isValidHexCode(colorStr) {
			col, err := hexToRGBA(colorStr)
			if err != nil {
				log.Printf("Error parsing color '%s': %v", colorStr, err)
				return
			}
			showWindowChan <- col

			// Reset the timer to hide the window every time a new color is copied
			go func() {
				time.Sleep(time.Duration(displayTime) * time.Second) // Wait for 3 seconds
				hideWindowChan <- struct{}{}
			}()
		}
	}()

	for !window.ShouldClose() {
		select {
		case col := <-showWindowChan:
			gl.ClearColor(col.r, col.g, col.b, col.a)
			window.Show() // Show the window when a new color is received
		case <-hideWindowChan:
			window.Hide() // Hide the window after 3 seconds
			// Close the program
			return
		default:
			// Non-blocking default case to keep the application responsive
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func hexToRGBA(hexStr string) (color, error) {
	// Ensure hex code is in a full 6 or 8 character format by expanding shorthand notation
	if len(hexStr) == 4 { // #RGB
		expanded := "#"
		for _, char := range hexStr[1:] {
			expanded += fmt.Sprintf("%c%c", char, char)
		}
		hexStr = expanded
	} else if len(hexStr) == 5 { // #RGBA
		expanded := "#"
		for _, char := range hexStr[1:] {
			expanded += fmt.Sprintf("%c%c", char, char)
		}
		hexStr = expanded
	}

	var r, g, b, a uint64
	var err error

	// Parse the possibly expanded hex code
	if len(hexStr) == 7 { // Now #RRGGBB
		_, err = fmt.Sscanf(hexStr, "#%02x%02x%02x", &r, &g, &b)
		a = 0xff // Full opacity
	} else if len(hexStr) == 9 { // Now #RRGGBBAA
		_, err = fmt.Sscanf(hexStr, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	} else {
		return color{}, fmt.Errorf("invalid hex color format: %s", hexStr)
	}

	if err != nil {
		return color{}, err
	}

	// Convert to float32 and normalize (0-1 range)
	return color{
		r: float32(r) / 255.0,
		g: float32(g) / 255.0,
		b: float32(b) / 255.0,
		a: float32(a) / 255.0,
	}, nil
}

func isValidHexCode(s string) bool {
	match, _ := regexp.MatchString("^#([A-Fa-f0-9]{3}([A-Fa-f0-9]{1})?|[A-Fa-f0-9]{6}([A-Fa-f0-9]{2})?)$", s)
	return match
}
