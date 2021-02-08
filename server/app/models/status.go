package models

// Status 微博
type Status struct {
  BaseModel
  Content string `gorm:"type:text;not null"`
  UserID  uint   `gorm:"not null"`
}
