package db

import "context"

const createProjectSQL = `
insert into Projects(Name,Description,CompanyId) values (?,?,?)
`

func (q *Queries) CreateProjectByAssociatedCompany(ctx context.Context, c CreateProjectTxParams) error {

	_, err := q.db.ExecContext(ctx, createProjectSQL, c.ProjectName, c.Description, c.AssociatedCompany)
	if err != nil {
		return err
	}

	return nil
}

const findProjectByAssociatedCompany = `
	select ProjectId,Name,Description,CompanyId from Projects where CompanyId = ?
`

func (q *Queries) FindProjectsByAssociatedCompany(ctx context.Context, companyId int) ([]Project, error) {
	var projects []Project

	result, err := q.db.QueryContext(ctx, findProjectByAssociatedCompany, companyId)
	if err != nil {
		return []Project{}, err
	}

	for result.Next() {
		var project Project
		if err := result.Scan(&project.Id, &project.ProjectName, &project.Description, &project.AssociatedCompany); err != nil {
			return []Project{}, err
		}

		projects = append(projects, project)
	}
	return projects, nil
}

const findSingleProjectByAssociatedCompany = `
	select ProjectId,Name,Description from Projects where CompanyId = ? and ProjectId = ?
`

func (q *Queries) FindSingleProjectByAssociatedCompany(ctx context.Context, companyId int, projectId int) (Project, error) {
	var project Project

	err := q.db.QueryRowContext(ctx, findSingleProjectByAssociatedCompany, companyId, projectId).Scan(&project.Id, &project.ProjectName, &project.Description)
	if err != nil {
		return project, err
	}

	return project, nil
}
