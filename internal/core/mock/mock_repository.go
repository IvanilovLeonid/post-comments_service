package mock

import (
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"social-comments/internal/core/domain"
)

// MockPostRepository is a mock of PostRepository interface
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPostRepository) Create(ctx context.Context, post domain.Post) (domain.Post, error) {
	ret := m.ctrl.Call(m, "Create", ctx, post)
	ret0, _ := ret[0].(domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPostRepositoryMockRecorder) Create(ctx, post interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostRepository)(nil).Create), ctx, post)
}

// GetByID mocks base method
func (m *MockPostRepository) GetByID(ctx context.Context, id int) (domain.Post, error) {
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPostRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPostRepository)(nil).GetByID), ctx, id)
}

// GetAll mocks base method
func (m *MockPostRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Post, error) {
	ret := m.ctrl.Call(m, "GetAll", ctx, limit, offset)
	ret0, _ := ret[0].([]domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPostRepositoryMockRecorder) GetAll(ctx, limit, offset interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPostRepository)(nil).GetAll), ctx, limit, offset)
}
