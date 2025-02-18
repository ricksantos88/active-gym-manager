package usecase

import (
	"active-gym-manager/internal/domain/model"
	"active-gym-manager/internal/repository"
)

type ContactTypeUsecase struct {
	repository repository.ContactTypeRepository
}

func NewContactTypeUsecase(ctr repository.ContactTypeRepository) ContactTypeUsecase {
	return ContactTypeUsecase{
		repository: ctr,
	}
}

func (ctu *ContactTypeUsecase) FindAll() ([]model.ContactType, error) {
	contactTypes, err := ctu.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return contactTypes, nil
}
