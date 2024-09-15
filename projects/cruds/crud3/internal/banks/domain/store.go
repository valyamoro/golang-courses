package domain

import "database/sql"

type Store interface {
	getAll() ([]Bank, error)
	get(id int) (*Bank, error)
	create(bank Bank) (int, error)
	deleteAll() error
	update(bank Bank) (*Bank, error)
	delete(id int) error
}

type BankStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *BankStore {
	return &BankStore{db: db}
}

func (s *BankStore) getAll() ([]Bank, error) {
	var banks []Bank
	if err := s.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, ErrDbQuery{Err: errors.Wrap(err, "BankStore.getAll() error")}
	}
	if banks == nil {
		return []Bank{}, nil
	}

	return banks, nil
}

func (s *BankStore) create(bank Bank) (int, error) {
	result, err := s.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, ErrDbQuery{Err: errors.Wrap(err, "")}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, ErrDbNotSupported{Err: errors.Wrap(err, "BankStore.create() error")}
	}
	return int(lastID), nil
}

func (s *BankStore) deleteAll() error {
	if, err := s.db.Exec("TRUNCATE table banks"); err != nil {
		return ErrDbQuery{Err: errors.Wrap(err, "BankStore.deleteAll() error")}
	}

	return nil
}

func (s *BankStore) delete(id int) error {
	resm err := s.db.Exec("DELETE FROM banks WHERE id=?", id)
	if err != nil {
		return ErrDbQuery{Err: errors.Wrap(err, "BankStore.delete() error")}
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrDbQuery{Err: errors.Wrap(err, "BankStore.delete() RowsAffected error")}
	}

	if affect == 0 {
		return ErrEntityNotFound{Err: errors.Wrap(err, "BankStore.delete() NotFoudn error")}
	}

	return nil
}

func (s *BankStore) update(bank Bank) (*Bank, error) {
	res, err := s.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, ErrDbQuery{Err: errors.Wrap(err, "BankStore.update() error")}
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, ErrDbQuery{Err: errors.Wrap(err, "BankStore.update() RowsAffected error")}
	}

	if affect == 0 {
		return nil, ErrEntityNotFound{Err: errors.Wrap(err, "BankStore.update() NotFound error")}
	}

	return &bank, nil
}
