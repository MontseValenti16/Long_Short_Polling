package application

import (
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type ViewStudentUseCase struct {
	StudentRepo repositories.StudentRepository
}

func NewViewStudentUseCase(repo repositories.StudentRepository) *ViewStudentUseCase {
	return &ViewStudentUseCase{StudentRepo: repo}
}

func (uc *ViewStudentUseCase) Execute() ([]entities.Student, error) {
	return uc.StudentRepo.GetAllStudents()
}
