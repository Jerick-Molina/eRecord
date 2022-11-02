package db

import (
	"context"
	"database/sql"
	"eRecord/internal/security"
	"eRecord/util"
	"fmt"
)

type Record struct {
	db *sql.DB
	*Queries
}

func NewRecord(db *sql.DB) *Record {
	return &Record{
		db:      db,
		Queries: New(db),
	}
}

func (record *Record) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := record.db.BeginTx(ctx, nil)
	if err != nil {

		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {

			return err
		}
		return err
	}
	return tx.Commit()
}

type CreateAccountWithCompanyParams struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	CompanyName string `json:"CompanyName"`
}

//Create company and account transaction
func (record *Record) CreateStarterAccountTx(ctx context.Context, arg CreateAccountWithCompanyParams) (string, error) {
	var token string
	var uniqueCode string

	//Create Company
	err := record.execTx(ctx, func(*Queries) error {
		var err error
		var validUniqueId bool
		var compId int
		for validUniqueId == false {
			uniqueCode = util.RandomChars(10)
			_, err = record.Queries.FindCompanyWithUniqueId(ctx, uniqueCode)
			if err != nil {
				if err == sql.ErrNoRows {

					err = nil
					break
				}

				return err
			}

		}

		if err != nil {
			fmt.Println("1")

			return err
		}

		err = record.Queries.CreateCompany(ctx, arg.CompanyName, uniqueCode)
		if err != nil {
			fmt.Println("2")

			return err
		}
		compId, err = record.Queries.FindCompanyWithUniqueId(ctx, uniqueCode)
		if err != nil {
			return err
		}
		acc := CreateAccountParams{
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
			Email:     arg.Email,
			Password:  util.HashPassword(arg.Password),
			CompanyId: compId,
			Role:      "Owner",
		}
		err = record.Queries.AccountCreateValidation(ctx, acc.Email)

		if err != nil {
			fmt.Println("3")

			return err
		}
		err = record.Queries.CreateAccount(ctx, acc)
		if err != nil {
			fmt.Println(err)

			return nil
		}
		result, err := record.Queries.SignInValidation(ctx, acc.Email, acc.Password)

		if err != nil {

			fmt.Println("5")
			return err
		}

		token, err = security.CreateAccessToken(result.Id, acc.Role, acc.CompanyId)
		if err != nil {
			fmt.Println("6")
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println("7")
		return "", err
	}

	//Set the user to that company
	// Make it the owner

	return token, nil
}

type CreateAccountWithCompanyToken struct {
	FirstName          string `json:"FirstName"`
	LastName           string `json:"LastName"`
	Email              string `json:"Email"`
	Password           string `json:"Password"`
	CompanyInviteToken string `json:"CompanyInviteToken"`
}

// Join company and create account transaction
func (record *Record) CreateAccountAndJoinCompanyTx(ctx context.Context, args CreateAccountWithCompanyToken) (string, error) {

	var token string

	err := record.execTx(ctx, func(q *Queries) error {

		claims, err := security.TokenReader(args.CompanyInviteToken)

		invCode := fmt.Sprintf("%v", claims["invitationCode"])

		if err != nil {
			return err
		}
		result, err := record.ValidateInvitationToken(ctx, invCode)

		var acc = CreateAccountParams{
			FirstName: args.FirstName,
			LastName:  args.LastName,
			Email:     args.Email,
			Password:  util.HashPassword(args.Password),
			Role:      result.Role,
			CompanyId: result.CompanyId,
		}
		err = record.Queries.AccountCreateValidation(ctx, acc.Email)

		if err != nil {
			fmt.Println(err)
			return err
		}
		err = record.CreateAccount(ctx, acc)
		if err != nil {
			return err
		}
		token, err = security.CreateAccessToken(result.Id, acc.Role, acc.CompanyId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

// TX : PROJECTS

type CreateProjectTxParams struct {
	ProjectName       string `json:"Name"`
	Description       string `json:"Descrption"`
	AssociatedCompany int    `"json:"CompanyId""`
}

func (record *Record) CreateProjectTx(ctx context.Context, args CreateProjectTxParams) error {

	if err := record.execTx(ctx, func(q *Queries) error {

		err := q.CreateProjectByAssociatedCompany(ctx, args)
		if err != nil {

			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (record *Record) FindProjectByAssociatedCompanyTx(ctx context.Context, companyId int) ([]Project, error) {
	var projects []Project
	var err error

	err = record.execTx(ctx, func(q *Queries) error {
		projects, err = q.FindProjectsByAssociatedCompany(ctx, companyId)
		if err != nil {

			return err
		}
		return nil
	})
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (record *Record) FindSingleProjectTx(ctx context.Context, companyId int, projectId int) (error, any) {
	var tkts []Ticket
	var usrs []Account
	var project Project
	var err error

	var array []any
	err = record.execTx(ctx, func(q *Queries) error {

		tkts, err = q.FindTicketsByAssociatedProject(ctx, companyId, projectId)
		if err != nil {
			return err
		}
		array = append(array, tkts)
		usrs, err = q.FindAllUsersAssociatedByProjectId(ctx, companyId)
		if err != nil {
			return err
		}
		array = append(array, usrs)

		project, err = q.FindSingleProjectByAssociatedCompany(ctx, companyId, projectId)
		if err != nil {
			return err
		}
		array = append(array, project)
		return nil
	})
	if err != nil {
		return err, nil
	}

	return err, array
}

//TX : TICKETS

//TX: DATA

// func (record *Record) GetDashboardDataTx(ctx context.Context)
func (record *Record) CreateTicketDashboardTx(ctx context.Context, companyId int) (error, any) {
	var err error
	var usrs []Account
	var projects []Project

	var array []any
	err = record.execTx(ctx, func(q *Queries) error {
		projects, err = q.FindProjectsByAssociatedCompany(ctx, companyId)
		if err != nil {
			fmt.Println("Err?")
			return err
		}
		array = append(array, projects)

		usrs, err = q.FindAllUsersAssociatedByCompanyId(ctx, companyId)
		if err != nil {
			fmt.Println("Err?")
			return err
		}
		array = append(array, usrs)

		return nil
	})
	if err != nil {
		return err, nil
	}

	return err, array
}
func (record *Record) FindTicketByAssociatedCompanyTx(ctx context.Context, companyId int) ([]Ticket, error) {
	var tkts []Ticket
	var err error

	err = record.execTx(ctx, func(q *Queries) error {

		tkts, err = q.FindTicketsByAssociatedCompany(ctx, companyId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return []Ticket{}, err
	}

	return tkts, nil
}

func (record *Record) FindTicketsDashboardTx(ctx context.Context, companyId int) (error, any) {
	var tkts []Ticket
	var usrs []Account
	var projects []Project
	var err error

	var array []any
	err = record.execTx(ctx, func(q *Queries) error {

		tkts, err = q.FindTicketsByAssociatedCompany(ctx, companyId)
		if err != nil {
			return err
		}
		array = append(array, tkts)

		usrs, err = q.FindAllUsersAssociatedByCompanyId(ctx, companyId)
		if err != nil {
			return err
		}
		array = append(array, usrs)
		projects, err = q.FindProjectsByAssociatedCompany(ctx, companyId)
		if err != nil {
			return err
		}
		array = append(array, projects)
		return nil
	})
	if err != nil {
		return err, nil
	}

	return err, array
}
