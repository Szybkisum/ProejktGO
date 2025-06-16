package core

import (
	"math"
)

func filterFoxInterests(source Entity, nearbyEntities []Entity) (rabbit *Rabbit, fox *Fox) {
	closestFoxDistanceSq := math.MaxFloat64
	closestRabbitDistanceSq := math.MaxFloat64

	processor := func(candidate Entity, distSq float64) {
		switch identifiedEntity := candidate.(type) {
		case *Rabbit:
			if closestRabbitDistanceSq >= distSq {
				closestRabbitDistanceSq = distSq
				rabbit = identifiedEntity
			}
		case *Fox:
			if closestFoxDistanceSq >= distSq && identifiedEntity.IsReadyToReproduce() && identifiedEntity != source {
				closestFoxDistanceSq = distSq
				fox = identifiedEntity
			}
		}
	}
	processNearbyEntities(source, nearbyEntities, processor)
	return
}

func filterRabbitInterests(source Entity, nearbyEntities []Entity) (rabbit *Rabbit, fox *Fox, grass *Grass) {
	closestRabbitDistanceSq := math.MaxFloat64
	closestFoxDistanceSq := math.MaxFloat64
	closestGrassDistanceSq := math.MaxFloat64

	processor := func(candidate Entity, distSq float64) {
		switch identifiedEntity := candidate.(type) {
		case *Rabbit:
			if closestRabbitDistanceSq >= distSq && identifiedEntity.IsReadyToReproduce() && identifiedEntity != source {
				closestRabbitDistanceSq = distSq
				rabbit = identifiedEntity
			}
		case *Fox:
			if closestFoxDistanceSq >= distSq {
				closestFoxDistanceSq = distSq
				fox = identifiedEntity
			}
		case *Grass:
			if closestGrassDistanceSq >= distSq {
				closestGrassDistanceSq = distSq
				grass = identifiedEntity
			}
		}
	}
	processNearbyEntities(source, nearbyEntities, processor)
	return
}

func processNearbyEntities(source Entity, nearbyEntities []Entity, processor func(candidate Entity, distSq float64)) {
	var otherPos *Position
	var dist float64

	for _, entity := range nearbyEntities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(source.GetPosition())
		processor(entity, dist)
	}
}
