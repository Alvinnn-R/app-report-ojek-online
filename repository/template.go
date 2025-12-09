package repository

import (
	"errors"
	"session-14/model"
	"time"
	"context"

	"github.com/jackc/pgx/v5"
)

type RepositoryTemplate struct {
	DB *pgx.Conn
}

func NewRepoTemplate(db *pgx.Conn) RepositoryTemplate {
	return RepositoryTemplate{
		DB: db,
	}
}

func (repo *RepositoryTemplate) Create(template *model.Template) error {
	query := `INSERT INTO templates (username, password, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4) RETURNING id`

	now := time.Now()
	err := repo.DB.QueryRow(context.Background(), query,
		template.Username,
		template.Password,
		now,
		now,
	).Scan(&template.ID)

	if err != nil {
		return err
	}

	template.CreatedAt = now
	template.UpdatedAt = now
	return nil
}

func (repo *RepositoryTemplate) List() ([]*model.Template, error) {
	query := `SELECT id, username, password, created_at, updated_at, deleted_at 
	          FROM templates WHERE deleted_at IS NULL`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*model.Template

	for rows.Next() {
		var t model.Template
		err := rows.Scan(&t.ID, &t.Username, &t.Password, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, err
		}
		templates = append(templates, &t)
	}

	return templates, nil
}

func (repo *RepositoryTemplate) Update(id int, template *model.Template) error {
	query := `UPDATE templates 
	          SET username = $1, password = $2, updated_at = $3 
	          WHERE id = $4 AND deleted_at IS NULL`

	_, err := repo.DB.Exec(context.Background(), query,
		template.Username,
		template.Password,
		time.Now(),
		id,
	)
	return err
}

func (repo *RepositoryTemplate) Delete(id int) error {
	query := `DELETE FROM templates WHERE id = $1 `
	_, err := repo.DB.Exec(context.Background(), query, id)
	return err
}

func (repo *RepositoryTemplate) CheckEmail(email string) error {
	query := `SELECT * FROM users WHERE email = $1 RETURNING ID`
	var id int
	err := repo.DB.QueryRow(context.Background(), query,
		email,
	).Scan(&id)

	if err != nil {
		return err
	}

	if id < 0 {
		return errors.New("not found")
	}
	return nil
}
