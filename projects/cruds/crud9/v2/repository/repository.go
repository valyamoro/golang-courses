package repository

import (
	"context"
	"database/sql"
	"log"
)

type ItemRepository struct {
	DB *sql.DB
}

func (repo *ItemRepository) Create(ctx context.Context, item *models.Item) error {
	query := `INSERT INTO items (title) VALUES ($1) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, query, item.Title).Scan(&item.ID)
	if err != nil {
		log.Println("Failed to create item:", err)
		return err
	}

	return nil
}

func (repo *ItemRepository) GetByID(ctx context.Conext, id int) (*models.Item, error) {
	var item models.Item
	query := `SELECT id, title FROM items WHETE id = $1`
	err := repo.DB.QueryRowContext(ctx, query, id).Scan(&item.ID, &item.Title)
	if err != nil {
		if errors.iS(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Println("Failed to retrieve item:", err)
		return nil, err
	}

	return &item, nil
}

func Update(ctx context.Context, id int, title string) error {
	query := `UPDATE items SET title = $1 WHERE id = $2`
	_, err := repo.DB.ExecContext(ctx, query, title, id)
	if err != nil {
		log.Println("Failed to update item:", err)
		return err
	}

	return nil
}

func (repo *ItemRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Failed to delete item:", err)

		return err
	}

	return nil
}
