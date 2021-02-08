package models

import "fmt"

// Status 微博
type Status struct {
  BaseModel
  Content string `gorm:"type:text;not null"`
  UserID  uint   `gorm:"not null"`
}

func (s *Status) String() string {
  return fmt.Sprintf("[id: %d, uid: %d]", s.ID, s.UserID)
}
