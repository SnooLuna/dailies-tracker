package main

import (
	"context"
	"fmt"
	"os"
)

// App struct
type App struct {
	dailiesService DailiesService
	ctx            context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dailiesService = DailiesService{}
}

func (a *App) LoadCommissions(dateMs int64) []Commission {
	dailies, err := a.dailiesService.LoadCommissionsForDate(dateMs)
	if err != nil {
		fmt.Println("Error loading dailies:", err)
		return nil
	}
	fmt.Println("Loaded dailies:", dailies)
	return dailies
}

func (a *App) CreateTask(description string, rewardsJson string) (Commission, error) {
	return a.dailiesService.CreateNewCommission(description, rewardsJson)
}

func (a *App) CompleteTask(id int) error {
	return a.dailiesService.CompleteCommission(id)
}

func (a *App) DeleteTask(id int) error {
	return a.dailiesService.DeleteCommission(id)
}

func (a *App) IsDev() bool {
	return len(os.Getenv("DEBUG")) != 0 || len(os.Getenv("DEV")) != 0
}

func (a *App) GetLocale() string {
	envVar := os.Getenv("LANG")
	if len(envVar) == 0 {
		return "en-US.UTF-8"
	}
	return envVar
}
