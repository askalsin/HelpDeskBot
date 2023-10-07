package database

import (
	"context"
	"database/sql"

	log "codeberg.org/kalsin/UtelBot/pkg/logging"
	"codeberg.org/kalsin/UtelBot/pkg/types"
)

func GetUserData(userID int64) *types.User {
	user := types.User{}

	err = conn.QueryRow(context.Background(),
		`SELECT id, chat_id, name, type_account, phone_number, email, organization, address
		 FROM "public.users"
		 WHERE chat_id=$1
		`, userID).Scan(&user.ID, &user.ChatID, &user.Name, &user.TypeAccount,
		&user.PhoneNumber, &user.Email, &user.Organization, &user.Address)
	if err != nil {
		log.Warning.Printf("QueryRow failed or user don't exist (func GetUserData()): %s\n", err)
		return nil
	}

	return &user
}

func NewUser(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`INSERT INTO "public.users" (chat_id, name, phone_number, email, organization, address) 
		 VALUES ($1,$2,$3,$4,$5,$6)`,
		user.ChatID, user.Name, user.PhoneNumber, user.Email, user.Organization, user.Address)

	if err != nil {
		log.Error.Printf("error create new user (func NewUser()): %s\n", err)
		return err
	}

	return nil
}

func DeleteUser(chatID int64) error {
	_, err = conn.Exec(context.Background(),
		`DELETE FROM "public.issues" WHERE chat_id=($1)`, chatID)

	if err != nil {
		log.Error.Printf("Error delete issues (func DeleteUser()): %s\n", err)
		return err
	}

	_, err = conn.Exec(context.Background(),
		`DELETE FROM "public.users" WHERE chat_id=($1)`, chatID)

	if err != nil {
		log.Error.Printf("error delete user (func DeleteUser()): %s\n", err)
		return err
	}

	return nil
}

func ChangeName(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.users" SET name=($1) WHERE chat_id=($2)`, user.Name, user.ChatID)

	if err != nil {
		log.Error.Printf("error change name (func ChangeName()): %s\n", err)
		return err
	}

	return nil
}

func ChangeOrganization(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.users" SET organization=($1) WHERE chat_id=($2)`,
		user.Organization, user.ChatID)

	if err != nil {
		log.Error.Printf("error change organization (func ChangeOrganization()): %s\n", err)
		return err
	}

	return nil
}

func ChangeAddressOrganization(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.users" SET address=($1) WHERE chat_id=($2)`, user.Address, user.ChatID)

	if err != nil {
		log.Error.Printf("error change address organization (func ChangeAddressOrganization()): %s\n", err)
		return err
	}

	return nil
}

func ChangeEmail(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.users" SET email=($1) WHERE chat_id=($2)`, user.Email, user.ChatID)

	if err != nil {
		log.Error.Printf("error change address organization (func ChangeAddressOrganization()): %s\n", err)
		return err
	}

	return nil
}

func ChangePhoneNumber(user *types.User) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.users" SET phone_number=($1) WHERE chat_id=($2)`,
		user.PhoneNumber, user.ChatID)

	if err != nil {
		log.Error.Printf("error change phone_number (func ChangePhoneNumber()): %s\n", err)
		return err
	}

	return nil
}

func NewIssue(order *types.Order) error {
	_, err = conn.Exec(context.Background(),
		`INSERT INTO "public.issues" (chat_id, issue, status) 
		 VALUES ($1,$2,$3)`,
		order.ClientChatID, order.IssueID, order.Status)

	if err != nil {
		log.Error.Printf("error add new issue (func NewIssue()): %s\n", err)
		return err
	}

	return nil
}

func GetIssuesID(userID int64) ([]string, error) {
	var issues []string

	rows, err := conn.Query(context.Background(),
		`SELECT issue FROM "public.issues" WHERE chat_id=$1`, userID)
	if err != nil {
		log.Warning.Printf("(func GetIssuesID()): %s\n", err)
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Error.Println("error while iterating dataset")
			return nil, err
		}
		issues = append(issues, values[0].(string))
	}

	return issues, nil
}

func GetIssues() map[string]types.Issue {
	issues := make(map[string]types.Issue)

	rows, err := conn.Query(context.Background(),
		`SELECT chat_id, issue, assignee, status FROM "public.issues"`)
	if err != nil {
		log.Warning.Printf("(func GetIssues()): %s\n", err)
		return nil
	}

	for rows.Next() {
		var issue types.Issue
		var nullstring sql.NullString

		err := rows.Scan(&issue.ChatID, &issue.IssueID, &nullstring, &issue.Status)
		if err != nil {
			log.Error.Println(err)
			return nil
		}
		if nullstring.Valid {
			issue.Assignee = nullstring.String
		}

		issues[issue.IssueID] = issue
	}

	return issues
}

func NewGroup(chatID int64) error {
	_, err = conn.Exec(context.Background(),
		`INSERT INTO "public.groups" (chat_id)
		 VALUES ($1)`, chatID)

	if err != nil {
		log.Error.Printf("error add new group (func NewGroup()): %s\n", err)
		return err
	}

	return nil
}

func GetGroupsChatID() []int64 {
	rows, err := conn.Query(context.Background(),
		`SELECT chat_id FROM "public.groups"`)
	if err != nil {
		log.Warning.Printf("(func GetGroupsChatID()): %s\n", err)
	}

	var chatsID []int64
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Error.Println("error while iterating dataset")
			return nil
		}

		chatID := values[0].(int64)
		chatsID = append(chatsID, chatID)
	}

	return chatsID
}

func ChangeAssignee(issue *types.Issue) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.issues" SET assignee=($1) WHERE issue=($2)`, issue.Assignee, issue.IssueID)

	if err != nil {
		log.Error.Printf("error change assignee (func ChangeAssignee()): %s\n", err)
		return err
	}

	return nil
}

func ChangeStatus(issue *types.Issue) error {
	_, err = conn.Exec(context.Background(),
		`UPDATE "public.issues" SET status=($1) WHERE issue=($2)`, issue.Status, issue.IssueID)

	if err != nil {
		log.Error.Printf("error change assignee (func ChangeAssignee()): %s\n", err)
		return err
	}

	return nil
}

func DeleteIssue(issueID string) error {
	_, err = conn.Exec(context.Background(),
		`DELETE FROM "public.issues" WHERE issue=($1)`, issueID)

	if err != nil {
		log.Error.Printf("Error delete issue (func DeleteIssue()): %s\n", err)
		return err
	}

	return nil
}