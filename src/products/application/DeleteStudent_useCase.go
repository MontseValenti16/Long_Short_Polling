package application

import (
	"errors"
	"arquitecturahex/src/products/domain/repositories"
)

type DeleteStudentUseCase struct {
	StudentRepo repositories.StudentRepository
}

func NewDeleteStudentUseCase(repo repositories.StudentRepository) *DeleteStudentUseCase {
	return &DeleteStudentUseCase{StudentRepo: repo}
}

func (uc *DeleteStudentUseCase) Execute(id int) error {
	exists, err := uc.StudentRepo.ExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("el estudiante no existe")
	}

	return uc.StudentRepo.Delete(id)
}
