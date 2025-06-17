package main

import (
	"ProjektGO/pkg/config"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func integerRangeValidator(min, max int) func(string) error {
	return func(text string) error {
		val, err := strconv.Atoi(text)
		if err != nil {
			return errors.New("musi być liczbą całkowitą")
		}
		if val < min || val > max {
			return fmt.Errorf("wartość musi być między %d a %d", min, max)
		}
		return nil
	}
}

func floatRangeValidator(min, max float64) func(string) error {
	return func(text string) error {
		val, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return errors.New("wartość musi być liczbą")
		}
		if val < min || val > max {
			return fmt.Errorf("wartość musi być między %.2f a %.2f", min, max)
		}
		return nil
	}
}

func main() {
	myApp := app.New()

	currentConfig := config.ReadConfig()
	uiModel := config.NewUIModel(currentConfig)

	myWindow := myApp.NewWindow("Ustawienia Symulacji")

	initialRabbitsEntry := widget.NewEntryWithData(uiModel.InitialRabbits)
	initialRabbitsEntry.Validator = integerRangeValidator(0, 2000)

	initialFoxesEntry := widget.NewEntryWithData(uiModel.InitialFoxes)
	initialFoxesEntry.Validator = integerRangeValidator(0, 2000)

	initialGrassEntry := widget.NewEntryWithData(uiModel.InitialGrass)
	initialGrassEntry.Validator = integerRangeValidator(0, 2000)

	rabbitSpeedEntry := widget.NewEntryWithData(uiModel.RabbitSpeed)
	rabbitSpeedEntry.Validator = floatRangeValidator(0.1, 10.0)

	rabbitSeeingRangeEntry := widget.NewEntryWithData(uiModel.RabbitSeeingRange)
	rabbitSeeingRangeEntry.Validator = floatRangeValidator(0.1, 800.0)

	rabbitInitialEnergyEntry := widget.NewEntryWithData(uiModel.RabbitInitialEnergy)
	rabbitInitialEnergyEntry.Validator = integerRangeValidator(1, 10000)

	rabbitMaxEnergyEntry := widget.NewEntryWithData(uiModel.RabbitMaxEnergy)
	rabbitMaxEnergyEntry.Validator = integerRangeValidator(1, 1000)

	rabbitEnergyToReproduceEntry := widget.NewEntryWithData(uiModel.RabbitEnergyToReproduce)
	rabbitEnergyToReproduceEntry.Validator = integerRangeValidator(1, 10000)

	rabbitReproductionCooldownEntry := widget.NewEntryWithData(uiModel.RabbitReproductionCooldown)
	rabbitReproductionCooldownEntry.Validator = integerRangeValidator(50, 10000)

	foxSpeedEntry := widget.NewEntryWithData(uiModel.FoxSpeed)
	foxSpeedEntry.Validator = floatRangeValidator(0.1, 10.0)

	foxSeeingRangeEntry := widget.NewEntryWithData(uiModel.FoxSeeingRange)
	foxSeeingRangeEntry.Validator = floatRangeValidator(0.1, 800.0)

	foxInitialEnergyEntry := widget.NewEntryWithData(uiModel.FoxInitialEnergy)
	foxInitialEnergyEntry.Validator = integerRangeValidator(1, 10000)

	foxMaxEnergyEntry := widget.NewEntryWithData(uiModel.FoxMaxEnergy)
	foxMaxEnergyEntry.Validator = integerRangeValidator(1, 10000)

	foxEnergyToReproduceEntry := widget.NewEntryWithData(uiModel.FoxEnergyToReproduce)
	foxEnergyToReproduceEntry.Validator = integerRangeValidator(1, 10000)

	foxReproductionCooldownEntry := widget.NewEntryWithData(uiModel.FoxReproductionCooldown)
	foxReproductionCooldownEntry.Validator = integerRangeValidator(50, 10000)

	grassSpawnCountEntry := widget.NewEntryWithData(uiModel.GrassSpawnCount)
	grassSpawnCountEntry.Validator = integerRangeValidator(0, 10)

	grassSpawnIntervalEntry := widget.NewEntryWithData(uiModel.GrassSpawnInterval)
	grassSpawnIntervalEntry.Validator = integerRangeValidator(1, 1000)

	form := widget.NewForm(
		widget.NewFormItem("Początkowa liczba królików", initialRabbitsEntry),
		widget.NewFormItem("Początkowa liczba lisów", initialFoxesEntry),
		widget.NewFormItem("Początkowa liczba traw", initialGrassEntry),
	)

	rabbitAdvancedForm := widget.NewForm(
		widget.NewFormItem("Prędkość", rabbitSpeedEntry),
		widget.NewFormItem("Zasięg wzroku", rabbitSeeingRangeEntry),
		widget.NewFormItem("Energia początkowa", rabbitInitialEnergyEntry),
		widget.NewFormItem("Maksymalna energia", rabbitMaxEnergyEntry),
		widget.NewFormItem("Energia do rozmnażania", rabbitEnergyToReproduceEntry),
		widget.NewFormItem("Cooldown rozmnażania", rabbitReproductionCooldownEntry),
	)

	foxAdvancedForm := widget.NewForm(
		widget.NewFormItem("Prędkość", foxSpeedEntry),
		widget.NewFormItem("Zasięg wzroku", foxSeeingRangeEntry),
		widget.NewFormItem("Energia początkowa", foxInitialEnergyEntry),
		widget.NewFormItem("Maksymalna energia", foxMaxEnergyEntry),
		widget.NewFormItem("Energia do rozmnażania", foxEnergyToReproduceEntry),
		widget.NewFormItem("Cooldown rozmnażania", foxReproductionCooldownEntry),
	)

	grassAdvancedForm := widget.NewForm(
		widget.NewFormItem("Ilość nowej trawy", grassSpawnCountEntry),
		widget.NewFormItem("Częstotliwość (w klatkach)", grassSpawnIntervalEntry),
	)

	accordion := widget.NewAccordion(
		widget.NewAccordionItem("Ustawienia Królików", rabbitAdvancedForm),
		widget.NewAccordionItem("Ustawienia Lisów", foxAdvancedForm),
		widget.NewAccordionItem("Ustawienia Trawy", grassAdvancedForm),
	)

	startButton := widget.NewButton("Uruchom Symulację", func() {
		finalConfig := uiModel.ToConfig()

		config.SaveConfig(finalConfig)

		cmd := exec.Command("./simulation")
		err := cmd.Start()
		if err != nil {
			log.Fatal("Błąd uruchomienia symulacji:", err)
		}
		myApp.Quit()
	})

	onAnyFieldChanged := func(_ string) {
		if initialRabbitsEntry.Validate() == nil &&
			initialFoxesEntry.Validate() == nil &&
			initialGrassEntry.Validate() == nil &&
			rabbitAdvancedForm.Validate() == nil &&
			foxAdvancedForm.Validate() == nil &&
			grassAdvancedForm.Validate() == nil {
			startButton.Enable()
		} else {
			startButton.Disable()
		}
	}

	for _, item := range form.Items {
		item.Widget.(*widget.Entry).OnChanged = onAnyFieldChanged
	}
	for _, item := range rabbitAdvancedForm.Items {
		item.Widget.(*widget.Entry).OnChanged = onAnyFieldChanged
	}
	for _, item := range foxAdvancedForm.Items {
		item.Widget.(*widget.Entry).OnChanged = onAnyFieldChanged
	}
	for _, item := range grassAdvancedForm.Items {
		item.Widget.(*widget.Entry).OnChanged = onAnyFieldChanged
	}

	myWindow.SetContent(container.NewVBox(
		form,
		accordion,
		startButton,
	))
	onAnyFieldChanged("")
	myWindow.ShowAndRun()
}
