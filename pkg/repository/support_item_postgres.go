package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"support"
)

type SupportItemPostgres struct {
	db *sqlx.DB
}

func NewSupportItemPostgres(db *sqlx.DB) *SupportItemPostgres {
	return &SupportItemPostgres{db: db}
}

func (r *SupportItemPostgres) Create(item support.Support) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("[db] begin transaction error: %w", err)
	}
	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (number, name, date) VALUES ($1, $2, $3) RETURNING id", supportTable)
	row := tx.QueryRow(createItemQuery, item.Number, item.Name, item.Date)
	if err := row.Scan(&id); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return 0, fmt.Errorf("[db] transaction rollback error: %w", err)
		}
		return 0, fmt.Errorf("[db] creating support item error: %w", err)
	}
	return id, tx.Commit()
}

func (r *SupportItemPostgres) GetAll() ([]support.Support, error) {
	var items []support.Support
	query := fmt.Sprintf("SELECT * FROM %s", supportTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, fmt.Errorf("[db] getting support items error: %w", err)
	}
	return items, nil
}

func (r *SupportItemPostgres) GetItemById(itemId int) (support.Support, error) {
	var item support.Support
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", supportTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, fmt.Errorf("[db] getting support item by id error: %w", err)
	}
	return item, nil
}

func (r *SupportItemPostgres) Update(id int, input support.UpdateSupportItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Number != nil {
		setValues = append(setValues, fmt.Sprintf("number=$%d", argId))
		args = append(args, *input.Number)
		argId++
	}
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Date)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", supportTable, setQuery, argId)
	args = append(args, id)

	if _, err := r.db.Exec(query, args...); err != nil {
		return fmt.Errorf("[db] updating support item error: %w", err)
	}
	return nil
}
