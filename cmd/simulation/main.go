package main

import (
	"ProjektGO/pkg/config"
	"ProjektGO/pkg/core"
)

func main() {
	cfg := config.ReadConfig()
	core.RunSimulation(cfg)
}
