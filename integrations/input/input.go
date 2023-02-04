package input

type InputCreate struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
