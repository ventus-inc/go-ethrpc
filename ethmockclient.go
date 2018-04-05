package ethrpc

import (
	"github.com/stretchr/testify/mock"
)

type MockEthClient struct {
	mock.Mock
}

func (m *MockEthClient) Call(from string, to string, data string) (string, error) {
	args := m.Mock.Called(from, to, data)
	return args.Get(0).(string), args.Error(1)
}
