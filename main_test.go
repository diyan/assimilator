package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

func TestRunMainSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (t *testSuite) SetupTest() {
}

func (t *testSuite) TestSampleMainMethod() {
}
