package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	canvasSize := 500
	canvas := MakeCanvas(canvasSize, canvasSize)
	shape := MakeSphereTransformed(MakeShear(.25, 0, 0, 0, 0, 0))
	shape.material.Color = MakeColor(1, 0.2, 1)

	lightPosition := MakePoint(-10, 10, -10)
	lightColor := MakeColor(1, 1, 1)
	light := MakePointLight(lightPosition, lightColor)

	rayOrigin := MakePoint(0, 0, -5)
	wallZ := 10
	wallSize := 7

	pixelSize := float64(wallSize) / float64(canvasSize)
	halfWall := float64(wallSize) / 2

	for y := 0; y < canvas.Height; y++ {
		// We subtract Y from the max Y value of the wall since the
		// y-coordinates of the canvas are inverted compared to the Y values in
		// world space.
		worldY := halfWall - pixelSize*float64(y)
		for x := 0; x < canvas.Width; x++ {
			worldX := -halfWall + pixelSize*float64(x)

			targetPosition := MakePoint(worldX, worldY, float64(wallZ))

			ray := MakeRay(
				rayOrigin,
				targetPosition.Subtract(rayOrigin).Normalized(),
			)
			intersections := shape.Intersect(ray)

			if hit, exists := intersections.Hit(); exists {
				point := ray.Position(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := ray.Direction.Multiply(-1)

				color := Lighting(
					hit.Object.Material(),
					light,
					point,
					eye,
					normal,
				)

				canvas.SetPixel(x, y, color)
			}
		}
	}

	writeCanvasToFile(canvas, "output.ppm")
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
