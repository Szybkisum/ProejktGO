package config

import (
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type UIModel struct {
	InitialRabbits binding.String
	InitialFoxes   binding.String
	InitialGrass   binding.String

	RabbitSpeed                binding.String
	RabbitSeeingRange          binding.String
	RabbitInitialEnergy        binding.String
	RabbitMaxEnergy            binding.String
	RabbitEnergyToReproduce    binding.String
	RabbitReproductionCooldown binding.String

	FoxSpeed                binding.String
	FoxSeeingRange          binding.String
	FoxInitialEnergy        binding.String
	FoxMaxEnergy            binding.String
	FoxEnergyToReproduce    binding.String
	FoxReproductionCooldown binding.String

	GrassSpawnCount    binding.String
	GrassSpawnInterval binding.String
}

func NewUIModel(cfg *SimulationConfig) *UIModel {
	ir := binding.NewString()
	ir.Set(strconv.Itoa(cfg.InitialRabbits))

	ifr := binding.NewString()
	ifr.Set(strconv.Itoa(cfg.InitialFoxes))

	ig := binding.NewString()
	ig.Set(strconv.Itoa(cfg.InitialGrass))

	rs := binding.NewString()
	rs.Set(strconv.FormatFloat(cfg.RabbitParams.Speed, 'f', 2, 64))

	rsr := binding.NewString()
	rsr.Set(strconv.FormatFloat(cfg.RabbitParams.SeeingRange, 'f', 2, 64))

	rie := binding.NewString()
	rie.Set(strconv.Itoa(cfg.RabbitParams.InitialEnergy))

	rme := binding.NewString()
	rme.Set(strconv.Itoa(cfg.RabbitParams.MaxEnergy))

	rer := binding.NewString()
	rer.Set(strconv.Itoa(cfg.RabbitParams.EnergyToReproduce))

	rrc := binding.NewString()
	rrc.Set(strconv.Itoa(cfg.RabbitParams.ReproductionCooldown))

	fs := binding.NewString()
	fs.Set(strconv.FormatFloat(cfg.FoxParams.Speed, 'f', 2, 64))

	fsr := binding.NewString()
	fsr.Set(strconv.FormatFloat(cfg.FoxParams.SeeingRange, 'f', 2, 64))

	fie := binding.NewString()
	fie.Set(strconv.Itoa(cfg.FoxParams.InitialEnergy))

	fme := binding.NewString()
	fme.Set(strconv.Itoa(cfg.FoxParams.MaxEnergy))

	fer := binding.NewString()
	fer.Set(strconv.Itoa(cfg.FoxParams.EnergyToReproduce))

	frc := binding.NewString()
	frc.Set(strconv.Itoa(cfg.FoxParams.ReproductionCooldown))

	gsc := binding.NewString()
	gsc.Set(strconv.Itoa(cfg.GrassParams.GrassSpawnCount))

	gsi := binding.NewString()
	gsi.Set(strconv.Itoa(cfg.GrassParams.GrassSpawnInterval))

	return &UIModel{
		InitialRabbits:             ir,
		InitialFoxes:               ifr,
		InitialGrass:               ig,
		RabbitSpeed:                rs,
		RabbitSeeingRange:          rsr,
		RabbitInitialEnergy:        rie,
		RabbitMaxEnergy:            rme,
		RabbitEnergyToReproduce:    rer,
		RabbitReproductionCooldown: rrc,
		FoxSpeed:                   fs,
		FoxSeeingRange:             fsr,
		FoxInitialEnergy:           fie,
		FoxMaxEnergy:               fme,
		FoxEnergyToReproduce:       fer,
		FoxReproductionCooldown:    frc,
		GrassSpawnCount:            gsc,
		GrassSpawnInterval:         gsi,
	}
}

func (m *UIModel) ToConfig() *SimulationConfig {
	ir, _ := binding.StringToInt(m.InitialRabbits).Get()
	ifr, _ := binding.StringToInt(m.InitialFoxes).Get()
	ig, _ := binding.StringToInt(m.InitialGrass).Get()
	rs, _ := binding.StringToFloat(m.RabbitSpeed).Get()
	rsr, _ := binding.StringToFloat(m.RabbitSeeingRange).Get()
	rie, _ := binding.StringToInt(m.RabbitInitialEnergy).Get()
	rme, _ := binding.StringToInt(m.RabbitMaxEnergy).Get()
	rer, _ := binding.StringToInt(m.RabbitEnergyToReproduce).Get()
	rrc, _ := binding.StringToInt(m.RabbitReproductionCooldown).Get()
	fs, _ := binding.StringToFloat(m.FoxSpeed).Get()
	fsr, _ := binding.StringToFloat(m.FoxSeeingRange).Get()
	fie, _ := binding.StringToInt(m.FoxInitialEnergy).Get()
	fme, _ := binding.StringToInt(m.FoxMaxEnergy).Get()
	fer, _ := binding.StringToInt(m.FoxEnergyToReproduce).Get()
	frc, _ := binding.StringToInt(m.FoxReproductionCooldown).Get()
	gsc, _ := binding.StringToInt(m.GrassSpawnCount).Get()
	gsi, _ := binding.StringToInt(m.GrassSpawnInterval).Get()

	return &SimulationConfig{
		InitialRabbits: ir,
		InitialFoxes:   ifr,
		InitialGrass:   ig,
		RabbitParams: RabbitConfig{
			Speed:                rs,
			SeeingRange:          rsr,
			InitialEnergy:        rie,
			MaxEnergy:            rme,
			EnergyToReproduce:    rer,
			ReproductionCooldown: rrc,
		},
		FoxParams: FoxConfig{
			Speed:                fs,
			SeeingRange:          fsr,
			InitialEnergy:        fie,
			MaxEnergy:            fme,
			EnergyToReproduce:    fer,
			ReproductionCooldown: frc,
		},
		GrassParams: GrassConfig{
			GrassSpawnCount:    gsc,
			GrassSpawnInterval: gsi,
		},
	}
}
