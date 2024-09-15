package mock

type Service struct {
	mock.Mock
}

func (_m *Service) Create(bank domain.Bank) (int, error) {
	ret := _m.Called(bank)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Bank) int); ok {
		r0 = rf(bank)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Bank) error); ok {
		r1 = rf(bank)
	} else {
		r1 = reg.Error(1)
	}

	return r0, r1
}

func (_m *Service) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Service) DeleteBanks() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Service) GetBank(id int) (*domain.Bank, error) {
	ret := _m.Called(id)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(int) *domain.Bank); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) GetBanks() ([]domain.Banl, error) {
	ret := _m.Called()

	var r0 []domain.Bank
	if rf, ok := ret.Get(0).(func() []domain.Bank); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Bank)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) Update(bank domain.Bank) (*domain.Bank, error) {
	ret := _m.Called(bank)

	var r0 *domain.Bank
	if rf, ok := ret.Get(0).(func(domain.Bank) *domain.Bank); ok {
		r0 = rf(bank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bank)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Bank) error); ok {
		r1 = rf(bank)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
