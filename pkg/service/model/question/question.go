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
	key         string
	label       string
	order       uint32
	required    bool
	controlType string
}

// Key is for interface
func (qb *Base) Key() string {
	return qb.key
}

// Label is for interface
func (qb *Base) Label() string {
	return qb.label
}

// Order is for interface
func (qb *Base) Order() uint32 {
	return qb.order
}

// Required is for interface
func (qb *Base) Required() bool {
	return qb.required
}

// ControlType is for interface
func (qb *Base) ControlType() string {
	return qb.controlType
}
