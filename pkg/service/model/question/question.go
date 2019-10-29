package question

// Question is a base interface to confirm to a question UI box
type Question interface {
	Key() string
	Label() string
	Order() uint32
	Required() bool
	ControlType() string
}

// NewBase for creating a new base question
func NewBase(key, label, controlType string, order uint32, required bool) *Base {
	return &Base{key, label, order, required, controlType}
}

// Base is a base type to confirm to a question interface
type Base struct {
	Key         string `json:"key"`
	Label       string `json:"label"`
	Order       uint32 `json:"order"`
	Required    bool   `json:"required"`
	ControlType string `json:"controlType"`
}
