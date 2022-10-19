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

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Role, arg.CompanyId)
	if err != nil {
		return err
	}
	return nil
}

const searchAccountValidation = "select Email from Users where Email = ? "

func (q *Queries) EmailExistValidation(ctx context.Context, email string) error {
	var usr Account

	err := q.db.QueryRowContext(ctx, searchAccountValidation, email).Scan(&usr.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		} else {
			fmt.Println(err)
			return err
		}
	}

	return nil
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
			return Account{}, sql.ErrNoRows
		}
		return Account{}, err
	}

	//User is able to sign in
	return acc, nil
}
