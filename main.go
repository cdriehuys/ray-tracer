package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func main() {
	canvas := MakeCanvas(1000, 1000)
	radius := 400.0

	centerTranslate := MakeTranslation(float64(canvas.Width)/2, float64(canvas.Height)/2, 0)

	for i := 0; i < 60; i++ {
		point := MakePoint(0, radius, 0)
		rotation := MakeZRotation(2 * math.Pi / 60 * float64(i))
		point = centerTranslate.Multiply(rotation).TupleMultiply(point)

		xOrigin := int(math.Round(point.X))
		yOrigin := int(math.Round(point.Y))

		tickColor := MakeColor(.8, .8, .8)
		tickRadius := 1
		if i%5 == 0 {
			tickColor = MakeColor(.75, 0, 0)
			tickRadius = 2
		}

		for x := xOrigin - tickRadius; x < xOrigin+tickRadius; x++ {
			for y := yOrigin - tickRadius; y < yOrigin+tickRadius; y++ {
				canvas.SetPixel(x, y, tickColor)
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
