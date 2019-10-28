package question

// Textbox for basic textboxes
type Textbox struct {
	*Base
	kind string
}

// Type to get the type of text box
func (tb *Textbox) Type() string {
	return tb.kind
}

// NewTextBox for creating a new textbox
func NewTextBox(root *Base, kind string) *Textbox {
	return &Textbox{root, kind}
}
