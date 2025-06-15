package main

import (
	"math"
	"math/rand"
)

type LivingEntity struct {
	GameEntity
	Speed,
	SeeingRange,
	Energy,
	MaxEnergy,
	EnergyToReproduce float64
	ReproductionCooldown,
	MaxCooldown int
}

func (e *LivingEntity) RecoverFromReproduction() {
	if e.ReproductionCooldown > 0 {
		e.ReproductionCooldown--
	}
}

func (e *LivingEntity) BurnEnergy() {
	e.Energy--
}

func (e *LivingEntity) Metabolise() {
	e.BurnEnergy()
	if e.Energy <= 0 {
		e.Die()
	}
	e.RecoverFromReproduction()
}

func (e *LivingEntity) IsHungry() bool {
	return (e.Energy / e.MaxEnergy) < 0.3
}

func (e *LivingEntity) IsReadyToReproduce() bool {
	return e.ReproductionCooldown <= 0 && e.Energy >= e.EnergyToReproduce
}

func (e *LivingEntity) RecoverEnergy() {
	e.Energy = e.MaxEnergy
}

func (e *LivingEntity) StartReproductionCooldown() {
	e.ReproductionCooldown = e.MaxCooldown
}

func (e *LivingEntity) RandomMove() {
	deltaX := ((rand.Float64() * 2) - 1) * e.Speed
	deltaY := ((rand.Float64() * 2) - 1) * e.Speed
	e.Pos.Move(deltaX, deltaY)
}

func (e *LivingEntity) MoveToward(other Entity) {
	pos := e.GetPosition()
	targetPos := other.GetPosition()

	directionX := targetPos.X - pos.X
	directionY := targetPos.Y - pos.Y

	length := math.Sqrt(directionX*directionX + directionY*directionY)
	if length > 0 {
		moveX := (directionX / length) * e.Speed
		moveY := (directionY / length) * e.Speed
		pos.Move(moveX, moveY)
	}
}

func (e *LivingEntity) MoveAwayFrom(other Entity) {
	pos := e.GetPosition()
	otherPos := other.GetPosition()

	directionX := pos.X - otherPos.X
	directionY := pos.Y - otherPos.Y

	length := math.Sqrt(directionX*directionX + directionY*directionY)
	if length > 0 {
		moveX := (directionX / length) * e.Speed
		moveY := (directionY / length) * e.Speed
		pos.Move(moveX, moveY)
	}
}

func (e *LivingEntity) IsDangerouslyClose(other Entity) bool {
	pos := e.GetPosition()
	otherPos := other.GetPosition()

	distSq := pos.CalculateDistanceSquared(otherPos)
	halfSeeingRange := e.SeeingRange / 2
	return distSq < halfSeeingRange*halfSeeingRange
}
