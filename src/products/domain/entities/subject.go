package entities

type Subject struct {
    ID     int    `json:"id" gorm:"primaryKey"`
    Name   string `json:"name"`
    Credit int    `json:"credit"`
}