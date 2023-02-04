package unit_test

import (
	"fmt"
	"testing"

	"github.com/restuhaqza/tdd-in-go/unit"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	math unit.IMath
}

// entrypoint test
func TestMath(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	fmt.Println("SetupSuite() executed")
	s.math = unit.New()
}

func (s *testSuite) TearDownSuite() {
	fmt.Println("TearDownSuite() executed")
}

func (s *testSuite) SetupTest() {
	fmt.Println("SetupTest() executed")
}

func (s *testSuite) TearDownTest() {
	fmt.Println("TearDownTest() executed")
}

func (s *testSuite) TestAdd() {
	result := s.math.Add(1, 2)

	s.Equal(3, result)
}

func (s *testSuite) TestSubtract() {
	result := s.math.Subtract(2, 1)

	s.Equal(1, result)
}

func (s *testSuite) TestMultiply() {
	result := s.math.Multiply(2, 2)

	s.Equal(4, result)
}

func (s *testSuite) TestDivide() {
	result := s.math.Divide(4, 2)

	s.Equal(2, result)
}
