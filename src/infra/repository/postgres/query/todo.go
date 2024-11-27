package query

import "fmt"

type TodoQueryBuilder interface {
	Select() TodoQuerySelectBuilder
	Insert() TodoQueryInsertBuilder
	Update() TodoQueryUpdateBuilder
}

type todoQueryBuilder struct{}

func Todo() TodoQueryBuilder {
	return &todoQueryBuilder{}
}

type TodoQuerySelectBuilder interface {
	All() string
	FindByID() string
}

type todoQuerySelectBuilder struct{}

func (*todoQueryBuilder) Select() TodoQuerySelectBuilder {
	return &todoQuerySelectBuilder{}
}

func (*todoQuerySelectBuilder) All() string {
	return "SELECT * FROM todo WHERE is_active = true;"
}

func (*todoQuerySelectBuilder) FindByID() string {
	return "SELECT * FROM todo WHERE id = $1;"
}

type todoQueryInsertBuilder struct {
}

type TodoQueryInsertBuilder interface {
	Insert() string
}

func (*todoQueryBuilder) Insert() TodoQueryInsertBuilder {
	return &todoQueryInsertBuilder{}
}

func (t todoQueryInsertBuilder) Insert() string {
	return fmt.Sprint(`INSERT INTO todo (description, is_active) VALUES ($1, $2) RETURNING id;`)
}

type todoQueryUpdateBuilder struct {
}

type TodoQueryUpdateBuilder interface {
	Update() string
	ChangeStatus() string
}

func (*todoQueryBuilder) Update() TodoQueryUpdateBuilder {
	return &todoQueryUpdateBuilder{}
}

func (t todoQueryUpdateBuilder) Update() string {
	return fmt.Sprint(`UPDATE todo 
							SET description = COALESCE(NULLIF($1, ''), description)
							WHERE id = $2;`)
}

func (t todoQueryUpdateBuilder) ChangeStatus() string {
	return fmt.Sprint(`UPDATE todo 
							SET is_active = COALESCE($1, is_active)
							WHERE id = $2;`)
}
