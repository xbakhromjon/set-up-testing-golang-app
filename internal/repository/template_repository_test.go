package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/domain"
	"log"
)

func (s *PostgresTestSuite) TestTemplateRepository_Save() {
	repository := NewTemplateRepository(s.db)
	got, err := repository.Save(domain.Template{
		Id:   1,
		Name: "name",
	})
	assert.Nil(s.T(), err)
	want := int64(1)
	assert.Equal(s.T(), want, got)
}

func (s *PostgresTestSuite) TestTemplateRepository_FindById() {
	repository := NewTemplateRepository(s.db)

	want := domain.Template{
		Id:   1,
		Name: "name",
	}
	// create template
	id, err := repository.Save(want)
	log.Println("Id:", id)
	assert.Nil(s.T(), err)

	// find by id
	got, err := repository.FindById(id)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), want, *got)
}
