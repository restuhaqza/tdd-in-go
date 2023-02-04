package integrations

import (
	"github.com/gin-gonic/gin"
	"github.com/restuhaqza/tdd-in-go/integrations/db"
	"github.com/restuhaqza/tdd-in-go/integrations/handler"
	"github.com/restuhaqza/tdd-in-go/integrations/model"
)

func HTTPServer() *gin.Engine {

	// init gin engine
	r := gin.Default()

	// connect to database
	db, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	// load model
	noteModel := model.NewNoteModel(db)

	// load controller
	noteHandler := handler.NewNoteHandler(noteModel)

	// register endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/notes", noteHandler.GetNotes)
	r.POST("/notes", noteHandler.AddNote)

	return r
}
