package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

type projectile struct {
	Position Tuple
	Velocity Tuple
}

type environment struct {
	Gravity Tuple
	Wind    Tuple
}

func main() {
	canvas := MakeCanvas(900, 550)

	start := MakePoint(0, 1, 0)
	velocity := MakeVector(1, 1.8, 0).Normalized().Multiply(11.25)
	proj := projectile{start, velocity}
	env := environment{MakeVector(0, -0.1, 0), MakeVector(-0.02, 0, 0)}
	tickCount := 0
	summarize(tickCount, proj)

	for proj.Position.Y >= 0 {
		proj = tick(env, proj)
		tickCount += 1
		summarize(tickCount, proj)
		recordProjectileLocation(&canvas, proj)
	}

	writeCanvasToFile(canvas, "output.ppm")
}

func tick(env environment, proj projectile) projectile {
	return projectile{
		Position: proj.Position.Add(proj.Velocity),
		Velocity: proj.Velocity.Add(env.Gravity).Add(env.Wind),
	}
}

func summarize(tickCount int, proj projectile) {
	log.Printf(
		"[%5d] Projectile at (%3.2f, %3.2f, %3.2f); Velocity (%3.2f, %3.2f, %3.2f)\n",
		tickCount,
		proj.Position.X,
		proj.Position.Y,
		proj.Position.Z,
		proj.Velocity.X,
		proj.Velocity.Y,
		proj.Velocity.Z,
	)
}

func recordProjectileLocation(canvas *Canvas, proj projectile) {
	color := MakeColor(1, 0, 0)
	width := 1

	x := int(math.Round(proj.Position.X))
	y := int(math.Round(float64(canvas.Height) - 1 - proj.Position.Y))

	for xDraw := x - width; xDraw < x+width; xDraw++ {
		for yDraw := y - width; yDraw < y+width; yDraw++ {
			if xDraw >= 0 && xDraw < canvas.Width && yDraw >= 0 && yDraw < canvas.Height {
				canvas.SetPixel(xDraw, yDraw, color)
			}
		}
	}
}

func writeCanvasToFile(canvas Canvas, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create '%s': %v", filePath, err)
	}
	log.Printf("Created output file '%s'", filePath)

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing '%s': %v", filePath, err)
		}

		log.Printf("Closed file '%s'", filePath)
	}()

	// Buffer writes to the file to increase performance. Otherwise each
	// incremental write is done straight to disk and kills performance. This
	// reduced write times from ~5 sec to < 1 sec.
	fileWriter := bufio.NewWriter(file)
	defer func() {
		if err := fileWriter.Flush(); err != nil {
			log.Fatalf("Error flushing file writer: %v", err)
		}
	}()

	log.Println("Writing canvas to PPM...")
	if err := WriteCanvasToPPM(canvas, fileWriter); err != nil {
		log.Fatalf("Error writing PPM to '%s': %v", filePath, err)
	}
	log.Println("Finished writing canvas to PPM.")
}
