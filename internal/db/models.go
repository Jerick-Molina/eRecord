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
	Id                int    `json:"ProjectId"`
	ProjectName       string `json:"Name"`
	Description       string `json:"Description"`
	AssociatedCompany int    `json:"CompanyId"`
}

type Ticket struct {
	Id                int    `json:"TicketId"`
	TicketName        string `json:"Name"`
	Status            string `json:"Status"`
	AssignedUser      int    `json:"AssignedTo"`
	CreatedBy         int    `json:"CreatedById"`
	AssociatedCompany int    `json:"CompanyId"`
	Description       string `json:"Description"`
	SeverityLevel     string `json:"Priority"`
	AssignedProject   int    `json:"ProjectId"`
}

type AccessTokenClaims struct {
	UserId    string
	Role      string
	CompanyId string
}

type InviteTokenClaims struct {
	InvitationCode string
	GivenRole      string
}
