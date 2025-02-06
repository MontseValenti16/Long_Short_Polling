package application

import (
	"errors"
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type UpdateSubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewUpdateSubjectUseCase(repo repositories.SubjectRepository) *UpdateSubjectUseCase {
	return &UpdateSubjectUseCase{SubjectRepo: repo}
}

func (uc *UpdateSubjectUseCase) Execute(id int, subject entities.Subject) error {
	if subject.Name == "" {
		return errors.New("el nombre de la materia es obligatorio")
	}

	exists, err := uc.SubjectRepo.ExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("la materia no existe")
	}

	return uc.SubjectRepo.Update(id, subject)
}
