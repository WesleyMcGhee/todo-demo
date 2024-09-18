package models

type Todo struct {
  ID uint `gorm:"primaryKey"`
  Name string
  Urgency string
  Completed bool
}
