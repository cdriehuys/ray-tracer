package main

import "fmt"

type projectile struct {
	Position Tuple
	Velocity Tuple
}

type environment struct {
	Gravity Tuple
	Wind    Tuple
}

func main() {
	proj := projectile{MakePoint(0, 1, 0), MakeVector(0, 10, 0)}
	env := environment{MakeVector(0, -0.1, 0), MakeVector(-0.01, 0, 0)}
	tickCount := 0
	summarize(tickCount, proj)

	for proj.Position.Y >= 0 {
		proj = tick(env, proj)
		tickCount += 1
		summarize(tickCount, proj)
	}
}

func tick(env environment, proj projectile) projectile {
	return projectile{
		Position: proj.Position.Add(proj.Velocity),
		Velocity: proj.Velocity.Add(env.Gravity).Add(env.Wind),
	}
}

func summarize(tickCount int, proj projectile) {
	fmt.Printf(
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
