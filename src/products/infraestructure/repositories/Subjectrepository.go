package repositories

import (
    "arquitecturahex/src/products/domain/entities"
    "arquitecturahex/src/products/infraestructure/db"
)

type SubjectRepositoryImpl struct{}

func NewSubjectRepositoryImpl() *SubjectRepositoryImpl {
    return &SubjectRepositoryImpl{}
}

func (repo *SubjectRepositoryImpl) Save(subject entities.Subject) error {
    query := "INSERT INTO subjects (name, credit) VALUES (?, ?)"
    _, err := db.DB.Exec(query, subject.Name, subject.Credit)
    return err
}

func (repo *SubjectRepositoryImpl) GetAllSubjects() ([]entities.Subject, error) {
    query := "SELECT id, name, credit FROM subjects"
    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var subjects []entities.Subject
    for rows.Next() {
        var subject entities.Subject
        if err := rows.Scan(&subject.ID, &subject.Name, &subject.Credit); err != nil {
            return nil, err
        }
        subjects = append(subjects, subject)
    }
    return subjects, nil
}

func (repo *SubjectRepositoryImpl) Update(id int, subject entities.Subject) error {
    query := "UPDATE subjects SET name = ?, credit = ? WHERE id = ?"
    _, err := db.DB.Exec(query, subject.Name, subject.Credit, id)
    return err
}

func (repo *SubjectRepositoryImpl) Delete(id int) error {
    query := "DELETE FROM subjects WHERE id = ?"
    _, err := db.DB.Exec(query, id)
    return err
}

func (repo *SubjectRepositoryImpl) ExistsById(id int) (bool, error) {
    query := "SELECT COUNT(*) FROM subjects WHERE id = ?"
    var count int
    err := db.DB.QueryRow(query, id).Scan(&count)
    return count > 0, err
}

func (repo *SubjectRepositoryImpl) ExistsByName(name string) (bool, error) {
    query := "SELECT COUNT(*) FROM subjects WHERE name = ?"
    var count int
    err := db.DB.QueryRow(query, name).Scan(&count)
    return count > 0, err
}