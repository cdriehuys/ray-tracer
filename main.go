package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}

		defer pprof.StopCPUProfile()
	}

	world := createWorld()

	canvasSize := 500
	camera := MakeCamera(canvasSize, canvasSize/2, math.Pi/3)

	from := MakePoint(0, 1.5, -5)
	to := MakePoint(0, 1, 0)
	up := MakeVector(0, 1, 0)
	camera.Transform = ViewTransform(from, to, up)

	log.Println("Rendering world...")
	canvas := Render(camera, world)
	log.Println("Finished rendering world.")

	writeCanvasToFile(canvas, "output.ppm")
}

func createWorld() World {
	log.Println("Constructing world...")
	wallMaterial := MakeMaterial()
	wallMaterial.Color = MakeColor(1, 0.9, 0.9)
	wallMaterial.Specular = 0

	floor := MakeSphereTransformed(
		MakeScale(10, 0.01, 10),
	)
	floor.material = wallMaterial

	leftWall := MakeSphereTransformed(
		MakeTranslation(0, 0, 5).
			Multiply(MakeYRotation(-math.Pi / 4)).
			Multiply(MakeXRotation(math.Pi / 2)).
			Multiply(MakeScale(10, 0.01, 10)),
	)
	leftWall.material = wallMaterial

	rightWall := MakeSphereTransformed(
		MakeTranslation(0, 0, 5).
			Multiply(MakeYRotation(math.Pi / 4)).
			Multiply(MakeXRotation(math.Pi / 2)).
			Multiply(MakeScale(10, 0.01, 10)),
	)
	rightWall.material = wallMaterial

	middle := MakeSphereTransformed(
		MakeTranslation(-0.5, 1, 0.5),
	)
	middleMaterial := MakeMaterial()
	middleMaterial.Color = MakeColor(0.1, 1, 0.5)
	middleMaterial.Diffuse = 0.7
	middleMaterial.Specular = 0.3
	middle.material = middleMaterial

	right := MakeSphereTransformed(
		MakeTranslation(1.5, 0.5, -0.5).Multiply(MakeScale(0.5, 0.5, 0.5)),
	)
	rightMaterial := MakeMaterial()
	rightMaterial.Color = MakeColor(0.5, 1, 0.1)
	rightMaterial.Diffuse = 0.7
	rightMaterial.Specular = 0.3
	right.material = rightMaterial

	left := MakeSphereTransformed(
		MakeTranslation(-1.5, 0.33, -0.75).
			Multiply(MakeScale(0.33, 0.33, 0.33)),
	)
	leftMaterial := MakeMaterial()
	leftMaterial.Color = MakeColor(1, 0.8, 0.1)
	leftMaterial.Diffuse = 0.7
	leftMaterial.Specular = 0.3
	left.material = leftMaterial

	light := MakePointLight(MakePoint(-10, 10, -10), MakeColor(1, 1, 1))

	world := MakeWorld()
	world.Light = light
	world.Objects = []Object{floor, leftWall, rightWall, middle, right, left}

	log.Println("Finished constructing world.")

	return world
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
