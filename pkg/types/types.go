package types

type User struct {
	Stage        string
	Name         string `db:"name"`
	Surname      string `db:"surname"`
	MiddleName   string `db:"middle_name"`
	TypeAccount  string `db:"type_account"`
	PhoneNumber  string `db:"phone_number"`
	Email        string `db:"email"`
	Organization string `db:"organization"`
	Address      string `db:"address"`
	ID           int64  `db:"id"`
	ChatID       int64  `db:"chat_id"`
	Registered   bool
}

type Order struct {
	Organization             string
	ProblemDescription       string
	ContactPerson            string
	PhoneNumberContactPerson string
	Address                  string
	Email                    string
	IssueID                  string
	ClientChatID             int64
	Status                   string
}

type Issue struct {
	IssueID  string `db:"issue"`
	Status   string `db:"status"`
	Assignee string `db:"assignee"`
	ChatID   int64  `db:"chat_id"`
}
