package db

import (
	"context"
	"database/sql"
	"fmt"
)

const createAccount = `-- Creates Account based on the Company's invite code
insert into Users(
	FirstName,
	LastName,
	Email,
	Password,
	Role,
	CompanyId
) values(?,?,?,?,?,?);
`

type CreateAccountParams struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

//TODO :  Return a key to the user after the account has been created!
func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Role, arg.CompanyId)
	if err != nil {
		return err
	}
	return nil
}

const searchAccountValidation = "select Email from Users where Email = ? "

func (q *Queries) AccountCreateValidation(ctx context.Context, email string) error {
	var usr Account

	err := q.db.QueryRowContext(ctx, searchAccountValidation, email).Scan(&usr.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("Email already exist")
}

const validSignInCredentials = `
select UserId,Role,CompanyId from Users where Email = ? and Password = ?
`

func (q *Queries) SignInValidation(ctx context.Context, email string, password string) (Account, error) {

	var acc Account

	err := q.db.QueryRowContext(ctx, validSignInCredentials, email, password).Scan(&acc.Id, &acc.Role, &acc.CompanyId)
	//User cannot sign in
	if err != nil {
		if err == sql.ErrNoRows {
			return acc, fmt.Errorf("Invalid Credidentials")
		}
		return acc, err
	}

	//User is able to sign in
	return acc, nil
}

const findAllUsersAssociatedByCompanyId = `
select UserId,FirstName,LastName,CompanyId from Users where CompanyId = ? `

func (q *Queries) FindAllUsersAssociatedByCompanyId(ctx context.Context, companyId int) ([]Account, error) {
	var accs []Account

	results, err := q.db.QueryContext(ctx, findAllUsersAssociatedByCompanyId, companyId)
	if err != nil {
		return accs, err
	}

	for results.Next() {
		var acc Account
		if err := results.Scan(&acc.Id, &acc.FirstName, &acc.LastName, &acc.CompanyId); err != nil {
			return accs, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}

const findAllUsersAssociatedByProjectId = `
select UserId,FirstName,LastName,Email,Role from Users where CompanyId = ?`

func (q *Queries) FindAllUsersAssociatedByProjectId(ctx context.Context, companyId int) ([]Account, error) {
	var accs []Account

	results, err := q.db.QueryContext(ctx, findAllUsersAssociatedByProjectId, companyId)
	if err != nil {
		return accs, err
	}

	for results.Next() {
		var acc Account
		if err := results.Scan(&acc.Id, &acc.FirstName, &acc.LastName, &acc.Email, &acc.Role); err != nil {
			return accs, err
		}
		accs = append(accs, acc)
	}

	return accs, nil
}
