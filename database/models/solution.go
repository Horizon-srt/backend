package models

import (
	"time"

	"gorm.io/gorm"
)

type Solution struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ProblemID   uint   `json:"problem_id"`
	Name        string `sql:"index" json:"name"`
	Author      string `sql:"index" json:"auther"`
	Description string `json:"description"`
	Likes       uint   `json:"likes"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (s Solution) GetID() uint {
	return s.ID
}

func (s Solution) GetProblemID() uint {
	return s.ProblemID
}
