
package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" 
    "github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width = 500
	height = 500
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	for !window.ShouldClose() {

	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err !=  nil {
		panic(err)
	} 

	glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window,err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil , nil )
	if
}
