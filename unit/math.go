package unit

// contract
type IMath interface {
	Add(a, b int) int
	Multiply(a, b int) int
	Subtract(a, b int) int
	Divide(a, b int) int
}

// struct for implementation
type Math struct{}

// singletone pattern initialization
func New() *Math {
	return &Math{}
}

func (m *Math) Add(a, b int) int {
	return a + b
}

func (m *Math) Subtract(a, b int) int {
	return a - b
}

func (m *Math) Multiply(a, b int) int {
	return a * b
}

func (m *Math) Divide(a, b int) int {
	return a / b
}
