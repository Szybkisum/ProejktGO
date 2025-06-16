package config

import (
	"encoding/json"
	"log"
	"os"
)

type RabbitConfig struct {
	Speed                float64
	SeeingRange          float64
	InitialEnergy        int
	MaxEnergy            int
	EnergyToReproduce    int
	ReproductionCooldown int
}

type FoxConfig struct {
	Speed                float64
	SeeingRange          float64
	InitialEnergy        int
	MaxEnergy            int
	EnergyToReproduce    int
	ReproductionCooldown int
}

type GrassConfig struct {
	GrassSpawnInterval int
	GrassSpawnCount    int
}

type SimulationConfig struct {
	InitialRabbits int
	InitialFoxes   int
	InitialGrass   int
	RabbitParams   RabbitConfig
	FoxParams      FoxConfig
	GrassParams    GrassConfig
}

func ReadConfig() *SimulationConfig {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("Nie można wczytać pliku config.json:", err)
		return NewDefaultConfig()
	}
	var cfg SimulationConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Println("Błąd dekodowania JSON:", err)
		return NewDefaultConfig()
	}
	return &cfg
}

func SaveConfig(cfg *SimulationConfig) {
	configData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatal("Błąd kodowania JSON:", err)
	}
	err = os.WriteFile("config.json", configData, 0644)
	if err != nil {
		log.Fatal("Błąd zapisu pliku konfiguracyjnego:", err)
	}
}

func NewDefaultConfig() *SimulationConfig {
	return &SimulationConfig{
		InitialRabbits: 100,
		InitialFoxes:   50,
		InitialGrass:   1000,
		RabbitParams: RabbitConfig{
			Speed:                1.5,
			SeeingRange:          50.0,
			InitialEnergy:        500,
			MaxEnergy:            1000,
			EnergyToReproduce:    700,
			ReproductionCooldown: 150,
		},
		FoxParams: FoxConfig{
			Speed:                2.0,
			SeeingRange:          75.0,
			InitialEnergy:        750,
			MaxEnergy:            1500,
			EnergyToReproduce:    1000,
			ReproductionCooldown: 600,
		},
		GrassParams: GrassConfig{
			GrassSpawnInterval: 2,
			GrassSpawnCount:    1,
		},
	}
}
