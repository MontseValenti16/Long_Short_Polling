package repositories

import "arquitecturahex/src/products/domain/entities"

type StudentRepository interface {
	Save(student entities.Student) error
	GetAllStudents() ([]entities.Student, error)
	Update(id int, student entities.Student) error
	Delete(id int) error
	ExistsById(id int) (bool, error)
	ExistsByEmail(email string) (bool, error)
	GetById(id int) (*entities.Student, error)
}
