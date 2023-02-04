package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuhaqza/tdd-in-go/integrations/input"
	"github.com/restuhaqza/tdd-in-go/integrations/model"
)

type noteHandler struct {
	noteModel model.NoteModel
}

func NewNoteHandler(noteModel model.NoteModel) *noteHandler {
	return &noteHandler{
		noteModel: noteModel,
	}
}

func (h *noteHandler) GetNotes(c *gin.Context) {
	notes, err := h.noteModel.FindNotes()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    notes,
	})
}

func (h *noteHandler) AddNote(c *gin.Context) {

	var inputBody input.InputCreate

	err := c.ShouldBindJSON(inputBody)

	fmt.Println(inputBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add note",
			"error":   err.Error(),
		})
		return
	}

	note, err := h.noteModel.AddNote(inputBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add note",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    note,
	})
}
