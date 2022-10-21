package db

import (
	"context"
	"database/sql"
	"fmt"
)

const createCompany = `insert into Company(Name,UniqueCode) values (?,?)
`

func (q *Queries) CreateCompany(context context.Context, companyName string, uniqueCode string) error {

	_, err := q.db.ExecContext(context, createCompany, companyName, uniqueCode)
	if err != nil {
		return err
	}

	return nil
}

const findCompanyWithUniqueId = `select CompanyId from Company where UniqueCode = ?
`

func (q *Queries) FindCompanyWithUniqueId(ctx context.Context, uniqueId string) (int, error) {
	var id int
	err := q.db.QueryRowContext(ctx, findCompanyWithUniqueId, uniqueId).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, sql.ErrNoRows
		}
		return -1, err
	}
	return id, nil
}

// TODO: Create a company key when a Authorized User makes a invite code that has 3 params
// What company
// EXP time
// Has key been used?
// func (q *Queries) JoinCompanyByInvitation(ctx context.Context, companyKey string) error {

// 	return nil
// }
const createToken = `
 insert into Invite(InvitationToken,WasUsed) values (?,?)
`

func (q *Queries) SaveToken(ctx context.Context, token string) error {

	_, err := q.db.ExecContext(ctx, createToken, token, false)
	if err != nil {
		return err
	}

	return nil
}

const validateInvitationToken = `
select InvitationToken,WasUsed from InviteCode where InvitationToken = ?
`

func (q *Queries) ValidateInvitationToken(ctx context.Context, token string) (Account, error) {
	var hasBeenUsed bool = false
	var acc Account
	err := q.db.QueryRowContext(ctx, validateInvitationToken, token).Scan(&hasBeenUsed, &acc.Role)
	if err != nil {
		return acc, err
	}

	if hasBeenUsed == true {
		return acc, fmt.Errorf("This invite code is not valid anymore")
	}

	return acc, nil
}

const deValidateInviteCode = `
 update InviteCode set WasUsed = ?  where InvitationToken = ?
`

func (q *Queries) deValidateInviteCode(ctx context.Context, token string) error {

	_, err := q.db.ExecContext(ctx, deValidateInviteCode, true, token)
	if err != nil {
		return err
	}

	return nil
}
