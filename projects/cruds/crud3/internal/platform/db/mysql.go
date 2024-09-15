package db

const mysql = "mysql"

func New(sqlConnection string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(mysql, sqlConnection)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	return db, nil
}
