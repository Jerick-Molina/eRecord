package db


var createProjectSQL = 
`insert into Projects(ProjectName,CreatedBy,CreatedById,AssociatedCompany) values (?,?,?,?)`
func(q *Queries) CreateProjectByAssociatedCompany(ctx content.Context, c Company) error {
	

	err := q.db.ExecContent(ctx,createProjectSQL,c.Project,c.CreatedByName,c.CreatedById,c.AssociatedCompany)
	if err != nil {
		return err 
	}

	return nil
}
func (q *Queries) 