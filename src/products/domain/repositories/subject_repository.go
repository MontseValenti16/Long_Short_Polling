package repositories

import "arquitecturahex/src/products/domain/entities"

type SubjectRepository interface {
    Save(subject entities.Subject) error
    GetAllSubjects() ([]entities.Subject, error)
    Update(id int, subject entities.Subject) error
    Delete(id int) error
    ExistsById(id int) (bool, error)
    ExistsByName(name string) (bool, error)
}