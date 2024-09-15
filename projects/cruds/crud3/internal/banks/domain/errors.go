package domain

import "fmt"

const (
	DbQueryFail     = "DB_QUERY_FAIL"
	DbNotSupported  = "DB_NOT_SUPPORTED"
	EntityNotExists = "ENTITY_NOT_EXISTS"
)

type ErrDbQuery struct {
	Err error
}

func (e ErrDbQuery) Error() string {
	return fmt.Sprintf("%s: %s", DbQueryFail, e.Err)
}

type ErrDbNotSupported struct {
	Err error
}

func (s ErrDbNotSupported) Error() string {
	return fmt.Sprintf("%s: %s", DbNotSupported, e.Err)
}

type ErrEntityNotFound struct {
	Err error
}

func (e ErrEntityNotFound) Error() string {
	return fmt.Sprintf("%s: %s", EntityNotExist, e.Err)
}
