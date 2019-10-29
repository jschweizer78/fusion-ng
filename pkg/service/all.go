package service

import "github.com/asdine/storm/v3"

const (
	// DBUsers database collections
	DBUsers = "users"
	// DBCustomers database collections
	DBCustomers = "customers"
)

// SrvAll to manager all APIs
type SrvAll struct {
	users     *SrvUsers
	customers *SrvCustomer
}

// NewSrvAll to manager all APIs
func NewSrvAll(db *storm.DB) *SrvAll {
	userNode := db.From(DBUsers)
	customerNode := db.From(DBCustomers)
	return &SrvAll{
		NewSrvUsers(userNode),
		NewSrvCustomer(customerNode),
	}
}

// Users To interact with users
func (sa *SrvAll) Users() *SrvUsers {
	return sa.users
}

// Customers To interact with users
func (sa *SrvAll) Customers() *SrvCustomer {
	return sa.customers
}
