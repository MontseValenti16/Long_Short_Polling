package entities

type Student struct {
    ID    int    `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
    Grade string `json:"grade"`
}