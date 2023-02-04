package model

import (
	"github.com/restuhaqza/tdd-in-go/integrations/entity"
	"github.com/restuhaqza/tdd-in-go/integrations/input"
	"gorm.io/gorm"
)

type NoteModel interface {
	FindNotes() ([]entity.Note, error)
	AddNote(input input.InputCreate) (entity.Note, error)
}

type noteModel struct {
	db *gorm.DB
}

func NewNoteModel(db *gorm.DB) *noteModel {
	return &noteModel{db: db}
}

func (m *noteModel) FindNotes() ([]entity.Note, error) {
	var notes []entity.Note

	if err := m.db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (m *noteModel) AddNote(input input.InputCreate) (entity.Note, error) {
	var note entity.Note

	note.Title = input.Title
	note.Body = input.Body

	err := m.db.Create(&note).Error

	if err != nil {
		return note, err
	}

	return note, nil
}
