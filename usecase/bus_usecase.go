package usecase

import (
	"context"

	"github.com/tockn/tamabus/domain/repository"
)

type BusUseCase interface {
	GetBusData(ctx context.Context) (*model.Bus, error)
}

type busUseCase struct {
	repository.BusRepository
}

func NewBusUseCase(r repository.BusRepository) BusUseCase {
	return &busUseCase{r}
}

func (b *busUseCase) GetBusData(ctx context.Context) (*model.Bus, error) {
}
