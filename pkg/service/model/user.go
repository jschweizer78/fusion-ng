package model

import "github.com/jschweizer78/fusion-ng/pkg/service/model/question"

// User of the application
type User struct {
	Email     string                   `json:"email" storm:"id"`
	Customers map[string]*UserCustomer `json:"customers"`
}

// UserQuestion meta data
type UserQuestion struct {
	Email     *question.Textbox  `json:"email"`
	Customers *question.DropDown `json:"customers"`
}

func NewUserQuestions() *UserQuestion {
	eBase := question.NewBase("email", "Email", "text", 0, true)
	email := question.NewTextBox(eBase, "email")
	cBase := question.NewBase("customers", "Customers", "dropdown", 1, false)
	cOptions := make([]*question.SelectOptions, 0, 10)
	cOptions = append(cOptions, &question.SelectOptions{
		Key:   "new",
		Value: "New",
	})
	customers := question.NewDropDown(cBase, cOptions)
	return &UserQuestion{email, customers}
}

// UserCustomer used to link an admin to a customer
type UserCustomer struct {
	CustomerID string `json:"customerID"`
	AuthToken  string `json:"authToken"`
}

// UserCustomerQuestion meta data
type UserCustomerQuestion struct {
	CustomerID *question.Textbox `json:"customerID"`
	AuthToken  *question.Textbox `json:"authToken"`
}
