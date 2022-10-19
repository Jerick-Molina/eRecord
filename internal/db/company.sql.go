package db

import (
	"context"
	"database/sql"
	"fmt"
)

const createCompany = `insert into Company(Name,UniqueId) values (?,?)
`

func (q *Queries) CreateCompany(context context.Context, companyName string, uniqueId string) error {

	_, err := q.db.ExecContext(context, createCompany, companyName, uniqueId)
	if err != nil {
		return err
	}

	return nil
}

const findCompanyWithUniqueId = `select CompanyId from Company where UniqueId = ?
`

func (q *Queries) FindCompanyWithUniqueId(ctx context.Context, uniqueId string) (int, error) {
	var id int
	response := q.db.QueryRowContext(ctx, findCompanyWithUniqueId, uniqueId).Scan(&id)
	if response != nil {
		fmt.Print(response)
		return -1, response
	}
	return id, nil
}
func (q *Queries) ValidateUniqueId(ctx context.Context, uniqueId string) error {
	var r string
	response := q.db.QueryRowContext(ctx, findCompanyWithUniqueId, uniqueId).Scan(&r)
	if response == sql.ErrNoRows {
		return response
	}
	return nil
}

// TODO: Create a company key when a Authorized User makes a invite code that has 3 params
// What company
// EXP time
// Has key been used?
func (q *Queries) JoinCompanyByInvitation(ctx context.Context, companyKey string) error {

	return nil
}
