package integrations_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/restuhaqza/tdd-in-go/integrations"
	"github.com/restuhaqza/tdd-in-go/integrations/db"
	"github.com/restuhaqza/tdd-in-go/integrations/db/seed"
	"github.com/restuhaqza/tdd-in-go/integrations/entity"
	"github.com/restuhaqza/tdd-in-go/integrations/input"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type httpServer struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

type responseConstruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func TestServer(t *testing.T) {
	suite.Run(t, new(httpServer))
}

// SetupSuite will be run before every test in the suite.
func (s *httpServer) SetupSuite() {
	s.server = integrations.HTTPServer()
	s.db = db.GetDB()
}

func (s *httpServer) SetupTest() {
	db.Clean()
	db.Migrate()
}

// TearDownTest will be run after every test in the suite.
func (s *httpServer) TearDownSuite() {
	// command for close the server
}

func (s *httpServer) performRequest(method, path string, bodyBytes []byte) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, bytes.NewReader(bodyBytes))

	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	s.server.ServeHTTP(w, req)
	return w
}

func (s *httpServer) TestPingEndpoint() {
	w := s.performRequest(http.MethodGet, "/ping", nil)
	fmt.Println(w.Body.String())

	s.Equal(200, w.Code)
	s.Equal("pong", w.Body.String())
}

func (s *httpServer) TestGetNotesEndpoint() {

	seed.SeedNote(s.db).Create()

	w := s.performRequest(http.MethodGet, "/notes", nil)

	type responseType struct {
		responseConstruct
		Data []entity.Note `json:"data"`
	}

	var responseBody responseType

	json.Unmarshal(w.Body.Bytes(), &responseBody)

	s.Equal(http.StatusOK, w.Code)
	s.Equal(1, len(responseBody.Data))
}

func (s *httpServer) TestAddNoteEndpoint() {
	var input input.InputCreate

	input.Title = "Title"
	input.Body = "Body"

	jsonBody, _ := json.Marshal(input)

	fmt.Println(string(jsonBody))

	w := s.performRequest(http.MethodPost, "/notes", jsonBody)

	fmt.Println(w.Body.String())
	s.Equal(200, w.Code)
}
