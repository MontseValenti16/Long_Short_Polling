package application

import (
	"errors"
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type CreateStudentUseCase struct {
	StudentRepo repositories.StudentRepository
}

func NewCreateStudentUseCase(repo repositories.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{StudentRepo: repo}
}

func (uc *CreateStudentUseCase) Execute(student entities.Student) error {
	if student.Name == "" || student.Email == "" {
		return errors.New("el nombre y el correo son obligatorios")
	}

	exists, err := uc.StudentRepo.ExistsByEmail(student.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("el correo ya está registrado")
	}

	return uc.StudentRepo.Save(student)
}
