package main

import (
	"math"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
)

func GetMousePosition() rl.Vector2 {

	pos := rl.GetMousePosition()

	// pos.X *= float32(screenWidth) / float32(rl.GetScreenWidth())
	// pos.Y *= float32(screenHeight) / float32(rl.GetScreenHeight())

	pos.X = float32(math.Round(float64(pos.X)))
	pos.Y = float32(math.Round(float64(pos.Y)))

	return pos

}

func GetWorldMousePosition() rl.Vector2 {

	pos := camera.Target

	mousePos := GetMousePosition()
	// mousePos.X -= screenWidth / 2
	// mousePos.Y -= screenHeight / 2

	mousePos.X -= float32(rl.GetScreenWidth() / 2)
	mousePos.Y -= float32(rl.GetScreenHeight() / 2)

	mousePos.X /= camera.Zoom
	mousePos.Y /= camera.Zoom

	pos.X += mousePos.X
	pos.Y += mousePos.Y

	return pos

}

var PrevMousePosition rl.Vector2 = rl.Vector2{}

func GetMouseDelta() rl.Vector2 {
	vec := raymath.Vector2Subtract(GetMousePosition(), PrevMousePosition)
	raymath.Vector2Scale(&vec, 1/camera.Zoom)
	return vec
}

func GetPath(folders ...string) string {

	// Running apps from Finder in MacOS makes the working directory the home directory, which is nice, because
	// now I have to make this function to do what should be done anyway and give me a relative path starting from
	// the executable so that I can load assets from the assets directory. :,)

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)

	if releaseMode == "false" {
		// Not in release mode, so current working directory is the root.
		exeDir, _ = os.Getwd()
	}

	return filepath.Join(exeDir, filepath.Join(folders...))

}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false
	}
	return true
}
