package application

import (
	"errors"
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type UpdateStudentUseCase struct {
	StudentRepo repositories.StudentRepository
}

func NewUpdateStudentUseCase(repo repositories.StudentRepository) *UpdateStudentUseCase {
	return &UpdateStudentUseCase{StudentRepo: repo}
}

func (uc *UpdateStudentUseCase) Execute(id int, student entities.Student) error {
	if student.Name == "" || student.Email == "" {
		return errors.New("el nombre y el correo son obligatorios")
	}

	exists, err := uc.StudentRepo.ExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("el estudiante no existe")
	}

	return uc.StudentRepo.Update(id, student)
}
