package seed

import (
	"github.com/restuhaqza/tdd-in-go/integrations/entity"
	"gorm.io/gorm"
)

type seedNote struct {
	db *gorm.DB
}

func SeedNote(db *gorm.DB) *seedNote {
	return &seedNote{db: db}
}

func (s *seedNote) Create() (entity.Note, error) {
	note := entity.Note{
		Title: "Title",
		Body:  "Content",
	}

	err := s.db.Create(&note).Error

	if err != nil {
		return note, err
	}

	return note, nil
}
