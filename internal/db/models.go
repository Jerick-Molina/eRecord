package db

type Account struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

type Company struct {
	Id          int    `json:"Id"`
	CompanyName string `json:"CompanyName"`
}

type InviteCode struct {
	CompanyId  int
	Role       string
	InviteCode string
}


type Project struct {

	Id int `json:CompanyId`
	ProjectName string`json:Name`
	CreatedByName string `json:CreatedByName`
	CreatedById  string `json:CreatedById`
	AssociatedCompany int `json:AssociatedCompany`

}

type Ticket struct {
	Id int `json:TicketId`
	TicketName string `json:Name`
	Description string `json:Description`
	SeverityLevel string `json:SeverityLevel`
	AssociatedCompany int `json:AssociatedCompany`
	AssignedUser int `json:AssignedUser`
}

type AccessTokenClaims {
	UserId string
	Role string
	CompanyId string
}

type InviteTokenClaims {
	InvitationCode string
	GivenRole string
}