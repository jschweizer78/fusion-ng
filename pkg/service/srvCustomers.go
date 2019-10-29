package service

import (
	"github.com/asdine/storm/v3"
	"github.com/jschweizer78/fusion-ng/pkg/service/model"
)

// SrvCustomer is a service for Customers
type SrvCustomer struct {
	DB       storm.Node
	MetaData *model.CustomerQuestion
}

// NewSrvCustomer for Customer CRUD
func NewSrvCustomer(db storm.Node) *SrvCustomer {
	return &SrvCustomer{
		DB:       db,
		MetaData: &model.CustomerQuestion{},
	}
}

// Name to comply with interface
func (sc *SrvCustomer) Name() string {
	return DBCustomers
}
