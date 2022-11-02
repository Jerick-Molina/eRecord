package db

import (
	"context"
	"database/sql"
	"fmt"
)

//HelperFunction
func InsertTicketToList(results *sql.Rows) ([]Ticket, error) {

	var tkts []Ticket
	for results.Next() {

		var tkt Ticket
		if err := results.Scan(&tkt.Id, &tkt.TicketName, &tkt.Description, &tkt.SeverityLevel, &tkt.AssociatedCompany, &tkt.AssignedUser, &tkt.CreatedBy, &tkt.AssignedProject, &tkt.Status); err != nil {

			return tkts, nil
		}

		tkts = append(tkts, tkt)
	}
	return tkts, nil
}

const createTicketAssignedToProjectSQL = `
insert Into Tickets(Name,Description,Priority,CompanyId,
	 AssignedTo,CreatedById,ProjectId,Status) values (?,?,?,?,?,?,?,?) 
`

func (q *Queries) CreateTicketAssignedToProject(ctx context.Context, tkt Ticket, compId int, usrId int) error {
	fmt.Println(tkt)
	fmt.Println(compId)
	fmt.Println(usrId)
	_, err := q.db.ExecContext(ctx, createTicketAssignedToProjectSQL,
		tkt.TicketName, tkt.Description, tkt.SeverityLevel, compId,
		tkt.AssignedUser, usrId, tkt.AssignedProject, "Open")

	if err != nil {
		return err
	}

	return err
}

const searchTicketByCompany = `select TicketId,Name,Description,Priority,CompanyId,
AssignedTo,CreatedById, ProjectId,Status  from Tickets where CompanyId = ?`

func (q *Queries) FindTicketsByAssociatedCompany(ctx context.Context, companyId int) ([]Ticket, error) {
	var tkts []Ticket

	results, err := q.db.QueryContext(ctx, searchTicketByCompany, companyId)
	if err != nil {
		return tkts, err
	}

	tkts, err = InsertTicketToList(results)
	if err != nil {
		return tkts, err
	}
	return tkts, nil
}

const searchTicketByProject = `select TicketId,Name,Description,Priority,CompanyId,
AssignedTo,CreatedById,ProjectId,Status from Tickets where CompanyId = ? and ProjectId = ?`

func (q *Queries) FindTicketsByAssociatedProject(ctx context.Context, companyId int, projectId int) ([]Ticket, error) {
	var tkts []Ticket

	results, err := q.db.QueryContext(ctx, searchTicketByProject, companyId, projectId)
	if err != nil {
		return tkts, err
	}

	tkts, err = InsertTicketToList(results)
	if err != nil {
		return tkts, err
	}
	return tkts, nil
}
