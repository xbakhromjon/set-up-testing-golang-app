package domain

// Template domain struct
type Template struct {
	Id   int64
	Name string
}

type TemplateRepository interface {
	// needed methods
	Save(template Template) (int64, error)
	FindById(id int64) (*Template, error)
}
