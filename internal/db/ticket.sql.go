package db


const createTicketAssignedToProjectSQL = `
insert Into Tickets(TicketName,Description,SeverityLevel,AssociatedCompany,
Assigned User) values (?,?,?,?,?) 
`

func (q *Queries) CreatedTicketAssignedToProject(ctx, tkt Ticket) error{


	err := q.db.ExecContent(ctx, createTicketAssignedToProjectSQL, tkt.TicketName,tkt.Description, tkt.SeverityLevel,tkt.Asssociate,tkt.AssignedUser)

	if err != nil {
		return  err
	}

	return err 
}

func (q *Queries) 