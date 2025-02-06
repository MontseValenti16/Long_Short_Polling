package application

import (
    "arquitecturahex/src/products/domain/entities"
    "arquitecturahex/src/products/domain/repositories"
)

type ViewStudentIDUseCase struct {
    StudentRepo repositories.StudentRepository
}

func NewViewStudentIDUseCase(repo repositories.StudentRepository) *ViewStudentIDUseCase {
    return &ViewStudentIDUseCase{StudentRepo: repo}
}

func (uc *ViewStudentIDUseCase) Execute(id int) (*entities.Student, error) {
    return uc.StudentRepo.GetById(id)
}