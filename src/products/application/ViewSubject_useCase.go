package application

import (
	"arquitecturahex/src/products/domain/entities"
	"arquitecturahex/src/products/domain/repositories"
)

type ViewSubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewViewSubjectUseCase(repo repositories.SubjectRepository) *ViewSubjectUseCase {
	return &ViewSubjectUseCase{SubjectRepo: repo}
}

func (uc *ViewSubjectUseCase) Execute() ([]entities.Subject, error) {
	return uc.SubjectRepo.GetAllSubjects()
}
