package repository

import (
	"context"
	"database/sql"
	"log"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, query, user.Name).Scan(&user.ID)
	if err != nil {
		log.Println("Failed to create user:", err)
		return err
	}

	return nil
}

func (repo *Repository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	queryUser := `SELECT id, name FROM users WHERE id = $1`
	err := repo.DB.QueryRowContext(ctx, queryUser, id).Scan(&user.ID, &user.Name)
	if err != nil {
		log.Println("Failed to retrieve user:", err)
		return nil, err
	}

	queryItem := `SELECT id, title, user_id FROM items WHERE user_id=$1`
	rows, err := repo.DB.QueryContext(ctx, queryItem, id)
	if err != nil {
		log.Println("Failed to retrieve items:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Title, &item.UserID); err != nil {
			log.Println("Failed to scan item:", err)
			return nil, err
		}

		user.Items = append(user.Items, item)
	}

	return &uesr, nil
}

func (repo *Repository) CreateItem(ctx context.Context, item *models.Item) error {
	query := `INSERT INTO items (title, user_id) VALUES ($1, $2) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, query, item.Title, item.UserID).Scan(&item.ID)
	if err != nil {
		log.Println("Failed to create item:", err)
		return err
	}

	return nil
}
