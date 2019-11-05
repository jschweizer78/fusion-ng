package service

import "github.com/asdine/storm/v3"

const (
	// DBUsers database collections
	DBUsers = "users"
	// DBCustomers database collections
	DBCustomers = "customers"
)

// Servicer a standard interface for Services
type Servicer interface {
	Name() string
}

// PaginationOptions for frontend to send pagintaion optons
type PaginationOptions struct {
	Skip  int `json:"skip"`
	Limit int `json:"Limit"`
}

// SrvAll to manager all APIs
type SrvAll struct {
	services map[string]Servicer
}

// NewSrvAll to manager all APIs
func NewSrvAll(db *storm.DB) *SrvAll {
	// get storm DB node (collection/table).
	userNode := db.From(DBUsers)
	userSrv := NewSrvUsers(userNode)

	customerNode := db.From(DBCustomers)
	customerSrv := NewSrvCustomer(customerNode)

	services := make(map[string]Servicer)

	services[DBUsers] = userSrv
	services[DBCustomers] = customerSrv
	return &SrvAll{services}
}

// Users To interact with users
func (sa *SrvAll) Users() *SrvUsers {
	base := sa.services[DBUsers]
	user := base.(*SrvUsers)
	return user
}

// Customers To interact with users
func (sa *SrvAll) Customers() *SrvCustomer {
	base := sa.services[DBCustomers]
	customer := base.(*SrvCustomer)
	return customer
}

// GetAll return the map of Servicer, then you will need to type assert
func (sa *SrvAll) GetAll() map[string]Servicer {
	return sa.services
}
