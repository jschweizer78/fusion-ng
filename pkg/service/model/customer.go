package model

import "github.com/jschweizer78/fusion-ng/pkg/service/model/question"

// Customer that we are managing
type Customer struct {
	Name    string   `json:"name"`
	Domain  string   `json:"domain" storm:"id"`
	Address string   `json:"address"`
	Sites   []*Sites `json:"sites"`
	Admins  []string `json:"managers"` // ID of administrators
}

// CustomerQuestion meta data
type CustomerQuestion struct {
	Name    *question.Textbox     `json:"name"`
	Domain  *question.Textbox     `json:"domain"`
	Address *question.Textbox     `json:"address"`
	Sites   *question.MultiSelect `json:"sites"`
	Admins  *question.MultiSelect `json:"managers"`
}

// Sites that we are managing
type Sites struct {
	Name      string `json:"name" storm:"name"`
	ShortName string `json:"shortName" storm:"id"`
	Address   string `json:"address" storm:"adress"`
	Subnet    string `json:"subnet"`
	IDFs      []*IDF
}

// IDF that we are managing
type IDF struct {
	Name     string `json:"name" storm:"id"`
	Building string `json:"building"`
	Room     string `json:"room"`
}
