package main

import "void_fleet/ecs"

type Player struct {
	ecs.Component
	ecs.SpeedComponent
	ecs.HealthComponent
	money      int
	kill       int
	reputation int
}
