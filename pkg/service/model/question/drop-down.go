package question

// DropDown a type of question
type DropDown struct {
	*Base
	Options []*SelectOptions `json:"options"`
}

// SelectOptions to provide options for a drop down
type SelectOptions struct {
	Key   string
	Value string
}

// MultiSelect a type of question
type MultiSelect struct {
	*DropDown
	Multiple bool
}

/*
//Options for reading the options
func (dd *DropDown) Options() []*SelectOptions {
	return dd.options
}
*/

// NewDropDown for creating a new drop down
func NewDropDown(root *Base, options []*SelectOptions) *DropDown {
	if options == nil {
		options = make([]*SelectOptions, 0, 10)
	}
	return &DropDown{root, options}
}

// NewMultiSelect for creating a new drop down
func NewMultiSelect(root *Base, options []*SelectOptions) *MultiSelect {
	if options == nil {
		options = make([]*SelectOptions, 0, 10)
	}
	return &MultiSelect{&DropDown{root, options}, true}
}
