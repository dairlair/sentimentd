package domain

import (
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
)

type App struct {
	brainRepository BrainRepositoryInterface
}

func NewApp(brainRepository BrainRepositoryInterface) *App {
	return &App{
		brainRepository: brainRepository,
	}
}