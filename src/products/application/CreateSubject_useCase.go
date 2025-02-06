package application

import (
	"errors"
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type CreateSubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewCreateSubjectUseCase(repo repositories.SubjectRepository) *CreateSubjectUseCase {
	return &CreateSubjectUseCase{SubjectRepo: repo}
}

func (uc *CreateSubjectUseCase) Execute(subject entities.Subject) error {
	if subject.Name == "" {
		return errors.New("el nombre de la materia es obligatorio")
	}

	return uc.SubjectRepo.Save(subject)
}
