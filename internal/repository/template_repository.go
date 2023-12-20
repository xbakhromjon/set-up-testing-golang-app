package repository

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"golang-project-template/internal/domain"
)

type templateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) domain.TemplateRepository {
	return &templateRepository{db: db}
}

func (t *templateRepository) Save(template domain.Template) (int64, error) {
	query, args, err := sq.
		Insert("template").Columns("id", "name").
		Values(template.Id, template.Name).
		ToSql()
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}
	_, err = t.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return template.Id, nil
}

func (t *templateRepository) FindById(id int64) (*domain.Template, error) {
	query, args, err := sq.Select("id, name").From("template").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	row := t.db.QueryRow(query, args...)

	var template domain.Template
	err = row.Scan(&template.Id, &template.Name)
	if err != nil {
		return nil, err
	}

	return &template, nil
}
