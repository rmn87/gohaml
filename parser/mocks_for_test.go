package parser

import (
	"github.com/realistschuckle/testify/mock"
)

type mockNode struct {
	mock.Mock
}

func (self *mockNode) Accept(visitor NodeVisitor) {
	self.Mock.Called(visitor)
}

type mockVisitor struct {
	mock.Mock
}

func (self *mockVisitor) VisitDoctype(node *DoctypeNode) {
	self.Mock.Called(node)
}

type mockRuneReader struct {
	mock.Mock
}

func (self *mockRuneReader) ReadRune() (rune, int, error) {
	args := self.Mock.Called()
	return args.Get(0).(rune), args.Int(1), args.Error(2)
}