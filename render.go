package main

import "log"

// Render a world using the view of a specific camera.
func Render(camera Camera, world World) Canvas {
	image := MakeCanvas(camera.Width, camera.Height)

	for y := 0; y < camera.Height; y++ {
		for x := 0; x < camera.Width; x++ {
			ray := camera.MakeRayForPixel(x, y)
			color := world.ColorAt(ray)

			image.SetPixel(x, y, color)
		}

		log.Printf("Rendered row %d of %d\n", y+1, camera.Height)
	}

	return image
}
