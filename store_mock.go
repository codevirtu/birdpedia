package main

import "github.com/stretchr/testify/mock"

// MockStore - The mock store contains additional methods for inspection
type MockStore struct {
	mock.Mock
}

// CreateBird -
func (m *MockStore) CreateBird(bird *Bird) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler test)
	*/
	rets := m.Called(bird)
	return rets.Error(0)
}

// GetBirds -
func (m *MockStore) GetBirds() ([]*Bird, error) {
	rets := m.Called()
	/*
		Since `rests.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to be the type we expect, which is this case is []*Bird
	*/
	return rets.Get(0).([]*Bird), rets.Error(1)
}

// InitMockStore -
func InitMockStore() *MockStore {
	/*
		Like the InitMockStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	s := new(MockStore)
	store = s
	return s
}
