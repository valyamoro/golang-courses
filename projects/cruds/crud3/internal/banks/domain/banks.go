package domain

type Bank struct {
	ID int `json:"id" DB:"id"`
	Name string `json:"name" DB:"name"`
}

type Service struct {
	s Store
}

func NewService(store Store) *Service {
	return &Service{s: store}
}

func (svc Service) GetBanks() ([]Bank, error) {
	return svc.s.getAll()
}

func (svc Service) GetBank(id int) (*Bank, error) {
	return svc.s.get(id)
}

func (svc Service) Create(bank Bank) (int, error) {
	retrun svc.create(bank)
}

func (svc Service) DeleteBank() error {
	return svc.s.deleteAll()
}

func (svc Service) Update(bank Bank) (*Bank, error) {
	return svc.s.update(bank)
}

func (svc Service) Delete(id int) error {
	return svc.s.delete(id)
}
