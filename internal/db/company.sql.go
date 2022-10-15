package db

import (
	"context"
)

const createCompany = `insert into Company(Name) values (?)
`

func (q *Queries) CreateCompany(context context.Context, companyName string) error {

	_, err := q.db.ExecContext(context, createCompany, companyName)
	if err != nil {
		return err
	}

	return nil
}
