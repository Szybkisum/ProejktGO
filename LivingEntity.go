package main

type LivingEntity struct {
    Speed,
    SeeingRange,
    Energy,
    MaxEnergy,
    EnergyToReproduce float64
    ReproductionCooldown,
    MaxCooldown int
    IsDead bool
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
    if (e.Energy <= 0) {
        e.Die()
    }
    e.RecoverFromReproduction()
}

func (e *LivingEntity) Die() {
    e.IsDead = true
}


func (e *LivingEntity) IsHungry() bool {
    return (e.Energy / e.MaxEnergy) < 0.3
}

func (e *LivingEntity) IsReadyToReproduce() bool {
    return e.Energy >= e.EnergyToReproduce
}


func (e *LivingEntity) RecoverEnergy() {
    e.Energy = e.MaxEnergy
}

func (e *LivingEntity) StartReproductionCooldown() {
    e.ReproductionCooldown = e.MaxCooldown
} 