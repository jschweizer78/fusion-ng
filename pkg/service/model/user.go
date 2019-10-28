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
