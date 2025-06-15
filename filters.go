package main

import (
	"math"
)

func getClosestFoxAndGrass(e Entity, entities []Entity) (fox *Fox, grass *Grass) {
	closestFoxDistanceSq := math.MaxFloat64
	closestGrassDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, entity := range entities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(e.GetPosition())
		switch eType := entity.(type) {
		case *Fox:
			if closestFoxDistanceSq >= dist {
				closestFoxDistanceSq = dist
				fox = eType
			}

		case *Grass:
			if closestGrassDistanceSq >= dist {
				closestGrassDistanceSq = dist
				grass = eType
			}

		default:
			continue
		}
	}
	return
}

func getClosestFoxAndRabbitReadyToReproduce(e Entity, entities []Entity) (fox *Fox, rabbit *Rabbit) {
	closestFoxDistanceSq := math.MaxFloat64
	closestRabbitDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, entity := range entities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(e.GetPosition())
		switch eType := entity.(type) {
		case *Fox:
			if closestFoxDistanceSq >= dist {
				closestFoxDistanceSq = dist
				fox = eType
			}

		case *Rabbit:
			if eType.IsReadyToReproduce() && closestRabbitDistanceSq >= dist && eType != e {
				closestRabbitDistanceSq = dist
				rabbit = eType
			}

		default:
			continue
		}
	}
	return
}

func getClosestRabbit(e Entity, entities []Entity) (rabbit *Rabbit) {
	closestRabbitDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, entity := range entities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(e.GetPosition())
		switch eType := entity.(type) {

		case *Rabbit:
			if closestRabbitDistanceSq >= dist {
				closestRabbitDistanceSq = dist
				rabbit = eType
			}

		default:
			continue
		}
	}
	return
}

func getClosestFox(e Entity, entities []Entity) (fox *Fox) {
	closestFoxDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, entity := range entities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(e.GetPosition())
		switch eType := entity.(type) {

		case *Fox:
			if closestFoxDistanceSq >= dist {
				closestFoxDistanceSq = dist
				fox = eType
			}

		default:
			continue
		}
	}
	return
}

func getClosestFoxReadyToReproduce(e Entity, entities []Entity) (fox *Fox) {
	closestFoxDistanceSq := math.MaxFloat64

	var otherPos *Position
	var dist float64

	for _, entity := range entities {
		otherPos = entity.GetPosition()
		dist = otherPos.CalculateDistanceSquared(e.GetPosition())
		switch eType := entity.(type) {

		case *Fox:
			if eType.IsReadyToReproduce() && closestFoxDistanceSq >= dist && eType != e {
				closestFoxDistanceSq = dist
				fox = eType
			}

		default:
			continue
		}
	}
	return
}
