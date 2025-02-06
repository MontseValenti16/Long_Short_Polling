package application

import (
	"errors"
	"arquitecturahex/src/products/domain/repositories"
)

type DeleteSubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewDeleteSubjectUseCase(repo repositories.SubjectRepository) *DeleteSubjectUseCase {
	return &DeleteSubjectUseCase{SubjectRepo: repo}
}

func (uc *DeleteSubjectUseCase) Execute(id int) error {
	exists, err := uc.SubjectRepo.ExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("la materia no existe")
	}

	return uc.SubjectRepo.Delete(id)
}
