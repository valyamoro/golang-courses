package mock

import "reflect"

type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

type MockServiceMockRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecoder{mock}

	return mock
}

func (m *MockService) EXCEPT() *MockServiceMockRecorder {
	return m.recorder
}

func (m *MockService) GetBanks() ([]domain.Bank, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanks")
	ret0, _ := ret[0].([]domain.Bank)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockServiceRecorder) GetBanks() *gomock.Call {
	mr.mock.ctrl.T.Helper()

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanks", reflect.TypeOf((*MockService)(nil).GetBanks))
}

func (m *MockService) GetBank(id int) (*domain.Bank, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBank", id)
	ret0, _ := ret[0].(*domain.Bank)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockServiceMockRecorder) GetBank(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBank", reflect.TypeOf((*MockService)(nil).GetBank), id)
}

func (m *MockService) Create(bank domain.Bank) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", bank)
	ret0, := ret[0].(int)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockServiceMockRecorder) Create(bank interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), bank)
}

func (m *MockService) DeleteBanks() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanks")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceMockRecorder) DeleteBanks() *gomock.Call {
	mr.mock.ctrl.T.Helper()

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanks", reflect.TypeOf((*MockService)(nil).DeleteBanks))
}

func (m *MockService) Update(bank domain.Bank) (*domain.Bank, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", bank)
	ret0, _ := ret[0].(*domain.Bank)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockServiceMockRecorder) Update(bank interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), bank)
}

func (m *MockService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()

	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}
